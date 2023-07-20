// SPDX-License-Identifier: MIT
pragma solidity >= 0.8.0;

import "ds-test/src/test.sol"; 
import {IStringStore} from "../contracts/interfaces/IStringStore.sol";

contract StringStoreTest is DSTest {

    IStringStore stringStore = IStringStore(0x0300000000000000000000000000000000000005);

    function step_getString() public {
        assertEq(stringStore.getString(), "Cornell");
    }

    function step_getSet() public {
        string memory newStr = "Apple";
        stringStore.setString(newStr);
        assertEq(stringStore.getString(), newStr);
    }

}