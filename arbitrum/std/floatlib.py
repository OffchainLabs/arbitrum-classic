import math
from ..annotation import modifies_stack
from .struct import Struct
from .locals import Locals
from .import tup

_mantbits = 52
_expbits = 11
_bias = -1023

_nan = (((1 << _expbits)-1) << _mantbits) + 1
_inf64 = ((1 << _expbits)-1) << _mantbits
_neg64 = _inf64 ^ (1 << 63)


def getnan():
    return _nan


def getinfinity():
    return _inf64


def getneginfinity():
    return _neg64


_unpackedFloatFields = ['sign', 'mant', 'exp', 'isinf', 'isnan']
_unpackedFloat = Struct("float", _unpackedFloatFields)


@modifies_stack(1, 0)
def toPythonFloat(vm):  # call from Python emulator only; will weird out the compiler
    unpack(vm)
    f = vm.stack[0]
    vm.pop()
    if f[4] != 0:
        return math.nan
    if f[3] != 0:
        if f[0] == 0:
            return math.inf
        else:
            return -math.inf

    floatval = 1 + f[1]/(2**_mantbits)
    expToUse = f[2]+_bias
    if expToUse < 0:
        floatval /= (2**(-expToUse))
    else:
        floatval *= (2*expToUse)
    if f[0] != 0:
        floatval = -floatval
    return floatval


# unpacks the low-order 64 bits of the (int) input into an _unpackedFloat
@modifies_stack(1, 1)
def unpack(vm):
    local_vars = Locals(vm, _unpackedFloatFields)

    vm.dup0()
    vm.dup0()
    vm.push(1 << (_mantbits+_expbits))
    vm.bitwise_and()
    vm.swap2()
    # f f sign
    vm.push((1 << _mantbits) - 1)
    vm.bitwise_and()
    vm.swap1()
    # f mant sign
    vm.push(2**_mantbits)
    vm.swap1()
    vm.div()
    vm.push((1 << _expbits) - 1)
    vm.bitwise_and()
    # exp mant sign
    vm.push(0)
    vm.dup0()
    local_vars.new()
    local_vars.set_val(['isinf', 'isnan', 'exp', 'mant', 'sign'])

    local_vars.get('exp')
    vm.ifelse(lambda vm: [
        local_vars.get('exp'),
        vm.push((1 << _expbits) - 1),
        vm.eq(),
        vm.ifelse(lambda vm: [
            # inf or nan case
            local_vars.get('mant'),
            vm.ifelse(lambda vm: [
                vm.push(1),
                local_vars.set_val('isnan'),
            ], lambda vm: [
                vm.push(1),
                local_vars.set_val('isinf'),
            ]),
        ], lambda vm: [
            # normal case
            local_vars.get(['mant', 'exp']),
            vm.push(1 << _mantbits),
            vm.bitwise_or(),
            local_vars.set_val('mant'),
            vm.push(_bias),
            vm.add(),
            local_vars.set_val('exp'),
        ])
    ], lambda vm: [
        # denormalized case
        local_vars.get('mant'),
        vm.ifelse(lambda vm: [
            local_vars.get(['exp', 'mant']),
            vm.push(_bias + 1),
            vm.add(),
            vm.swap1(),
            # mant exp
            vm.while_loop(lambda vm: [
                vm.dup0(),
                vm.push(1 << _mantbits),
                vm.gt(),
            ], lambda vm: [
                vm.push(2),
                vm.mul(),
                vm.swap1(),
                vm.push(-1),
                vm.add(),
                vm.swap1(),
            ]),
            # mant exp
            local_vars.set_val(['mant', 'exp']),
        ])
    ])
    vm.auxpop()   # return our local_vars
    # unpackedFloat


