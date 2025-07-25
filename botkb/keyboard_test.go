package botkb

import (
	"reflect"
	"testing"
)

func TestNewOnetimeKeyboard(t *testing.T) {
	tests := []struct {
		name    string
		buttons [][]Button
		want    *MessageKeyboard
	}{
		{
			name:    "Empty onetime keyboard",
			buttons: [][]Button{},
			want: &MessageKeyboard{
				kbType:    KeyboardTypeBottom,
				isOneTime: true,
				Buttons:   [][]Button{},
			},
		},
		{
			name: "Onetime keyboard with one row of data buttons",
			buttons: [][]Button{
				{
					&DataButton{Text: "Button 1", Data: "data1"},
					&DataButton{Text: "Button 2", Data: "data2"},
				},
			},
			want: &MessageKeyboard{
				kbType:    KeyboardTypeBottom,
				isOneTime: true,
				Buttons: [][]Button{
					{
						&DataButton{Text: "Button 1", Data: "data1"},
						&DataButton{Text: "Button 2", Data: "data2"},
					},
				},
			},
		},
		{
			name: "Onetime keyboard with multiple rows of mixed buttons",
			buttons: [][]Button{
				{
					&DataButton{Text: "Button 1", Data: "data1"},
					&UrlButton{Text: "URL 1", URL: "https://example.com"},
				},
				{
					&DataButton{Text: "Button 2", Data: "data2"},
				},
			},
			want: &MessageKeyboard{
				kbType:    KeyboardTypeBottom,
				isOneTime: true,
				Buttons: [][]Button{
					{
						&DataButton{Text: "Button 1", Data: "data1"},
						&UrlButton{Text: "URL 1", URL: "https://example.com"},
					},
					{
						&DataButton{Text: "Button 2", Data: "data2"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOnetimeKeyboard(tt.buttons...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOnetimeKeyboard() = %v, want %v", got, tt.want)
			}

			// Test the KeyboardType method
			if got.KeyboardType() != KeyboardTypeBottom {
				t.Errorf("MessageKeyboard.KeyboardType() = %v, want %v", got.KeyboardType(), KeyboardTypeBottom)
			}

			// Test the IsOneTime method
			if !got.IsOneTime() {
				t.Errorf("MessageKeyboard.IsOneTime() = %v, want true", got.IsOneTime())
			}
		})
	}
}

func TestNewMessageKeyboard(t *testing.T) {
	tests := []struct {
		name    string
		kbType  KeyboardType
		buttons [][]Button
		want    *MessageKeyboard
	}{
		{
			name:    "Empty keyboard with KeyboardTypeNone",
			kbType:  KeyboardTypeNone,
			buttons: [][]Button{},
			want: &MessageKeyboard{
				kbType:  KeyboardTypeNone,
				Buttons: [][]Button{},
			},
		},
		{
			name:    "Empty keyboard with KeyboardTypeInline",
			kbType:  KeyboardTypeInline,
			buttons: [][]Button{},
			want: &MessageKeyboard{
				kbType:  KeyboardTypeInline,
				Buttons: [][]Button{},
			},
		},
		{
			name:   "Keyboard with one row of data buttons",
			kbType: KeyboardTypeBottom,
			buttons: [][]Button{
				{
					&DataButton{Text: "Button 1", Data: "data1"},
					&DataButton{Text: "Button 2", Data: "data2"},
				},
			},
			want: &MessageKeyboard{
				kbType: KeyboardTypeBottom,
				Buttons: [][]Button{
					{
						&DataButton{Text: "Button 1", Data: "data1"},
						&DataButton{Text: "Button 2", Data: "data2"},
					},
				},
			},
		},
		{
			name:   "Keyboard with multiple rows of mixed buttons",
			kbType: KeyboardTypeInline,
			buttons: [][]Button{
				{
					&DataButton{Text: "Button 1", Data: "data1"},
					&UrlButton{Text: "URL 1", URL: "https://example.com"},
				},
				{
					&DataButton{Text: "Button 2", Data: "data2"},
				},
			},
			want: &MessageKeyboard{
				kbType: KeyboardTypeInline,
				Buttons: [][]Button{
					{
						&DataButton{Text: "Button 1", Data: "data1"},
						&UrlButton{Text: "URL 1", URL: "https://example.com"},
					},
					{
						&DataButton{Text: "Button 2", Data: "data2"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMessageKeyboard(tt.kbType, tt.buttons...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessageKeyboard() = %v, want %v", got, tt.want)
			}

			// Also test the KeyboardType method
			if got.KeyboardType() != tt.kbType {
				t.Errorf("MessageKeyboard.KeyboardType() = %v, want %v", got.KeyboardType(), tt.kbType)
			}
		})
	}
}
