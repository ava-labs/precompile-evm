// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface Isha256 {

    function hash_with_sha256(string memory value) external view returns (bytes32 hash);

}