# sign mant exp trunc ->
# (int whose low 64 bits are IEEE representation of float64)
@modifies_stack(4, 1)
def pack(vm):
    local_vars = Locals(
        vm,
        ['sign', 'mant', 'mant0', 'exp', 'exp0', 'trunc', 'trunc0']
    )
    local_vars.new()
    local_vars.set_val(['sign', 'mant', 'exp', 'trunc'])
    local_vars.get(['mant', 'exp', 'trunc'])
    local_vars.set_val(['mant0', 'exp0', 'trunc0'])

    local_vars.get('mant')
    vm.ifelse(lambda vm: [  # first for-loop in go code
        local_vars.get(['mant', 'exp']),
        vm.while_loop(lambda vm: [
            vm.push(1 << _mantbits),
            vm.dup1(),
            vm.slt(),
        ], lambda vm: [
            # mant exp
            vm.push(2),
            vm.mul(),
            vm.swap1(),
            vm.push(-1),
            vm.add(),
            vm.swap1(),
        ]),
        local_vars.get('trunc'),
        # trunc mant exp
        # second for-loop in go code
        vm.while_loop(lambda vm: [
            vm.dup1(),
            vm.push(4 << _mantbits),
            vm.slt(),
        ], lambda vm: [
            # trunc mant exp
            vm.dup1(),
            vm.push(1),
            vm.bitwise_and(),
            vm.bitwise_or(),
            # trunc' mant exp
            vm.swap1(),
            vm.push(2),
            vm.swap1(),
            vm.div(),
            # mant' trunc' exp
            vm.swap2(),
            vm.push(1),
            vm.add(),
            # exp' trunc' mant'
            vm.swap2(),
            vm.swap1(),
        ]),
        # trunc mant exp
        vm.dup1(),
        vm.push(2 << _mantbits),
        vm.slt(),
        vm.ifelse(lambda vm: [   # in go code, second top-level if statement
            # trunc mant exp
            vm.dup1(),
            vm.push(1),
            vm.bitwise_and(),
            # mant&1 trunc mant exp
            vm.dup2(),
            vm.push(2),
            vm.bitwise_and(),
            vm.dup2(),
            vm.bitwise_or(),
            vm.mul(),
            # (mant&1)&&(trunc||(mant&2)) trunc mant exp
            vm.ifelse(lambda vm: [
                # trunc mant exp
                vm.swap1(),
                vm.push(1),
                vm.add(),
                # mant trunc exp
                vm.dup0(),
                vm.push(4 << _mantbits),
                vm.slt(),
                vm.ifelse(lambda vm: [
                    # mant trunc exp
                    vm.push(2),
                    vm.swap1(),
                    vm.div(),
                    vm.swap2(),
                    vm.push(1),
                    vm.add(),
                    # exp trunc mant
                    vm.swap2(),
                ]),
                # mant trunc exp
                vm.swap1(),
            ]),
            # trunc mant exp
            vm.swap1(),
            vm.push(2),
            vm.swap1(),
            vm.div(),
            # mant trunc exp
            vm.swap2(),
            vm.push(1),
            vm.add(),
            # exp trunc mant
            vm.swap2(),
            vm.swap1(),
        ]),
        # trunc mant exp
        vm.dup2(),
        vm.push((1 << _expbits)-1+_bias),
        vm.slt(),
        vm.ifelse(lambda vm: [        # third top-level if statement in go code
            vm.pop(),
            vm.pop(),
            vm.pop(),
            local_vars.get('sign'),
            vm.push(_inf64),
            vm.bitwise_xor(),
        ], lambda vm: [
            # trunc mant exp
            vm.dup2(),
            vm.push(_bias+1),
            vm.sgt(),
            vm.ifelse(lambda vm: [   # last top-level if statement in go code
                # trunc mant exp
                vm.pop(),
                vm.pop(),
                # exp
                vm.push(_bias - _mantbits),
                vm.sgt(),
                vm.ifelse(lambda vm: [
                    local_vars.get('sign'),
                ], lambda vm: [
                    local_vars.get(['mant0', 'exp0', 'trunc0']),
                    # mant exp trunc
                    vm.while_loop(lambda vm: [
                        vm.dup1(),
                        vm.push(_bias),
                        vm.sgt(),
                    ], lambda vm: [
                        vm.swap2(),
                        vm.dup2(),
                        # mant trunc exp mant
                        vm.push(1),
                        vm.bitwise_and(),
                        vm.bitwise_or(),
                        # trunc exp mant
                        vm.swap2(),
                        vm.push(2),
                        vm.swap1(),
                        vm.div(),
                        # mant exp trunc
                        vm.swap1(),
                        vm.push(1),
                        vm.add(),
                        # exp mant trunc
                        vm.swap1(),
                        # mant exp trunc
                    ]),
                    # mant exp trunc
                    vm.dup2(),
                    vm.dup1(),
                    vm.push(2),
                    vm.bitwise_and(),
                    vm.bitwise_or(),
                    # (trunc | (mant&2)) mant exp trunc
                    vm.iszero(),
                    vm.iszero(),
                    vm.dup1(),
                    vm.push(1),
                    vm.bitwise_and(),
                    vm.bitwise_and(),
                    # (mant&1 != 0 && (trunc != 0 || mant&2 != 0)) mant exp trunc
                    vm.ifelse(lambda vm: [
                        vm.push(1),
                        vm.add(),
                    ]),
                    # mant exp trunc
                    vm.push(2),
                    vm.swap1(),
                    vm.div(),
                    vm.swap1(),
                    vm.push(1),
                    vm.add(),
                    # exp mant trunc
                    vm.dup1(),
                    vm.push(1 << _mantbits),
                    vm.sgt(),
                    vm.ifelse(lambda vm: [
                        vm.pop(),
                        vm.swap1(),
                        vm.pop(),
                        # mant
                        local_vars.get('sign'),
                        vm.bitwise_or(),
                    ], lambda vm: [
                        # exp mant trunc
                        vm.swap2(),
                        # trunc mant exp
                        vm.pop(),
                        # mant exp
                        vm.push((1 << _mantbits)-1),
                        vm.bitwise_and(),
                        vm.swap1(),
                        vm.push(-bias),
                        vm.add(),
                        vm.push(1 << _mantbits),
                        vm.mul(),
                        local_vars.get('sign'),
                        vm.bitwise_or(),
                        vm.bitwise_or(),
                    ]),
                ])
            ], lambda vm: [
                # trunc mant exp
                vm.pop(),
                # mant exp
                vm.push((1 << _mantbits)-1),
                vm.bitwise_and(),
                vm.swap1(),
                vm.push(-_bias),
                vm.add(),
                vm.push(1 << _mantbits),
                vm.mul(),
                local_vars.get('sign'),
                vm.bitwise_or(),
                vm.bitwise_or(),
            ]),
        ]),
    ], lambda vm: [
        local_vars.get('sign')
    ])
    local_vars.discard()


