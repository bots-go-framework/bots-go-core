package botkb

import (
	"testing"
)

func TestTextButton(t *testing.T) {
	tests := []struct {
		name     string
		button   TextButton
		wantText string
		wantType ButtonType
	}{
		{
			name:     "Basic TextButton",
			button:   TextButton{Text: "Click me"},
			wantText: "Click me",
			wantType: ButtonTypeText,
		},
		{
			name:     "Empty TextButton",
			button:   TextButton{Text: ""},
			wantText: "",
			wantType: ButtonTypeText,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.button.GetText(); got != tt.wantText {
				t.Errorf("TextButton.GetText() = %v, want %v", got, tt.wantText)
			}
			if got := tt.button.ButtonType(); got != tt.wantType {
				t.Errorf("TextButton.ButtonType() = %v, want %v", got, tt.wantType)
			}
		})
	}
}

func TestNewTextButton(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		wantText string
		wantType ButtonType
	}{
		{
			name:     "Basic NewTextButton",
			text:     "Click me",
			wantText: "Click me",
			wantType: ButtonTypeText,
		},
		{
			name:     "Empty text NewTextButton",
			text:     "",
			wantText: "",
			wantType: ButtonTypeText,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := NewTextButton(tt.text)

			if button.Text != tt.text {
				t.Errorf("NewTextButton().Text = %v, want %v", button.Text, tt.text)
			}

			if got := button.GetText(); got != tt.wantText {
				t.Errorf("TextButton.GetText() = %v, want %v", got, tt.wantText)
			}
			if got := button.ButtonType(); got != tt.wantType {
				t.Errorf("TextButton.ButtonType() = %v, want %v", got, tt.wantType)
			}
		})
	}
}

func TestDataButton(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		data     string
		wantText string
		wantType ButtonType
	}{
		{
			name:     "Basic DataButton",
			text:     "Click me",
			data:     "button_data",
			wantText: "Click me",
			wantType: ButtonTypeData,
		},
		{
			name:     "Empty text DataButton",
			text:     "",
			data:     "button_data",
			wantText: "",
			wantType: ButtonTypeData,
		},
		{
			name:     "Empty data DataButton",
			text:     "Click me",
			data:     "",
			wantText: "Click me",
			wantType: ButtonTypeData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := NewDataButton(tt.text, tt.data)

			if button.Text != tt.text {
				t.Errorf("NewDataButton().Text = %v, want %v", button.Text, tt.text)
			}
			if button.Data != tt.data {
				t.Errorf("NewDataButton().Data = %v, want %v", button.Data, tt.data)
			}

			if got := button.GetText(); got != tt.wantText {
				t.Errorf("DataButton.GetText() = %v, want %v", got, tt.wantText)
			}
			if got := button.ButtonType(); got != tt.wantType {
				t.Errorf("DataButton.ButtonType() = %v, want %v", got, tt.wantType)
			}
		})
	}
}

func TestUrlButton(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		url      string
		wantText string
		wantType ButtonType
	}{
		{
			name:     "Basic UrlButton",
			text:     "Visit website",
			url:      "https://example.com",
			wantText: "Visit website",
			wantType: ButtonTypeURL,
		},
		{
			name:     "Empty text UrlButton",
			text:     "",
			url:      "https://example.com",
			wantText: "",
			wantType: ButtonTypeURL,
		},
		{
			name:     "Empty URL UrlButton",
			text:     "Visit website",
			url:      "",
			wantText: "Visit website",
			wantType: ButtonTypeURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := NewUrlButton(tt.text, tt.url)

			if button.Text != tt.text {
				t.Errorf("NewUrlButton().Text = %v, want %v", button.Text, tt.text)
			}
			if button.URL != tt.url {
				t.Errorf("NewUrlButton().URL = %v, want %v", button.URL, tt.url)
			}

			if got := button.GetText(); got != tt.wantText {
				t.Errorf("UrlButton.GetText() = %v, want %v", got, tt.wantText)
			}
			if got := button.ButtonType(); got != tt.wantType {
				t.Errorf("UrlButton.ButtonType() = %v, want %v", got, tt.wantType)
			}
		})
	}
}

func TestSwitchInlineQueryButton(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		query    string
		wantText string
		wantType ButtonType
	}{
		{
			name:     "Basic SwitchInlineQueryButton",
			text:     "Search",
			query:    "query_string",
			wantText: "Search",
			wantType: ButtonTypeSwitchInlineQuery,
		},
		{
			name:     "Empty text SwitchInlineQueryButton",
			text:     "",
			query:    "query_string",
			wantText: "",
			wantType: ButtonTypeSwitchInlineQuery,
		},
		{
			name:     "Empty query SwitchInlineQueryButton",
			text:     "Search",
			query:    "",
			wantText: "Search",
			wantType: ButtonTypeSwitchInlineQuery,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := NewSwitchInlineQueryButton(tt.text, tt.query)

			if button.Text != tt.text {
				t.Errorf("NewSwitchInlineQueryButton().Text = %v, want %v", button.Text, tt.text)
			}
			if button.Query != tt.query {
				t.Errorf("NewSwitchInlineQueryButton().Query = %v, want %v", button.Query, tt.query)
			}

			if got := button.GetText(); got != tt.wantText {
				t.Errorf("SwitchInlineQueryButton.GetText() = %v, want %v", got, tt.wantText)
			}
			if got := button.ButtonType(); got != tt.wantType {
				t.Errorf("SwitchInlineQueryButton.ButtonType() = %v, want %v", got, tt.wantType)
			}
		})
	}
}

func TestSwitchInlineQueryCurrentChatButton(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		query    string
		wantText string
		wantType ButtonType
	}{
		{
			name:     "Basic SwitchInlineQueryCurrentChatButton",
			text:     "Search here",
			query:    "query_string",
			wantText: "Search here",
			wantType: ButtonTypeSwitchInlineQueryCurrentChat,
		},
		{
			name:     "Empty text SwitchInlineQueryCurrentChatButton",
			text:     "",
			query:    "query_string",
			wantText: "",
			wantType: ButtonTypeSwitchInlineQueryCurrentChat,
		},
		{
			name:     "Empty query SwitchInlineQueryCurrentChatButton",
			text:     "Search here",
			query:    "",
			wantText: "Search here",
			wantType: ButtonTypeSwitchInlineQueryCurrentChat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := NewSwitchInlineQueryCurrentChatButton(tt.text, tt.query)

			if button.Text != tt.text {
				t.Errorf("NewSwitchInlineQueryCurrentChatButton().Text = %v, want %v", button.Text, tt.text)
			}
			if button.Query != tt.query {
				t.Errorf("NewSwitchInlineQueryCurrentChatButton().Query = %v, want %v", button.Query, tt.query)
			}

			if got := button.GetText(); got != tt.wantText {
				t.Errorf("SwitchInlineQueryCurrentChatButton.GetText() = %v, want %v", got, tt.wantText)
			}
			if got := button.ButtonType(); got != tt.wantType {
				t.Errorf("SwitchInlineQueryCurrentChatButton.ButtonType() = %v, want %v", got, tt.wantType)
			}
		})
	}
}
