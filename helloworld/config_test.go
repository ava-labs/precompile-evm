// (c) 2022 Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package helloworld

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ethereum/go-ethereum/common"
)

func TestVerify(t *testing.T) {
	// We don't have a custom verification logic for HelloWorld
	// so we just test the allowlist verification logic with a nil custom verifyTests input.
	// VerifyPrecompileWithAllowListTests will add the allowlist verification logic tests for us.
	allowlist.VerifyPrecompileWithAllowListTests(t, Module, nil)
}

func TestEqualHelloWorldConfig(t *testing.T) {
	admins := []common.Address{allowlist.TestAdminAddr}
	enableds := []common.Address{allowlist.TestEnabledAddr}
	tests := map[string]testutils.ConfigEqualTest{
		"non-nil config and nil other": {
			Config:   NewConfig(big.NewInt(3), admins, enableds),
			Other:    nil,
			Expected: false,
		},
		"different type": {
			Config:   NewConfig(big.NewInt(3), admins, enableds),
			Other:    precompileconfig.NewNoopStatefulPrecompileConfig(),
			Expected: false,
		},
		"different timestamp": {
			Config:   NewConfig(big.NewInt(3), admins, nil),
			Other:    NewConfig(big.NewInt(4), admins, nil),
			Expected: false,
		},
		"different enabled": {
			Config:   NewConfig(big.NewInt(3), admins, nil),
			Other:    NewConfig(big.NewInt(3), admins, enableds),
			Expected: false,
		},
		"same config": {
			Config:   NewConfig(big.NewInt(3), admins, nil),
			Other:    NewConfig(big.NewInt(3), admins, nil),
			Expected: true,
		},
	}
	// EqualPrecompileWithAllowListTests will add the allowlist verification logic tests for us.
	// We also add the custom tests defined above as [tests] input to the function.
	allowlist.EqualPrecompileWithAllowListTests(t, Module, tests)
}
