// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;
import "@avalabs/subnet-evm-contract-examples/contracts/IAllowList.sol";

interface IHelloWorld is IAllowList {
  // sayHello returns the stored greeting string
  function sayHello() external view returns (string calldata result);

  // setGreeting  stores the greeting string
  function setGreeting(string calldata response) external;
}
