/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

contract DataPackets {
    mapping(uint => Packet) public packets;

    struct Packet {
        bytes data;
        address sender;
        uint blockHeight;
    }

    function sendDataPacket(bytes calldata encodedData, uint messageType) external {
        packets[messageType] = Packet(
            {
                data: encodedData,
                sender: msg.sender,
                blockHeight: block.number
            }
        );
    }
}
