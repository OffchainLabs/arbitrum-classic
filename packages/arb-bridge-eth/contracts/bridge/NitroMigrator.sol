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
import "@openzeppelin/contracts/proxy/ProxyAdmin.sol";

pragma solidity ^0.6.11;

interface INitroBridge is IBridge {
    function acceptFundsFromOldBridge() external payable;
}

interface INitroInbox is IInbox {
    function postUpgradeInit(INitroBridge) external;
}

interface INitroRollup {
    function bridge() external view returns (INitroBridge);

    function inbox() external view returns (INitroInbox);

    function setInbox(IInbox newInbox) external;
}

contract NitroMigrator is Ownable {
    uint8 internal constant L1MessageType_shutdownForNitro = 128;

    Inbox public inbox;
    SequencerInbox public sequencerInbox;
    Bridge public bridge;
    RollupEventBridge public rollupEventBridge;
    OldOutbox public outboxV1;
    Outbox public outboxV2;
    // assumed this contract is now the rollup admin
    RollupAdminFacet public rollup;

    INitroBridge public nitroBridge;

    /// @dev The nitro migration includes various steps.
    ///
    /// > Uninitialized is before the rollup contract addresses have been set.
    ///
    /// > Step0 is setup with contract deployments and integrity checks
    ///
    /// This is the setup for the upgrade, where all contracts are deployed but the upgrade migration hasn't started yet.
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
        Uninitialized,
        Step0,
        Step1,
        Step2,
        Step3
    }
    NitroMigrationSteps public latestCompleteStep;

    constructor() public Ownable() {
        latestCompleteStep = NitroMigrationSteps.Uninitialized;
    }

    function configureDeployment(
        Inbox _inbox,
        SequencerInbox _sequencerInbox,
        Bridge _bridge,
        RollupEventBridge _rollupEventBridge,
        OldOutbox _outboxV1,
        Outbox _outboxV2,
        RollupAdminFacet _rollup,
        INitroRollup nitroRollup,
        ProxyAdmin proxyAdmin
    ) external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Uninitialized, "WRONG_STEP");

        inbox = _inbox;
        sequencerInbox = _sequencerInbox;
        bridge = _bridge;
        rollupEventBridge = _rollupEventBridge;
        rollup = _rollup;
        outboxV1 = _outboxV1;
        outboxV2 = _outboxV2;

        nitroBridge = nitroRollup.bridge();

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
        require(Address.isContract(address(nitroBridge)), "NITRO_BRIDGE_NOT_CONTRACT");

        // Upgrade the classic inbox to the nitro inbox's impl,
        // and configure nitro to use the classic inbox's address.
        INitroInbox oldNitroInbox = nitroRollup.inbox();
        address nitroInboxImpl = proxyAdmin.getProxyImplementation(
            TransparentUpgradeableProxy(payable(address(oldNitroInbox)))
        );
        nitroBridge.setInbox(address(oldNitroInbox), false);
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(inbox))),
            nitroInboxImpl,
            abi.encodeWithSelector(INitroInbox.postUpgradeInit.selector, nitroBridge)
        );
        nitroRollup.setInbox(inbox);

        latestCompleteStep = NitroMigrationSteps.Step0;
    }

    /// @dev this assumes this contract owns the rollup/inboxes/bridge before this function is called (else it will revert)
    /// this will create the final input in the inbox, but there won't be the final assertion available yet.
    /// it is assumed that at this point the sequencer has stopped receiving txs and has posted its final batch on-chain
    /// Before this step the ownership of the Rollup and Bridge  must have been transferred to this contract
    // CHRIS: TODO: remove bridge data
    function nitroStep1(address[] calldata seqAddresses, bytes calldata bridgeData)
        external
        onlyOwner
    {
        require(latestCompleteStep == NitroMigrationSteps.Step0, "WRONG_STEP");

        // check that ownership of the bridge and rollup has been transferred
        require(rollup.owner() == address(this), "ROLLUP_OWNER_NOT_SET");
        require(bridge.owner() == address(this), "BRIDGE_OWNER_NOT_SET");

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
            // TODO: import nitro contracts and use interface
            (bool success, ) = bridge.executeCall(
                address(nitroBridge),
                bal,
                abi.encodeWithSelector(INitroBridge.acceptFundsFromOldBridge.selector)
            );
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

        // this speeds up the process allowing validators to post assertions more frequently
        rollup.setMinimumAssertionPeriod(4);

        // TODO: remove permissions from gas refunder to current sequencer inbox
        latestCompleteStep = NitroMigrationSteps.Step1;
    }

    /// @dev this assumes step 1 has executed succesfully and that a validator has made the final assertion that includes the inbox shutdownForNitro
    function nitroStep2(
        uint256 finalNodeNum,
        bool destroyAlternatives,
        bool destroyChallenges
    ) external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step1, "WRONG_STEP");
        rollup.shutdownForNitro(finalNodeNum, destroyAlternatives, destroyChallenges);
        bridge.setInbox(address(rollupEventBridge), false);
        latestCompleteStep = NitroMigrationSteps.Step2;
    }

    function nitroStep3() external onlyOwner {
        require(latestCompleteStep == NitroMigrationSteps.Step2, "WRONG_STEP");
        // CHRIS: TODO: destroying the node in the previous steps does not reset latestConfirmed/latestNodeCreated
        require(
            rollup.latestConfirmed() == rollup.latestNodeCreated(),
            "ROLLUP_SHUTDOWN_NOT_COMPLETE"
        );

        nitroBridge.setInbox(address(inbox), true);
        nitroBridge.setOutbox(address(outboxV1), true);
        nitroBridge.setOutbox(address(outboxV2), true);

        latestCompleteStep = NitroMigrationSteps.Step3;
    }

    /// @dev allows the owner to do arbitrary calls. This is useful in case an unexpected event
    /// happens and we need to react to it using the migrator.
    /// This should be enough to recover from any unexpected state since no external contracts rely
    /// on this contract's state (ie `latestCompleteStep`).
    /// If other contracts relied on this, we'd need to use a delegate call instead.
    function executeTransaction(
        bytes calldata data,
        address destination,
        uint256 amount
    ) external payable onlyOwner {
        if (data.length > 0) require(Address.isContract(destination), "NO_CODE_AT_ADDR");
        (bool success, ) = destination.call{ value: amount }(data);
        if (!success) {
            assembly {
                let ptr := mload(0x40)
                let size := returndatasize()
                returndatacopy(ptr, 0, size)
                revert(ptr, size)
            }
        }
    }

    function transferOtherContractOwnership(Ownable ownable, address newOwner)
        external
        payable
        onlyOwner
    {
        ownable.transferOwnership(newOwner);
    }
}
