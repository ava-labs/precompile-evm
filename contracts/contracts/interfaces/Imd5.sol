// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface Imd5 {

    function hash_with_md5(string memory value) external view returns (bytes16 hash);

}