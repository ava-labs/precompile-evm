// (c) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"github.com/ava-labs/subnet-evm/plugin/runner"

	_ "github.com/ava-labs/precompilevm/helloworld"
)

func main() {
	runner.Run()
}
