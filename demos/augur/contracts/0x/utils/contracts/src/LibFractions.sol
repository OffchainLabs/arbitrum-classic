pragma solidity 0.5.15;

import "../../../utils/contracts/src/LibSafeMath.sol";


library LibFractions {

    using LibSafeMath for uint256;

    /// @dev Safely adds two fractions `n1/d1 + n2/d2`
    /// @param n1 numerator of `1`
    /// @param d1 denominator of `1`
    /// @param n2 numerator of `2`
    /// @param d2 denominator of `2`
    /// @return numerator Numerator of sum
    /// @return denominator Denominator of sum
    function add(
        uint256 n1,
        uint256 d1,
        uint256 n2,
        uint256 d2
    )
        internal
        pure
        returns (
            uint256 numerator,
            uint256 denominator
        )
    {
        if (n1 == 0) {
            return (numerator = n2, denominator = d2);
        }
        if (n2 == 0) {
            return (numerator = n1, denominator = d1);
        }
        numerator = n1
            .safeMul(d2)
            .safeAdd(n2.safeMul(d1));
        denominator = d1.safeMul(d2);
        return (numerator, denominator);
    }

    /// @dev Rescales a fraction to prevent overflows during addition if either
    ///      the numerator or the denominator are > `maxValue`.
    /// @param numerator The numerator.
    /// @param denominator The denominator.
    /// @param maxValue The maximum value allowed for both the numerator and
    ///        denominator.
    /// @return scaledNumerator The rescaled numerator.
    /// @return scaledDenominator The rescaled denominator.
    function normalize(
        uint256 numerator,
        uint256 denominator,
        uint256 maxValue
    )
        internal
        pure
        returns (
            uint256 scaledNumerator,
            uint256 scaledDenominator
        )
    {
        // If either the numerator or the denominator are > `maxValue`,
        // re-scale them by `maxValue` to prevent overflows in future operations.
        if (numerator > maxValue || denominator > maxValue) {
            uint256 rescaleBase = numerator >= denominator ? numerator : denominator;
            rescaleBase = rescaleBase.safeDiv(maxValue);
            scaledNumerator = numerator.safeDiv(rescaleBase);
            scaledDenominator = denominator.safeDiv(rescaleBase);
        } else {
            scaledNumerator = numerator;
            scaledDenominator = denominator;
        }
        return (scaledNumerator, scaledDenominator);
    }

    /// @dev Rescales a fraction to prevent overflows during addition if either
    ///      the numerator or the denominator are > 2 ** 127.
    /// @param numerator The numerator.
    /// @param denominator The denominator.
    /// @return scaledNumerator The rescaled numerator.
    /// @return scaledDenominator The rescaled denominator.
    function normalize(
        uint256 numerator,
        uint256 denominator
    )
        internal
        pure
        returns (
            uint256 scaledNumerator,
            uint256 scaledDenominator
        )
    {
        return normalize(numerator, denominator, 2 ** 127);
    }

    /// @dev Safely scales the difference between two fractions.
    /// @param n1 numerator of `1`
    /// @param d1 denominator of `1`
    /// @param n2 numerator of `2`
    /// @param d2 denominator of `2`
    /// @param s scalar to multiply by difference.
    /// @return result `s * (n1/d1 - n2/d2)`.
    function scaleDifference(
        uint256 n1,
        uint256 d1,
        uint256 n2,
        uint256 d2,
        uint256 s
    )
        internal
        pure
        returns (uint256 result)
    {
        if (s == 0) {
            return 0;
        }
        if (n2 == 0) {
            return result = s
                .safeMul(n1)
                .safeDiv(d1);
        }
        uint256 numerator = n1
            .safeMul(d2)
            .safeSub(n2.safeMul(d1));
        uint256 tmp = numerator.safeDiv(d2);
        return s
            .safeMul(tmp)
            .safeDiv(d1);
    }
}