@modifies_stack(2, 1)
def add(vm):
    local_vars = Locals(
        vm,
        [
            'f', 'fs', 'fm', 'fe', 'fi', 'fn', 'g', 'gs',
            'gm', 'ge', 'gi', 'gn', 'shift', 'trunc'
        ]
    )
    local_vars.new()

    vm.dup1()
    # g f g
    local_vars.set_val(['g', 'f'])
    # g
    unpack(vm)
    tup.tbreak(5)(vm)
    local_vars.get('f')
    unpack(vm)
    tup.tbreak(5)(vm)
    local_vars.set_val(
        ['fs', 'fm', 'fe', 'fi', 'fn', 'gs', 'gm', 'ge', 'gi', 'gn']
    )

    local_vars.get(['fn', 'gn'])
    vm.bitwise_or()
    vm.ifelse(lambda vm: [
        vm.push(_nan),
    ], lambda vm: [
        local_vars.get(['fs', 'gs', 'fi', 'gi']),
        vm.eq(),
        vm.iszero(),
        vm.bitwise_and(),
        vm.bitwise_and(),
        # second if statement in go code
        vm.ifelse(lambda vm: [
            vm.push(_nan)
        ], lambda vm: [
            local_vars.get('fi'),
            vm.ifelse(lambda vm: [  # third if in go code
                local_vars.get('f')
            ], lambda vm: [
                local_vars.get('gi'),
                vm.ifelse(lambda vm: [  # fourth if in go code
                    local_vars.get('g'),
                ], lambda vm: [
                    local_vars.get(['fm', 'gm', 'fs', 'gs']),
                    vm.iszero(),
                    vm.swap1(),
                    vm.iszero(),
                    vm.bitwise_and(),
                    vm.bitwise_and(),
                    vm.bitwise_and(),
                    # fifth if in go code
                    vm.ifelse(lambda vm: [
                        local_vars.get('f'),
                    ], lambda vm: [
                        local_vars.get('fm'),
                        vm.iszero(),
                        # sixth if in go code
                        vm.ifelse(lambda vm: [
                            local_vars.get(['gm', 'g']),
                            vm.iszero(),
                            vm.ifelse(lambda vm: [
                                local_vars.get('gs'),
                                vm.bitwise_xor(),
                            ]),
                            # g
                        ], lambda vm: [
                            local_vars.get('gm'),
                            vm.iszero(),
                            # 7th if in top-level go code
                            vm.ifelse(lambda vm: [
                                local_vars.get('f'),
                            ], lambda vm: [
                                local_vars.get(
                                    ['fm', 'gm', 'fe', 'ge', 'fe', 'ge']
                                ),
                                vm.lt(),
                                vm.swap2(),
                                # ge fe fm<gm fe ge
                                vm.eq(),
                                vm.bitwise_and(),
                                # (fe==ge)&(fm<gm) fe ge
                                vm.swap2(),
                                # ge fe ()&()
                                vm.gt(),
                                vm.bitwise_or(),
                                vm.ifelse(lambda vm: [
                                    local_vars.get(
                                        ['g', 'f', 'gs', 'gm', 'ge', 'fs', 'fm', 'fe']
                                    ),
                                    local_vars.set_val(
                                        ['f', 'g', 'fs', 'fm', 'fe', 'gs', 'gm', 'ge']
                                    ),
                                ]),
                                local_vars.get(['fe', 'ge', 'fm', 'gm']),
                                vm.sub(),
                                local_vars.set_val('shift'),
                                # fm gm
                                vm.push(4),
                                vm.mul(),
                                vm.swap1(),
                                vm.push(4),
                                vm.mul(),
                                local_vars.set_val(['gm', 'fm']),
                                local_vars.get(['shift', 'gm']),
                                vm.push(2),
                                vm.exp(),
                                vm.push(-1),
                                vm.add(),
                                # (1<<shift)-1 gm
                                vm.bitwise_and(),
                                local_vars.set_val('trunc'),
                                local_vars.get(['shift', 'gm']),
                                vm.push(2),
                                vm.exp(),
                                vm.swap1(),
                                vm.div(),
                                local_vars.set_val('gm'),
                                local_vars.get(['fs', 'gs', 'fm', 'gm']),
                                vm.eq(),
                                vm.ifelse(lambda vm: [
                                    # fm gm
                                    vm.add(),
                                    local_vars.set_val('fm'),
                                ], lambda vm: [
                                    vm.sub(),
                                    local_vars.set_val('fm'),
                                    local_vars.get('trunc'),
                                    vm.ifelse(lambda vm: [
                                        local_vars.get('fm'),
                                        vm.push(-1),
                                        vm.add(),
                                        local_vars.set_val('fm'),
                                    ]),
                                ]),
                                local_vars.get('fm'),
                                vm.iszero(),
                                vm.ifelse(lambda vm: [
                                    vm.push(0),
                                    local_vars.set_val('fs'),
                                ]),
                                local_vars.get(['fe', 'fm', 'fs', 'trunc']),
                                vm.push(-2),
                                vm.add(),
                                vm.swap2(),
                                pack(vm),
                            ])
                        ])
                    ])
                ])
            ])
        ])
    ])
    local_vars.discard()


