// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface IMD5 {

    function hashWithMD5(string memory value) external view returns (bytes16 hash);

}