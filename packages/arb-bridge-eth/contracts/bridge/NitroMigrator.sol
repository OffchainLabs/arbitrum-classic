// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import "./Bridge.sol";
import "./Outbox.sol";
import "./Inbox.sol";
import "./SequencerInbox.sol";
import "./Old_Outbox/OldOutbox.sol";
import "../rollup/facets/RollupAdmin.sol";
import "../rollup/RollupEventBridge.sol";
import "../rollup/RollupLib.sol";
import "../libraries/NitroReadyQuery.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";

pragma solidity ^0.6.11;

contract NitroMigrator is Ownable {
    uint8 internal constant L1MessageType_shutdownForNitro = 128;

    Inbox public immutable inbox;
    SequencerInbox public immutable sequencerInbox;
    Bridge public immutable bridge;
    RollupEventBridge public immutable rollupEventBridge;
    OldOutbox public immutable outboxV1;
    Outbox public immutable outboxV2;
    // assumed this contract is now the rollup admin
    RollupAdminFacet public immutable rollup;

    address public immutable nitroBridge;
    address public immutable nitroOutbox;
    address public immutable nitroSequencerInbox;
    address public immutable nitroInboxLogic;

    // this is used track the message count in which the final inbox message was force included
    // initially set to max uint256. after step 1 its set to sequencer's inbox message count
    uint256 public messageCountWithHalt;

    /// @dev The nitro migration includes various steps.
    /// Step0 is setup with contract deployments and integrity checks
    /// Step1 is settling the current state of inputs to the system (bridge, delayed and sequencer inbox) in a consistent way
    /// Step2 is settling the current state of outputs to the system (rollup assertions) in a consistent way that includes state from step1
    /// Step3 is enabling the nitro chain from the state settled in Step1 and Step2
    /// Between steps 1 and 2 a validator needs to make the final assertion that includes the inbox shutdownForNitro message
    enum NitroMigrationSteps {
        Step0,
        Step1,
        Step2,
        Step3
    }
    NitroMigrationSteps public latestCompleteStep;

    constructor(
        Inbox _inbox,
        SequencerInbox _sequencerInbox,
        Bridge _bridge,
        RollupEventBridge _rollupEventBridge,
        OldOutbox _outboxV1,
        Outbox _outboxV2,
        RollupAdminFacet _rollup,
        address _nitroBridge,
        address _nitroOutbox,
        address _nitroSequencerInbox,
        address _nitroInboxLogic
    ) public Ownable() {
        inbox = _inbox;
        sequencerInbox = _sequencerInbox;
        bridge = _bridge;
        rollupEventBridge = _rollupEventBridge;
        rollup = _rollup;
        outboxV1 = _outboxV1;
        outboxV2 = _outboxV2;

        {
            // this contract is the rollup admin, and we want to check if the user facet is upgraded
            // so we deploy a new contract to ensure the query is dispatched to the user facet, not the admin
            NitroReadyQuery queryContract = new NitroReadyQuery();
            require(queryContract.isNitroReady(address(_rollup)), "ROLLUP_NOT_NITRO_READY");
        }

        require(_inbox.isNitroReady(), "INBOX_NOT_UPGRADED");
        require(_sequencerInbox.isNitroReady(), "SEQINBOX_NOT_UPGRADED");

        // we check that the new contracts that will receive permissions are actually contracts
        require(Address.isContract(_nitroBridge), "NITRO_BRIDGE_NOT_CONTRACT");
        require(Address.isContract(_nitroOutbox), "NITRO_OUTBOX_NOT_CONTRACT");
        require(Address.isContract(_nitroSequencerInbox), "NITRO_SEQINBOX_NOT_CONTRACT");
        require(Address.isContract(_nitroInboxLogic), "NITRO_INBOX_NOT_CONTRACT");

        nitroBridge = _nitroBridge;
        nitroOutbox = _nitroOutbox;
        nitroSequencerInbox = _nitroSequencerInbox;
        nitroInboxLogic = _nitroInboxLogic;

        // setting to max value means it won't be possible to execute step 2 before step 1
        messageCountWithHalt = type(uint256).max;
        latestCompleteStep = NitroMigrationSteps.Step0;
    }

    /// @dev this assumes this contract owns the rollup/inboxes/bridge before this function is called (else it will revert)
    /// this will create the final input in the inbox, but there won't be the final assertion available yet.
    /// it is assumed that at this point the sequencer has stopped receiving txs and has posted its final batch on-chain
    function nitroStep1(address[] calldata seqAddresses) external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step0, "WRONG_STEP");
        require(messageCountWithHalt == type(uint256).max, "STEP1_ALREADY_TRIGGERED");
        uint256 delayedMessageCount = inbox.shutdownForNitro();

        // the `bridge` won't have any enabled inboxes after nitroStep2, so force inclusion after this shouldn't be possible
        // the rollup event bridge will update the delayed accumulator after the final rollup shutdown events, but this
        // shouldn't be an issue
        bridge.setInbox(address(inbox), false);
        bridge.setOutbox(address(outboxV1), false);
        bridge.setOutbox(address(outboxV2), false);
        // we disable the rollupEventBridge later since its needed in order to create/confirm assertions
        // TODO: will the nitro node process these events from the rollup event bridge? probably not since these aren't force included.
        // is it a problem that we're dropping these delayed messages? probably not.

        bridge.setOutbox(address(this), true);

        {
            uint256 bal = address(bridge).balance;
            (bool success, ) = bridge.executeCall(nitroBridge, bal, "");
            require(success, "ESCROW_TRANSFER_FAIL");
        }

        bridge.setOutbox(address(this), false);

        // if the sequencer posted its final batch and was shutdown before `nitroStep1` there shouldn't be any reorgs
        // we could lock the sequencer inbox with `shutdownForNitro` but this wouldnt stop a reorg from accepting
        // txs in the RPC interface without posting a batch.
        // `nitroStep2` will only enforce inclusion of assertions that read up to this current point.
        sequencerInbox.shutdownForNitro(
            delayedMessageCount,
            bridge.inboxAccs(delayedMessageCount - 1),
            seqAddresses
        );

        // we can use this to verify in step 2 that the assertion includes the shutdownForNitro message
        messageCountWithHalt = sequencerInbox.messageCount();

        // TODO: remove permissions from gas refunder to current sequencer inbox
        latestCompleteStep = NitroMigrationSteps.Step1;
    }

    /// @dev this assumes step 1 has executed succesfully and that a validator has made the final assertion that includes the inbox shutdownForNitro
    function nitroStep2(
        bytes32[3] memory bytes32Fields,
        uint256[4] memory intFields,
        uint256 proposedBlock,
        bytes32 beforeSendAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths,
        uint256 afterSendCount,
        bytes32 afterLogAcc,
        uint256 afterLogCount
    ) external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step1, "WRONG_STEP");
        RollupLib.ExecutionState memory afterExecutionState = RollupLib.decodeExecutionState(
            bytes32Fields,
            intFields,
            proposedBlock,
            messageCountWithHalt
        );
        bytes32 expectedStateHash = RollupLib.stateHash(afterExecutionState);

        uint256 nodeNum = rollup.latestNodeCreated();
        // the actual nodehash doesn't matter, only its after state of execution
        bytes32 actualStateHash = rollup.getNode(nodeNum).stateHash();
        require(expectedStateHash == actualStateHash, "WRONG_STATE_HASH");

        rollup.forceConfirmNode(
            nodeNum,
            beforeSendAcc,
            sendsData,
            sendLengths,
            afterSendCount,
            afterLogAcc,
            afterLogCount
        );

        // TODO: we can forceCreate the assertion and have the rollup paused in step 1
        rollup.pause();
        // we could disable the rollup user facet so only the admin can interact with the rollup
        // would make the dispatch rollup revert when calling user facet. but easier to just pause it

        // TODO: ensure everyone is unstaked?
        // need to wait until last assertion beforeforce confirm assertion
        uint256 stakerCount = rollup.stakerCount();
        address[] memory stakers = new address[](stakerCount);
        for (uint64 i = 0; i < stakerCount; ++i) {
            stakers[i] = rollup.getStakerAddress(i);
        }
        // they now have withdrawable stake to claim
        // rollup doesn't need to be unpaused for this.
        rollup.forceRefundStaker(stakers);

        // TODO: forceResolveChallenge if any
        // TODO: double check that challenges can't be created and new stakes cant be added
        bridge.setInbox(address(rollupEventBridge), false);

        latestCompleteStep = NitroMigrationSteps.Step2;
    }

    function nitroStep3() external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step2, "WRONG_STEP");
        // enable new Bridge with funds (ie set old outboxes)
        // TODO: enable new elements of nitro chain (ie bridge, inbox, outbox, rollup, etc)
        // TODO: trigger inbox upgrade to new logic
        latestCompleteStep = NitroMigrationSteps.Step3;
    }
}
