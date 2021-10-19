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

#include <avm/machinestate/machineoperation.hpp>

#include <avm/machinestate/ecops.hpp>
#include <avm/machinestate/machinestate.hpp>

#include <PicoSHA2/picosha2.h>
#include <ethash/keccak.h>
#include <secp256k1_recovery.h>
#include <ethash/keccak.hpp>

#include <netdb.h>
#include <stdlib.h>
#include <fstream>
#include <iostream>

#include "tee/eigentee.h"
#include "tee/tee_task.hpp"

eigen_enclave_info_t* g_enclave_info = NULL;
eigen_auditor_set_t* g_auditors = NULL;
int32_t g_tms_port = 8082;

int submit_task(const char* method,
                const char* args,
                const char* uid,
                const char* token,
                char** output,
                size_t* output_size) {
    struct sockaddr_in tms_addr;
    char recvbuf[2048] = {0};
    int ret;
    struct hostent* hptr;
    const char* fns_hostname = "fns";

    if ((hptr = gethostbyname(fns_hostname)) == NULL) {
        printf("[TEESDK] gethostbyname error for host: %s\n", fns_hostname);
        return EXIT_FAILURE;
    }

    printf("[TEESDK] official hostname:%s\n", hptr->h_name);

    if (!hptr->h_addr_list[0]) {
        printf("[TEESDK] empty address\n");
        return EXIT_FAILURE;
    }

    tms_addr.sin_family = AF_INET;
    // tms_addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    tms_addr.sin_addr.s_addr = *(u_long*)hptr->h_addr_list[0];
    tms_addr.sin_port = htons(g_tms_port);

    printf("[+] This is a single-party task: %s\n", method);

    eigen_t* context = eigen_context_new(g_enclave_info, uid, token,
                                         (struct sockaddr*)&tms_addr);
    if (context == NULL) {
        return EXIT_FAILURE;
    }

    eigen_task_t* task = eigen_create_task(context, method);
    if (task == NULL) {
        return EXIT_FAILURE;
    }
    printf("args: %s, size=%lu\n", args, strlen(args));
    // BUG result truncating
    ret = eigen_task_invoke_with_payload(task, args, strlen(args), recvbuf,
                                         sizeof(recvbuf));
    if (ret <= 0) {
        return EXIT_FAILURE;
    }

    printf("Response: %s\n", recvbuf);
    *output_size = strlen(recvbuf);
    *output = (char*)malloc(strlen(recvbuf) + 1);
    memset(*output, 0, *output_size);
    memcpy(*output, recvbuf, *output_size);

    eigen_task_free(task);
    eigen_context_free(context);
    return 0;
}

int init(const char* pub, const char* pri, const char* conf, int32_t port1) {
    eigen_init();

    g_auditors = eigen_auditor_set_new();
    eigen_auditor_set_add_auditor(g_auditors, pub, pri);

    if (g_auditors == NULL) {
        return EXIT_FAILURE;
    }

    g_enclave_info = eigen_enclave_info_load(g_auditors, conf);

    if (g_enclave_info == NULL) {
        return EXIT_FAILURE;
    }
    g_tms_port = port1;

    return 0;
}

int release() {
    eigen_enclave_info_free(g_enclave_info);
    eigen_auditor_set_free(g_auditors);
    return 0;
}

static bool is_file_exist(const char* fileName) {
    std::ifstream infile(fileName);
    return infile.good();
}

// Many opcode implementations were inspired from the Apache 2.0 licensed EVM
// implementation https://github.com/ethereum/evmone

using namespace intx;

namespace {
template <typename T>
static T shrink(uint256_t i) {
    return static_cast<T>(i & std::numeric_limits<T>::max());
}
}  // namespace

