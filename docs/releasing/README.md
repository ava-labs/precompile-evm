# Releasing

## When to release

- When a [Subnet-EVM](https://github.com/ava-labs/avalanchego/releases) release is made
- If there is a significant code change or bugfix in this repository

## Procedure

In this section, we create a release `v0.4.0`, upgrading the subnet-evm Go dependency to `v0.7.4` and the AvalancheGo version to the one used by subnet-evm, `v1.13.0`.

We therefore assign these environment variables to simplify copying instructions:

```bash
export VERSION=v0.4.0
export SUBNET_EVM_VERSION=v0.7.4
```

1. Create your branch, usually from the tip of the `main` branch:

    ```bash
    git fetch origin main:main
    git checkout main
    git checkout -b "releases/$VERSION"
    ```

1. Modify the [plugin/main.go](../../plugin/main.go) `Version` global string constant and set it to the desired `$VERSION`.
1. Upgrade the [subnet-evm Go dependency](https://github.com/ava-labs/subnet-evm/releases), for example to version `v0.7.4`:

    ```bash
    go get "github.com/ava-labs/subnet-evm@$SUBNET_EVM_VERSION"
    go mod tidy
    ```

    This will also upgrade the AvalancheGo version to the one used by subnet-evm.
1. Get the Go version from the `go.mod` file:

    ```bash
    grep -oE "^go (1\.[0-9]+)" go.mod | cut -d' ' -f2
    ```

    and update the [.devcontainer/devcontainer.json](../../.devcontainer/devcontainer.json) file with the Go version at `.features.["ghcr.io/devcontainers/features/go:1.version"].version`, for example:

    ```json
    {
        "features": {
            "ghcr.io/devcontainers/features/go:1": {"version": 1.23},
        }
    }
    ```

1. Get the AvalancheGo version from the `go.mod` file:

    ```bash
    go list -m all | grep github.com/ava-labs/avalanchego | awk '{print $2}'
    ```

    and update the [.devcontainer/devcontainer.json](../../.devcontainer/devcontainer.json) file with it at `.build.args.AVALANCHEGO_VERSION`, for example:

    ```json
    {
        "build": {
            "args": {
                "AVALANCHEGO_VERSION": "v1.13.0"
            }
        },
    }
    ```

1. Add an entry in the object in [compatibility.json](../../compatibility.json), adding the target release `$VERSION` as key and the AvalancheGo RPC chain VM protocol version as value, to the `"rpcChainVMProtocolVersion"` JSON object. For example, we would add:

    ```json
    "v0.4.0": 39,
    ```

    💁 If you are unsure about the RPC chain VM protocol version, set the version to `0`, for example `"v0.4.0": 0`, and then run:

    ```bash
    go test -run ^TestCompatibility$ github.com/ava-labs/precompile-evm/plugin
    ```

    This will fail with an error similar to:

    ```text
    compatibility.json has precompile-evm version v0.4.0 stated as compatible with RPC chain VM protocol version 0 but AvalancheGo protocol version is 39
    ```

    This message can help you figure out what the correct RPC chain VM protocol version (here `39`) has to be in compatibility.json for your current release. Alternatively, you can refer to the [Avalanchego repository `version/compatibility.json` file](https://github.com/ava-labs/avalanchego/blob/main/version/compatibility.json) to find the RPC chain VM protocol version matching the AvalancheGo version we use here.
1. Specify the AvalancheGo compatibility in the [README.md relevant section](../../README.md#avalanchego-compatibility). For example we would add:

    ```text
    ...
    [v0.4.0] AvalancheGo@v1.12.2/1.13.0-fuji/1.13.0 (Protocol Version: 39)
    ```

1. Commit your changes and push the branch

    ```bash
    git add .
    git commit -S -m "chore: release $VERSION"
    git push -u origin "releases/$VERSION"
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
    git rebase "releases/$VERSION"
    # Fix eventual conflicts
    git push --force hello-world-example
    ```

1. [Wait for the checks](https://github.com/ava-labs/precompile-evm/pull/12/checks) of the `hello-world-example` branch [PR](https://github.com/ava-labs/precompile-evm/pull/12) to pass. **Never merge this PR**. You can also use `gh` with:

    ```bash
    gh pr checks 12 --watch
    ```

1. Squash and merge your release branch into `main`, for example:

    ```bash
    gh pr merge "releases/$VERSION" --squash --delete-branch --subject "chore: release $VERSION" --body "\n- Bump subnet-evm from v0.7.3 to v0.7.4\n- Update AvalancheGo from v1.12.3 to v1.13.0"
    ```

1. Create and push a tag from the `main` branch:

    ```bash
    git fetch origin main:main
    git checkout main
    # Double check the tip of the main branch is the expected commit
    # of the squashed release branch
    git log -1
    git tag -s "$VERSION"
    git push origin "$VERSION"
    ```

1. Create a new release on Github, either using:
    - the [Github web interface](https://github.com/ava-labs/subnet-evm/releases/new)
        1. In the "Choose a tag" box, select the tag previously created `$VERSION` (`v0.4.0`)
        1. Pick the previous tag, for example as `v0.3.1`.
        1. Set the "Release title" to `$VERSION` (`v0.4.0`)
        1. Set the description using this format:

            ```markdown
            # AvalancheGo Compatibility

            The plugin version is unchanged at 39 and is compatible with AvalancheGo version v1.13.0.

            # Breaking changes

            # Features

            # Fixes

            # Documentation

            ```

        1. Only tick the box "Set as the latest release"
        1. Click on the "Create release" button
    - the Github CLI `gh`:

        ```bash
        PREVIOUS_VERSION=v0.3.1
        NOTES="# AvalancheGo Compatibility

        The plugin version is unchanged at 39 and is compatible with AvalancheGo version v1.13.0.

        # Breaking changes

        # Features

        # Fixes

        # Documentation

        "
        gh release create "$VERSION" --notes-start-tag "$PREVIOUS_VERSION" --notes-from-tag "$VERSION" --title "$VERSION" --notes "$NOTES" --verify-tag
        ```
