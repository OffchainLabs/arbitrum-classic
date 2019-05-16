# API for bignum library:
#
# zero(vm)                   -> bignum(0)
# fromint(vm)                int -> bignum(int)
#       assumes -(2^126) < int < 2^126
# fromPythonInt(vm, pint)    -> bignum(pint)
# toPythonInt(vm, pint)      bignum -> pint
#       works only on the Python emulator; is a no-op if used with the compiler
# shiftleft(vm)              bignum bitsToShift -> resultBignum
# shiftright(vm)             bignum bitsToShift -> resultBignum
# negate(vm)                 bignum -> -bignum
# add(vm)                    bn1 bn2 -> bn1+bn2
# subtract(vm)               bn1 bn2 -> bn1-bn2
# multiply(vm)               bn1 bn2 -> bn1*bn2
# divmod(vm)                 bn1 bn2 -> bn1//bn2 bn1%bn2
#       assumes bn1>=0, bn2>0
# mod(vm)                           bn1 bn2 -> bn1%bn2
#       assumes bn2>0
# intmultiply(vm)                   int bn  -> int*bn
#       assumes -(2^126) < int < (2^126)
# modadd(vm)                        bn1 bn2 bn3 -> (bn1+bn2)%bn3
#       assumes bn3>0
# modmul(vm)                        bn1 bn2 bn3 -> (bn1*bn2)%bn3
#       assumes bn3>0
# modpow(vm)                        bn1 bn2 bn3 -> (bn1^bn2)%bn3
#       assumes bn1>=0, bn2>=0, bn3>0
# modinv(vm)                        bn1 bn2 -> bn3 such that (bn1*bn3)%bn2 = bignum(1)
#       assumes bn2 > 0
# eq, lt, gt, geq, leq              bn1 bn2 -> int(0 or 1)


from .import bigtuple
from .import random
from .import tup
from .locals import Locals
from .struct import Struct
from ..annotation import modifies_stack
from ..vm import VM
from .. import value

bignum = Struct("bignum", ['val', 'size', 'ispositive'])
#    invariant: for all i>size, chunk[i] returns 0
#    slots in arry hold values mod 2^126

_CHUNK_BITS = 126
_CHUNK_MOD = 2**_CHUNK_BITS


def make_zero():
    return value.Tuple([value.Tuple([]), 0, 1])


def make_from_int(pint):
    if pint < 0:
        return negate_static(make_from_int(-pint))

    val = make_zero()
    i = 0
    while pint > 0:
        val = setchunk_static(val, i, pint % _CHUNK_MOD)
        i += 1
        pint //= _CHUNK_MOD
    return val


def negate_static(val):
    return val.set_tup_val(2, int(not val[2]))


def setchunk_static(bignum_val, chunk_num, val):
    if chunk_num + 1 > bignum_val[1]:
        bignum_val = bignum_val.set_tup_val(1, chunk_num + 1)

    return bignum_val.set_tup_val(
        0,
        bigtuple.set_static(bignum_val[0], chunk_num, val)
    )


def to_python_int(big):
    acc = 0
    val = big[0]
    size = big[1]
    ispositive = big[2]
    for i in range(size):
        sub_val = bigtuple.get_static(val, i)
        if isinstance(sub_val, value.Tuple):
            sub_val = 0
        acc += sub_val*(_CHUNK_MOD**i)
    if not ispositive:
        acc *= -1
    return acc


@modifies_stack(0, 1)
def zero(vm):
    vm.push(make_zero())


@modifies_stack(1, 1)
def fromint(vm):
    vm.push(1)
    vm.swap1()
    vm.push(1)
    vm.swap1()
    # val 1 1
    vm.push(0)
    bigtuple.new(vm)
    bigtuple.set_val(vm)
    # bigtuple 1 1
    tup.make(3)(vm)


def fromPythonInt(vm, pint):
    if pint < 0:
        fromPythonInt(vm, -pint)
        negate(vm)
    else:
        zero(vm)
        i = 0
        while pint > 0:
            # bignum
            vm.push(pint % _CHUNK_MOD)
            vm.swap1()
            vm.push(i)
            vm.swap1()
            setchunk(vm)
            i = i+1
            pint = pint // _CHUNK_MOD


def toPythonInt(vm):
    if isinstance(vm, VM):  # this requires that we're in the python emulator
        acc = 0
        size = vm.stack[0][1]
        ispositive = vm.stack[0][2]
        vm.tgetn(0)
        # bigtuple
        for i in range(size):
            vm.push(i)
            vm.dup1()
            bigtuple.get(vm)
            if isinstance(vm.stack[0], value.Tuple):
                vm.pop()
                vm.push(0)
            acc = acc + vm.stack[0]*(_CHUNK_MOD**i)
            vm.pop()
        vm.pop()
        if ispositive == 1:
            return acc
        else:
            return -acc
    else:
        vm.pop()
        return 0


@modifies_stack(2, 1)
def getchunk(vm):
    # bignum index
    vm.swap1()
    vm.dup1()
    # bignum index bignum
    bignum.get('size')(vm)
    # size index bignum
    vm.dup1()
    # index size index bignum
    vm.lt()
    # indexValid index bignum
    vm.ifelse(lambda vm: [
        # index bignum
        vm.swap1(),
        bignum.get('val')(vm),
        bigtuple.get(vm),
        vm.dup0(),
        vm.tnewn(0),
        vm.eq(),
        vm.ifelse(lambda vm: [
            vm.pop(),
            vm.push(0),
        ]),
    ], lambda vm: [
        # index bignum
        vm.pop(),
        vm.pop(),
        vm.push(0),
    ])
    # result
    vm.dup0()
    tup.make(0)(vm)
    vm.eq()
    # result==[] result
    vm.ifelse(lambda vm: [
        vm.pop(),
        vm.push(0),
    ])


