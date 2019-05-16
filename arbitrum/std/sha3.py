from ..annotation import modifies_stack, modifies_stack_unchecked
from . import tup, byterange, bitwise
from .. import value

# Usage:
#    call ctx_new to get a new sha3Ctx object (usable for SHA3-256
#    or Keccak-256 hash)
#         arguments: none
#         returns: sha3ctx object
#    for each byte of the thing you want to hash, call ctx_pushbyte
#         arguments: sha3ctx byteToPush
#         returns: updatedSha3ctx
#    call ctx_finish to get the result of the SHA3-256 hash
#         argument: sha3ctx
#         returns: hash value as Integer
#    or call keccak_ctx_finish to get the result of the Keccak-256 hash
#         argument: sha3ctx
#         returns: hash value as Integer

# This code is based on the public-domain code at
# https://github.com/brainhub/SHA3IUF/blob/master/sha3.c

# sha3ctx = [sha3accumulator blockBuffer numBytesInBlockBuffer]

TT256M1 = 2 ** 256 - 1


@modifies_stack(0, 1)  # -> sha3ctx
def ctx_new(vm):
    vm.push(0)
    make_buf_1600(vm)
    make_buf_1600(vm)
    tup.make(3)(vm)


@modifies_stack([value.ValueType(), value.IntType()], 1)  # sha3ctx byte -> updatedsha3ctx
def ctx_pushbyte(vm):
    vm.cast(value.TupleType(3))
    vm.swap1()
    vm.dup1()
    # ctx byte ctx
    vm.dup0()
    vm.tgetn(2)
    vm.cast(value.IntType())
    # bytesInBuf ctx byte ctx
    vm.swap1()
    vm.tgetn(1)
    # blockBuffer bytesInBuf byte ctx
    set_byte_in_buf(vm)
    # blockBuffer ctx
    vm.swap1()
    vm.tsetn(1)
    # ctx
    vm.dup0()
    vm.tgetn(2)
    vm.cast(value.IntType())
    vm.push(1)
    vm.add()
    # updatedNumBytes ctx
    vm.dup0()
    vm.push(136)
    vm.eq()
    # updatedNumBytes==136 updatedNumBytes ctx
    vm.ifelse(lambda vm: [
        # 136 ctx
        vm.pop(),
        # ctx
        tup.tbreak(3)(vm),
        # shaAccum blockBuffer 136
        absorb_block(vm),
        # updatedShaAccum 136
        vm.swap1(),
        vm.pop(),
        # updatedShaAccum
        vm.push(0),
        vm.swap1(),
        make_buf_1600(vm),
        vm.swap1(),
        # updatedShaAccum emptyBlockBuf 0
        tup.make(3)(vm),
        # updatedCtx
    ], lambda vm: [
        # updNumByte ctx
        vm.swap1(),
        vm.tsetn(2),
        # updatedCtx
    ])


@modifies_stack(1, 1)  # sha3ctx -> hashValue
def ctx_finish(vm):
    vm.dup0()
    vm.tgetn(2)
    vm.push(135)
    vm.eq()
    # numBytesInBuf==135 ctx
    vm.ifelse(lambda vm: [
        # ctx
        vm.push(0x86),
        vm.swap1(),
        ctx_pushbyte(vm),
        # ctx
    ], lambda vm: [
        vm.push(6),
        vm.swap1(),
        ctx_pushbyte(vm),
        vm.while_loop(lambda vm: [
            # ctx
            vm.dup0(),
            vm.tgetn(2),
            vm.push(135),
            vm.eq(),
            vm.iszero(),
            # ctx.numBytes!=135 ctx
        ], lambda vm: [
            vm.push(0),
            vm.swap1(),
            ctx_pushbyte(vm),
        ]),
        vm.push(0x80),
        vm.swap1(),
        ctx_pushbyte(vm),
    ])
    vm.tgetn(0)
    vm.tgetn(0)
    bitwise.flip_endianness(vm)


