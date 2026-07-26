package botkb

import (
	"encoding/json"
	"testing"
)

func TestNewCopyTextButton(t *testing.T) {
	button := NewCopyTextButton("Copy invite", "https://t.me/SneatBot?start=pref_123")

	if button.ButtonType() != ButtonTypeCopyText {
		t.Fatalf("ButtonType() = %v, want %v", button.ButtonType(), ButtonTypeCopyText)
	}
	if button.GetText() != "Copy invite" {
		t.Fatalf("GetText() = %q, want %q", button.GetText(), "Copy invite")
	}
	if button.CopyText != "https://t.me/SneatBot?start=pref_123" {
		t.Fatalf("CopyText = %q", button.CopyText)
	}
}

func TestInlineButtonAppearance(t *testing.T) {
	button := NewDataButton("Start", "start")
	button.Style = ButtonStyleSuccess
	button.IconCustomEmojiID = "emoji-123"

	var provider InlineButtonAppearanceProvider = button
	appearance := provider.GetInlineButtonAppearance()
	if appearance.Style != ButtonStyleSuccess {
		t.Errorf("Style = %q, want %q", appearance.Style, ButtonStyleSuccess)
	}
	if appearance.IconCustomEmojiID != "emoji-123" {
		t.Errorf("IconCustomEmojiID = %q, want %q", appearance.IconCustomEmojiID, "emoji-123")
	}
}

func TestInlineButtonAppearanceJSONIsBackwardCompatible(t *testing.T) {
	button := NewUrlButton("Open", "https://example.com")
	data, err := json.Marshal(button)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(data), `{"text":"Open","url":"https://example.com"}`; got != want {
		t.Fatalf("JSON = %s, want %s", got, want)
	}

	button.Style = ButtonStylePrimary
	data, err = json.Marshal(button)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(data), `{"style":"primary","text":"Open","url":"https://example.com"}`; got != want {
		t.Fatalf("styled JSON = %s, want %s", got, want)
	}
}

func TestLegacyButtonStyleNamesUseCurrentTelegramWireValues(t *testing.T) {
	if ButtonStylePositive != ButtonStyleSuccess {
		t.Errorf("ButtonStylePositive = %q, want %q", ButtonStylePositive, ButtonStyleSuccess)
	}
	if ButtonStyleDestructive != ButtonStyleDanger {
		t.Errorf("ButtonStyleDestructive = %q, want %q", ButtonStyleDestructive, ButtonStyleDanger)
	}
	if ButtonStyleDefault != "" {
		t.Errorf("ButtonStyleDefault = %q, want empty", ButtonStyleDefault)
	}
}

func TestReplyTextButtonSupportsAppearance(t *testing.T) {
	button := NewTextButton("Play")
	button.Style = ButtonStylePrimary
	button.IconCustomEmojiID = "emoji-play"

	var provider InlineButtonAppearanceProvider = button
	appearance := provider.GetInlineButtonAppearance()
	if appearance.Style != ButtonStylePrimary {
		t.Errorf("Style = %q, want %q", appearance.Style, ButtonStylePrimary)
	}
	if appearance.IconCustomEmojiID != "emoji-play" {
		t.Errorf("IconCustomEmojiID = %q, want %q", appearance.IconCustomEmojiID, "emoji-play")
	}
}
