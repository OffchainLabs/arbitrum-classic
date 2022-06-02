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

    /// @dev The nitro migration includes various steps.
    ///
    /// > Step0 is setup with contract deployments and integrity checks
    ///
    /// This is the setup for the upgrade, where all contracts are deployed but the upgrade migration hasn't started yet.
    /// Before Step1 the ownership of the Inbox / Rollup / Outbox / Bridge / SequencerInbox must all be transferred to this contract
    /// The sequencer should stop receiving messages over RPC and post its final batch before Step1 is called.
    ///
    /// > Step1 is settling the current state of inputs to the system (bridge, delayed and sequencer inbox) in a consistent way
    ///
    /// The validator must now post the final assertion that executes all messages included in Step1 (ie the shutdownForNitro message)
    ///
    /// > Step2 is settling the current state of outputs to the system (rollup assertions) in a consistent way that includes state from step1. This pauses most of the rollup functionality
    ///
    /// The validator is now able to confirm all pending nodes between latestConfirmed and latestCreated. Step3 is only possible after this happens.
    ///
    /// > Step3 is enabling the nitro chain from the state settled in Step1 and Step2.
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
            require(
                queryContract.isNitroReady(address(_rollup)) == uint8(0xa4b1),
                "USER_ROLLUP_NOT_NITRO_READY"
            );
        }
        // this returns a different magic value so we can differentiate the user and admin facets
        require(_rollup.isNitroReady() == uint8(0xa4b2), "ADMIN_ROLLUP_NOT_NITRO_READY");

        require(_inbox.isNitroReady() == uint8(0xa4b1), "INBOX_NOT_UPGRADED");
        require(_sequencerInbox.isNitroReady() == uint8(0xa4b1), "SEQINBOX_NOT_UPGRADED");

        // we check that the new contracts that will receive permissions are actually contracts
        require(Address.isContract(_nitroBridge), "NITRO_BRIDGE_NOT_CONTRACT");
        require(Address.isContract(_nitroOutbox), "NITRO_OUTBOX_NOT_CONTRACT");
        require(Address.isContract(_nitroSequencerInbox), "NITRO_SEQINBOX_NOT_CONTRACT");
        require(Address.isContract(_nitroInboxLogic), "NITRO_INBOX_NOT_CONTRACT");

        nitroBridge = _nitroBridge;
        nitroOutbox = _nitroOutbox;
        nitroSequencerInbox = _nitroSequencerInbox;
        nitroInboxLogic = _nitroInboxLogic;

        latestCompleteStep = NitroMigrationSteps.Step0;
    }

    /// @dev this assumes this contract owns the rollup/inboxes/bridge before this function is called (else it will revert)
    /// this will create the final input in the inbox, but there won't be the final assertion available yet.
    /// it is assumed that at this point the sequencer has stopped receiving txs and has posted its final batch on-chain
    function nitroStep1(address[] calldata seqAddresses) external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step0, "WRONG_STEP");
        uint256 delayedMessageCount = inbox.shutdownForNitro();

        // the `bridge` won't have any enabled inboxes after nitroStep2, so force inclusion after this shouldn't be possible
        // the rollup event bridge will update the delayed accumulator after the final rollup shutdown events, but this
        // shouldn't be an issue
        bridge.setInbox(address(inbox), false);
        bridge.setOutbox(address(outboxV1), false);
        bridge.setOutbox(address(outboxV2), false);

        // we disable the rollupEventBridge later since its needed in order to create/confirm assertions
        // the rollup event bridge will still add messages to the Bridge's accumulator, but these will never be included into the sequencer inbox
        // it is not a problem that these messages will be lost, as long as classic shutdown and nitro boot are deterministic

        bridge.setOutbox(address(this), true);

        {
            uint256 bal = address(bridge).balance;
            (bool success, ) = bridge.executeCall(nitroBridge, bal, "");
            require(success, "ESCROW_TRANSFER_FAIL");
        }

        bridge.setOutbox(address(this), false);

        // if the sequencer posted its final batch and was shutdown before `nitroStep1` there shouldn't be any reorgs
        // even though we remove the seqAddr from the sequencer inbox with `shutdownForNitro` this wouldnt stop a reorg from
        // the sequencer accepting txs in the RPC interface without posting a batch.
        // `nitroStep2` will only enforce inclusion of assertions that read up to this current point.
        sequencerInbox.shutdownForNitro(
            delayedMessageCount,
            bridge.inboxAccs(delayedMessageCount - 1),
            seqAddresses
        );

        // TODO: remove permissions from gas refunder to current sequencer inbox
        latestCompleteStep = NitroMigrationSteps.Step1;
    }

    /// @dev this assumes step 1 has executed succesfully and that a validator has made the final assertion that includes the inbox shutdownForNitro
    function nitroStep2(uint256 finalNodeNum) external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step1, "WRONG_STEP");
        rollup.shutdownForNitro(finalNodeNum);
        latestCompleteStep = NitroMigrationSteps.Step2;
    }

    function nitroStep3() external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step2, "WRONG_STEP");
        require(
            rollup.latestConfirmed() == rollup.latestNodeCreated(),
            "ROLLUP_SHUTDOWN_NOT_COMPLETE"
        );
        bridge.setInbox(address(rollupEventBridge), false);

        // enable new Bridge with funds (ie set old outboxes)
        // TODO: enable new elements of nitro chain (ie bridge, inbox, outbox, rollup, etc)
        // TODO: trigger inbox upgrade to new logic

        latestCompleteStep = NitroMigrationSteps.Step3;
    }
}