@modifies_stack(1, 1)
def negate(vm):
    vm.push(1 << (_mantbits+_expbits))
    vm.bitwise_xor()


@modifies_stack(2, 1)
def sub(vm):
    vm.swap1()
    negate(vm)
    add(vm)


@modifies_stack(2, 1)
def mul(vm):
    local_vars = Locals(
        vm,
        [
            'f', 'fs', 'fm', 'fe', 'fi', 'fn', 'g', 'gs',
            'gm', 'ge', 'gi', 'gn', 'shift', 'trunc'
        ]
    )
    local_vars.new()

    vm.dup1()
    # g f g
    local_vars.set_val(['g', 'f'])
    # g
    unpack(vm)
    tup.tbreak(5)(vm)
    local_vars.get('f')
    unpack(vm)
    tup.tbreak(5)(vm)
    local_vars.set_val(
        ['fs', 'fm', 'fe', 'fi', 'fn', 'gs', 'gm', 'ge', 'gi', 'gn']
    )

    local_vars.get(['fn', 'gn'])
    vm.bitwise_or()
    vm.ifelse(lambda vm: [
        vm.push(_nan),
    ], lambda vm: [
        local_vars.get(['fi', 'gi']),
        vm.bitwise_and(),
        vm.ifelse(lambda vm: [
            local_vars.get(['f', 'gs']),
            vm.bitwise_xor(),
        ], lambda vm: [
            local_vars.get(['gm', 'fi', 'gi', 'fm']),
            vm.iszero(),
            vm.bitwise_and(),
            # (gm==0)&fi gi fm
            vm.swap2(),
            vm.iszero(),
            vm.bitwise_and(),
            vm.bitwise_or(),
            vm.ifelse(lambda vm: [
                vm.push(_nan),
            ], lambda vm: [
                local_vars.get('fm'),
                vm.ifelse(lambda vm: [
                    local_vars.get('gm'),
                    vm.ifelse(lambda vm: [
                        local_vars.get(['fm', 'gm']),
                        vm.mul(),
                        vm.dup0(),
                        vm.push((1 << (_mantbits-1))-1),
                        vm.bitwise_and(),
                        vm.swap1(),
                        # prod trunc
                        vm.push(1 << (_mantbits-1)),
                        vm.swap1(),
                        vm.div(),
                        # mant trunc
                        local_vars.get(['fe', 'ge']),
                        vm.add(),
                        vm.push(-1),
                        vm.add(),
                        vm.swap1(),
                        # mant fe+ge-1 trunc
                        local_vars.get(['fs', 'gs']),
                        vm.bitwise_xor(),
                        pack(vm),
                    ], lambda vm: [
                        local_vars.get(['g', 'fs']),
                        vm.bitwise_or(),
                    ])
                ], lambda vm: [
                    local_vars.get(['f', 'gs']),
                    vm.bitwise_xor(),
                ])
            ])
        ])
    ])
    local_vars.discard()


