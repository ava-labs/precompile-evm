// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface ICalculator {

    function add(uint value1, uint value2) external view returns(uint result);

    function nextTwo(uint value1) external view returns(uint result1, uint result2);

    function repeat(uint times, string memory text) external view returns(string memory result);

}