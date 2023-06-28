package helloworld

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/subnet-evm/precompile/testutils"

	"github.com/ethereum/go-ethereum/common"
)

// TestVerify tests the verification of Config.
func TestVerify(t *testing.T) {
	admins := []common.Address{allowlist.TestAdminAddr}
	enableds := []common.Address{allowlist.TestEnabledAddr}
	tests := map[string]testutils.ConfigVerifyTest{
		"valid config": {
			Config:        NewConfig(big.NewInt(3), admins, enableds),
			ExpectedError: "",
		},
		// CUSTOM CODE STARTS HERE
		// Add your own Verify tests here, e.g.:
		// "your custom test name": {
		// 	Config: NewConfig(big.NewInt(3), admins, enableds),
		// 	ExpectedError: ErrYourCustomError.Error(),
		// },
		// We don't have a custom verification logic for HelloWorld
		// so we just test the allowlist verification logic with a nil custom verifyTests input.
		// VerifyPrecompileWithAllowListTests will add the allowlist verification logic tests for us.
	}
	// Verify the precompile with the allowlist.
	// This adds allowlist verify tests to your custom tests
	// and runs them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist verify tests.
	allowlist.VerifyPrecompileWithAllowListTests(t, Module, tests)
}

// TestEqual tests the equality of Config with other precompile configs.
func TestEqual(t *testing.T) {
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
			Config:   NewConfig(big.NewInt(3), admins, enableds),
			Other:    NewConfig(big.NewInt(4), admins, enableds),
			Expected: false,
		},
		"same config": {
			Config:   NewConfig(big.NewInt(3), admins, enableds),
			Other:    NewConfig(big.NewInt(3), admins, enableds),
			Expected: true,
		},
		// CUSTOM CODE STARTS HERE
		// Add your own Equal tests here
		"different enabled": {
			Config:   NewConfig(big.NewInt(3), admins, nil),
			Other:    NewConfig(big.NewInt(3), admins, enableds),
			Expected: false,
		},
	}
	// Run allow list equal tests.
	// This adds allowlist equal tests to your custom tests
	// and runs them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist equal tests.
	allowlist.EqualPrecompileWithAllowListTests(t, Module, tests)
}
