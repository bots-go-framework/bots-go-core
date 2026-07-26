package botkb

type ButtonType int

type Button interface {
	GetText() string
	ButtonType() ButtonType
}

const (
	ButtonTypeText ButtonType = iota
	ButtonTypeData
	ButtonTypeURL
	ButtonTypeSwitchInlineQuery
	ButtonTypeSwitchInlineQueryCurrentChat
	ButtonTypeCopyText
)

var _ Button = (*TextButton)(nil)
var _ Button = (*DataButton)(nil)
var _ Button = (*UrlButton)(nil)
var _ Button = (*SwitchInlineQueryButton)(nil)
var _ Button = (*SwitchInlineQueryCurrentChatButton)(nil)
var _ Button = (*CopyTextButton)(nil)

// ButtonStyle controls the color Telegram uses for an inline keyboard button.
//
// An empty style preserves the client-specific Telegram appearance. The named
// styles require Telegram Bot API 9.4 or newer.
type ButtonStyle string

const (
	ButtonStyleDefault ButtonStyle = ""
	ButtonStylePrimary ButtonStyle = "primary"
	ButtonStyleSuccess ButtonStyle = "success"
	ButtonStyleDanger  ButtonStyle = "danger"

	// ButtonStylePositive is kept for source compatibility. Telegram calls
	// this style "success".
	ButtonStylePositive ButtonStyle = ButtonStyleSuccess

	// ButtonStyleDestructive is kept for source compatibility. Telegram calls
	// this style "danger".
	ButtonStyleDestructive ButtonStyle = ButtonStyleDanger
)

// InlineButtonAppearance contains optional presentation metadata shared by
// Telegram keyboard buttons. Embed it in custom button types to let the
// Telegram renderer carry the same appearance metadata through. The historical
// name is retained because this metadata was first exposed for inline buttons.
type InlineButtonAppearance struct {
	Style             ButtonStyle `json:"style,omitempty"`
	IconCustomEmojiID string      `json:"icon_custom_emoji_id,omitempty"`
}

// GetInlineButtonAppearance allows platform renderers to read optional
// appearance metadata without coupling botkb to a Telegram API package.
func (a InlineButtonAppearance) GetInlineButtonAppearance() InlineButtonAppearance {
	return a
}

type InlineButtonAppearanceProvider interface {
	GetInlineButtonAppearance() InlineButtonAppearance
}

func NewDataButton(text, data string) *DataButton {
	return &DataButton{Text: text, Data: data}
}

type TextButton struct {
	InlineButtonAppearance
	Text string `json:"text"`
}

func NewTextButton(text string) *TextButton {
	return &TextButton{Text: text}
}

func (t TextButton) GetText() string {
	return t.Text
}

func (t TextButton) ButtonType() ButtonType {
	return ButtonTypeText
}

type DataButton struct {
	InlineButtonAppearance
	Text string `json:"text"`
	Data string `json:"data"`
}

func (b DataButton) GetText() string {
	return b.Text
}

func (DataButton) ButtonType() ButtonType {
	return ButtonTypeData
}

type UrlButton struct {
	InlineButtonAppearance
	Text string `json:"text"`
	URL  string `json:"url"`
}

func NewUrlButton(text, url string) *UrlButton {
	return &UrlButton{Text: text, URL: url}
}

func (UrlButton) ButtonType() ButtonType {
	return ButtonTypeURL
}

func (b UrlButton) GetText() string {
	return b.Text
}

type SwitchInlineQueryButton struct {
	InlineButtonAppearance
	Text  string `json:"text"`
	Query string `json:"query"`
}

func NewSwitchInlineQueryButton(text, query string) *SwitchInlineQueryButton {
	return &SwitchInlineQueryButton{Text: text, Query: query}
}

func (SwitchInlineQueryButton) ButtonType() ButtonType {
	return ButtonTypeSwitchInlineQuery
}

func (b SwitchInlineQueryButton) GetText() string {
	return b.Text
}

type SwitchInlineQueryCurrentChatButton struct {
	InlineButtonAppearance
	Text  string `json:"text"`
	Query string `json:"query"`
}

func NewSwitchInlineQueryCurrentChatButton(text, query string) *SwitchInlineQueryCurrentChatButton {
	return &SwitchInlineQueryCurrentChatButton{Text: text, Query: query}
}

func (SwitchInlineQueryCurrentChatButton) ButtonType() ButtonType {
	return ButtonTypeSwitchInlineQueryCurrentChat
}

func (b SwitchInlineQueryCurrentChatButton) GetText() string {
	return b.Text
}

// CopyTextButton is an inline keyboard button that copies CopyText to the
// user's clipboard without sending a callback query to the bot.
type CopyTextButton struct {
	InlineButtonAppearance
	Text     string `json:"text"`
	CopyText string `json:"copy_text"`
}

func NewCopyTextButton(text, copyText string) *CopyTextButton {
	return &CopyTextButton{Text: text, CopyText: copyText}
}

func (CopyTextButton) ButtonType() ButtonType {
	return ButtonTypeCopyText
}

func (b CopyTextButton) GetText() string {
	return b.Text
}
