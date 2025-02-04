// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface IVerifyBLS {

    function verifySignatureBLS(string calldata message, bytes calldata signature, bytes calldata publicKey) external view returns(bool result);
}