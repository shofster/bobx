package misc

import (
	"fmt"
	"testing"
)

func TestSoundex(t *testing.T) {
	var tests = []struct {
		word   string
		result string
	}{
		{"", "000"},
		{"ab", "A100"},
		{"aaaaaaaaaaeeeeiiiiooooouuuu123456789", "A000"},
		{"Iowa", "I000"},
		{"texas", "T220"},
		{"TEXAS   ", "T220"},
		{"Villa Grove, Colorado 77123", "V426"},
		{"Vla Grov, Colodaro", "V426"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Word: %s ", tt.word), func(t *testing.T) {
			word := SoundexCode(tt.word)
			if word != tt.result {
				t.Errorf("TestSoundex: expected %s, got %s", tt.result, word)
			}
		})
	}

}
