// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package precompile

import (
	"os"
	"testing"

	ginkgo "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"

	// Import the solidity package, so that ginkgo maps out the tests declared within the package
	"github.com/ava-labs/precompile-evm/tests/precompile/solidity"
)

func TestE2E(t *testing.T) {
	if basePath := os.Getenv("TEST_SOURCE_ROOT"); basePath != "" {
		err := os.Chdir(basePath)
		require.NoError(t, err)
	}
	gomega.RegisterFailHandler(ginkgo.Fail)
	solidity.RegisterAsyncTests()
	ginkgo.RunSpecs(t, "precompile-evm precompile ginkgo test suite")
}
