// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface ICounter {

    function getCounter() external view returns (uint value);
    function incrementCounter() external;
    function setCounter(uint value) external;

}