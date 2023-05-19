package helloworld

import (
	"testing"

	"github.com/ava-labs/subnet-evm/core/state"
	"github.com/ava-labs/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	testGreeting = "test"
	tests        = map[string]testutils.PrecompileTest{
		"set greeting from no role fails": {
			Caller:     allowlist.TestNoRoleAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSetGreeting("test")
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    false,
			ExpectedErr: ErrCannotSetGreeting.Error(),
		},
		"set greeting from enabled address": {
			Caller:     allowlist.TestEnabledAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSetGreeting(testGreeting)
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    false,
			ExpectedRes: []byte{},
			AfterHook: func(t testing.TB, state contract.StateDB) {
				greeting := GetGreeting(state)
				require.Equal(t, greeting, testGreeting)
			},
		},
		"set greeting from admin address": {
			Caller:     allowlist.TestAdminAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSetGreeting(testGreeting)
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    false,
			ExpectedRes: []byte{},
			AfterHook: func(t testing.TB, state contract.StateDB) {
				greeting := GetGreeting(state)
				require.Equal(t, greeting, testGreeting)
			},
		},
		"get default hello from non-enabled address": {
			Caller:     allowlist.TestNoRoleAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSayHello()
				require.NoError(t, err)

				return input
			},
			Config:      NewConfig(common.Big0, nil, nil), // give a zero config for immediate activation
			SuppliedGas: SayHelloGasCost,
			ReadOnly:    true,
			ExpectedRes: func() []byte {
				res, err := PackSayHelloOutput(defaultGreeting)
				if err != nil {
					panic(err)
				}
				return res
			}(),
		},
		"store greeting then say hello from non-enabled address": {
			Caller: allowlist.TestNoRoleAddr,
			BeforeHook: func(t testing.TB, state contract.StateDB) {
				allowlist.SetDefaultRoles(Module.Address)(t, state)
				StoreGreeting(state, testGreeting)
			},
			InputFn: func(t testing.TB) []byte {
				input, err := PackSayHello()
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SayHelloGasCost,
			ReadOnly:    true,
			ExpectedRes: func() []byte {
				res, err := PackSayHelloOutput(testGreeting)
				if err != nil {
					panic(err)
				}
				return res
			}(),
		},
		"set a very long greeting from enabled address": {
			Caller:     allowlist.TestEnabledAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				longString := "a very long string that is longer than 32 bytes and will cause an error"
				input, err := PackSetGreeting(longString)
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    false,
			ExpectedErr: ErrInputExceedsLimit.Error(),
		},
		"readOnly setFeeConfig with noRole fails": {
			Caller:     allowlist.TestNoRoleAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSetGreeting(testGreeting)
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    true,
			ExpectedErr: vmerrs.ErrWriteProtection.Error(),
		},
		"readOnly setFeeConfig with enabled role fails": {
			Caller:     allowlist.TestEnabledAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSetGreeting(testGreeting)
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    true,
			ExpectedErr: vmerrs.ErrWriteProtection.Error(),
		},
		"readOnly setFeeConfig with admin role fails": {
			Caller:     allowlist.TestAdminAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSetGreeting(testGreeting)
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost,
			ReadOnly:    true,
			ExpectedErr: vmerrs.ErrWriteProtection.Error(),
		},
		"insufficient gas setFeeConfig from admin": {
			Caller:     allowlist.TestAdminAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				input, err := PackSetGreeting(testGreeting)
				require.NoError(t, err)

				return input
			},
			SuppliedGas: SetGreetingGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
	}
)

// TestRun tests the Run function of the precompile contract.
// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
func TestHelloWorld(t *testing.T) {
	// RunPrecompileWithAllowListTests will add the allowlist verification logic tests for us.
	// We also add the custom tests defined above as [tests] input to the function.
	allowlist.RunPrecompileWithAllowListTests(t, Module, state.NewTestStateDB, tests)
}

func BenchmarkHelloWorld(b *testing.B) {
	// BenchPrecompileWithAllowList will add the allowlist benchmarks for us.
	// We also add the custom tests defined above as [tests] input to the function.
	allowlist.BenchPrecompileWithAllowList(b, Module, state.NewTestStateDB, tests)
}
