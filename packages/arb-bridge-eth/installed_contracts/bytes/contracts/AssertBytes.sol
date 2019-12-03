/*
 * @title Solidity Bytes Assertion Library 
 * @author Gonçalo Sá <goncalo.sa@consensys.net>
 *
 * @dev A Solidity library built to complete assertions in Solidity unit tests.
 *      This library is compliant with the test event convention that the Truffle suite uses.
 */

 pragma solidity ^0.5.0;


library AssertBytes {
    // Event to maintain compatibility with Truffle's Assertion Lib
    event TestEvent(bool indexed result, string message);

    /*
        Function: equal(bytes memory, bytes memory)

        Assert that two tightly packed bytes arrays are equal.

        Params:
            A (bytes) - The first bytes.
            B (bytes) - The second bytes.
            message (string) - A message that is sent if the assertion fails.

        Returns:
            result (bool) - The result.
    */
    function _equal(bytes memory _a, bytes memory _b) internal pure returns (bool) {
        bool returnBool = true;

        assembly {
            let length := mload(_a)

            // if lengths don't match the arrays are not equal
            switch eq(length, mload(_b))
            case 1 {
                // cb is a circuit breaker in the for loop since there's 
                //  no said feature for inline assembly loops
                // cb = 1 - don't breaker
                // cb = 0 - break
                let cb := 1

                let mc := add(_a, 0x20)
                let end := add(mc, length)

                for {
                    let cc := add(_b, 0x20)
                // the next line is the loop condition:
                // while(uint(mc < end) + cb == 2)
                } eq(add(lt(mc, end), cb), 2) {
                    mc := add(mc, 0x20)
                    cc := add(cc, 0x20)
                } {
                    // if any of these checks fails then arrays are not equal
                    if iszero(eq(mload(mc), mload(cc))) {
                        // unsuccess:
                        returnBool := 0
                        cb := 0
                    }
                }
            }
            default {
                // unsuccess:
                returnBool := 0
            }
        }

        return returnBool;
    }

    function equal(bytes memory _a, bytes memory _b, string memory message) internal returns (bool) {
        bool returnBool = _equal(_a, _b);

        _report(returnBool, message);

        return returnBool;
    }

    function notEqual(bytes memory _a, bytes memory _b, string memory message) internal returns (bool) {
        bool returnBool = _equal(_a, _b);

        _report(!returnBool, message);

        return !returnBool;
    }

    /*
        Function: equal(bytes storage, bytes memory)

        Assert that two tightly packed bytes arrays are equal.

        Params:
            A (bytes) - The first bytes.
            B (bytes) - The second bytes.
            message (string) - A message that is sent if the assertion fails.

        Returns:
            result (bool) - The result.
    */
    function _equalStorage(bytes storage _a, bytes memory _b) internal view returns (bool) {
        bool returnBool = true;

        assembly {
            // we know _a_offset is 0
            let fslot := sload(_a_slot)
            let slength := div(and(fslot, sub(mul(0x100, iszero(and(fslot, 1))), 1)), 2)
            let mlength := mload(_b)

            // if lengths don't match the arrays are not equal
            switch eq(slength, mlength)
            case 1 {
                // slength can contain both the length and contents of the array
                // if length < 32 bytes so let's prepare for that
                // v. http://solidity.readthedocs.io/en/latest/miscellaneous.html#layout-of-state-variables-in-storage
                if iszero(iszero(slength)) {
                    switch lt(slength, 32)
                    case 1 {
                        // blank the last byte which is the length
                        fslot := mul(div(fslot, 0x100), 0x100)

                        if iszero(eq(fslot, mload(add(_b, 0x20)))) {
                            // unsuccess:
                            returnBool := 0
                        }
                    }
                    default {
                        // cb is a circuit breaker in the for loop since there's 
                        //  no said feature for inline assembly loops
                        // cb = 1 - don't breaker
                        // cb = 0 - break
                        let cb := 1

                        // get the keccak hash to get the contents of the array
                        mstore(0x0, _a_slot)
                        let sc := keccak256(0x0, 0x20)
                        
                        let mc := add(_b, 0x20)
                        let end := add(mc, mlength)

                        // the next line is the loop condition:
                        // while(uint(mc < end) + cb == 2)
                        for {} eq(add(lt(mc, end), cb), 2) {
                            sc := add(sc, 1)
                            mc := add(mc, 0x20)
                        } {
                            if iszero(eq(sload(sc), mload(mc))) {
                                // unsuccess:
                                returnBool := 0
                                cb := 0
                            }
                        }
                    }
                }
            }
            default {
                // unsuccess:
                returnBool := 0
            }
        }

        return returnBool;
    }

    function equalStorage(bytes storage _a, bytes memory _b, string memory message) internal returns (bool) {
        bool returnBool = _equalStorage(_a, _b);

        _report(returnBool, message);

        return returnBool;
    }

    function notEqualStorage(bytes storage _a, bytes memory _b, string memory message) internal returns (bool) {
        bool returnBool = _equalStorage(_a, _b);

        _report(!returnBool, message);

        return !returnBool;
    }

    /********** Maintaining compatibility with Truffle's Assert.sol ***********/
    /******************************** internal ********************************/

    /*
        Function: _report

        Internal function for triggering <TestEvent>.

        Params:
            result (bool) - The test result (true or false).
            message (string) - The message that is sent if the assertion fails.
    */

    function _report(bool result, string memory message) internal {
        if (result)
            emit TestEvent(true, "");
        else
            emit TestEvent(false, message);
    }
}
