# Helloworld

There are some must-be-done changes waiting in the generated file.

- Each place requiring you to add your code is marked with `// CUSTOM CODE`
- Add your precompile where the comment `// ADD YOUR PRECOMPILE HERE` is present, to activate your precompile.

For testing, you can refer to other precompile tests in [contract_test.go](contract_test.go) and [config_test.go](config_test.go).

The [hello world precompile tutorial](https://docs.avax.network/subnets/hello-world-precompile-tutorial) should guide you on precompile development.

## General guidelines for precompile development

- In the generated [`module.go`](module.go):
  - Set a suitable config key, for example

    ```go
    const ConfigKey = "yourPrecompileConfig"
    ```

  - Set a suitable contract address, for example:

    ```go
    var ContractAddress = common.HexToAddress("ASUITABLEHEXADDRESS")
    ```

- Only modify code after `// CUSTOM CODE STARTS HERE`. Modifying code outside of these areas should be done with caution and with a good understanding of how changes may impact the EVM.
- Set gas costs in the generated [`contract.go`](contract.go) file, for example:

    ```go
    const (
        // Gas costs for each function. These are set to 1 by default.
        // You should set a gas cost for each function in your contract.
        // Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
        // There are some predefined gas costs in contract/utils.go that you can use.
        // This contract also uses AllowList precompile.
        // You should also increase gas costs of functions that read from AllowList storage.
        SayHelloGasCost    uint64 = contract.ReadGasCostPerSlot
        SetGreetingGasCost uint64 = contract.WriteGasCostPerSlot + allowlist.ReadAllowListGasCost
    )
    ```

- Force import your precompile package in `precompile/registry/registry.go`, for example with:

    ```go
    import (
        _ "github.com/ava-labs/precompile-evm/helloworld"
    )
    ```

- Add your config unit tests in [`config_test.go`](config_test.go)
- Add your contract unit tests in [`contract_test.go`](contract_test.go)
- You can add a full-fledged VM test for your precompile in [`plugin/vm/vm_test.go`](plugin/vm/vm_test.go). See existing precompile tests for examples.
- Add your Solidity interface and test contract to [`contracts/contracts`](../contracts/contracts/)
- Write Solidity contract tests for your precompile in [`contracts/contracts/test`](../contracts/contracts/test)
- Write TypeScript DS-Test counterparts for your Solidity tests in [`contracts/test`](../contracts/test)
- Create your genesis with your precompile enabled in [`tests/precompile/genesis/`](../tests/precompile/genesis)
- Create e2e test for your Solidity test in [`tests/precompile/solidity/suites.go`](../tests/precompile/solidity/suites.go)
- Run your e2e precompile Solidity tests with `./scripts/run_ginkgo.sh`