@modifies_stack(2, 1)
def div(vm):
    local_vars = Locals(
        vm,
        [
            'f', 'fs', 'fm', 'fe', 'fi', 'fn', 'g',
            'gs', 'gm', 'ge', 'gi', 'gn', 'shift', 'trunc'
        ]
    )
    local_vars.new()

    vm.dup1()
    # g f g
    local_vars.set_val(['g', 'f'])
    # g
    unpack(vm)
    print('broken g:', vm.stack[0])
    tup.tbreak(5)(vm)
    local_vars.get('f')
    unpack(vm)
    print('broken f:', vm.stack[0])
    tup.tbreak(5)(vm)
    local_vars.set_val(
        ['fs', 'fm', 'fe', 'fi', 'fn', 'gs', 'gm', 'ge', 'gi', 'gn']
    )

    local_vars.get(['fi', 'gi', 'fn', 'gn'])
    vm.bitwise_and()
    vm.bitwise_or()
    vm.bitwise_or()
    vm.ifelse(lambda vm: [   # first two cases in go code
        vm.push(_nan),
    ], lambda vm: [
        local_vars.get(['fi', 'gi', 'fm', 'gm']),
        vm.bitwise_or(),
        vm.bitwise_or(),
        vm.bitwise_or(),
        vm.iszero(),
        # third case in go code
        vm.ifelse(lambda vm: [
            vm.push(_nan),
        ], lambda vm: [
            local_vars.get(['gi', 'gm', 'fi']),
            vm.iszero(),
            vm.swap1(),
            vm.iszero(),
            vm.bitwise_and(),
            vm.bitwise_or(),
            # 4th case in go code
            vm.ifelse(lambda vm: [
                local_vars.get(['fs', 'gs']),
                vm.push(_inf64),
                vm.bitwise_xor(),
                vm.bitwise_xor(),
            ], lambda vm: [
                local_vars.get(['fm', 'gi']),
                vm.iszero(),
                vm.bitwise_or(),
                # fifth case in go code
                vm.ifelse(lambda vm: [
                    local_vars.get(['fs', 'gs']),
                    vm.bitwise_xor(),
                ], lambda vm: [
                    local_vars.get(['fm', 'gm']),
                    vm.push(1 << (_mantbits+2)),
                    vm.mul(),
                    # fm<<shift gm
                    vm.dup1(),
                    vm.dup1(),
                    vm.mod(),
                    # fms%gm fms gm
                    vm.swap2(),
                    vm.swap1(),
                    vm.div(),
                    # fms//gm fms%gm
                    vm.push(2),
                    local_vars.get(['fe', 'ge']),
                    vm.sub(),
                    vm.sub(),
                    vm.swap1(),
                    local_vars.get(['fs', 'gs']),
                    vm.bitwise_xor(),
                    pack(vm),
                ])
            ])
        ])
    ])
    local_vars.discard()