@modifies_stack(3, 1)   # bn chunkNum val -> updatedBn
def setchunk(vm):
    # update size if needed
    vm.dup0()
    bignum.get('size')(vm)
    # bnsize bn chunkNum val
    vm.dup2()
    vm.push(1)
    vm.add()
    # chunkNum+1 bnsize bn chunkNum val
    vm.gt()
    vm.ifelse(lambda vm: [
        # bn chunkNum val
        vm.dup1(),
        vm.push(1),
        vm.add(),
        vm.swap1(),
        bignum.set_val('size')(vm),
    ])

    # update the chunk
    # bn chunkNum val
    vm.swap2()
    vm.dup2()
    # bn val chunkNum bn
    vm.swap2()
    vm.swap1()
    vm.swap2()
    # bn chunkNum val bn
    bignum.get('val')(vm)
    bigtuple.set_val(vm)
    # updatedVal bn
    vm.swap1()
    bignum.set_val('val')(vm)


@modifies_stack(1, 1)
def trim(vm):
    local_vars = Locals(vm, ['i', 'bn'])
    # bn
    vm.push(1)
    vm.dup1()
    bignum.get('size')(vm)
    vm.sub()
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get(['bn', 'i', 'i']),
        getchunk(vm),
        vm.iszero(),
        # chunk==0 i
        vm.swap1(),
        vm.push(-1),
        vm.slt(),
        vm.bitwise_and(),
    ], lambda vm: [
        local_vars.get('i'),
        vm.push(-1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    vm.auxpop()
    tup.tbreak(2)(vm)
    # i bn
    vm.push(1)
    vm.add()
    vm.swap1()
    bignum.set_val('size')(vm)


@modifies_stack(1, 1)   # bignum -> lengthinbits
def bitlength(vm):
    trim(vm)
    vm.dup0()
    # bignum bignum
    bignum.get('size')(vm)
    # bnsizechunks bignum
    vm.dup0()
    vm.ifelse(lambda vm: [
        # bnsizechunks bignum
        vm.push(-1),
        vm.add(),
        vm.dup0(),
        vm.push(_CHUNK_BITS),
        vm.mul(),
        # bnFullChunkBits bnchunks-1 bignum
        vm.swap2(),
        # bignum bnchunks-1 bnFullChunkBits
        getchunk(vm),
        # chunk[0] bnFullChunkBits
        bitlength_chunk(vm),
        vm.add(),
    ], lambda vm: [
        # bnsizechunks(=0) bignum
        vm.pop(),
        vm.pop(),
        vm.push(0),
    ])


@modifies_stack(1, 1)   # chunk -> lengthinbits
def bitlength_chunk(vm):
    vm.push(0)
    vm.swap1()
    bitlength_chunk2(vm, 128)


def bitlength_chunk2(vm, size):  # size must be power of 2
    # chunk soFar
    if size == 1:
        vm.ifelse(lambda vm: [
            vm.push(1),
            vm.add(),
        ])
    else:
        vm.dup0()
        vm.push((1 << (size//2))-1)
        vm.lt()
        # ((1<<(size//2))-1)<chunk chunk soFar
        vm.ifelse(lambda vm: [
            vm.swap1(),
            vm.push(size//2),
            vm.add(),
            vm.swap1(),
            # chunk soFar'
            vm.push(1 << (size//2)),
            vm.swap1(),
            vm.div(),
        ])
        # chunk soFar
        bitlength_chunk2(vm, size//2)


@modifies_stack(2, 1)   # val numChunks -> val[0..(numChunks-1)]
def loworderwords(vm):
    local_vars = Locals(vm, ['result', 'i', 'num', 'limit'])
    vm.push(0)
    zero(vm)
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get(['i', 'limit']),
        vm.lt(),
    ], lambda vm: [
        local_vars.get(['num', 'i', 'i', 'result', 'i']),
        getchunk(vm),
        vm.swap2(),
        setchunk(vm),
        local_vars.set_val('result'),

        vm.push(1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    local_vars.discard('result')


@modifies_stack(2, 1)   # bignum shiftBits
def shiftleft(vm):
    local_vars = Locals(vm, ['res', 'i', 'bnsize', 'bn', 'blockCount'])

    vm.swap1()
    vm.dup0()
    vm.push(_CHUNK_BITS)
    vm.swap1()
    vm.mod()
    # bitsCount shiftBits bignum
    vm.swap1()
    vm.push(_CHUNK_BITS)
    vm.swap1()
    vm.div()
    # chunksCount bitsCount bignum
    vm.swap2()
    vm.swap1()
    # bitscount bignum blocksCount
    vm.push(2)
    vm.exp()
    intmultiply(vm)
    # bignum' blocksCount
    vm.dup1()
    vm.ifelse(lambda vm: [
        vm.dup0(),
        bignum.get('size')(vm),
        vm.push(0),
        zero(vm),
        local_vars.make(),
        vm.while_loop(lambda vm: [
            local_vars.get(['i', 'bnsize']),
            vm.lt(),
        ], lambda vm: [
            local_vars.get(['bn', 'i', 'i', 'blockCount', 'res', 'i']),
            getchunk(vm),
            # bn[i] i blockCount res i
            vm.swap2(),
            vm.add(),
            # i+blockCount bn[i] res i
            vm.swap1(),
            vm.swap2(),
            # res i+blockCount bn[i] i
            setchunk(vm),
            # res i
            local_vars.set_val('res'),

            # i
            vm.push(1),
            vm.add(),
            local_vars.set_val('i'),
        ]),
        local_vars.discard('res')
    ], lambda vm: [
        # bignum' blockcount
        vm.swap1(),
        vm.pop(),
    ])
    trim(vm)


@modifies_stack(2, 1)  # bignum shiftBits -> shiftedBignum
def shiftright(vm):
    vm.swap1()
    vm.dup0()
    # shiftBits shiftBits bignum
    vm.push(_CHUNK_BITS)
    vm.swap1()
    vm.mod()
    # sb%chunkBits shiftBits bignum
    vm.dup0()
    vm.ifelse(lambda vm: [
        # shiftbits%chunkbits shiftbits bignum
        vm.push(_CHUNK_BITS),
        vm.sub(),
        # reverseshiftbits shiftbits bignum
        vm.dup0(),
        vm.swap2(),
        # shiftbits reverseshiftbits rsb bignum
        vm.add(),
        vm.swap2(),
        # bignum reverseshiftbits modshiftbits
        vm.swap1(),
        vm.push(2),
        vm.exp(),
        intmultiply(vm),
        vm.swap1(),
    ], lambda vm: [
        vm.pop(),
        # shiftbits bignum
    ])

    vm.push(_CHUNK_BITS)
    vm.swap1()
    vm.div()
    # shiftchunks bignum
    vm.swap1()
    vm.dup0()
    bignum.get('size')(vm)
    # bnsize bignum shiftchunks
    vm.dup2()
    vm.swap1()
    vm.sub()
    # limit bignum shiftchunks
    zero(vm)
    vm.push(0)
    local_vars = Locals(vm, ['i', 'result', 'limit', 'bn', 'shiftchunks'])
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get(['i', 'limit']),
        vm.slt(),
    ], lambda vm: [
        local_vars.get(['i', 'shiftchunks', 'bn']),
        vm.add(),
        vm.swap1(),
        getchunk(vm),
        # chunk
        local_vars.get(['result', 'i']),
        # result i chunk
        setchunk(vm),
        local_vars.set_val('result'),

        local_vars.get('i'),
        vm.push(1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    local_vars.discard('result')


@modifies_stack(2, 1)
def sizeoflarger(vm):
    # bn1 bn2
    bignum.get('size')(vm)
    vm.swap1()
    bignum.get('size')(vm)
    _max2(vm)


@modifies_stack(2, 1)
def _max2(vm):
    # v1 v2
    vm.dup1()
    vm.dup1()
    vm.lt()
    # v1<v2 v1 v2
    vm.ifelse(lambda vm: [
        # v1 v2
        vm.pop(),
    ], lambda vm: [
        vm.swap1(),
        vm.pop(),
    ])


@modifies_stack(1, 1)  # bignum -bignum
def negate(vm):
    vm.dup0()
    bignum.get('ispositive')(vm)
    # ispositive bignum
    vm.iszero()
    vm.swap1()
    bignum.set_val('ispositive')(vm)


@modifies_stack(2, 1)  # bn1 bn2 -> bn1+bn2   (assume both bn1, bn2 >= 0)
def add_bothpositive(vm):
    local_vars = Locals(vm, ['i', 'result', 'carry', 'limit', 'bn1', 'bn2'])
    # bn1 bn2
    vm.dup1()
    vm.dup1()
    sizeoflarger(vm)
    # size bn1 bn2
    vm.push(0)
    zero(vm)
    # result 0 size bn1 bn2
    vm.push(0)
    # 0 result 0 size bn1 bn2
    local_vars.make()
    vm.while_loop(lambda vm: [
        local_vars.get(['i', 'limit']),
        vm.lt(),
    ], lambda vm: [
        local_vars.get(['bn1', 'i', 'i', 'bn2', 'carry']),
        getchunk(vm),
        # val1 i bn2 carry
        vm.swap2(),
        getchunk(vm),
        # val2 val1 carry
        vm.add(),
        vm.add(),
        # newval
        vm.dup0(),
        vm.push(_CHUNK_MOD),
        vm.swap1(),
        vm.div(),
        # newval//chunkMod newval
        local_vars.set_val('carry'),
        # newval
        vm.push(_CHUNK_MOD),
        vm.swap1(),
        vm.mod(),
        # truncatedNewval
        local_vars.get(['result', 'i']),
        setchunk(vm),
        local_vars.set_val('result'),

        local_vars.get('i'),
        vm.push(1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    local_vars.get('carry')
    vm.push(0)
    vm.dup1()
    # carry 0 carry
    vm.gt()
    vm.ifelse(lambda vm: [
        local_vars.get(['result', 'i']),
        setchunk(vm),
        # result
    ], lambda vm: [
        vm.pop(),
        local_vars.get('result'),
    ])
    local_vars.discard()


@modifies_stack(2, 1)   # bn1 bn2 -> difference
def subtract_bothpositive(vm):
    vm.dup1()
    vm.dup1()
    ltbothpositive(vm)
    # bn1<bn2 bn1 bn2
    vm.ifelse(lambda vm: [
        vm.swap1(),
        subtract_allpositive(vm),
        negate(vm),
    ], lambda vm: [
        subtract_allpositive(vm),
    ])


@modifies_stack(2, 1)   # bn1 bn2 -> difference
def subtract_allpositive(vm):  # bn1 >= bn2 >= 0
    # set up local_vars
    local_vars = Locals(vm, ['i', 'borrow', 'limit', 'bn1', 'bn2'])
    vm.dup0()
    bignum.get('size')(vm)
    vm.push(0)
    vm.push(0)
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get(['i', 'limit']),
        vm.lt(),
    ], lambda vm: [
        local_vars.get(['bn2', 'i', 'borrow']),
        getchunk(vm),
        # bn2[i] borrow
        local_vars.get(['bn1', 'i']),
        getchunk(vm),
        # bn1[i] bn2[i] borrow
        vm.sub(),
        vm.sub(),
        # diff
        vm.dup0(),
        vm.push(0),
        vm.sgt(),
        # 0>diff diff
        vm.ifelse(lambda vm: [
            vm.push(1),
            local_vars.set_val('borrow'),
            vm.push(_CHUNK_MOD),
            vm.add(),
        ], lambda vm: [
            vm.push(0),
            local_vars.set_val('borrow'),
        ]),
        # diff
        local_vars.get(['bn1', 'i']),
        # bn1 i diff
        setchunk(vm),
        local_vars.set_val('bn1'),

        local_vars.get('i'),
        vm.push(1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    local_vars.discard('bn1')


# bn1 bn2 modulus -> (bn1+bn2)%modulus (assume bn1,bn2>=0; modulus>0)
@modifies_stack(3, 1)
def modadd(vm):
    add(vm)
    modallpositive(vm)


@modifies_stack(2, 1)   # bn1 bn2 -> sum
def add(vm):
    vm.dup1()
    bignum.get('ispositive')(vm)
    vm.dup1()
    bignum.get('ispositive')(vm)
    # ispositive1 ispositive2 bn1 bn2
    # ispositive1 ispositive2 bn1 bn2
    vm.ifelse(lambda vm: [
        # ispositive2 bn1 bn2
        vm.ifelse(lambda vm: [
            add_bothpositive(vm),
            vm.push(1),
            vm.swap1(),
            bignum.set_val('ispositive')(vm),
        ], lambda vm: [
            vm.swap1(),
            negate(vm),
            vm.swap1(),
            subtract_bothpositive(vm),
        ]),
    ], lambda vm: [
        vm.push(1),
        vm.eq(),
        # ispositive2 bn1 bn2
        vm.ifelse(lambda vm: [
            # bn1 bn2
            negate(vm),
            vm.swap1(),
            subtract_bothpositive(vm),
        ], lambda vm: [
            negate(vm),
            vm.swap1(),
            negate(vm),
            add_bothpositive(vm),
            negate(vm),
        ])
    ])


@modifies_stack(2, 1)   # bn1 bn2 -> sum
def subtract(vm):
    vm.swap1()
    negate(vm)
    add(vm)


# bn1 bn2 modulus -> (bn1*bn2)%modulus (assume bn1*bn2>=0, modulus>0)
@modifies_stack(3, 1)
def modmul(vm):
    # TODO: make this more efficient by working with smaller intermediate values
    multiply(vm)
    modallpositive(vm)


@modifies_stack(2, 1)  # bn1 bn2 -> bn1*bn2
def multiply(vm):
    vm.dup1()
    bignum.get('ispositive')(vm)
    vm.dup1()
    bignum.get('ispositive')(vm)
    vm.eq()
    # samesign bn1 bn2
    vm.swap2()
    multiplyignoringsign(vm)
    # product samesign
    vm.swap1()
    vm.iszero()
    vm.ifelse(lambda vm: [
        negate(vm)
    ])


@modifies_stack(2, 1)  # bn1 bn2 -> bn1*bn2 (assume bn1, bn2 both >= 0)
def multiplyignoringsign(vm):
    local_vars = Locals(
        vm,
        ['result', 'scratch', 'i', 'j', 'size1', 'size2', 'bn1', 'bn2']
    )

    vm.dup1()
    bignum.get('size')(vm)
    vm.dup1()
    bignum.get('size')(vm)
    vm.push(0)
    vm.push(0)
    zero(vm)
    zero(vm)
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get(['i', 'size1']),
        vm.lt()
    ], lambda vm: [
        vm.push(0),
        local_vars.set_val('j'),
        zero(vm),
        local_vars.set_val('scratch'),
        vm.while_loop(lambda vm: [
            local_vars.get(['j', 'size2']),
            vm.lt()
        ], lambda vm: [
            local_vars.get(['bn1', 'i']),
            getchunk(vm),
            local_vars.get(['bn2', 'j']),
            getchunk(vm),
            vm.mul(),
            local_vars.get(['i', 'j', 'scratch']),
            # i j scratch bn1[i]*bn2[j]
            vm.add(),
            vm.swap1(),
            setchunk(vm),
            local_vars.set_val('scratch'),

            local_vars.get('j'),
            vm.push(1),
            vm.add(),
            local_vars.set_val('j'),
        ]),
        local_vars.get(['result', 'scratch']),
        add_bothpositive(vm),
        local_vars.set_val('result'),

        local_vars.get('i'),
        vm.push(1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    local_vars.discard('result')


@modifies_stack(2, 1)    # int bignum -> int*bignum
def intmultiply(vm):
    vm.dup1()
    bignum.get('ispositive')(vm)
    vm.dup1()
    vm.push(-1)
    # -1 int ispos(bignum) int bignum
    vm.slt()
    # ispos(int) ispos(bignum) int bignum
    vm.dup0()
    vm.iszero()
    # isneg(int) ispos(int) ispos(bignum) int bignum
    vm.ifelse(lambda vm: [
        vm.swap2(),
        vm.push(0),
        vm.sub(),
        vm.swap2(),
    ])
    vm.eq()
    # samesign abs(int) bignum
    vm.swap2()
    # bignum abs(int) samesign
    intmultiplyignoringsign(vm)
    # product samesign
    bignum.set_val('ispositive')(vm)


@modifies_stack(2, 1)   # bignum int -> bignum*int   (assume int>=0)
def intmultiplyignoringsign(vm):
    local_vars = Locals(vm, ['carry', 'i', 'limit', 'bn', 'int'])
    # bignum int
    vm.dup0()
    bignum.get('size')(vm)
    vm.push(0)
    vm.push(0)
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get(['i', 'limit']),
        vm.lt()
    ], lambda vm: [
        local_vars.get(['bn', 'i', 'int', 'carry']),
        getchunk(vm),
        vm.mul(),
        vm.add(),
        # prodWithCarry
        vm.push(_CHUNK_MOD),
        vm.dup1(),
        vm.div(),
        local_vars.set_val('carry'),
        # prodWithCarry
        vm.push(_CHUNK_MOD),
        vm.swap1(),
        vm.mod(),
        local_vars.get(['bn', 'i']),
        setchunk(vm),
        local_vars.set_val('bn'),

        local_vars.get('i'),
        vm.push(1),
        vm.add(),
        local_vars.set_val('i'),
    ])

    local_vars.get('carry')
    vm.ifelse(lambda vm: [
        local_vars.get(['bn', 'i', 'carry']),
        setchunk(vm),
        local_vars.set_val('bn'),
    ])

    local_vars.discard('bn')


@modifies_stack(3, 1)  # x y m -> (x^y)%m   (assume x,y >= 0; m>0)
def modpow(vm):
    # x y m
    vm.dup2()
    vm.swap1()
    modallpositive(vm)
    modpow2(vm)


@modifies_stack(3, 1)  # x y m -> (x^y)%m   (assume x,y >= 0; m>0, x<m)
def modpow2(vm):
    vm.dup1()
    # y x y m
    zero(vm)
    lt(vm)
    vm.ifelse(lambda vm: [
        # x y m
        vm.dup2(),
        vm.dup2(),
        # y m x y m
        vm.push(1),
        vm.swap1(),
        shiftright(vm),
        # y//2 m x y m
        vm.dup2(),
        modpow2(vm),
        # x^(y//2)%m x y m
        vm.swap2(),
        # y x x^(y//2)%m m
        vm.push(0),
        vm.swap1(),
        getchunk(vm),
        # y[0] x x^(y//2)%m m
        vm.push(1),
        vm.bitwise_and(),
        # y[0]%2 x x^(y//2)%m m
        vm.ifelse(lambda vm: [
            # x x^(y//2)%m m
            vm.dup2(),
            # m x x^(y//2)%m m
            vm.swap1(),
            vm.swap2(),
            # x^(y//2)%m m x m
            vm.dup0(),
            modmul(vm),
            # x^(2*(y//2))%m x m
            modmul(vm),
        ], lambda vm: [
            vm.pop(),
            vm.dup0(),
            modmul(vm),
        ])
    ], lambda vm: [
        # x 0 m
        vm.pop(),
        vm.pop(),
        vm.pop(),
        vm.push(1),
        fromint(vm),
    ])


@modifies_stack(2, 1)   # x y -> x%y (assume y>0)
def mod(vm):
    mod_modpositive(vm)


@modifies_stack(2, 1)   # x y -> x%y (assume y>0)
def mod_modpositive(vm):
    vm.dup0()
    zero(vm)
    leq(vm)
    vm.ifelse(lambda vm: [
        # x y
        modallpositive(vm),
    ], lambda vm: [
        vm.dup1(),
        vm.swap1(),
        negate(vm),
        modallpositive(vm),
        # (-x)%y y
        vm.dup0(),
        zero(vm),
        vm.eq(),
        vm.ifelse(lambda vm: [
            vm.swap1(),
            vm.pop(),
            # bignum(0)
        ], lambda vm: [
            vm.swap1(),
            subtract(vm),
        ]),
    ])


@modifies_stack(2, 1)  # x y -> x%y (assume x>=0, y>0)
def modallpositive(vm):
    divmodallpositive(vm)
    vm.pop()


@modifies_stack(2, 2)  # x y -> x//y x%y  (assume x>=0, y>0)
def divmodallpositive(vm):
    trim(vm)
    vm.swap1()
    trim(vm)
    # denom num
    vm.dup0()
    div_initscale(vm)
    vm.dup0()
    vm.auxpush()  # push initscale onto the auxstack
    # initscale denom num
    vm.dup0()
    vm.swap2()
    # denom initscale initscale num
    shiftleft(vm)
    # scaledDenom initscale num
    vm.swap2()
    # num initscale scaledDenom
    shiftleft(vm)
    # scaledNum scaledDenom
    divmod2(vm)
    # q r'
    vm.swap1()
    # r' q
    vm.auxpop()
    # initscale r' q
    vm.swap1()
    shiftright(vm)
    trim(vm)

    # r q
    vm.swap1()


@modifies_stack(2, 2)   # num denom -> quotient remainder
def divmod2(vm):
    local_vars = Locals(
        vm,
        ['qp', 'rp', 'shiftbits', 'shiftwords', 'num', 'denom']
    )

    trim(vm)
    vm.swap1()
    trim(vm)
    vm.swap1()

    vm.dup1()
    bignum.get('size')(vm)
    vm.dup1()
    bignum.get('size')(vm)
    vm.dup1()
    vm.dup1()
    # numsize denomsize numsize denomsize num denom
    vm.lt()
    vm.ifelse(lambda vm: [
        vm.pop(),
        vm.pop(),
        vm.swap1(),
        vm.pop(),
        # num
        zero(vm),
        # 0 num
    ], lambda vm: [
        # numsize denomsize num denom
        vm.eq(),
        vm.ifelse(lambda vm: [
            # num denom
            vm.dup1(),
            vm.dup1(),
            ltbothpositive(vm),
            vm.ifelse(lambda vm: [
                # num denom
                vm.swap1(),
                vm.pop(),
                zero(vm),
                # 0 num
            ], lambda vm: [
                # num denom
                subtract(vm),
                vm.push(1),
                fromint(vm),
                # 1 num-denom
            ]),
        ], lambda vm: [
            # num denom
            vm.dup1(),
            bignum.get('size')(vm),
            vm.dup1(),
            bignum.get('size')(vm),
            # numsize denomsize num denom

            vm.dup1(),
            vm.push(1),
            vm.add(),
            vm.dup1(),
            vm.eq(),
            # numsize==denomsize+1 numsize denomsize num denom
            vm.ifelse(lambda vm: [
                vm.pop(),
                vm.pop(),
                divmod3(vm),
            ], lambda vm: [
                vm.sub(),
                vm.push(-1),
                vm.add(),
                # shiftwords num denom
                vm.dup0(),
                vm.push(_CHUNK_BITS),
                vm.mul(),
                # shiftbits shiftwords num denom
                vm.push(0),
                vm.dup0(),
                local_vars.make(),
                local_vars.get(['num', 'shiftbits', 'denom']),
                shiftright(vm),
                # num' denom
                divmod3(vm),
                # q' r'
                local_vars.set_val(['qp', 'rp']),
                local_vars.get(['num', 'shiftwords', 'denom']),
                loworderwords(vm),
                # s denom
                local_vars.get(['rp', 'shiftbits']),
                shiftleft(vm),
                add(vm),
                divmod2(vm),
                # q r
                local_vars.get(['qp', 'shiftbits']),
                shiftleft(vm),
                add(vm),
                # quot rem
                local_vars.discard(),
            ])
        ])
    ])
    vm.swap1()
    trim(vm)
    vm.swap1()


@modifies_stack(2, 2)  # num denom -> quotient remainder
def divmod3(vm):
    local_vars = Locals(vm, ['t', 'q', 'num', 'denom'])

    vm.dup1()
    vm.push(_CHUNK_MOD)
    # _chunkMod denom num denom
    intmultiply(vm)
    vm.dup0()
    vm.dup2()
    # num denom*_chunkMod denom*_chunkMod num denom
    geq(vm)
    vm.ifelse(lambda vm: [
        # denom*_chunkMod num denom
        vm.swap1(),
        subtract(vm),
        # num-(denom<<_chunkBits) denom
        divmod3(vm),
        # q r
        vm.push(_CHUNK_BITS),
        vm.push(1),
        fromint(vm),
        shiftleft(vm),
        add(vm),
    ], lambda vm: [
        vm.pop(),
        # num denom
        vm.dup1(),
        vm.dup1(),
        divmod_approxquotient(vm),
        # q num denom
        vm.push(_CHUNK_MOD - 1),
        vm.dup1(),
        # q _chunkMod-1 q num denom
        vm.gt(),
        vm.ifelse(lambda vm: [
            # q num denom
            vm.pop(),
            vm.push(_CHUNK_MOD - 1),
        ]),
        # q num denom
        vm.push(0),
        # local_vars: ['t', 'q', 'num', 'denom']
        local_vars.make(),

        local_vars.get(['q', 'denom']),
        intmultiply(vm),
        local_vars.set_val('t'),

        local_vars.get(['t', 'num']),
        gt(vm),
        vm.ifelse(lambda vm: [
            local_vars.get(['q', 't', 'denom']),
            vm.push(-1),
            vm.add(),
            local_vars.set_val('q'),
            # t denom
            subtract(vm),
            local_vars.set_val('t'),
        ]),

        local_vars.get(['t', 'num']),
        gt(vm),
        vm.ifelse(lambda vm: [
            local_vars.get(['q', 't', 'denom']),
            vm.push(-1),
            vm.add(),
            local_vars.set_val('q'),
            # t denom
            subtract(vm),
            local_vars.set_val('t'),
        ]),

        local_vars.get(['num', 't', 'q']),
        subtract(vm),
        trim(vm),
        vm.swap1(),
        fromint(vm),
        local_vars.discard()
        # quotient remainder
    ])


@modifies_stack(2, 1)  # num denom -> approxquot
def divmod_approxquotient(vm):
    vm.swap1()
    vm.dup0()
    bignum.get('size')(vm)
    vm.push(-1)
    vm.add()
    # size(denom)-1 denom num
    vm.swap1()
    getchunk(vm)
    vm.swap1()
    # num approxdenom
    vm.dup0()
    bignum.get('size')(vm)
    # size(num) num approxdenom
    vm.dup1()
    vm.dup1()
    # size(num) num size(num) num approxdenom
    vm.push(-1)
    vm.add()
    vm.swap1()
    getchunk(vm)
    vm.push(_CHUNK_MOD)
    vm.mul()
    # _chunkMod*num[-1] size(num) num approxdenom
    vm.swap2()
    vm.swap1()
    # size(num) num _chunkmod*num[-1] approxdenom
    vm.push(-2)
    vm.add()
    vm.swap1()
    getchunk(vm)
    vm.add()
    vm.div()


@modifies_stack(1, 1)  # denom -> bitsToShift
def div_initscale(vm):
    vm.dup0()
    bignum.get('size')(vm)
    vm.push(-1)
    vm.add()
    vm.swap1()
    # denom size(denom)-1
    getchunk(vm)
    # topchunk
    vm.push(0)
    # i topchunk
    vm.while_loop(lambda vm: [
        vm.dup1(),
        vm.push(2**125),
        vm.bitwise_and(),
        vm.iszero(),
    ], lambda vm: [
        # i topchunk
        vm.push(1),
        vm.add(),
        vm.swap1(),
        vm.push(2),
        vm.mul(),
        vm.swap1(),
    ])
    # i topchunk
    vm.swap1()
    vm.pop()


@modifies_stack(2, 2)  # x y -> x//y x%y. (assume x>=0, y>0)
def divmodallpositive_save(vm):
    local_vars = Locals(vm, ['m', 'partial', 'x', 'y'])

    zero(vm)
    vm.dup0()
    local_vars.make()
    local_vars.get(['x', 'y'])
    lt(vm)
    vm.ifelse(lambda vm: [
        local_vars.get('x'),
        zero(vm),
        # 0 x
    ], lambda vm: [
        local_vars.get(['x', 'y']),
        quotlowerbound(vm),
        vm.dup0(),
        vm.push(1),
        fromint(vm),
        geq(vm),
        # 1>=m m
        vm.ifelse(lambda vm: [
            vm.pop(),
            local_vars.get(['x', 'y', 'y']),
            subtract(vm),
            divmodallpositive(vm),
            vm.push(1),
            fromint(vm),
            add(vm),
        ], lambda vm: [
            local_vars.set_val('m'),
            local_vars.get(['m', 'y', 'x']),
            multiply(vm),
            vm.swap1(),
            # x m*y
            divmodallpositive(vm),
            # q' r'
            local_vars.set_val('partial'),
            # r'
            local_vars.get('y'),
            vm.swap1(),
            # r' y
            divmodallpositive(vm),
            # q'' r
            local_vars.get(['partial', 'm']),
            multiply(vm),
            add(vm),
            # q r
        ])
    ])
    local_vars.discard()


@modifies_stack(2, 1)   # x y -> lowerbound(x/y). (assume x>=y)
def quotlowerbound(vm):
    local_vars = Locals(vm, ['lb', 'prevlb', 'x', 'y'])

    vm.push(1)
    fromint(vm)
    vm.push(2)
    fromint(vm)
    local_vars.make()
    vm.while_loop(lambda vm: [
        local_vars.get(['lb', 'y', 'x']),
        multiply(vm),
        leq(vm),
    ], lambda vm: [
        local_vars.get('lb'),
        vm.dup0(),
        local_vars.set_val('prevlb'),
        vm.dup0(),
        multiply(vm),
        local_vars.set_val('lb'),
    ])
    local_vars.discard('prevlb')


@modifies_stack(2, 1)   # bn1 bn2 -> bn1<bn2
def eq(vm):
    local_vars = Locals(vm, ['eqsofar', 'i', 'bn1', 'bn2'])

    vm.dup1()
    vm.dup1()
    sizeoflarger(vm)
    # size bn1 bn2
    vm.push(-1)
    vm.add()
    vm.push(1)
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get(['i', 'eqsofar']),
        vm.push(-1),
        vm.slt(),
        vm.bitwise_and(),
    ], lambda vm: [
        local_vars.get(['bn1', 'i', 'i', 'bn2']),
        getchunk(vm),
        vm.swap2(),
        getchunk(vm),
        vm.eq(),
        vm.iszero(),
        vm.ifelse(lambda vm: [
            vm.push(0),
            local_vars.set_val('eqsofar'),
        ]),

        local_vars.get('i'),
        vm.push(-1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    local_vars.discard('eqsofar')


@modifies_stack(2, 1)  # bn1 bn2 -> bn1<bn2
def lt(vm):
    vm.dup1()
    bignum.get('ispositive')(vm)
    vm.dup1()
    bignum.get('ispositive')(vm)
    # ispos(bn1) ispos(bn2) bn1 bn2
    vm.ifelse(lambda vm: [
        vm.ifelse(lambda vm: [
            ltbothpositive(vm),
        ], lambda vm: [
            vm.pop(),
            vm.pop(),
            vm.push(0),
        ])
    ], lambda vm: [
        vm.ifelse(lambda vm: [
            vm.pop(),
            vm.pop(),
            vm.push(1),
        ], lambda vm: [
            negate(vm),
            vm.swap1(),
            negate(vm),
            ltbothpositive(vm),
        ])
    ])


@modifies_stack(2, 1)  # bn1 bn2 -> bn1<bn2
def gt(vm):
    vm.swap1()
    lt(vm)


@modifies_stack(2, 1)  # bn1 bn2 -> bn1>=bn2
def geq(vm):
    lt(vm)
    vm.iszero()


@modifies_stack(2, 1)  # bn1 bn2 -> bn1<=bn2
def leq(vm):
    vm.swap1()
    geq(vm)


# bn1 bn2 -> bn1<bn2. (assume both bn1,bn2 non-negative)
@modifies_stack(2, 1)
def ltbothpositive(vm):
    local_vars = Locals(vm, ['undecided', 'islt', 'i', 'bn1', 'bn2'])

    vm.dup1()
    vm.dup1()
    sizeoflarger(vm)
    # size bn1 bn2
    vm.push(-1)
    vm.add()
    vm.push(0)
    vm.push(1)
    local_vars.make()

    vm.while_loop(lambda vm: [
        local_vars.get('i'),
        vm.push(-1),
        vm.slt(),
        local_vars.get('undecided'),
        vm.bitwise_and(),
    ], lambda vm: [
        local_vars.get(['bn2', 'i', 'i', 'bn1']),
        getchunk(vm),
        vm.swap2(),
        getchunk(vm),
        # bn1[i] bn2[i]
        vm.dup1(),
        vm.dup1(),
        vm.lt(),
        vm.ifelse(lambda vm: [
            vm.push(1),
            vm.push(0),
            local_vars.set_val(['undecided', 'islt']),
        ]),
        vm.gt(),
        vm.ifelse(lambda vm: [
            vm.push(0),
            local_vars.set_val('undecided'),
        ]),
        local_vars.get('i'),
        vm.push(-1),
        vm.add(),
        local_vars.set_val('i'),
    ])
    local_vars.discard('islt')


@modifies_stack(2, 3)   # a b -> gcd x y
def egcd(vm):
    local_vars = Locals(vm, ['g', 'x', 'y', 'a', 'b'])
    vm.dup0()
    zero(vm)
    vm.eq()
    vm.ifelse(lambda vm: [
        vm.pop(),
        zero(vm),
        vm.push(1),
        fromint(vm),
        vm.swap2(),
        # b 0 1
    ], lambda vm: [
        # a b
        vm.push(0),
        vm.dup0(),
        vm.dup0(),
        local_vars.make(),

        local_vars.get(['b', 'a', 'a']),
        modallpositive(vm),
        egcd(vm),
        local_vars.set_val(['g', 'y', 'x']),

        local_vars.get(['b', 'a', 'y', 'x', 'y']),
        divmodallpositive(vm),
        vm.swap1(),
        vm.pop(),
        multiply(vm),
        vm.swap1(),
        subtract(vm),

        local_vars.discard('g'),
        # g x-(b//a)*y y
    ])


# a m -> b such that (a*b)%m==1, or Error if none exists
@modifies_stack(2, 1)
def modinv(vm):
    vm.dup1()
    vm.swap1()
    # a m m
    egcd(vm)
    # g x y m
    vm.swap2()
    vm.pop()
    vm.swap1()
    # g x m
    vm.push(1)
    fromint(vm)
    vm.eq()
    # g==1 x m
    vm.ifelse(lambda vm: [
        mod_modpositive(vm)
    ], lambda vm: [
        deliberate_error(vm)
    ])


@modifies_stack(2, 2)   # bn rand -> randbn rand
def randomgen_pos_lessthan(vm):
    vm.push(1)
    fromint(vm)
    vm.swap1()
    subtract(vm)
    randomgen_lessthan(vm)
    vm.push(1)
    fromint(vm)
    add(vm)


@modifies_stack(2, 2)   # bn rand -> randbn rand
def randomgen_lessthan(vm):
    vm.swap1()
    vm.dup1()
    # bn rand bn
    bitlength(vm)
    randomgen(vm)
    # randbn rand bn
    vm.while_loop(lambda vm: [
        vm.dup2(),
        vm.dup1(),
        # randbn bn randbn rand bn
        geq(vm),
    ], lambda vm: [
        # randbn rand bn
        vm.pop(),
        vm.dup1(),
        bitlength(vm),
        randomgen(vm),
    ])
    # randbn rand bn
    vm.swap2()
    vm.pop()
    # rand randbn
    vm.swap1()


@modifies_stack(2, 2)   # nbits rand -> bignum rand
def randomgen(vm):
    local_vars = Locals(vm, ['ret', 'i', 'bitsleft', 'rand'])
    vm.push(0)
    zero(vm)
    local_vars.make()

    vm.while_loop(lambda vm: [
        vm.push(0),
        local_vars.get('bitsleft'),
        vm.sgt(),
    ], lambda vm: [
        vm.push(_CHUNK_BITS),
        local_vars.get('bitsleft'),
        vm.slt(),
        vm.ifelse(lambda vm: [
            # add the final, partial chunk
            local_vars.get(['bitsleft', 'rand']),
            vm.push(2),
            vm.exp(),
            # 2**bitsleft rand
            vm.swap1(),
            random.getmodn(vm),
            # val rand'
            local_vars.get(['ret', 'i']),
            setchunk(vm),
            # ret' rand'
            vm.push(0),
            local_vars.set_val(['bitsleft', 'ret', 'rand']),
        ], lambda vm: [
            # add a full chunk
            vm.push(_CHUNK_MOD),
            local_vars.get('rand'),
            random.getmodn(vm),
            # chunk gen'
            local_vars.get(['ret', 'i']),
            setchunk(vm),
            # ret' gen'
            local_vars.get(['i', 'bitsleft']),
            vm.push(1),
            vm.add(),
            vm.swap1(),
            vm.push(-_CHUNK_BITS),
            vm.add(),
            # bitsleft' i' ret' gen'
            local_vars.set_val(['bitsleft', 'i', 'ret', 'rand']),
        ])
    ])
    local_vars.get(['ret', 'rand'])
    local_vars.discard()


@modifies_stack(2, 2)   # nbits rand -> bignum rand
def randomgen_odd(vm):
    randomgen(vm)
    # bignum rand
    vm.push(0)
    vm.dup1()
    getchunk(vm)
    # chunk[0] bignum rand
    vm.push(1)
    vm.bitwise_or()
    vm.swap1()
    vm.push(0)
    vm.swap1()
    setchunk(vm)


@modifies_stack(1, 1)   # bignum -> mrctx
def _millerrabin_makectx(vm):
    local_vars = Locals(vm, ['a', 'n', 'r', 'd', 'one', 'mone', 'looksprime'])
    local_vars.new()

    vm.push(-1)
    vm.push(0)
    vm.push(1)
    fromint(vm)
    # bn(1) 0 -1 bignum
    local_vars.set_val(['one', 'looksprime', 'r', 'n'])
    local_vars.get(['n', 'one'])
    subtract(vm)
    vm.dup0()
    local_vars.set_val(['mone', 'd'])

    vm.while_loop(lambda vm: [
        vm.push(0),
        local_vars.get('d'),
        getchunk(vm),
        # d[0]
        vm.push(1),
        vm.bitwise_and(),
        vm.iszero(),
    ], lambda vm: [
        local_vars.get(['r', 'd']),
        vm.push(1),
        vm.add(),
        # r' d
        vm.swap1(),
        vm.push(1),
        vm.swap1(),
        shiftright(vm),
        # d' r'
        local_vars.set_val(['d', 'r']),
    ])

    # nonstandard move here--return our local_vars
    vm.auxpop()


@modifies_stack(2, 2)   # mrctx rand -> looksprime rand
def _millerrabin_step(vm):
    local_vars = Locals(vm, ['a', 'n', 'r', 'd', 'one', 'mone', 'looksprime'])
    vm.auxpush()    # nonstandard move -- take our local_vars as arg

    local_vars.get('n')
    randomgen_pos_lessthan(vm)
    local_vars.set_val('a')

    local_vars.get(['a', 'd', 'n', 'one', 'mone'])
    modpow(vm)
    vm.dup0()
    local_vars.set_val('a')
    # a' bn(1) bn(n-1)
    vm.swap1()
    vm.dup1()
    # a' bn(1) a' bn(n-1)
    eq(vm)
    vm.swap2()
    eq(vm)
    vm.bitwise_or()
    vm.ifelse(lambda vm: [
        # looks prime
        vm.push(1),
        local_vars.set_val('looksprime'),
    ], lambda vm: [
        # still looks composite
        vm.while_loop(lambda vm: [
            local_vars.get('r'),
            vm.push(0),
            vm.slt(),
        ], lambda vm: [
            local_vars.get(['a', 'n', 'mone']),
            vm.dup0(),
            modmul(vm),
            # a' bn(-1)
            vm.dup0(),
            # a' a' bn(-1)
            vm.swap2(),
            # bn(-1) a' a'
            eq(vm),
            vm.ifelse(lambda vm: [
                # looks prime
                # a'
                vm.pop(),
                vm.push(-1),
                vm.push(1),
                local_vars.set_val(['looksprime', 'r']),
            ], lambda vm: [
                # still looks composite
                # a'
                local_vars.get('r'),
                vm.push(-1),
                vm.add(),
                # r-1 a'
                local_vars.set_val(['r', 'a']),
            ])
        ])
    ])
    local_vars.discard('looksprime')


@modifies_stack(3, 1)   # bignum rand bitsOfConfidence -> isprime
def isprime(vm):
    local_vars = Locals(vm, ['looksprime', 'mrctx', 'rand', 'confNeeded'])
    _millerrabin_makectx(vm)
    vm.push(1)
    local_vars.make()

    vm.while_loop(lambda vm: [
        vm.push(0),
        local_vars.get('confNeeded'),
        vm.sgt(),
    ], lambda vm: [
        local_vars.get(['mrctx', 'rand']),
        _millerrabin_step(vm),
        vm.ifelse(lambda vm: [
            # mr says looks prime
            # rand
            local_vars.get('confNeeded'),
            vm.push(-2),
            vm.add(),
            local_vars.set_val(['confNeeded', 'rand']),
        ], lambda vm: [
            # mr says composite
            # rand
            vm.push(0),
            vm.dup0(),
            local_vars.set_val(['looksprime', 'confNeeded', 'rand']),
        ])
    ])
    local_vars.discard('looksprime')


@modifies_stack(0, 0)
def deliberate_error(vm):
    vm.push(0)
    vm.push(1)
    vm.div()
