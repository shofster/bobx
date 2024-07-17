package misc

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	var tests = []struct {
		header   string
		password string
		text     string
		length   int
		xtra     string
	}{
		{"message for Fred", "password1", "now is the time for all good men", 60, ""},
		{"message for Sam", "password2", "now is the time for all good men", 60, "xx"},
	}

	for _, tt := range tests {
		t.Logf("Test: %s / %s", tt.header, tt.password)
		digest, err := EncryptAEAD([]byte(tt.text), CreateEncryptionKey([]byte(tt.password)),
			[]byte(tt.header))
		if err != nil {
			t.Error("* test error", err)
			continue
		}
		if len(digest) != tt.length {
			t.Error("* Encrypt error", len(digest), tt.length)
			continue
		}
		if tt.xtra == "" {
			text, err := DecryptAEAD(digest, CreateEncryptionKey([]byte(tt.password)),
				[]byte(tt.header))
			if err != nil {
				t.Error("* Decrypt error", err)
				continue
			}
			if string(text) != tt.text {
				t.Error("* Encrypt error", string(text), tt.text)
				continue
			}
		} else {
			hd := tt.header + tt.xtra // force incorrect authentication
			_, err := DecryptAEAD(digest, CreateEncryptionKey([]byte(tt.password)),
				[]byte(hd))
			if err == nil {
				t.Error("* Decrypt  should fail error", err)
				continue
			}
		}

	}
}
