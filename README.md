# Package `github.com/bots-go-framework/bots-go-core`

Types and interfaces referenced in bot API clients (like [bots-api-telegram](https://github.com/bots-go-framework/bots-api-telegram)) &amp; [Bots FW](https://github.com/bots-go-framework/bots-fw)

<!-- dev-approach:v1 -->
## Our approach to development

We build with our own tooling:

- **[SpecScore](https://specscore.md)** — specify requirements as `SpecScore.md` artifacts
- **[SpecStudio](https://specscore.studio)** — author & manage specs across their lifecycle
- **[inGitDB](https://ingitdb.com)** — store structured data in Git where applicable
- **[DALgo](https://dalgo.io)** — data access layer for Go
- **[cover100.dev](https://cover100.dev)** — drive toward 100% test coverage
- **[DataTug](https://datatug.io)** — query & explore data
<!-- /dev-approach -->

## [keyboard.go](keyboard.go)

```go
package botsgocore

type KeyboardType int

type Keyboard interface {
	KeyboardType() KeyboardType
}

```

### Keyboard types

- `KeyboardTypeNone`
- `KeyboardTypeReply`
- `KeyboardTypeInline`
- `KeyboardTypeBottom`
- `KeyboardTypeForceReply`

## Bot API clients that use this Go module

- [Telegram API client](https://github.com/bots-go-framework/bots-api-telegram)
