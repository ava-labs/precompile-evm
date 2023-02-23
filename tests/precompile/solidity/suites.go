// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Implements solidity tests.
package solidity

import (
	"context"

	"github.com/ava-labs/avalanchego/api/health"
	"github.com/ava-labs/subnet-evm/tests/utils"
	ginkgo "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("[Precompiles]", ginkgo.Ordered, func() {
	ginkgo.It("ping the network", ginkgo.Label("setup"), func() {
		client := health.NewClient(utils.DefaultLocalNodeURI)
		healthy, err := client.Readiness(context.Background())
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(healthy.Healthy).Should(gomega.BeTrue())
	})
})

var _ = ginkgo.Describe("[Precompiles]", ginkgo.Ordered, func() {
	// Each ginkgo It node specifies the name of the genesis file (in ./tests/precompile/genesis/)
	// to use to launch the subnet and the name of the TS test file to run on the subnet (in ./contract-examples/tests/)

	// ADD YOUR PRECOMPILE HERE
	/*
		ginkgo.It("your precompile", ginkgo.Label("Precompile"), ginkgo.Label("YourPrecompile"), func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			// Specify the name shared by the genesis file in ./tests/precompile/genesis/{your_precompile}.json
			// and the test file in ./contract-examples/tests/{your_precompile}.ts
			utils.ExecuteHardHatTestOnNewBlockchain(ctx, "your_precompile")
		})
	*/
})