@modifies_stack(1, [value.IntType()])  # sha3ctx -> hashValue
def keccak_ctx_finish(vm):
    vm.cast(value.TupleType(3))
    vm.dup0()
    vm.tgetn(2)
    vm.push(135)
    vm.eq()
    # numBytesInBuf==135 ctx
    vm.ifelse(lambda vm: [
        # ctx
        vm.push(0x81),
        vm.swap1(),
        ctx_pushbyte(vm),
        # ctx
    ], lambda vm: [
        vm.push(1),
        vm.swap1(),
        ctx_pushbyte(vm),
        vm.while_loop(lambda vm: [
            # ctx
            vm.dup0(),
            vm.cast(value.TupleType(3)),
            vm.tgetn(2),
            vm.push(135),
            vm.eq(),
            vm.iszero(),
            # ctx.numBytes!=135 ctx
        ], lambda vm: [
            vm.push(0),
            vm.swap1(),
            ctx_pushbyte(vm),
        ]),
        vm.push(0x80),
        vm.swap1(),
        ctx_pushbyte(vm),
    ])
    vm.cast(value.TupleType(3))
    vm.tgetn(0)
    vm.cast(value.TupleType(3))
    vm.tgetn(0)
    vm.cast(value.IntType())
    bitwise.flip_endianness(vm)


_ROUND_CONST = [
    0x0000000000000001, 0x0000000000008082,
    0x800000000000808a, 0x8000000080008000,
    0x000000000000808b, 0x0000000080000001,
    0x8000000080008081, 0x8000000000008009,
    0x000000000000008a, 0x0000000000000088,
    0x0000000080008009, 0x000000008000000a,
    0x000000008000808b, 0x800000000000008b,
    0x8000000000008089, 0x8000000000008003,
    0x8000000000008002, 0x8000000000000080,
    0x000000000000800a, 0x800000008000000a,
    0x8000000080008081, 0x8000000000008080,
    0x0000000080000001, 0x8000000080008008
]

_PILN = [
    10, 7, 11, 17, 18, 3, 5, 16, 8, 21, 24, 4,
    15, 23, 19, 13, 12, 2, 20, 14, 22, 9, 6, 1
]

_ROTC = [
    1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 2,
    14, 27, 41, 56, 8, 25, 43, 62, 18, 39, 61, 20, 44
]

_BLOCK_FOR_EMPTY_MESSAGE = [6, 0, 0, 0, 1 << 63, 0, 0]


# blockBuffer byteNum byte -> updatedBlockBuffer
# assume the slot is zero-filled
@modifies_stack([value.ValueType(), value.IntType(), value.IntType()], 1)
def set_byte_in_buf(vm):
    vm.cast(value.TupleType(7))
    vm.dup1()
    vm.push(32)
    vm.swap1()
    vm.div()
    # slotNum blockBuf byteNum byte
    vm.dup1()
    vm.swap1()
    vm.tget()
    # slotVal blockBuf byteNum byte
    vm.dup2()
    vm.push(32)
    vm.swap1()
    vm.mod()
    # subSlotNum slotVal blockBuf byteNum byte
    tup.make(5)(vm)
    # [subSlotNum slotVal blockBuf byteNum byte]
    vm.dup0()
    vm.tgetn(4)
    # byte [subSlotNum slotVal blockBuf byteNum byte]
    vm.dup1()
    vm.tgetn(0)
    # subSlotNum byte [subSlotNum slotVal blockBuf byteNum byte]
    vm.dup2()
    vm.tgetn(1)
    # slotVal subSlotNum byte [subSlotNum slotVal blockBuf byteNum byte]
    vm.cast(value.IntType())
    set_byte_in_word(vm)
    # updatedSlotVal [subSlotNum slotVal blockBuf byteNum byte]
    vm.swap1()
    vm.dup0()
    vm.tgetn(2)
    vm.swap1()
    vm.tgetn(3)
    # byteNum blockBuf updatedSlotVal
    vm.push(32)
    vm.swap1()
    vm.div()
    # slotNum blockBuf updatedSlotVal
    vm.tset()
    # updatedBlockBuf


# word slotNum byte -> updatedWord
@modifies_stack([value.IntType(), value.IntType(), value.IntType()], 1)
def set_byte_in_word(vm):  # assume the slot is zero-filled
    vm.swap2()
    # byte slotNum word
    vm.while_loop(lambda vm: [
        vm.dup1(),
        vm.iszero(),
        vm.iszero(),
        # slotNum!=0 vyte slotNum word
    ], lambda vm: [
        # byte slotNum word
        vm.push(256),
        vm.mul(),
        # updByte slotNum word
        vm.swap1(),
        vm.push(-1 & TT256M1),
        vm.add(),
        vm.swap1(),
    ])
    # shiftedByte 0 word
    vm.swap1()
    vm.pop()
    # shiftedByte word
    vm.bitwise_or()
    # updatedWord


