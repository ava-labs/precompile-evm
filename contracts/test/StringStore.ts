// (c) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

import { ethers } from "hardhat"
import { test } from "@avalabs/subnet-evm-contracts"
import { factory } from "typescript"

const STRINGSTORE_ADDRESS = "0x0300000000000000000000000000000000000005"

describe("StringStoreTest", function() {
    this.timeout("30s")

    beforeEach("Setup DS-Test", async function () {
        const stringStorePromise = ethers.getContractAt("IStringStore", STRINGSTORE_ADDRESS)

        return ethers.getContractFactory("StringStoreTest").then(factory => factory.deploy())
        .then(contract => {
          this.testContract = contract
          return contract.deployed().then(() => contract)
        })
    })

    test("Testing get function", "step_getString")

    test("Testing get and set function", "step_getSet")
})