namespace machineoperation {

uint256_t& assumeInt(value& val) {
    auto aNum = std::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

const uint256_t& assumeInt(const value& val) {
    auto aNum = std::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

const CodePointStub& assumeCodePoint(const value& val) {
    auto cp = std::get_if<CodePointStub>(&val);
    if (!cp) {
        throw bad_pop_type{};
    }
    return *cp;
}

uint64_t assumeInt64(uint256_t& val) {
    if (val > std::numeric_limits<uint64_t>::max()) {
        throw int_out_of_bounds{};
    }

    return static_cast<uint64_t>(val);
}

const Tuple& assumeTuple(const value& val) {
    auto tup = std::get_if<Tuple>(&val);
    if (!tup) {
        throw bad_pop_type{};
    }
    return *tup;
}

Tuple& assumeTuple(value& val) {
    auto tup = std::get_if<Tuple>(&val);
    if (!tup) {
        throw bad_pop_type{};
    }
    return *tup;
}

Buffer& assumeBuffer(value& val) {
    auto buf = std::get_if<Buffer>(&val);
    if (!buf) {
        throw bad_pop_type{};
    }
    return *buf;
}

void add(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum + bNum;
    m.stack.popClear();
    ++m.pc;
}

void mul(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum * bNum;
    m.stack.popClear();
    ++m.pc;
}

void sub(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum - bNum;
    m.stack.popClear();
    ++m.pc;
}

void div(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = aNum / bNum;
    }
    m.stack.popClear();
    ++m.pc;
}

void sdiv(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = sdivrem(aNum, bNum).quot;
    }
    m.stack.popClear();
    ++m.pc;
}

void mod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = aNum % bNum;
    }
    m.stack.popClear();
    ++m.pc;
}

void smod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = sdivrem(aNum, bNum).rem;
    }
    m.stack.popClear();
    ++m.pc;
}

void addmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[2] = addmod(aNum, bNum, cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void mulmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[2] = mulmod(aNum, bNum, cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void exp(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = exp(aNum, bNum);
    m.stack.popClear();
    ++m.pc;
}

void signExtend(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (aNum < 31) {
        auto sign_bit = 8 * narrow_cast<uint8_t>(aNum) + 7;
        auto sign_mask = uint256_t{1} << sign_bit;
        auto value_mask = sign_mask - 1;
        auto is_neg = (bNum & sign_mask) != 0;
        m.stack[1] = is_neg ? bNum | ~value_mask : bNum & value_mask;
    }
    m.stack.popClear();
    ++m.pc;
}

void lt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum < bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void gt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum > bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void slt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 0 : 1;
        } else {
            m.stack[1] = aNum < bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void sgt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 1 : 0;
        } else {
            m.stack[1] = aNum > bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void eq(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aVal = m.stack[0];
    auto& bVal = m.stack[1];
    m.stack[1] = aVal == bVal ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void iszero(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = aNum == 0 ? 1 : 0;
    ++m.pc;
}

void bitwiseAnd(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum & bNum;
    m.stack.popClear();
    ++m.pc;
}

void bitwiseOr(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum | bNum;
    m.stack.popClear();
    ++m.pc;
}

void bitwiseXor(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum ^ bNum;
    m.stack.popClear();
    ++m.pc;
}

void bitwiseNot(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = ~aNum;
    ++m.pc;
}

void byte(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (aNum >= 32) {
        m.stack[1] = 0;
    } else {
        const auto shift = 256 - 8 - 8 * shrink<uint8_t>(aNum);
        const auto mask = uint256_t(255) << shift;
        m.stack[1] = (bNum & mask) >> shift;
    }
    m.stack.popClear();
    ++m.pc;
}

void shl(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = bNum << aNum;
    m.stack.popClear();
    ++m.pc;
}

void shr(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = bNum >> aNum;
    m.stack.popClear();
    ++m.pc;
}

void sar(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if ((bNum & (uint256_t{1} << 255)) == 0) {
        shr(m);
        return;
    }

    if (aNum >= 256) {
        m.stack[1] = ~uint256_t{0};
    } else {
        m.stack[1] = (bNum >> aNum) | (~uint256_t{0} << (256 - aNum));
    }
    m.stack.popClear();
    ++m.pc;
}

void hashOp(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = ::hash_value(m.stack[0]);
    ++m.pc;
}

struct ValueTypeVisitor {
    ValueTypes operator()(const uint256_t&) const { return NUM; }
    ValueTypes operator()(const CodePointStub&) const { return CODEPT; }
    ValueTypes operator()(const Tuple&) const { return TUPLE; }
    ValueTypes operator()(const std::shared_ptr<HashPreImage>&) const {
        return TUPLE;
    }
    ValueTypes operator()(const Buffer&) const { return BUFFER; }
};

void typeOp(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = std::visit(ValueTypeVisitor{}, m.stack[0]);
    ++m.pc;
}

void ethhash2Op(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    std::array<unsigned char, 64> inData{};
    auto it = to_big_endian(aNum, inData.begin());
    to_big_endian(bNum, it);

    auto hash_val = ethash::keccak256(inData.data(), inData.size());
    m.stack[1] = be::load<uint256_t>(hash_val);

    m.stack.popClear();
    ++m.pc;
}

namespace internal {
void encodeKeccakState(const Tuple& tup, uint64_t* state) {
    if (tup.tuple_size() != 7) {
        throw bad_pop_type{};
    }

    for (uint64_t i = 0; i < 6; ++i) {
        intx::be::unsafe::store(reinterpret_cast<uint8_t*>(&state[i * 4]),
                                bswap(assumeInt(tup.get_element_unsafe(i))));
    }
    // Handle last val separately
    state[24] = static_cast<uint64_t>(assumeInt(tup.get_element_unsafe(6)));
}

Tuple decodeKeccakState(const uint64_t* state) {
    return Tuple(bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[0]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[4]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[8]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[12]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[16]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[20]))),
                 uint256_t{state[24]});
}
}  // namespace internal

