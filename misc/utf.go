package misc

import (
	"bytes"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
	"golang.org/x/text/transform"
	"io"
)

/*

  File:    endian.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: convert endian (little or big) to GO int

Encoding	 Representation (hexadecimal)
UTF-8        EF BB BF
UTF-16 (BE)  FE FF
UTF-16 (LE)  FF FE
UTF-32 (BE)  00 00 FE FF
UTF-32 (LE)  FF FE 00 00

"od -t x1 <file-name> | less" to see individual hex bytes

*/

type UTFtype int

const (
	UTF_8 = iota
	UTF_8BOM
	UTF_16LE
	UTF_16BE
	UTF_32LE
	UTF_32BE
)

func (u UTFtype) String() string {
	switch u {
	case UTF_8:
		return "UTF_8"
	case UTF_8BOM:
		return "UTF_8BOM"
	case UTF_16LE:
		return "UTF_16LE"
	case UTF_16BE:
		return "UTF_16BE"
	case UTF_32LE:
		return "UTF_32LE"
	case UTF_32BE:
		return "UTF_32BE"
	}
	return "UTF_UNKNOWN"
}

// UTFbom - decodes the first bytes for BOMmv
func UTFbom(s []byte) UTFtype {
	switch {
	case len(s) > 3 && bytes.Equal(s[:4], []byte{0xFF, 0xFE, 0x00, 0x00}):
		return UTF_32LE
	case len(s) > 3 && bytes.Equal(s[:4], []byte{0x00, 0x00, 0xFE, 0xFF}):
		return UTF_32BE
	case len(s) > 2 && bytes.Equal(s[:3], []byte{0xEF, 0xBB, 0xBF}):
		return UTF_8BOM
	case len(s) > 1 && bytes.Equal(s[:2], []byte{0xFF, 0xFE}):
		return UTF_16LE
	case len(s) > 1 && bytes.Equal(s[:2], []byte{0xFE, 0xFF}):
		return UTF_16BE
	default:
		return UTF_8
	}
}

func LittleEndianUTF32String(utf []byte) (string, error) {
	// Create a UTF-32 decoder
	decoder := utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM).NewDecoder()
	// Decode the UTF-32 bytes to UTF-8
	utf8Data, err := io.ReadAll(transform.NewReader(bytes.NewReader(utf), decoder))
	if err != nil {
		return "unable to UTF32toUTF8", err
	}
	return string(utf8Data), nil
}

func BigEndianUTF32String(utf []byte) (string, error) {
	// Create a UTF-32 decoder
	decoder := utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM).NewDecoder()
	// Decode the UTF-32 bytes to UTF-8
	utf8Data, err := io.ReadAll(transform.NewReader(bytes.NewReader(utf), decoder))
	if err != nil {
		return "unable to UTF32toUTF8", err
	}
	return string(utf8Data), nil
}

func LittleEndianUTF16String(utf []byte) (string, error) {
	// Create a UTF-16 decoder
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	// Decode the UTF-16 bytes to UTF-8
	utf8Data, err := io.ReadAll(transform.NewReader(bytes.NewReader(utf), decoder))
	if err != nil {
		return "unable to UTF16toUTF8", err
	}
	return string(utf8Data), nil
}

func BigEndianUTF16String(utf []byte) (string, error) {
	// Create a UTF-16 decoder
	decoder := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder()
	// Decode the UTF-16 bytes to UTF-8
	utf8Data, err := io.ReadAll(transform.NewReader(bytes.NewReader(utf), decoder))
	if err != nil {
		return "unable to UTF16toUTF8", err
	}
	return string(utf8Data), nil
}
