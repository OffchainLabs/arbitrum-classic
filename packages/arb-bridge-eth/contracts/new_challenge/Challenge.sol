// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

pragma solidity ^0.5.11;

// What’s in an initial assertion:

// Preconditions:
//     N_reads_pre [total number of messages ever read from inbox], and Inbox_hash(N_reads_pre),
//     N_arrivals [total number of messages that ever arrived in the inbox], and Inbox_hash(N_arrivals),
//     H_state_pre
//     Size: N_steps

// Postconditions:
//     N_reads_post,
//     Inbox_hash(N_reads_post),
//     H_state_post,
//     H_outputs
// */

// contract Challenge {
//     // Precondition Data
//     bytes32 beforeInboxAcc;
//     uint256 beforeInboxCount;

//     bytes32 topInboxAcc;
//     uint256 topInboxCount;

//     bytes32 beforeStateHash;

//     // Postcondition Data

//     bytes32 afterInboxAcc;
//     uint256 afterInboxCount;

//     bytes32 afterStateHash;

//     uint256 numSteps;

//     /*
//     Inbox inconsistency objection: challenger claims that there is no sequence of
//     N_arrivals-N_reads_post messages that chains from the Inbox_hash(N_reads_post)
//     claimed in the postcondition to Inbox_hash(N_arrivals)

//     [switch to inbox inconsistency challenge protocol]
//     */
//     function inconsistentInbox(bytes32[] memory alternativeAfterInboxAcc) {
//         uint256 messagesCount = topInboxCount - afterInboxCount;
//     }

//     /*
//     Halt objection:  accepts inbox claim but asserts that machine halts after < N_steps

//     Challenger makes a bisected assertion with same preconditions, fewer steps

//     H_state_post must be 0 or 1 (halted states)
//     */
//     function machineShouldHaveHalted(
//         bytes32 _machineHash,
//         bytes32 _inboxAcc,
//         bytes32 _messageAcc,
//         bytes32 _logAcc,
//         uint256 _gasUsed,
//         uint256 _inboxCount,
//         uint256 _messageCount,
//         uint256 _logCount,
//         uint256 _numSteps,
//         bytes32 _afterInboxHash,
//         uint256 _afterInboxCount
//     ) {
//         require(_machineHash == 0 || _machineHash == 1);
//         require(_numSteps < numSteps);

//         bytes32 assertionHash = Assertion.hashAssertion(
//             _machineHash,
//             _inboxAcc,
//             _messageAcc,
//             _logAcc,
//             _gasUsed,
//             _inboxCount,
//             _messageCount,
//             _logCount
//         );
//     }

//     /*
//     Inbox exhaustion objection: accepts consistency of inbox claim but asserts that
//     execution would consume more messages than have arrived in the inbox

//     Challenger makes a bisected assertion with same preconditions, fewer steps,
//     N_reads_post == N_arrivals (and equal Inbox_hashes)

//     Challenger proves that next instruction at final state hash is an inbox instruction
//     or an inboxpeek instruction with the pending message empty
//     */
//     function inboxExhausted() {}

//     /*
//     Execution objection: accepts inbox claim and number of steps, disagrees about
//     correctness of postconditions

//     Challenger makes a bisected assertion with same precondition and size but
//     different postcondition hash
//     */
//     function incorrectExecution() {}
// }
