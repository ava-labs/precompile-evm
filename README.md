# Precompile-EVM

Precompile-EVM is a repository for registering precompiles to Subnet-EVM without forking the Subnet-EVM codebase. Subnet-EVM supports registering external precompiles through `precompile/modules` package. By importing Subnet-EVM as a library, you can register your own precompiles to Subnet-EVM and build it together with Subnet-EVM.

## Environment Setup

To effectively build, run, and test Precompile-EVM, the following is a (non-exhaustive) list of dependencies that you will need:

- Golang
- Node.js
- [AvalancheGo](https://github.com/ava-labs/avalanchego)

To get started easily, we provide a Dev Container specification, that can be used using GitHub Codespace or locally using Docker and VS Code. DevContainers are a concept that utilizes containerization (via Docker containers) to create consistent and isolated development environment. We can access this environment through VS code, which allows for the development experience to feel as if you were developing locally..

### Dev Container in Codespace

Codespaces is a development environment service offered by GitHub that allows developers to write, run, test, and debug their code directly on a cloud machine provided by GitHub. The developer can edit the code through a VS Code running in the browser or locally.

To run a Codespace click on the **Code** and switch to the **Codespaces** tab. There, click **Create Codespace on branch [...]**.

### Local Dev Container

In order to run the Dev Container locally:

- Install VS Code, Docker and the [Dev Container Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
- Clone the Repository
- Open the Container by issuing the Command "Dev Containers: Reopen in Container" in the VS Code command palette (on Mac-OS, run [Cmd + Shift + P]).

## Learn about Precompile-EVM

To get a comprehensive introduction to Precompile-EVM, take the Avalanche Academy course on [Customizing the EVM](https://academy.avax.com/course/customizing-evm).

## Hello World Example

### 1. Clone the Repo

```bash
git clone https://github.com/ava-labs/precompile-evm.git
cd precompile-evm/
```

### 2. Checkout the `hello-world-example` Branch

```bash
git checkout hello-world-example
```

### 3. Install NodeJS Dependencies

First you have to `cd contracts/` and run `npm install` to get the dependencies.

```bash
cd contracts/
npm install
```

### 4. Create a New Contract

`hello-world-example` branch has already a precompile contract called `HelloWorld.sol`. All necessary files were already created for you. You can check existing files and see how a fully implemented precompile should look like. If you'd like to redo steps to create a new precompile contract, you can follow the steps below.

Copy the existing `IHelloWorld.sol` interface to a new file called `IHolaMundo.sol`.

```bash
cd .. # change directory back to the root of the repo
```
```bash
cp contracts/contracts/interfaces/IHelloWorld.sol contracts/contracts/interfaces/IHolaMundo.sol
```

### 5. Generate an ABI

Now generate a `.abi` from a `.sol` using `solc`.

Passing in the following flags

- `--abi`
  - ABI specification of the contracts.
- `--base-path path`
  - Use the given path as the root of the source tree instead of the root of the filesystem.
- `--include-path path`
  - Make an additional source directory available to the default import callback. Use this option if you want to import contracts whose location is not fixed in relation to your main source tree, e.g. third-party libraries installed using a package manager. Can be used multiple times. Can only be used if base path has a non-empty value.
- `--output-dir path`
  - If given, creates one file per output component and contract/file at the specified directory.
- `--overwrite`
  - Overwrite existing files (used together with `--output-dir`).

```bash
cd contracts/ # change directory to the contracts/ directory
```
```bash
npx solc --abi contracts/interfaces/IHolaMundo.sol --output-dir abis --base-path . --include-path ./node_modules
```

Rename the files for easier readability
```bash
mv abis/@avalabs_subnet-evm-contracts_contracts_interfaces_IAllowList_sol_IAllowList.abi abis/IAllowList.abi
```
```bash
mv abis/contracts_interfaces_IHolaMundo_sol_IHelloWorld.abi abis/IHolaMundo.abi
```

### 6. Generate Precompile Files

First, you need to create your precompile contract interface in the `contracts` directory and build the ABI. Then you can generate your precompile files with `./scripts/generate_precompile.sh --abi {abiPath} --out {outPath}`. This script installs the `precompilegen` tool from Subnet-EVM and runs it to generate your precompile.

```bash
cd .. # change directory back to the root directory of the repo
```
```bash
./scripts/generate_precompile.sh --abi contracts/abis/IHolaMundo.abi --out holamundo/
```
```bash
Using branch: hello-world-example
installing precompilegen from Subnet-EVM v0.5.2
generating precompile with Subnet-EVM v0.5.2
Precompile files generated successfully at:  holamundo/
```

Confirm that the new `holamundo/` directory has the appropriate files.

```bash
ls -lh holamundo
```
```bash
-rw-r--r--  1 user group 2.3K Jul  5 13:26 README.md
-rw-r--r--  1 user group 2.3K Jul  5 13:26 config.go
-rw-r--r--  1 user group 2.8K Jul  5 13:26 config_test.go
-rw-r--r--  1 user group 963B Jul  5 13:26 contract.abi
-rw-r--r--  1 user group 8.1K Jul  5 13:26 contract.go
-rw-r--r--  1 user group 8.3K Jul  5 13:26 contract_test.go
-rw-r--r--  1 user group 2.7K Jul  5 13:26 module.go
```

### 7. Register Precompile

In `plugin/main.go` Subnet-EVM is already imported and ready to be Run from the main package.

All you need to do is explicitly register your precompiles to Subnet-EVM in `plugin/main.go` and build it together with Subnet-EVM.

Precompiles generated by `precompilegen` tool have a self-registering mechanism in their `module.go/init()` function.

```go
package main
import (
	"fmt"
	"github.com/ava-labs/avalanchego/version"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/plugin/runner"
	// Each precompile generated by the precompilegen tool has a self-registering init function
	// that registers the precompile with the subnet-evm. Importing the precompile package here
	// will cause the precompile to be registered with the subnet-evm.
	_ "github.com/ava-labs/precompile-evm/helloworld"
	// ADD YOUR PRECOMPILE HERE
	//_ "github.com/ava-labs/precompile-evm/holamundo"
)
```

### 8. Build

You can build your precompile and Subnet-EVM with
```bash
./scripts/build.sh
```
This script builds Subnet-EVM, and your precompile together and generates a binary file. The binary file is compatible with AvalancheGo plugins.

### 9. Run

You can now run Precompile-EVM by using the Avalanche-CLI

```bash
avalanche blockchain create myblockchain --custom --vm $AVALANCHEGO_PLUGIN_PATH/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy --genesis ./.devcontainer/genesis-example.json
```

Then launch the blockchain with your custom VM:

```bash
avalanche blockchain deploy myblockchain
```

### Test

You can create contract tests in `contracts/test` with the Hardhat test framework. These can be run by adding ginkgko test cases in `tests/precompile/solidity/suites.go` and a suitable genesis file in `tests/precompile/genesis`. You can install AvalancheGo binaries with `./scripts/install_avalanchego_release.sh` then run the tests with `./scripts/run_ginkgo.sh`

## Changing Versions

In order to upgrade the Subnet-EVM version, you need to change the version in `go.mod` and `scripts/versions.sh`. You can also change the AvalancheGo version through `scripts/versions.sh` as well. Then you can run `./scripts/build.sh` to build the plugin with the new version.

## AvalancheGo Compatibility

```text
[v0.2.0] AvalancheGo@v1.11.0-v1.11.1 (Protocol Version: 33)
<<<<<<< HEAD
[v0.2.1] AvalancheGo@v1.11.3-v1.11.9 (Protocol Version: 35)
[v0.2.2] AvalancheGo@v1.11.3-v1.11.9 (Protocol Version: 35)
[v0.2.3] AvalancheGo@v1.11.3-v1.11.9 (Protocol Version: 35)
[v0.2.4] AvalancheGo@v1.11.11 (Protocol Version: 37)
```
=======
[v0.2.1] AvalancheGo@v1.11.3-v1.11.7 (Protocol Version: 35)
[v0.2.2] AvalancheGo@v1.11.3-v1.11.7 (Protocol Version: 35)
[v0.2.3] AvalancheGo@v1.11.3-v1.11.7 (Protocol Version: 35)
```
>>>>>>> 9823d9df75a3db1a5466075297b4b03f088ff390