@modifies_stack(2, 1)   # state newblock -> state'
def absorb_block(vm):
    vm.cast(value.TupleType(7))
    vm.swap1()
    vm.cast(value.TupleType(7))
    vm.swap1()
    for i in range(7):
        vm.dup1()
        vm.dup1()
        # state newblock state newblock
        vm.tgetn(i)
        vm.cast(value.IntType())
        vm.swap1()
        vm.tgetn(i)
        vm.cast(value.IntType())
        vm.bitwise_xor()
        vm.swap1()
        vm.tsetn(i)
    # state newblock
    vm.swap1()
    vm.pop()
    permutation(vm)
    # state'


# buf1600 -> buf1600
@modifies_stack(1, 1)
def permutation(vm):
    for round_num in range(24):
        vm.push(_ROUND_CONST[round_num])
        vm.swap1()
        _round(vm)


# buf1600 roundConst -> buf1600
@modifies_stack(2, 1)
def _round(vm):
    theta(vm)
    rhopi(vm)
    chi(vm)
    iota(vm)


@modifies_stack_unchecked(1, 1)
def theta(vm):
    vm.cast(value.TupleType(7))
    make_buf_320(vm)
    # bc s
    vm.swap1()
    # s bc
    for i in range(5):
        vm.push(0)
        for j in range(0, 25, 5):
            # acc s bc
            vm.cast(value.IntType())
            vm.dup1()
            getword(vm, i+j)
            # s[i+j] acc s bc
            vm.bitwise_xor()
            # acc s bc
        # acc s bc
        vm.swap1()
        vm.swap2()
        # bc acc s
        setword(vm, i)
        # bc s
        vm.swap1()
    vm.swap1()
    # bc s
    for i in range(5):
        vm.dup0()
        getword(vm, (i+4) % 5)
        # bc[.] bc s
        vm.cast(value.IntType())
        vm.dup1()
        getword(vm, (i+1) % 5)
        vm.cast(value.IntType())
        rotl64(vm, 1)
        vm.bitwise_xor()
        # t bc s
        vm.swap1()
        vm.swap2()
        # s t bc
        for j in range(0, 25, 5):
            jpi = j+i
            vm.dup1()
            vm.cast(value.IntType())
            vm.dup1()
            # s t s t bc
            getword(vm, jpi)
            # s[jpi] t s t bc
            vm.cast(value.IntType())
            vm.swap1()
            vm.cast(value.IntType())
            vm.swap1()
            vm.bitwise_xor()
            # s[jpi]^t s t bc
            vm.swap1()
            setword(vm, jpi)
            # s t bc
        vm.swap1()
        vm.pop()
        vm.swap1()
        # bc s
    vm.pop()
    # s


@modifies_stack_unchecked(1, 1)
def rhopi(vm):
    # s
    vm.dup0()
    getword(vm, 1)
    # t s
    for i in range(24):
        j = _PILN[i]
        # t s
        vm.dup1()
        getword(vm, j)
        # bc0 t s
        vm.swap1()
        # t bc0 s
        rotl64(vm, _ROTC[i])
        # trot bc0 s
        vm.swap1()
        vm.swap2()
        # s trot bc0
        setword(vm, j)
        # s bc0    { bc0 becomes t here}
        vm.swap1()
        # t s
    # t s
    vm.pop()
    # s


@modifies_stack_unchecked(1, 1)
def chi(vm):
    # s
    make_buf_320(vm)
    # bc s
    for j in range(0, 25, 5):
        for i in range(5):
            # bc s
            vm.dup1()
            getword(vm, j+i)
            # s[j+i] bc s
            vm.swap1()
            setword(vm, i)
        # bc s
        for i in range(5):
            vm.dup0()
            getword(vm, (i+1) % 5)
            vm.push(_MASK_64)
            vm.bitwise_xor()
            # ~bc[(i+1)%5] bc s
            vm.dup1()
            getword(vm, (i+2) % 5)
            vm.bitwise_and()
            # _ bc s
            vm.swap1()
            vm.swap2()
            # s _ bc
            vm.swap1()
            vm.dup1()
            # s _ s bc
            getword(vm, j+i)
            vm.bitwise_xor()
            # xorResult s bc
            vm.swap1()
            setword(vm, j+i)
            # s bc
            vm.swap1()
        # bc s
    # bc s
    vm.pop()
    # s


@modifies_stack_unchecked(2, 1)
def iota(vm):
    # s roundConst
    vm.swap1()
    vm.dup1()
    vm.tgetn(0)
    # s[0] roundConst s
    vm.bitwise_xor()
    vm.swap1()
    vm.tsetn(0)
    # s


