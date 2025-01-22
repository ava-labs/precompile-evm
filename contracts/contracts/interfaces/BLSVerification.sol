// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface BLSVerification {
    function verifySignature(
        bytes memory signature,
        bytes memory publicKey,
        bytes memory message
    ) external view returns (bool isValid);
}
