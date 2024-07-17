package strings

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

/*

  File:    pretty.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: Various string & string list functions
*/

// PrettyStringSize makes a fixed length string
// if too long, truncates
// too short adds spaces to end
func PrettyStringSize(in string, size int) string {
	in = strings.TrimSpace(in)
	switch {
	case len(in) < size:
		return fmt.Sprintf(fmt.Sprintf("%%-%ds", size), in)
	case len(in) > size:
		return in[:size]
	}
	return in
}

// PrettyLines breaks a string into lines of maxs length
func PrettyLines(s string, maxs int) []string {
	var lineMax = func(i int) int {
		if i < maxs {
			return i
		}
		return maxs
	}
	var lines []string
	for len(s) > 0 {
		size := lineMax(len(s))
		line := s[0:size]
		lines = append(lines, line)
		s = s[size:] // [0:] gives empty slice
	}
	return lines
}

// PrettyMemoryInfo generates a ?meaningful string of memory use.
func PrettyMemoryInfo() string {
	var stats runtime.MemStats
	var mb uint64 = 1024 * 1024
	runtime.ReadMemStats(&stats)
	s := fmt.Sprintf("Mallocs: %d, Frees: %d, Live: %d", stats.Mallocs, stats.Frees,
		stats.Mallocs-stats.Frees)
	// Sys is the total bytes of memory obtained from the OS.
	// HeapAlloc is bytes of allocated heap objects.
	s += fmt.Sprintf(", Memory- sys: %d, heap: %d (MB)", stats.Sys/mb, stats.HeapAlloc/mb)
	return s
}

// PrettyFileInfo builds a formatted string of a FileInfo type
func PrettyFileInfo(fi os.FileInfo, dtFormat string) string {
	name := fi.Name()
	dt := fi.ModTime().Format(dtFormat)
	m := fmt.Sprintf("%s", fi.Mode())
	f := "%4s %10d %14s %s"
	s := fmt.Sprintf(f, m[0:4], fi.Size(), dt, name)
	return s
}

var RFCDateFormatTypes = []string{
	"Default", "UNIX", "RFC822", "RFC822Z", "RFC1123",
	"RFC1123Z", "RFC3339"}

//goland:noinspection ALL
func DtFormat(dtFormat string) (dfmt string) {
	switch strings.ToUpper(dtFormat) {
	case "UNIX":
		dfmt = time.UnixDate
	case "RFC822":
		dfmt = time.RFC822
	case "RFC822Z":
		dfmt = time.RFC822Z
	case "RFC1123":
		dfmt = time.RFC1123
	case "RFC1123Z":
		dfmt = time.RFC1123Z
	case "RFC3339":
		dfmt = time.RFC3339
	default:
		dfmt = "01/02/06 15:04"
	}
	return
}

// PrettyUint builds a string  from a uint4. put a "sep" (probably a comma) between 3 digits.
func PrettyUint(n uint64, sep string) string {
	if n == 0 {
		return "0"
	}
	thousands := make(StringList, 0)
	for n > 0 {
		q := n / 1000
		r := n - q*1000
		d := ""
		if q == 0 {
			d = fmt.Sprintf("%d", r)
		} else {
			d = fmt.Sprintf("%03d", r)
		}
		thousands = Prepend(thousands, d, 0)
		n = q
	}
	s := strings.Join(thousands, sep)
	return s
}

// Roman generates a string of roman numerals - up to 3999.
// 4000 and greater would require unicode characters 216x thru 218x.
func Roman(num int) (string, error) {
	var thousands = [...]string{"m", "mm", "mmm"}
	var hundreds = [...]string{"c", "cc", "ccc", "cd", "d", "dc", "dcc", "dccc", "cm"}
	var tens = [...]string{"x", "xx", "xxx", "xl", "l", "lx", "lxx", "lxxx", "xc"}
	var ones = [...]string{"i", "ii", "iii", "iv", "v", "vi", "vii", "viii", "ix"}
	if num > 3999 {
		return fmt.Sprintf("*%d*", num), errors.New("Invalid Roman")
	}
	s := ""
	k := num / 1000
	h := num % 1000 / 100
	t := num % 100 / 10
	o := num % 10
	//	log.Printf("     %d %d %d %d\n", k, h, t, o)
	if k > 0 {
		s += thousands[k-1]
	}
	if h > 0 {
		s += hundreds[h-1]
	}
	if t > 0 {
		s += tens[t-1]
	}
	if o > 0 {
		s += ones[o-1]
	}
	return s, nil
}

/*
Table of Roman numerals in Unicode:
Value 	1	2	3	4	5	6	7	8	9	10	11	12	50	100	500	1,000
U+216x	Ⅰ	Ⅱ	Ⅲ	Ⅳ	Ⅴ	Ⅵ	Ⅶ	Ⅷ	Ⅸ	Ⅹ	Ⅺ	Ⅻ	Ⅼ	 Ⅽ	 Ⅾ	  Ⅿ
U+217x	ⅰ	ⅱ	ⅲ	ⅳ	ⅴ	ⅵ	ⅶ	ⅷ	ⅸ	ⅹ	ⅺ	ⅻ	ⅼ	 ⅽ	 ⅾ	  ⅿ

2180 ↀ ROMAN NUMERAL ONE THOUSAND
2181 ↁ ROMAN NUMERAL FIVE THOUSAND
2182 ↂ ROMAN NUMERAL TEN THOUSAND
2183 Ↄ ROMAN NUMERAL REVERSED ONE HUNDRED

2187 ↇ ROMAN NUMERAL FIFTY THOUSAND
2188 ↈ ROMAN NUMERAL ONE HUNDRED THOUSAND

1100    MC
9,999   IXCMXCIX      Bar over IX
9,997	IXCMXCVII
9,998	IXCMXCVIII
10,000	X             Bar over X
10,001	XI
*/