void keccakF(MachineState& m) {
    m.stack.prepForMod(1);
    auto& tup = assumeTuple(m.stack[0]);
    uint64_t state[25];

    internal::encodeKeccakState(tup, state);

    ethash_keccakf1600(state);

    m.stack[0] = internal::decodeKeccakState(state);
    ++m.pc;
}

namespace internal {
uint256_t sha256_block(const uint256_t& digest_int,
                       std::array<uint8_t, 64>& input_data) {
    uint32_t digest_data[8];
    picosha2::word_t digest[8];
    intx::be::unsafe::store(reinterpret_cast<uint8_t*>(&digest_data),
                            bswap(digest_int));
    for (int i = 0; i < 8; ++i) {
        digest[7 - i] = digest_data[i];
    }

    picosha2::detail::hash256_block(digest, input_data.begin(),
                                    input_data.end());

    for (int i = 0; i < 8; ++i) {
        digest_data[7 - i] = static_cast<uint32_t>(digest[i]);
    }

    return bswap(intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const uint8_t*>(&digest_data)));
}
}  // namespace internal

void sha256F(MachineState& m) {
    m.stack.prepForMod(3);
    auto& digest_int = assumeInt(m.stack[0]);
    auto& input_first_int = assumeInt(m.stack[1]);
    auto& input_second_int = assumeInt(m.stack[2]);

    std::array<uint8_t, 64> input_data{};
    intx::be::unsafe::store(input_data.data(), input_first_int);
    intx::be::unsafe::store(input_data.data() + 32, input_second_int);

    input_second_int = internal::sha256_block(digest_int, input_data);

    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void pop(MachineState& m) {
    m.stack.popClear();
    ++m.pc;
}

void spush(MachineState& m) {
    value copiedStatic = m.static_val;
    m.stack.push(std::move(copiedStatic));
    ++m.pc;
}

void rpush(MachineState& m) {
    value copiedRegister = m.registerVal;
    m.stack.push(std::move(copiedRegister));
    ++m.pc;
}

void rset(MachineState& m) {
    m.stack.prepForMod(1);
    m.registerVal = m.stack[0];
    m.stack.popClear();
    ++m.pc;
}

void jump(MachineState& m) {
    m.stack.prepForMod(1);
    auto& target = assumeCodePoint(m.stack[0]);
    m.pc = target.pc;
    m.stack.popClear();
}

void cjump(MachineState& m) {
    m.stack.prepForMod(2);
    auto& target = assumeCodePoint(m.stack[0]);
    auto& cond = assumeInt(m.stack[1]);
    if (cond != 0) {
        m.pc = target.pc;
    } else {
        ++m.pc;
    }
    m.stack.popClear();
    m.stack.popClear();
}

void stackEmpty(MachineState& m) {
    if (m.stack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

void pcPush(MachineState& m) {
    m.stack.push(CodePointStub{m.pc, m.loadCurrentInstruction()});
    ++m.pc;
}

void auxPush(MachineState& m) {
    m.stack.prepForMod(1);
    m.auxstack.push(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

void auxPop(MachineState& m) {
    m.auxstack.prepForMod(1);
    m.stack.push(std::move(m.auxstack[0]));
    m.auxstack.popClear();
    ++m.pc;
}

void auxStackEmpty(MachineState& m) {
    if (m.auxstack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

void errPush(MachineState& m) {
    m.stack.push(m.errpc);
    ++m.pc;
}

void errSet(MachineState& m) {
    m.stack.prepForMod(1);
    auto codePointVal = std::get_if<CodePointStub>(&m.stack[0]);
    if (!codePointVal) {
        m.state = Status::Error;
    } else {
        m.errpc = *codePointVal;
    }
    m.stack.popClear();
    ++m.pc;
}

void dup0(MachineState& m) {
    value valACopy = m.stack[0];
    m.stack.push(std::move(valACopy));
    ++m.pc;
}

void dup1(MachineState& m) {
    value valBCopy = m.stack[1];
    m.stack.push(std::move(valBCopy));
    ++m.pc;
}

void dup2(MachineState& m) {
    value valCCopy = m.stack[2];
    m.stack.push(std::move(valCCopy));
    ++m.pc;
}

void swap1(MachineState& m) {
    m.stack.prepForMod(2);
    std::swap(m.stack[0], m.stack[1]);
    ++m.pc;
}

void swap2(MachineState& m) {
    m.stack.prepForMod(3);
    std::swap(m.stack[0], m.stack[2]);
    ++m.pc;
}

void tget(MachineState& m) {
    m.stack.prepForMod(2);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.stack[1]);
    m.stack[1] = tup.get_element(index);
    m.stack.popClear();
    ++m.pc;
}

void tset(MachineState& m) {
    m.stack.prepForMod(3);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.stack[1]);
    tup.set_element(index, std::move(m.stack[2]));
    m.stack[2] = std::move(tup);
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void xget(MachineState& m) {
    m.stack.prepForMod(1);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.auxstack[0]);
    m.stack[0] = tup.get_element(index);
    ++m.pc;
}

void xset(MachineState& m) {
    m.stack.prepForMod(2);
    m.auxstack.prepForMod(1);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.auxstack[0]);
    tup.set_element(index, std::move(m.stack[1]));
    m.auxstack[0] = std::move(tup);
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void tlen(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = assumeTuple(m.stack[0]).tuple_size();
    ++m.pc;
}

namespace {
uint256_t parseSignature(MachineState& m) {
    std::array<unsigned char, 64> sig_raw{};
    auto it = to_big_endian(assumeInt(m.stack[0]), sig_raw.begin());
    to_big_endian(assumeInt(m.stack[1]), it);
    auto recovery_int = assumeInt(m.stack[2]);
    auto message = be::store<ethash::hash256>(assumeInt(m.stack[3]));

    if (recovery_int != 0 && recovery_int != 1) {
        return 0;
    }

    static secp256k1_context* context = secp256k1_context_create(
        SECP256K1_CONTEXT_SIGN | SECP256K1_CONTEXT_VERIFY);

    secp256k1_ecdsa_recoverable_signature sig;
    int parsed_sig = secp256k1_ecdsa_recoverable_signature_parse_compact(
        context, &sig, sig_raw.data(), static_cast<int>(recovery_int));
    if (!parsed_sig) {
        return 0;
    }

    secp256k1_pubkey pubkey;
    if (!secp256k1_ecdsa_recover(context, &pubkey, &sig, message.bytes)) {
        return 0;
    }

    std::array<unsigned char, 65> pubkey_raw{};
    size_t output_length = pubkey_raw.size();
    int serialized_pubkey = secp256k1_ec_pubkey_serialize(
        context, pubkey_raw.data(), &output_length, &pubkey,
        SECP256K1_EC_UNCOMPRESSED);
    if (!serialized_pubkey) {
        return 0;
    }
    // Skip header byte
    auto hash_val = ethash::keccak256(pubkey_raw.data() + 1, 64);
    std::fill(&hash_val.bytes[0], &hash_val.bytes[12], 0);
    return be::load<uint256_t>(hash_val);
}
}  // namespace

void ec_recover(MachineState& m) {
    m.stack.prepForMod(4);

    m.stack[3] = parseSignature(m);
    m.stack.popClear();
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void ec_add(MachineState& m) {
    m.stack.prepForMod(4);
    auto& aVal = assumeInt(m.stack[0]);
    auto& bVal = assumeInt(m.stack[1]);
    auto& cVal = assumeInt(m.stack[2]);
    auto& dVal = assumeInt(m.stack[3]);

    auto ret = ecadd({aVal, bVal}, {cVal, dVal});

    if (std::holds_alternative<std::string>(ret)) {
        m.state = Status::Error;
        return;
    }

    G1Point ans = std::get<G1Point>(ret);
    cVal = ans.x;
    dVal = ans.y;
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void ec_mul(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aVal = assumeInt(m.stack[0]);
    auto& bVal = assumeInt(m.stack[1]);
    auto& cVal = assumeInt(m.stack[2]);

    auto ret = ecmul({aVal, bVal}, cVal);

    if (std::holds_alternative<std::string>(ret)) {
        m.state = Status::Error;
        return;
    }

    G1Point ans = std::get<G1Point>(ret);
    bVal = ans.x;
    cVal = ans.y;
    m.stack.popClear();
    ++m.pc;
}

void ec_pairing(MachineState& m) {
    m.stack.prepForMod(1);

    std::vector<std::pair<G1Point, G2Point>> points;

    const Tuple* val = &assumeTuple(m.stack[0]);
    for (int i = 0; i < max_ec_pairing_points; i++) {
        if (val->tuple_size() == 0) {
            break;
        }
        if (val->tuple_size() != 2) {
            throw bad_pop_type{};
        }
        auto& next = assumeTuple(val->get_element_unsafe(0));
        val = &assumeTuple(val->get_element_unsafe(1));

        if (next.tuple_size() != 6) {
            throw bad_pop_type{};
        }

        G1Point g1{assumeInt(next.get_element_unsafe(0)),
                   assumeInt(next.get_element_unsafe(1))};
        G2Point g2{assumeInt(next.get_element_unsafe(2)),
                   assumeInt(next.get_element_unsafe(3)),
                   assumeInt(next.get_element_unsafe(4)),
                   assumeInt(next.get_element_unsafe(5))};
        points.emplace_back(g1, g2);
    }
    if (val->tuple_size() != 0) {
        throw bad_pop_type{};
    }

    auto ret = ecpairing(points);
    if (std::holds_alternative<std::string>(ret)) {
        m.state = Status::Error;
        return;
    }

    m.stack[0] = std::get<bool>(ret) ? 1 : 0;
    ++m.pc;
}

uint64_t ec_pairing_variable_gas_cost(const MachineState& m) {
    // The fixed cost of the the pairing opcode is applied elsewhere
    uint64_t gas_cost = 0;
    if (m.stack.stacksize() == 0) {
        return gas_cost;
    }
    try {
        const value* val = &m.stack[0];
        for (int i = 0; i < max_ec_pairing_points; i++) {
            const Tuple* tup = &assumeTuple(*val);
            if (tup->tuple_size() == 0) {
                break;
            }
            if (tup->tuple_size() != 2) {
                throw bad_pop_type{};
            }
            val = &tup->get_element_unsafe(1);
            gas_cost += ec_pair_gas_cost;
        }
    } catch (const std::exception&) {
    }

    return gas_cost;
}

BlockReason breakpoint(MachineState& m) {
    ++m.pc;
    return BreakpointBlocked{};
}

void log(MachineState& m) {
    m.stack.prepForMod(1);
    m.addProcessedLog(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

static std::string get_string_from_tuple(const value& val) {
    const Tuple* tuple = &assumeTuple(val);
    value v = tuple->get_element(2);
    value vlen = tuple->get_element(0);
    value voffset = tuple->get_element(1);
    uint256_t len = assumeInt(vlen);
    uint256_t offset = assumeInt(voffset);
    Buffer buf = assumeBuffer(v);

    std::string data;

    for (uint64_t i = 0; i < len; i++) {
        data.push_back(buf.get(i + uint64_t(offset)));
    }

    return data;
}

static Buffer convert_string_to_buffer(const std::string& s) {
    std::vector<uint8_t> v(s.length());
    std::copy(s.begin(), s.end(), v.begin());
    return Buffer::fromData(v);
}

int get_arity(const std::string& op) {
    return int(op[op.size() - 1] - '0');
}

void ecall(MachineState& m) {
    m.stack.prepForMod(1);
    std::cerr << "[INFO] ecall called: " << m.stack[0] << std::endl;
    const Tuple* tuple = &assumeTuple(m.stack[0]);
    std::string operator_string = get_string_from_tuple(tuple->get_element(0));
    std::string arg1 = get_string_from_tuple(tuple->get_element(1));
    std::string arg2 = get_string_from_tuple(tuple->get_element(2));
    std::string arg3 = get_string_from_tuple(tuple->get_element(3));

    std::cerr << "ECALL operator = " << operator_string << ", arg1 = " << arg1
              << ", arg2 = " << arg2 << ", arg3 = " << arg3 << std::endl;
#if 0
    // FIXME: This code wiil be removed in the future, it is used for early debug
    const char *pub = "/app/release/services/auditors/godzilla/godzilla.public.der";
    const char *pri = "/app/release/services/auditors/godzilla/godzilla.sign.sha256";
    const char *conf = "/app/release/services/enclave_info.toml";
    const char *method = "echo";
    const char *args = "Hello Eigen";
    const char *uid = "uid";
    const char *token = "token";
#else
    const char* pub = getenv("TEESDK_PUB");
    const char* pri = getenv("TEESDK_PRI");
    const char* conf = getenv("TEESDK_CONF");
    const char* uid = getenv("TEESDK_UID");
    const char* token = getenv("TEESDK_TOKEN");
#endif
    // If env is set with invalid value, fast failue
    // Just check the file exists
    if (!is_file_exist(pub) || !is_file_exist(pri) || !is_file_exist(conf)) {
        std::cerr << "[TEESDK] ENV is not set with valid value, fail:"
                  << " TEESDK_PUB: " << pub << " TEESDK_PRI: " << pri
                  << " TEESDK_CONF: " << conf << std::endl;
        m.stack.popClear();
        // Returns a empty Buffer
        m.stack.push(Tuple(0, Buffer()));
        ++m.pc;
        return;
    }

    char* output = NULL;  // malloc from `submit_task`
    size_t outputsize = 0;
    int32_t port = 8082;
    int result = 0;
    static bool init_flag = false;

    if (!init_flag) {
        result = init(pub, pri, conf, port);
        init_flag = true;
    }

    if (result != 0) {
        std::cerr << "[TEESDK] init fail: " << result << std::endl;
    }

    std::string method_string = "operator";
    std::ostringstream oss;

    int arity = get_arity(operator_string);
    oss << operator_string << "," << arity << ",";

    switch (arity) {
        case 1:
            oss << arg1;
            break;
        case 2:
            oss << arg1 << "," << arg2;
            break;
        case 3:
            oss << arg1 << "," << arg2 << "," << arg3;
            break;
        default:
            std::cerr << "[TEESDK] arity of " << operator_string << " is "
                      << arity << ", it is not supported now" << std::endl;
            m.stack.popClear();
            // Returns a empty Buffer
            m.stack.push(Tuple(0, Buffer()));
            ++m.pc;
            return;
    }

    if (result == 0) {
        result = submit_task(method_string.c_str(), oss.str().c_str(), uid,
                             token, &output, &outputsize);

        if (result != 0) {
            std::cerr << "[TEESDK] submit_task fail: " << result << std::endl;
            m.stack.popClear();
            // Returns a empty Buffer
            m.stack.push(Tuple(0, Buffer()));
            ++m.pc;
            return;
        }
    }

    m.stack.popClear();
    std::string s(output, outputsize);
    Tuple res = Tuple(uint256_t(s.length()), convert_string_to_buffer(s));
    m.stack.push(res);
    ++m.pc;
}

void debug(MachineState& m) {
    m.stack.prepForMod(1);
    std::cerr << "[DEBUG] stack[0] = " << m.stack[0] << std::endl;
    m.context.debug_prints.push_back(m.stack.pop());
    ++m.pc;
}

void send(MachineState& m) {
    m.stack.prepForMod(2);

    auto msg_size = assumeInt64(assumeInt(m.stack[0]));
    Buffer& buf = assumeBuffer(m.stack[1]);

    // Note: the last msg_size == 0 check is implied by the buf.lastIndex()
    // check, but it's additionally specified for clarity and in case the
    // lastIndex method is refactored out.
    if (msg_size > send_size_limit || buf.lastIndex() >= msg_size ||
        msg_size == 0) {
        m.state = Status::Error;
        std::cerr << "Send failure: over size limit" << std::endl;
        return;
    }

    auto vec = std::vector<uint8_t>(msg_size);
    for (uint64_t i = 0; i < msg_size; i++) {
        vec[i] = buf.get(i);
    }
    m.addProcessedSend(std::move(vec));
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

BlockReason inboxOp(MachineState& m) {
    MachineMessage next_message;
    if (!m.context.inboxEmpty()) {
        next_message = m.context.popInbox();
    } else {
        return InboxBlocked();
    }

    m.addProcessedMessage(next_message);
    m.stack.push(next_message.message.toTuple());
    ++m.pc;
    return NotBlocked{};
}

void setgas(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.arb_gas_remaining = aNum;
    m.stack.popClear();
    ++m.pc;
}

void pushgas(MachineState& m) {
    auto& gas = m.arb_gas_remaining;
    m.stack.push(gas);
    ++m.pc;
}

void errcodept(MachineState& m) {
    m.stack.push(m.code->addSegment());
    ++m.pc;
}

void pushinsn(MachineState& m) {
    m.stack.prepForMod(2);
    auto target = std::get_if<CodePointStub>(&m.stack[1]);
    if (!target) {
        m.state = Status::Error;
        return;
    }
    auto& op_int = assumeInt(m.stack[0]);
    auto op = static_cast<uint8_t>(op_int);
    m.stack[1] = m.code->addOperation(target->pc, {static_cast<OpCode>(op)});
    m.stack.popClear();
    ++m.pc;
}

void pushinsnimm(MachineState& m) {
    m.stack.prepForMod(3);
    auto target = std::get_if<CodePointStub>(&m.stack[2]);
    if (!target) {
        m.state = Status::Error;
        return;
    }
    auto& op_int = assumeInt(m.stack[0]);
    auto op = static_cast<uint8_t>(op_int);
    m.stack[2] = m.code->addOperation(
        target->pc, {static_cast<OpCode>(op), std::move(m.stack[1])});
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

BlockReason sideload(MachineState& m) {
    m.stack.prepForMod(1);
    auto& block_num = assumeInt(m.stack[0]);
    m.output.last_sideload = block_num;
    if (!m.context.sideloads.empty()) {
        m.stack[0] = m.context.sideloads.back().toTuple();
        m.context.sideloads.pop_back();
    } else {
        if (m.context.stop_on_sideload && !m.context.first_instruction) {
            return SideloadBlocked{block_num};
        }
        m.stack[0] = Tuple();
    }
    ++m.pc;
    return NotBlocked{};
}

void newbuffer(MachineState& m) {
    m.stack.prepForMod(0);
    m.stack.push(Buffer{});
    ++m.pc;
}

void getbuffer8(MachineState& m) {
    m.stack.prepForMod(2);
    auto offset = assumeInt64(assumeInt(m.stack[0]));
    Buffer& md = assumeBuffer(m.stack[1]);
    auto res = uint256_t(md.get(offset));
    m.stack.popClear();
    m.stack.popClear();
    m.stack.push(res);
    ++m.pc;
}

void getbuffer64(MachineState& m) {
    m.stack.prepForMod(2);
    auto offset = assumeInt64(assumeInt(m.stack[0]));
    Buffer& md = assumeBuffer(m.stack[1]);
    if (offset + 7 < offset)
        throw int_out_of_bounds{};
    uint64_t res = 0;
    for (int i = 0; i < 8; i++) {
        res = res << 8U;
        res = res | md.get(offset + i);
    }
    m.stack.popClear();
    m.stack.popClear();
    m.stack.push(uint256_t(res));
    ++m.pc;
}

void getbuffer256(MachineState& m) {
    m.stack.prepForMod(2);
    auto offset = assumeInt64(assumeInt(m.stack[0]));
    Buffer& md = assumeBuffer(m.stack[1]);
    if (offset + 31 < offset)
        throw int_out_of_bounds{};
    uint256_t res = 0;
    std::vector<uint8_t> data(32);
    if ((offset + 31) % ALIGN < offset % ALIGN) {
        data = md.get_many(offset, ALIGN - (offset % ALIGN));
        auto data2 = md.get_many(offset + ALIGN - (offset % ALIGN),
                                 32 - (ALIGN - (offset % ALIGN)));
        data.insert(data.end(), data2.begin(), data2.end());
    } else {
        data = md.get_many(offset, 32);
    }
    for (int i = 0; i < 32; i++) {
        res = res << 8;
        res = res | data[i];
    }
    m.stack.popClear();
    m.stack.popClear();
    m.stack.push(res);
    ++m.pc;
}

void setbuffer8(MachineState& m) {
    m.stack.prepForMod(3);
    auto offset = assumeInt64(assumeInt(m.stack[0]));
    auto val_int = assumeInt(m.stack[1]);
    if (val_int > std::numeric_limits<uint8_t>::max()) {
        throw int_out_of_bounds{};
    }
    auto val = static_cast<uint8_t>(val_int);
    Buffer& md = assumeBuffer(m.stack[2]);
    auto res = md.set(offset, val);
    m.stack.popClear();
    m.stack.popClear();
    m.stack.popClear();
    m.stack.push(res);
    ++m.pc;
}

void setbuffer64(MachineState& m) {
    m.stack.prepForMod(3);
    auto offset = assumeInt64(assumeInt(m.stack[0]));
    auto val = assumeInt64(assumeInt(m.stack[1]));
    if (offset + 7 < offset)
        throw int_out_of_bounds{};
    // The initial value is copied here, there might be a way to optimize that
    // away
    Buffer res = assumeBuffer(m.stack[2]);
    m.stack.popClear();
    m.stack.popClear();
    m.stack.popClear();
    for (int i = 0; i < 8; i++) {
        res = res.set(offset + 7 - i, val & 0xffU);
        val = val >> 8U;
    }
    m.stack.push(res);
    ++m.pc;
}

void setbuffer256(MachineState& m) {
    m.stack.prepForMod(3);
    auto offset = assumeInt64(assumeInt(m.stack[0]));
    if (offset + 31 < offset)
        throw int_out_of_bounds{};
    auto val = assumeInt(m.stack[1]);
    // The initial value is copied here, there might be a way to optimize that
    // away
    Buffer res = assumeBuffer(m.stack[2]);
    m.stack.popClear();
    m.stack.popClear();
    m.stack.popClear();
    auto buf = std::vector<uint8_t>(32);
    for (int i = 0; i < 32; i++) {
        buf[31 - i] = static_cast<uint8_t>(val & 0xff);
        val = val >> 8;
    }

    if ((offset + 31) % ALIGN < offset % ALIGN) {
        auto data1 = std::vector<uint8_t>(
            buf.begin(), buf.begin() + (ALIGN - (offset % ALIGN)));
        auto data2 = std::vector<uint8_t>(
            buf.begin() + (ALIGN - (offset % ALIGN)), buf.end());
        res = res.set_many(offset, data1);
        res = res.set_many(offset + ALIGN - (offset % ALIGN), data2);
    } else {
        res = res.set_many(offset, buf);
    }
    m.stack.push(res);
    ++m.pc;
}

}  // namespace machineoperation
