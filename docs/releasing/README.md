# Releasing

## When to release

- When a [Subnet-EVM](https://github.com/ava-labs/avalanchego/releases) release is made
- If there is a significant code change or bugfix in this repository

## Procedure

In this section, we create a release `v0.4.0`. We therefore assign these environment variables to simplify copying instructions:

```bash
export VERSION=v0.4.0
export GITHUB_USER=username
```

`GITHUB_USER` is the username of the person creating the release. This is simply used to prefix branches.

1. Create your branch, usually from the tip of the `main` branch:

    ```bash
    git fetch origin main:main
    git checkout main
    git checkout -b "$GITHUB_USER/releases/$VERSION"
    ```

1. Modify the [plugin/main.go](../../plugin/main.go) `Version` global string constant and set it to the desired `$VERSION`.
1. Upgrade the [subnet-evm Go dependency](https://github.com/ava-labs/subnet-evm/releases), which will likely also upgrade the AvalancheGo dependency version:

    ```bash
    go get github.com/ava-labs/subnet-evm
    go mod tidy
    ```

1. Add an entry in the object in [compatibility.json](../../compatibility.json), adding the target release `$VERSION` as key and the AvalancheGo RPC chain VM protocol version as value, to the `"rpcChainVMProtocolVersion"` JSON object. For example, we would add:

    ```json
    "v0.4.0": 39,
    ```

    💁 If you are unsure about the RPC chain VM protocol version:

    1. Check [go.mod](../../go.mod) and spot the version used for `github.com/ava-labs/avalanchego`. For example `v1.13.0`.
    1. Refer to the [Avalanchego repository `version/compatibility.json` file](https://github.com/ava-labs/avalanchego/blob/main/version/compatibility.json) to find the RPC chain VM protocol version matching the AvalancheGo version we use here. In our case, we use an AvalancheGo version `v1.13.0`, so the RPC chain VM protocol version is `39`:

        ```json
        {
            "39": [
                "v1.12.2",
                "v1.13.0"
            ],
        }
        ```

    Finally, check the RPC chain VM protocol version compatibility is setup properly by running:

    ```bash
    go test -run ^TestCompatibility$ github.com/ava-labs/precompile-evm/plugin
    ```

1. Specify the AvalancheGo compatibility in the [README.md relevant section](../../README.md#avalanchego-compatibility). For example we would add:

    ```text
    ...
    [v0.4.0] AvalancheGo@v1.12.2/1.13.0-fuji/1.13.0 (Protocol Version: 39)
    ```

1. Commit your changes and push the branch

    ```bash
    git add .
    git commit -S -m "chore: release $VERSION"
    git push -u origin "$GITHUB_USER/releases/$VERSION"
    ```

1. Create a pull request (PR) from your branch targeting main, for example using [`gh`](https://cli.github.com/):

    ```bash
    gh pr create --repo github.com/ava-labs/precompile-evm --base main --title "chore: release $VERSION"
    ```

1. Wait for the PR checks to pass
1. Update the `hello-world-example` branch to be rebased on your release branch:

    ```bash
    git fetch origin hello-world-example:hello-world-example
    git checkout hello-world-example
    git rebase "$GITHUB_USER/releases/$VERSION"
    # Fix eventual conflicts
    git push --force hello-world-example
    ```

1. [Wait for the checks](https://github.com/ava-labs/precompile-evm/pull/12/checks) of the `hello-world-example` branch [PR](https://github.com/ava-labs/precompile-evm/pull/12) to pass. **Never merge this PR**.
1. Squash and merge your release branch into `main`:

    ```bash
    git checkout main
    git merge --squash "$GITHUB_USER/releases/$VERSION"
    git commit -S -m "chore: release $VERSION"
    git push
    ```

1. Create a release branch from the `main` branch, this time without your username prefix:

    ```bash
    git checkout -b "releases/$VERSION"
    ```

    This is to avoid creating a release targeting the `main` branch, which may contain new commits merged whilst the release is being created.
1. Create a new release through the [Github web interface](https://github.com/ava-labs/subnet-evm/releases/new)
    1. In the "Choose a tag" box, enter `$VERSION` (`v0.4.0`)
    1. In the "Target", pick the branch `releases/$VERSION` (`releases/v0.4.0`)
    1. Pick the previous release, for example as `v0.3.0`.
    1. Set the "Release title" to `$VERSION` (`v0.4.0`)
    1. Set the description (breaking changes, features, fixes, documentation)
    1. Only tick the box "Set as the latest release"
    1. Click on the "Create release" button
