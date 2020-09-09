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
library Precompiles {
    function keccakF(uint256[25] memory a) internal pure returns (uint256[25] memory) {
        uint256[5] memory c;
        uint256[5] memory d;
        //uint D_0; uint D_1; uint D_2; uint D_3; uint D_4;
        uint256[25] memory b;

        uint256[24] memory rc = [
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

            c[0] = a[0] ^ a[1] ^ a[2] ^ a[3] ^ a[4];
            c[1] = a[5] ^ a[6] ^ a[7] ^ a[8] ^ a[9];
            c[2] = a[10] ^ a[11] ^ a[12] ^ a[13] ^ a[14];
            c[3] = a[15] ^ a[16] ^ a[17] ^ a[18] ^ a[19];
            c[4] = a[20] ^ a[21] ^ a[22] ^ a[23] ^ a[24];

            /*
            for( x = 0 ; x < 5 ; x++ ) {
                D[x] = C[(x+4)%5]^((C[(x+1)%5] * 2)&0xffffffffffffffff | (C[(x+1)%5]/(2**63)));
            }*/

            d[0] = c[4] ^ (((c[1] * 2) & 0xffffffffffffffff) | (c[1] / (2**63)));
            d[1] = c[0] ^ (((c[2] * 2) & 0xffffffffffffffff) | (c[2] / (2**63)));
            d[2] = c[1] ^ (((c[3] * 2) & 0xffffffffffffffff) | (c[3] / (2**63)));
            d[3] = c[2] ^ (((c[4] * 2) & 0xffffffffffffffff) | (c[4] / (2**63)));
            d[4] = c[3] ^ (((c[0] * 2) & 0xffffffffffffffff) | (c[0] / (2**63)));

            /*
            for( x = 0 ; x < 5 ; x++ ) {
                for( y = 0 ; y < 5 ; y++ ) {
                    A[5*x+y] = A[5*x+y] ^ D[x];
                }
            }*/

            a[0] = a[0] ^ d[0];
            a[1] = a[1] ^ d[0];
            a[2] = a[2] ^ d[0];
            a[3] = a[3] ^ d[0];
            a[4] = a[4] ^ d[0];
            a[5] = a[5] ^ d[1];
            a[6] = a[6] ^ d[1];
            a[7] = a[7] ^ d[1];
            a[8] = a[8] ^ d[1];
            a[9] = a[9] ^ d[1];
            a[10] = a[10] ^ d[2];
            a[11] = a[11] ^ d[2];
            a[12] = a[12] ^ d[2];
            a[13] = a[13] ^ d[2];
            a[14] = a[14] ^ d[2];
            a[15] = a[15] ^ d[3];
            a[16] = a[16] ^ d[3];
            a[17] = a[17] ^ d[3];
            a[18] = a[18] ^ d[3];
            a[19] = a[19] ^ d[3];
            a[20] = a[20] ^ d[4];
            a[21] = a[21] ^ d[4];
            a[22] = a[22] ^ d[4];
            a[23] = a[23] ^ d[4];
            a[24] = a[24] ^ d[4];

            /*Rho and pi steps*/
            b[0] = a[0];
            b[8] = (((a[1] * (2**36)) & 0xffffffffffffffff) | (a[1] / (2**28)));
            b[11] = (((a[2] * (2**3)) & 0xffffffffffffffff) | (a[2] / (2**61)));
            b[19] = (((a[3] * (2**41)) & 0xffffffffffffffff) | (a[3] / (2**23)));
            b[22] = (((a[4] * (2**18)) & 0xffffffffffffffff) | (a[4] / (2**46)));
            b[2] = (((a[5] * (2**1)) & 0xffffffffffffffff) | (a[5] / (2**63)));
            b[5] = (((a[6] * (2**44)) & 0xffffffffffffffff) | (a[6] / (2**20)));
            b[13] = (((a[7] * (2**10)) & 0xffffffffffffffff) | (a[7] / (2**54)));
            b[16] = (((a[8] * (2**45)) & 0xffffffffffffffff) | (a[8] / (2**19)));
            b[24] = (((a[9] * (2**2)) & 0xffffffffffffffff) | (a[9] / (2**62)));
            b[4] = (((a[10] * (2**62)) & 0xffffffffffffffff) | (a[10] / (2**2)));
            b[7] = (((a[11] * (2**6)) & 0xffffffffffffffff) | (a[11] / (2**58)));
            b[10] = (((a[12] * (2**43)) & 0xffffffffffffffff) | (a[12] / (2**21)));
            b[18] = (((a[13] * (2**15)) & 0xffffffffffffffff) | (a[13] / (2**49)));
            b[21] = (((a[14] * (2**61)) & 0xffffffffffffffff) | (a[14] / (2**3)));
            b[1] = (((a[15] * (2**28)) & 0xffffffffffffffff) | (a[15] / (2**36)));
            b[9] = (((a[16] * (2**55)) & 0xffffffffffffffff) | (a[16] / (2**9)));
            b[12] = (((a[17] * (2**25)) & 0xffffffffffffffff) | (a[17] / (2**39)));
            b[15] = (((a[18] * (2**21)) & 0xffffffffffffffff) | (a[18] / (2**43)));
            b[23] = (((a[19] * (2**56)) & 0xffffffffffffffff) | (a[19] / (2**8)));
            b[3] = (((a[20] * (2**27)) & 0xffffffffffffffff) | (a[20] / (2**37)));
            b[6] = (((a[21] * (2**20)) & 0xffffffffffffffff) | (a[21] / (2**44)));
            b[14] = (((a[22] * (2**39)) & 0xffffffffffffffff) | (a[22] / (2**25)));
            b[17] = (((a[23] * (2**8)) & 0xffffffffffffffff) | (a[23] / (2**56)));
            b[20] = (((a[24] * (2**14)) & 0xffffffffffffffff) | (a[24] / (2**50)));

            /*Xi state*/
            /*
            for( x = 0 ; x < 5 ; x++ ) {
                for( y = 0 ; y < 5 ; y++ ) {
                    A[5*x+y] = B[5*x+y]^((~B[5*((x+1)%5)+y]) & B[5*((x+2)%5)+y]);
                }
            }*/

            a[0] = b[0] ^ ((~b[5]) & b[10]);
            a[1] = b[1] ^ ((~b[6]) & b[11]);
            a[2] = b[2] ^ ((~b[7]) & b[12]);
            a[3] = b[3] ^ ((~b[8]) & b[13]);
            a[4] = b[4] ^ ((~b[9]) & b[14]);
            a[5] = b[5] ^ ((~b[10]) & b[15]);
            a[6] = b[6] ^ ((~b[11]) & b[16]);
            a[7] = b[7] ^ ((~b[12]) & b[17]);
            a[8] = b[8] ^ ((~b[13]) & b[18]);
            a[9] = b[9] ^ ((~b[14]) & b[19]);
            a[10] = b[10] ^ ((~b[15]) & b[20]);
            a[11] = b[11] ^ ((~b[16]) & b[21]);
            a[12] = b[12] ^ ((~b[17]) & b[22]);
            a[13] = b[13] ^ ((~b[18]) & b[23]);
            a[14] = b[14] ^ ((~b[19]) & b[24]);
            a[15] = b[15] ^ ((~b[20]) & b[0]);
            a[16] = b[16] ^ ((~b[21]) & b[1]);
            a[17] = b[17] ^ ((~b[22]) & b[2]);
            a[18] = b[18] ^ ((~b[23]) & b[3]);
            a[19] = b[19] ^ ((~b[24]) & b[4]);
            a[20] = b[20] ^ ((~b[0]) & b[5]);
            a[21] = b[21] ^ ((~b[1]) & b[6]);
            a[22] = b[22] ^ ((~b[2]) & b[7]);
            a[23] = b[23] ^ ((~b[3]) & b[8]);
            a[24] = b[24] ^ ((~b[4]) & b[9]);

            /*Last step*/
            a[0] = a[0] ^ rc[i];
        }

        return a;
    }
}
