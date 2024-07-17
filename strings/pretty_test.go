package strings

import (
	"fmt"
	"os"
	"testing"
)

const dataFile = "../data/go.mod"

func TestPrettySize(t *testing.T) {
	var sizeTests = []struct {
		in     string
		size   int
		result string
	}{
		{"", 4, "    "},
		{"1234", 4, "1234"},
		{"12345", 4, "1234"},
	}
	for _, tt := range sizeTests {
		t.Run(fmt.Sprintf("StringSize : %d ", tt.size), func(t *testing.T) {
			s := PrettyStringSize(tt.in, tt.size)
			if s != tt.result {
				t.Errorf("Expected %s , got %s", tt.result, s)
			}
		})
	}
	var lineTests = []struct {
		in    string
		size  int
		count int
	}{
		{"123456789", 3, 3},
		{"1234", 3, 2},
		{"1", 3, 1},
	}
	for _, tt := range lineTests {
		t.Run(fmt.Sprintf("Line Size : %d ", tt.size), func(t *testing.T) {
			lines := PrettyLines(tt.in, tt.size)
			if len(lines) != tt.count {
				t.Errorf("Expected %d , got %d", tt.count, len(lines))
			}
		})
	}
}

func TestPretty(t *testing.T) {
	mi := PrettyMemoryInfo()
	t.Logf("PrettyMemoryInfo:  %s", mi)

	fi, err := os.Stat(dataFile)
	if err != nil {
		t.Errorf("PrettyFileInfo: unable to stat file %s", dataFile)
	} else {
		p := PrettyFileInfo(fi, DtFormat(RFCDateFormatTypes[0]))
		t.Logf("PrettyFileInfo: %s", p)
	}

	p := PrettyUint(123456789, ",")
	t.Logf("PrettyUint: got %s", p)
	if p != "123,456,789" {
		t.Errorf("PrettyUint: expected 213,456,789, got %s", p)
	}
}

func TestDTformat(t *testing.T) {
	var tests = []struct {
		name   string
		result string
	}{
		{"Default", "01/02/06 15:04"},
		{"UNIX", "Mon Jan _2 15:04:05 MST 2006"},
		{"RFC822", "02 Jan 06 15:04 MST"},
		{"RFC1123Z", "Mon, 02 Jan 2006 15:04:05 -0700"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("RFC Format: %s ", tt.name), func(t *testing.T) {
			s := DtFormat(tt.name)
			if s != tt.result {
				t.Errorf("Expected %s , got %s", tt.result, s)
			} else {
				t.Logf("GO Formatting string \"%s\"", s)
				fi, err := os.Stat(dataFile)
				if err != nil {
					t.Errorf("TestDTformat: unable to stat file %s", dataFile)
				} else {
					p := PrettyFileInfo(fi, s)
					t.Logf("PrettyFileInfo: %s", p)
				}
			}
		})
	}

}

/*
Roman numerals from 1 to 30 are
1 = I, 2 = II, 3 = III, 4 = IV, 5 = V,
6 = VI, 7 = VII, 8 = VIII, 9 = IX, 10 = X,
11 = XI, 12 = XII, 13 = XIII, 14 = XIV, 15 = XV,
16 = XVI, 17 = XVII, 18 = XVIII, 19 = XIX, 20 = XX,
21 = XXI
*/

func TestRoman(t *testing.T) {
	var tests = []struct {
		num    int
		result string
	}{
		{6, "vi"},
		{7, "vii"},
		{8, "viii"},
		{9, "ix"},
		{10, "x"},
		{11, "xi"},
		{15, "xv"},
		{16, "xvi"},
		{17, "xvii"},
		{18, "xviii"},
		{19, "xix"},
		{20, "xx"},
		{21, "xxi"},
		{99, "xcix"},
		{1579, "mdlxxix"},
		{2024, "mmxxiv"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Number: %d ", tt.num), func(t *testing.T) {
			s, err := Roman(tt.num)
			if err != nil {
				t.Errorf("TestRoman: error %s", err.Error())
			} else {
				// t.Logf("TestRoman code %d, got %s", tt.num, s)
				if s != tt.result {
					t.Errorf("Expected %s , got %s", tt.result, s)
				}
			}
		})
	}

}
