// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;
import "@avalabs/subnet-evm-contracts/contracts/interfaces/IAllowList.sol";

interface IHelloWorld is IAllowList {
  event GreetingChanged(address indexed sender, string oldGreeting, string newGreeting);
  // sayHello returns the stored greeting string
  function sayHello() external view returns (string calldata result);

  // setGreeting  stores the greeting string
  function setGreeting(string calldata response) external;
}
