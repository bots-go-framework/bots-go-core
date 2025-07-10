package botkb

type ButtonType int

type Button interface {
	ButtonType() ButtonType
}

const (
	ButtonTypeData ButtonType = iota
	ButtonTypeURL
)

var _ Button = (*DataButton)(nil)
var _ Button = (*UrlButton)(nil)

type DataButton struct {
	Text string
	Data string
}

func (DataButton) ButtonType() ButtonType {
	return ButtonTypeData
}

type UrlButton struct {
	Text string
	URL  string
}

func (UrlButton) ButtonType() ButtonType {
	return ButtonTypeURL
}
