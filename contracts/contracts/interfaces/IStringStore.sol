// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface IStringStore {

    function getString() external view returns (string memory value);
    function setString(string memory value) external;

}