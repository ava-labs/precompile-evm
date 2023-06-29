// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Implements solidity tests.
package solidity

import (
	"github.com/ava-labs/subnet-evm/tests/utils"
	ginkgo "github.com/onsi/ginkgo/v2"
)

var _ = ginkgo.Describe("[Precompiles]", ginkgo.Ordered, func() {
	utils.RegisterPingTest()
	// Each ginkgo It node specifies the name of the genesis file (in ./tests/precompile/genesis/)
	// to use to launch the subnet and the name of the TS test file to run on the subnet (in ./contract-examples/tests/)

	// ADD YOUR PRECOMPILE HERE
	/*
		ginkgo.It("your precompile", ginkgo.Label("Precompile"), ginkgo.Label("YourPrecompile"), func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			// Specify the name shared by the genesis file in ./tests/precompile/genesis/{your_precompile}.json
			// and the test file in ./contracts/tests/{your_precompile}.ts
			// If you want to use a different test command and genesis path than the defaults, you can
			// use the utils.RunTestCMD. See utils.RunDefaultHardhatTests for an example.
			utils.RunDefaultHardhatTests(ctx, "your_precompile")
		})
	*/
})
