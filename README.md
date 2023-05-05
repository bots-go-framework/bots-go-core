# Package `github.com/bots-go-framework/bots-go-core`

Types and interfaces referenced in bot API clients (like [bots-api-telegram](https://github.com/bots-go-framework/bots-api-telegram)) &amp; [Bots FW](https://github.com/bots-go-framework/bots-fw)

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
