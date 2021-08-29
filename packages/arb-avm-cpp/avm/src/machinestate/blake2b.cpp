/*
 * Copyright 2021, Offchain Labs, Inc.
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

/*
 *   Based off BLAKE2 implementation Written in 2012 by Samuel Neves <sneves@dei.uc.pt>
 *   Samuel's work is available under Creative-Commons 0 license in:
 *   https://github.com/BLAKE2/BLAKE2/blob/master/ref/blake2b-ref.c
 */

#include <avm/machinestate/blake2b.hpp>

static const std::array<uint64_t, 8> blake2b_IV =
{
    0x6a09e667f3bcc908ULL, 0xbb67ae8584caa73bULL,
    0x3c6ef372fe94f82bULL, 0xa54ff53a5f1d36f1ULL,
    0x510e527fade682d1ULL, 0x9b05688c2b3e6c1fULL,
    0x1f83d9abfb41bd6bULL, 0x5be0cd19137e2179ULL
};

static const std::array<std::array<uint8_t, 16>, 10> blake2b_sigma =
{{
    {  0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15 } ,
    { 14, 10,  4,  8,  9, 15, 13,  6,  1, 12,  0,  2, 11,  7,  5,  3 } ,
    { 11,  8, 12,  0,  5,  2, 15, 13, 10, 14,  3,  6,  7,  1,  9,  4 } ,
    {  7,  9,  3,  1, 13, 12, 11, 14,  2,  6,  5, 10,  4,  0, 15,  8 } ,
    {  9,  0,  5,  7,  2,  4, 10, 15, 14,  1, 11, 12,  6,  8,  3, 13 } ,
    {  2, 12,  6, 10,  0, 11,  8,  3,  4, 13,  7,  5, 15, 14,  1,  9 } ,
    { 12,  5,  1, 15, 14, 13,  4, 10,  0,  7,  6,  3,  9,  2,  8, 11 } ,
    { 13, 11,  7, 14, 12,  1,  3,  9,  5,  0, 15,  4,  8,  6,  2, 10 } ,
    {  6, 15, 14,  9, 11,  3,  0,  8, 12,  2, 13,  7,  1,  4, 10,  5 } ,
    { 10,  2,  8,  4,  7,  6,  1,  5, 15, 11,  9, 14,  3, 12, 13 , 0 }
}};

static inline uint64_t rotr64( const uint64_t w, const unsigned c )
{
  return ( w >> c ) | ( w << ( 64 - c ) );
}

int blake2b_F(uint32_t rounds,
    const std::array<uint64_t, 2> &c,
    bool final,
    const std::array<uint64_t, 16> &m,
    std::array<uint64_t, 8> &h)
{
    std::array<uint64_t, 16> v;

    for(int i = 0; i < 8; ++i )
        v[i] = h[i];

    v[ 8] = blake2b_IV[0];
    v[ 9] = blake2b_IV[1];
    v[10] = blake2b_IV[2];
    v[11] = blake2b_IV[3];
    v[12] = c[0] ^ blake2b_IV[4];
    v[13] = c[1] ^ blake2b_IV[5];
    if (final) {
        v[14] = ~blake2b_IV[6];
    } else {
        v[14] = blake2b_IV[6];
    }
    v[15] = blake2b_IV[7];

#define G(r,i,a,b,c,d) \
    do { \
    a = a + b + m[blake2b_sigma[r % 10][2*i+0]]; \
    d = rotr64(d ^ a, 32); \
    c = c + d; \
    b = rotr64(b ^ c, 24); \
    a = a + b + m[blake2b_sigma[r % 10][2*i+1]]; \
    d = rotr64(d ^ a, 16); \
    c = c + d; \
    b = rotr64(b ^ c, 63); \
    } while(0)

    for (uint32_t r = 0; r < rounds; r++) {
        G(r,0,v[ 0],v[ 4],v[ 8],v[12]);
        G(r,1,v[ 1],v[ 5],v[ 9],v[13]);
        G(r,2,v[ 2],v[ 6],v[10],v[14]);
        G(r,3,v[ 3],v[ 7],v[11],v[15]);
        G(r,4,v[ 0],v[ 5],v[10],v[15]);
        G(r,5,v[ 1],v[ 6],v[11],v[12]);
        G(r,6,v[ 2],v[ 7],v[ 8],v[13]);
        G(r,7,v[ 3],v[ 4],v[ 9],v[14]);
    }

    for(int i = 0; i < 8; ++i )
        h[i] = h[i] ^ v[i] ^ v[i + 8];

#undef G
    return 0;
}
