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
)

var _ Button = (*TextButton)(nil)
var _ Button = (*DataButton)(nil)
var _ Button = (*UrlButton)(nil)
var _ Button = (*SwitchInlineQueryButton)(nil)
var _ Button = (*SwitchInlineQueryCurrentChatButton)(nil)

func NewDataButton(text, data string) *DataButton {
	return &DataButton{Text: text, Data: data}
}

type TextButton struct {
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
