/*
Implementation by Gilles Van Assche, hereby denoted as "the implementer".

For more information, feedback or questions, please refer to our website:
https://keccak.team/

To the extent possible under law, the implementer has waived all copyright
and related or neighboring rights to the source code in this file.
http://creativecommons.org/publicdomain/zero/1.0/
*/

#ifndef _SimpleFIPS202_h_
#define _SimpleFIPS202_h_

#include "KeccakSpongeWidth1600.h"
#include <string.h>

/** Implementation of the SHAKE128 extendable output function (XOF) [FIPS 202].
  * @param  output          Pointer to the output buffer.
  * @param  outputByteLen   The desired number of output bytes.
  * @param  input           Pointer to the input message.
  * @param  inputByteLen    The length of the input message in bytes.
  * @return 0 if successful, 1 otherwise.
  */
int SHAKE128(unsigned char *output, size_t outputByteLen, const unsigned char *input, size_t inputByteLen);

/** Implementation of the SHAKE256 extendable output function (XOF) [FIPS 202].
  * @param  output          Pointer to the output buffer.
  * @param  outputByteLen   The desired number of output bytes.
  * @param  input           Pointer to the input message.
  * @param  inputByteLen    The length of the input message in bytes.
  * @return 0 if successful, 1 otherwise.
  */
int SHAKE256(unsigned char *output, size_t outputByteLen, const unsigned char *input, size_t inputByteLen);

/** Implementation of SHA3-224 [FIPS 202].
  * @param  output          Pointer to the output buffer (28 bytes).
  * @param  input           Pointer to the input message.
  * @param  inputByteLen    The length of the input message in bytes.
  * @return 0 if successful, 1 otherwise.
  */
int SHA3_224(unsigned char *output, const unsigned char *input, size_t inputByteLen);

/** Implementation of SHA3-256 [FIPS 202].
  * @param  output          Pointer to the output buffer (32 bytes).
  * @param  input           Pointer to the input message.
  * @param  inputByteLen    The length of the input message in bytes.
  * @return 0 if successful, 1 otherwise.
  */
int SHA3_256(unsigned char *output, const unsigned char *input, size_t inputByteLen);

/** Implementation of SHA3-384 [FIPS 202].
  * @param  output          Pointer to the output buffer (48 bytes).
  * @param  input           Pointer to the input message.
  * @param  inputByteLen    The length of the input message in bytes.
  * @return 0 if successful, 1 otherwise.
  */
int SHA3_384(unsigned char *output, const unsigned char *input, size_t inputByteLen);

/** Implementation of SHA3-512 [FIPS 202].
  * @param  output          Pointer to the output buffer (64 bytes).
  * @param  input           Pointer to the input message.
  * @param  inputByteLen    The length of the input message in bytes.
  * @return 0 if successful, 1 otherwise.
  */
int SHA3_512(unsigned char *output, const unsigned char *input, size_t inputByteLen);

#endif
