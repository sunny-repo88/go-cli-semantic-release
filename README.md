# example-go-application

This example repository shows how to use go-semantic-release to release a Go application with GitHub Actions.

Releases are created automatically when a commit is pushed to the main branch. All releases can be found on the [releases page](https://github.com/go-semantic-release/example-go-application/releases).

The binaries can be downloaded directly from the release page or via [get-release.xyz](https://get-release.xyz).
```bash
curl -SL "https://get-release.xyz/go-semantic-release/example-go-application/$(go env GOOS)/$(go env GOARCH)" -o ./example-go-application.tar.gz
tar -xzf ./example-go-application.tar.gz
```