@modifies_stack(2, 2)   # f g -> cmp isnan
def cmp(vm):
    local_vars = Locals(
        vm,
        [
            'f', 'fs', 'fm', 'fe', 'fi', 'fn', 'g',
            'gs', 'gm', 'ge', 'gi', 'gn', 'shift', 'trunc'
        ]
    )
    local_vars.new()

    vm.dup1()
    # g f g
    local_vars.set_val(['g', 'f'])
    # g
    unpack(vm)
    tup.tbreak(5)(vm)
    local_vars.get('f')
    unpack(vm)
    tup.tbreak(5)(vm)
    local_vars.set_val(
        ['fs', 'fm', 'fe', 'fi', 'fn', 'gs', 'gm', 'ge', 'gi', 'gn']
    )

    local_vars.get(['fn', 'gn'])
    vm.bitwise_or()
    vm.ifelse(lambda vm: [
        # nan case
        vm.push(1),
        vm.push(0),
    ], lambda vm: [
        local_vars.get(['fi', 'gi', 'fm', 'gm']),
        vm.bitwise_or(),
        vm.bitwise_or(),
        vm.bitwise_or(),
        vm.ifelse(lambda vm: [
            # not first two cases in go code
            local_vars.get(['fs', 'gs']),
            vm.eq(),
            vm.ifelse(lambda vm: [
                # not first four cases in go code
                # same sign, non NaN
                vm.push(0),
                local_vars.get(['f', 'g']),
                vm.dup1(),
                vm.dup1(),
                # f g f g 0
                vm.eq(),
                vm.ifelse(lambda vm: [
                    vm.pop(),
                    vm.pop(),
                    # 0
                    vm.dup0(),
                ], lambda vm: [
                    # f g 0
                    vm.lt(),
                    vm.ifelse(lambda vm: [
                        vm.push(-1),
                    ], lambda vm: [
                        vm.push(1)
                    ]),
                    local_vars.get('fs'),
                    vm.ifelse(lambda vm: [
                        vm.push(0),
                        vm.sub(),
                    ])
                ])
            ], lambda vm: [
                vm.push(0),
                local_vars.get(['fs', 'gs']),
                vm.lt(),
                vm.ifelse(lambda vm: [
                    # 0
                    vm.push(1),
                ], lambda vm: [
                    # 0
                    vm.push(-1),
                ])
            ])
        ], lambda vm: [
            # +-0 == +-0 case
            vm.push(0),
            vm.dup0(),
        ]),
    ])
    local_vars.discard()


@modifies_stack(2, 1)
def eq(vm):
    cmp(vm)
    vm.bitwise_or()
    vm.iszero()


@modifies_stack(2, 1)
def gt(vm):
    cmp(vm)
    # cmp nan
    vm.push(1)
    vm.sgt()
    vm.bitwise_or()
    vm.iszero()


def geq(vm):
    cmp(vm)
    vm.push(0)
    vm.sgt()
    vm.bitwise_or()
    vm.iszero()


@modifies_stack(2, 1)
def lt(vm):
    cmp(vm)
    # cmp nan
    vm.push(-1)
    vm.slt()
    vm.bitwise_or()
    vm.iszero()


@modifies_stack(2, 1)
def leq(vm):
    cmp(vm)
    vm.push(0)
    vm.slt()
    vm.bitwise_or()
    vm.iszero()


@modifies_stack(1, 1)
def fromint(vm):
    vm.push((1 << 64)-1)
    vm.bitwise_and()

    vm.dup0()
    vm.push(1 << 63)
    vm.bitwise_and()
    # sign val
    vm.swap1()
    vm.dup1()
    # sign val sign
    vm.ifelse(lambda vm: [
        vm.push(0),
        vm.sub(),
    ])
    # val sign
    vm.push(0)
    vm.swap2()
    # sign val 0
    vm.push(_mantbits)
    vm.swap2()
    vm.swap1()
    # sign mant _mantbits 0
    pack(vm)
