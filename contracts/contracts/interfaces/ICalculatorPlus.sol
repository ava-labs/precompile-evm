// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface ICalculatorPlus {

    function powOfThree(uint256 base) external view returns(uint256 secondPow, uint256 thirdPow, uint256 fourthPow);

    function moduloPlus(uint256 dividend, uint256 divisor) external view returns(uint256 multiple, uint256 remainder);

    function simplFrac(uint256 numerator, uint256 denominator) external view returns(uint256 simplNum, uint256 simplDenom);

}