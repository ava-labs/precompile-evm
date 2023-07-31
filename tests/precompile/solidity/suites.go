// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Implements solidity tests.
package solidity

import (
	"github.com/onsi/ginkgo/v2"
)

// Registers the Asynchronized Precompile Tests
// Before running the tests, this function creates all subnets given in the genesis files
// and then runs the hardhat tests for each one asynchronously if called with `ginkgo run -procs=`.
func RegisterAsyncTests() {
	/* Uncomment these if you want to use default hardhat tests
	// Tests here assumes that the genesis files are in ./tests/precompile/genesis/
	// with the name {precompile_name}.json
	genesisFiles, err := utils.GetFilesAndAliases("./tests/precompile/genesis/*.json")
	if err != nil {
		ginkgo.AbortSuite("Failed to get genesis files: " + err.Error())
	}
	if len(genesisFiles) == 0 {
		ginkgo.AbortSuite("No genesis files found")
	}
	subnetsSuite := utils.CreateSubnetsSuite(genesisFiles)
	*/
	_ = ginkgo.Describe("[Asynchronized Precompile Tests]", func() {
		// Uncomment below and register the ping test first
		// utils.RegisterPingTest()

		// ADD YOUR PRECOMPILE HERE
		/*
			ginkgo.It("your precompile", ginkgo.Label("Precompile"), ginkgo.Label("YourPrecompile"), func() {
				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()

				// Specify the name shared by the genesis file in ./tests/precompile/genesis/{your_precompile}.json
				// and the test file in ./contracts/tests/{your_precompile}.ts
				blockchainID := subnetsSuite.GetBlockchainID("{your_precompile}")
				runDefaultHardhatTests(ctx, blockchainID, "{your_precompile}")
			})
		*/
	})
}

/* Uncomment this if you want to use default hardhat tests
//	Default parameters are:
//
// 1. Hardhat contract environment is located at ./contracts
// 2. Hardhat test file is located at ./contracts/test/<test>.ts
// 3. npx is available in the ./contracts directory
func runDefaultHardhatTests(ctx context.Context, blockchainID, testName string) {
	cmdPath := "./contracts"
	// test path is relative to the cmd path
	testPath := fmt.Sprintf("./test/%s.ts", testName)
	utils.RunHardhatTests(ctx, blockchainID, cmdPath, testPath)
}
*/
