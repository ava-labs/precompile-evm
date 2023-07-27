// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package sha256

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	HashWithSHA256GasCost uint64 = 1 /* SET A GAS COST HERE */
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = abi.JSON
	_ = errors.New
	_ = big.NewInt
)

// Singleton StatefulPrecompiledContract and signatures.
var (

	// Sha256RawABI contains the raw ABI of Sha256 contract.
	//go:embed contract.abi
	Sha256RawABI string

	Sha256ABI = contract.ParseABI(Sha256RawABI)

	Sha256Precompile = createSha256Precompile()
)

// UnpackHashWithSHA256Input attempts to unpack [input] into the string type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackHashWithSHA256Input(input []byte) (string, error) {
	res, err := Sha256ABI.UnpackInput("hashWithSHA256", input)
	if err != nil {
		return "", err
	}
	unpacked := *abi.ConvertType(res[0], new(string)).(*string)
	return unpacked, nil
}

// PackHashWithSha256 packs [value] of type string into the appropriate arguments for hashWithSHA256.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackHashWithSHA256(value string) ([]byte, error) {
	return Sha256ABI.Pack("hashWithSHA256", value)
}

// PackHashWithSHA256Output attempts to pack given hash of type [32]byte
// to conform the ABI outputs.
func PackHashWithSHA256Output(hash [32]byte) ([]byte, error) {
	return Sha256ABI.PackOutput("hashWithSHA256", hash)
}

func hashWithSHA256(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, HashWithSHA256GasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the HashWithSHA256Input.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackHashWithSHA256Input(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	var output [32]byte // CUSTOM CODE FOR AN OUTPUT

	output = sha256.Sum256([]byte(inputStruct))

	packedOutput, err := PackHashWithSHA256Output(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createSha256Precompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createSha256Precompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"hashWithSHA256": hashWithSHA256,
	}

	for name, function := range abiFunctionMap {
		method, ok := Sha256ABI.Methods[name]
		if !ok {
			panic(fmt.Errorf("given method (%s) does not exist in the ABI", name))
		}
		functions = append(functions, contract.NewStatefulPrecompileFunction(method.ID, function))
	}
	// Construct the contract with no fallback function.
	statefulContract, err := contract.NewStatefulPrecompileContract(nil, functions)
	if err != nil {
		panic(err)
	}
	return statefulContract
}