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

///      This algorithm has been extracted from the implementation of smart pool (https://github.com/smartpool)
library Keccak {
    function keccak_f(uint256[25] memory A)
        internal
        pure
        returns (uint256[25] memory)
    {
        uint256[5] memory C;
        uint256[5] memory D;
        //uint D_0; uint D_1; uint D_2; uint D_3; uint D_4;
        uint256[25] memory B;

        uint256[24] memory RC = [
            uint256(0x0000000000000001),
            0x0000000000008082,
            0x800000000000808A,
            0x8000000080008000,
            0x000000000000808B,
            0x0000000080000001,
            0x8000000080008081,
            0x8000000000008009,
            0x000000000000008A,
            0x0000000000000088,
            0x0000000080008009,
            0x000000008000000A,
            0x000000008000808B,
            0x800000000000008B,
            0x8000000000008089,
            0x8000000000008003,
            0x8000000000008002,
            0x8000000000000080,
            0x000000000000800A,
            0x800000008000000A,
            0x8000000080008081,
            0x8000000000008080,
            0x0000000080000001,
            0x8000000080008008
        ];

        for (uint256 i = 0; i < 24; i++) {
            /*
            for( x = 0 ; x < 5 ; x++ ) {
                C[x] = A[5*x]^A[5*x+1]^A[5*x+2]^A[5*x+3]^A[5*x+4];
            }*/

            C[0] = A[0] ^ A[1] ^ A[2] ^ A[3] ^ A[4];
            C[1] = A[5] ^ A[6] ^ A[7] ^ A[8] ^ A[9];
            C[2] = A[10] ^ A[11] ^ A[12] ^ A[13] ^ A[14];
            C[3] = A[15] ^ A[16] ^ A[17] ^ A[18] ^ A[19];
            C[4] = A[20] ^ A[21] ^ A[22] ^ A[23] ^ A[24];

            /*
            for( x = 0 ; x < 5 ; x++ ) {
                D[x] = C[(x+4)%5]^((C[(x+1)%5] * 2)&0xffffffffffffffff | (C[(x+1)%5]/(2**63)));
            }*/

            D[0] =
                C[4] ^
                (((C[1] * 2) & 0xffffffffffffffff) | (C[1] / (2**63)));
            D[1] =
                C[0] ^
                (((C[2] * 2) & 0xffffffffffffffff) | (C[2] / (2**63)));
            D[2] =
                C[1] ^
                (((C[3] * 2) & 0xffffffffffffffff) | (C[3] / (2**63)));
            D[3] =
                C[2] ^
                (((C[4] * 2) & 0xffffffffffffffff) | (C[4] / (2**63)));
            D[4] =
                C[3] ^
                (((C[0] * 2) & 0xffffffffffffffff) | (C[0] / (2**63)));

            /*
            for( x = 0 ; x < 5 ; x++ ) {
                for( y = 0 ; y < 5 ; y++ ) {
                    A[5*x+y] = A[5*x+y] ^ D[x];
                }
            }*/

            A[0] = A[0] ^ D[0];
            A[1] = A[1] ^ D[0];
            A[2] = A[2] ^ D[0];
            A[3] = A[3] ^ D[0];
            A[4] = A[4] ^ D[0];
            A[5] = A[5] ^ D[1];
            A[6] = A[6] ^ D[1];
            A[7] = A[7] ^ D[1];
            A[8] = A[8] ^ D[1];
            A[9] = A[9] ^ D[1];
            A[10] = A[10] ^ D[2];
            A[11] = A[11] ^ D[2];
            A[12] = A[12] ^ D[2];
            A[13] = A[13] ^ D[2];
            A[14] = A[14] ^ D[2];
            A[15] = A[15] ^ D[3];
            A[16] = A[16] ^ D[3];
            A[17] = A[17] ^ D[3];
            A[18] = A[18] ^ D[3];
            A[19] = A[19] ^ D[3];
            A[20] = A[20] ^ D[4];
            A[21] = A[21] ^ D[4];
            A[22] = A[22] ^ D[4];
            A[23] = A[23] ^ D[4];
            A[24] = A[24] ^ D[4];

            /*Rho and pi steps*/
            B[0] = A[0];
            B[8] = (((A[1] * (2**36)) & 0xffffffffffffffff) | (A[1] / (2**28)));
            B[11] = (((A[2] * (2**3)) & 0xffffffffffffffff) | (A[2] / (2**61)));
            B[19] = (((A[3] * (2**41)) & 0xffffffffffffffff) |
                (A[3] / (2**23)));
            B[22] = (((A[4] * (2**18)) & 0xffffffffffffffff) |
                (A[4] / (2**46)));
            B[2] = (((A[5] * (2**1)) & 0xffffffffffffffff) | (A[5] / (2**63)));
            B[5] = (((A[6] * (2**44)) & 0xffffffffffffffff) | (A[6] / (2**20)));
            B[13] = (((A[7] * (2**10)) & 0xffffffffffffffff) |
                (A[7] / (2**54)));
            B[16] = (((A[8] * (2**45)) & 0xffffffffffffffff) |
                (A[8] / (2**19)));
            B[24] = (((A[9] * (2**2)) & 0xffffffffffffffff) | (A[9] / (2**62)));
            B[4] = (((A[10] * (2**62)) & 0xffffffffffffffff) |
                (A[10] / (2**2)));
            B[7] = (((A[11] * (2**6)) & 0xffffffffffffffff) |
                (A[11] / (2**58)));
            B[10] = (((A[12] * (2**43)) & 0xffffffffffffffff) |
                (A[12] / (2**21)));
            B[18] = (((A[13] * (2**15)) & 0xffffffffffffffff) |
                (A[13] / (2**49)));
            B[21] = (((A[14] * (2**61)) & 0xffffffffffffffff) |
                (A[14] / (2**3)));
            B[1] = (((A[15] * (2**28)) & 0xffffffffffffffff) |
                (A[15] / (2**36)));
            B[9] = (((A[16] * (2**55)) & 0xffffffffffffffff) |
                (A[16] / (2**9)));
            B[12] = (((A[17] * (2**25)) & 0xffffffffffffffff) |
                (A[17] / (2**39)));
            B[15] = (((A[18] * (2**21)) & 0xffffffffffffffff) |
                (A[18] / (2**43)));
            B[23] = (((A[19] * (2**56)) & 0xffffffffffffffff) |
                (A[19] / (2**8)));
            B[3] = (((A[20] * (2**27)) & 0xffffffffffffffff) |
                (A[20] / (2**37)));
            B[6] = (((A[21] * (2**20)) & 0xffffffffffffffff) |
                (A[21] / (2**44)));
            B[14] = (((A[22] * (2**39)) & 0xffffffffffffffff) |
                (A[22] / (2**25)));
            B[17] = (((A[23] * (2**8)) & 0xffffffffffffffff) |
                (A[23] / (2**56)));
            B[20] = (((A[24] * (2**14)) & 0xffffffffffffffff) |
                (A[24] / (2**50)));

            /*Xi state*/
            /*
            for( x = 0 ; x < 5 ; x++ ) {
                for( y = 0 ; y < 5 ; y++ ) {
                    A[5*x+y] = B[5*x+y]^((~B[5*((x+1)%5)+y]) & B[5*((x+2)%5)+y]);
                }
            }*/

            A[0] = B[0] ^ ((~B[5]) & B[10]);
            A[1] = B[1] ^ ((~B[6]) & B[11]);
            A[2] = B[2] ^ ((~B[7]) & B[12]);
            A[3] = B[3] ^ ((~B[8]) & B[13]);
            A[4] = B[4] ^ ((~B[9]) & B[14]);
            A[5] = B[5] ^ ((~B[10]) & B[15]);
            A[6] = B[6] ^ ((~B[11]) & B[16]);
            A[7] = B[7] ^ ((~B[12]) & B[17]);
            A[8] = B[8] ^ ((~B[13]) & B[18]);
            A[9] = B[9] ^ ((~B[14]) & B[19]);
            A[10] = B[10] ^ ((~B[15]) & B[20]);
            A[11] = B[11] ^ ((~B[16]) & B[21]);
            A[12] = B[12] ^ ((~B[17]) & B[22]);
            A[13] = B[13] ^ ((~B[18]) & B[23]);
            A[14] = B[14] ^ ((~B[19]) & B[24]);
            A[15] = B[15] ^ ((~B[20]) & B[0]);
            A[16] = B[16] ^ ((~B[21]) & B[1]);
            A[17] = B[17] ^ ((~B[22]) & B[2]);
            A[18] = B[18] ^ ((~B[23]) & B[3]);
            A[19] = B[19] ^ ((~B[24]) & B[4]);
            A[20] = B[20] ^ ((~B[0]) & B[5]);
            A[21] = B[21] ^ ((~B[1]) & B[6]);
            A[22] = B[22] ^ ((~B[2]) & B[7]);
            A[23] = B[23] ^ ((~B[3]) & B[8]);
            A[24] = B[24] ^ ((~B[4]) & B[9]);

            /*Last step*/
            A[0] = A[0] ^ RC[i];
        }

        return A;
    }
}
