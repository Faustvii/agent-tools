# agent-tools

Small static command-line tools for coding agents, such as
[kritika](https://github.com/home-operations/kritika)'s review runner.

| Tool | What it does |
| ---- | ------------ |
| [`konflate`](cmd/konflate/main.go) | Prints a pull request's rendered Flux diff from a [konflate](https://github.com/home-operations/konflate) server as plain text: `konflate <pr> [resource]`, with the server in `KONFLATE_URL` or `--server`. |

## Install with mise

Each release attaches one archive per platform, `agent-tools-linux-<arch>.tar.gz`,
with every tool inside. Pick the commands with `filter_bins`:

```toml
[tools]
"github:Faustvii/agent-tools" = { version = "1.0.1", filter_bins = ["konflate"] }

[env]
KONFLATE_URL = "https://konflate.example.com"
```

The binaries are static (`CGO_ENABLED=0`), linux amd64 and arm64 only.

## Adding a tool

Put it under `cmd/<name>/`, add a build with that `id` and `binary` to
`.goreleaser.yaml`, and list it in the table above. Every tool shares the
repository's version, which release-please bumps from conventional commits.

## Development

[mise](https://mise.jdx.dev/) pins Go, goreleaser, actionlint and shellcheck in
`.mise.toml`, and CI uses the same file. `mise install`, then:

- `mise run lint`: `go vet`, `goreleaser check` and `actionlint`
- `mise run test`: the Go tests
- `mise run build`: every release archive into `dist/`, unpublished
