package misc

import (
	"os"
	"testing"
)

func TestUTF8(t *testing.T) {
	f, err := os.Open("../data/utf8.txt")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = f.Close()
	}()
	b := make([]byte, 31)
	if n, err := f.Read(b); err != nil || n != 31 {
		t.Errorf("read bytes count: expected 31 got %d", n)
	}
	bom := UTFbom(b)
	if UTF_8 != bom {
		t.Errorf("UTF type error: expected UTF_8 got %s", bom)
	}
	txt := string(b[0:])

	if "∮ E⋅da = Q,  n → ∞, ∑" != txt {
		t.Errorf("UTF value error: expected \"∮ E⋅da = Q,  n → ∞, ∑\" got \"%s\"", txt)
	}
}

func TestUTF16BE(t *testing.T) {
	f, err := os.Open("../data/utf16be.txt")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = f.Close()
	}()
	b := make([]byte, 14)
	if n, err := f.Read(b); err != nil || n != 14 {
		t.Errorf("read bytes count: expected 12 got %d", n)
	}
	bom := UTFbom(b)
	if UTF_16BE != bom {
		t.Errorf("UTF type error: expected UTF_16BE got %s", bom)
	}
	txt, _ := BigEndianUTF16String(b[2:])

	if "0 HEAD" != txt {
		t.Errorf("UTF value error: expected \"0 HEAD\" got \"%s\"", txt)
	}
}

func TestUTF16LE(t *testing.T) {
	f, err := os.Open("../data/utf16le.txt")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = f.Close()
	}()
	b := make([]byte, 36)
	if n, err := f.Read(b); err != nil || n != 36 {
		t.Errorf("read bytes count: expected 36 got %d", n)
	}
	bom := UTFbom(b)
	if UTF_16LE != bom {
		t.Errorf("UTF type error: expected UTF_16LE got %s", bom)
	}
	txt, _ := LittleEndianUTF16String(b[2:])

	if "première is first" != txt {
		t.Errorf("UTF value error: expected \"première is first\" got \"%s\"", txt)
	}
}

func TestUTF32BE(t *testing.T) {
	f, err := os.Open("../data/ascii32be.txt.txt")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = f.Close()
	}()
	b := make([]byte, 44)
	if n, err := f.Read(b); err != nil || n != 44 {
		t.Errorf("read bytes count: expected 44 got %d", n)
	}
	bom := UTFbom(b)
	if UTF_32BE != bom {
		t.Errorf("UTF type error: expected UTF_32BE got %s", bom)
	}
	txt, _ := BigEndianUTF32String(b[4:])

	if "ascii file" != txt {
		t.Errorf("UTF value error: expected \"ascii file\" got \"%s\"", txt)
	}
}

func TestUTF32LE(t *testing.T) {
	f, err := os.Open("../data/ascii32le.txt.txt")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = f.Close()
	}()
	b := make([]byte, 44)
	if n, err := f.Read(b); err != nil || n != 44 {
		t.Errorf("read bytes count: expected 44 got %d", n)
	}
	bom := UTFbom(b)
	if UTF_32LE != bom {
		t.Errorf("UTF type error: expected UTF_32LE got %s", bom)
	}
	txt, _ := LittleEndianUTF32String(b[4:])

	if "ascii file" != txt {
		t.Errorf("UTF value error: expected \"ascii file\" got \"%s\"", txt)
	}
}
