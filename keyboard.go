package botsgocore

// KeyboardType defines keyboard type
type KeyboardType int

//goland:noinspection GoUnusedConst
const (
	// KeyboardTypeNone for no keyboard
	// Used by: Telegram
	KeyboardTypeNone KeyboardType = iota

	// KeyboardTypeHide commands to hide keyboard
	// Used by: Telegram
	KeyboardTypeHide

	// KeyboardTypeInline for inline keyboard
	// Used by: Telegram
	KeyboardTypeInline

	// KeyboardTypeBottom for bottom keyboard
	// Used by: Telegram
	KeyboardTypeBottom

	// KeyboardTypeForceReply to force reply from a user
	// Used by: Telegram
	KeyboardTypeForceReply
)

// Keyboard defines keyboard
type Keyboard interface {
	// KeyboardType defines keyboard type
	KeyboardType() KeyboardType
}