@modifies_stack(0, 1)
def make_buf_320(vm):
    vm.push(value.Tuple([0, 0]))


@modifies_stack(0, 1)
def make_buf_1600(vm):
    vm.push(value.Tuple([0, 0, 0, 0, 0, 0, 0]))


_TWO_TO_64 = 65536*65536*65536*65536
_TWO_TO_128 = _TWO_TO_64*_TWO_TO_64
_TWO_TO_192 = _TWO_TO_128*_TWO_TO_64
_MASK_64 = _TWO_TO_64-1
_MASK_128 = _MASK_64*_TWO_TO_64
_MASK_192 = _MASK_64*_TWO_TO_128
_MASK_256 = _MASK_64*_TWO_TO_192


def rotl64(vm, amount):
    # x
    vm.push(_MASK_64)
    vm.bitwise_and()
    vm.dup0()
    # x x
    vm.push(1 << amount)
    vm.mul()
    vm.push(_MASK_64)
    vm.bitwise_and()
    # r1 x
    vm.swap1()
    vm.push(1 << (64 - amount))
    vm.swap1()
    vm.cast(value.IntType())
    vm.div()
    # r2 r1
    vm.bitwise_or()
    # r


def getword(vm, word_num):
    # buf1600 -> word
    assert word_num >= 0
    assert word_num < 25
    vm.cast(value.TupleType(7))
    vm.tgetn(word_num//4)
    section = word_num % 4
    if section == 1:
        vm.push(_TWO_TO_64)
        vm.swap1()
        vm.cast(value.IntType())
        vm.div()
    elif section == 2:
        vm.push(_TWO_TO_128)
        vm.swap1()
        vm.cast(value.IntType())
        vm.div()
    elif section == 3:
        vm.push(_TWO_TO_192)
        vm.swap1()
        vm.cast(value.IntType())
        vm.div()
    vm.cast(value.IntType())
    vm.push(_MASK_64)
    vm.bitwise_and()


def setword(vm, word_num):
    # buf1600 word -> buf1600
    assert word_num >= 0
    assert word_num < 25

    vm.swap1()
    vm.cast(value.TupleType(7))
    vm.dup1()
    # buf1600 word buf1600
    getword(vm, word_num)
    # oldWord word buf1600
    vm.cast(value.IntType())
    vm.swap1()
    vm.cast(value.IntType())
    vm.swap1()
    vm.bitwise_xor()
    # oldWord^word buf1600
    section = word_num % 4
    if section == 1:
        vm.push(_TWO_TO_64)
        vm.mul()
    elif section == 2:
        vm.push(_TWO_TO_128)
        vm.mul()
    elif section == 3:
        vm.push(_TWO_TO_192)
        vm.mul()
    # xorWord buf1600
    vm.dup1()
    vm.tgetn(word_num // 4)
    # fullword xorWord buf1600
    vm.bitwise_xor()
    vm.swap1()
    vm.tsetn(word_num // 4)

@modifies_stack([byterange.typ, value.IntType()], [value.IntType()])
def hash_byterange(vm):
    # bytearray length
    vm.push(0)
    ctx_new(vm)
    tup.make(4)(vm)
    # [ctx, i, bytearray, length]
    vm.while_loop(lambda vm: [
        vm.dup0(),
        vm.tgetn(3),
        vm.dup1(),
        vm.tgetn(1),
        vm.lt()
    ], lambda vm: [
        vm.dup0(),
        vm.tgetn(1),
        vm.dup1(),
        vm.tgetn(2),
        byterange.get8(vm),
        # val [ctx, i, bytearray, length]
        vm.dup1(),
        vm.tgetn(0),
        ctx_pushbyte(vm),
        vm.swap1(),
        vm.tsetn(0),

        vm.dup0(),
        vm.tgetn(1),
        vm.push(1),
        vm.add(),
        vm.swap1(),
        vm.tsetn(1)
    ])
    vm.tgetn(0)
    keccak_ctx_finish(vm)

def print_nist_style(vm):
    # buf
    result = ""
    for i in range(25):
        vm.dup0()
        getword(vm, i)
        # word buf
        word = vm.stack[0]
        result += _print64(word)
        if i % 2 == 1:
            result += "\n"
        vm.pop()
    # buf
    print(result)


def _print64(word):
    ret = ""
    for i in range(0, 64, 8):
        val = (word >> i) & 0xff
        ret = ret + format(val, '02x') + " "
    return ret
