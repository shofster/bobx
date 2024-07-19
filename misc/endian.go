package misc

/*

  File:    endian.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: convert 2, 4, 8 byte positive integer to int64.
*/

// LittleEndianToUint64 - build Int from slice bytes in Intel order.
func LittleEndianToInt(bytes []uint8) int64 {
	s := 0
	var v int64
	for i := 0; i < len(bytes); i++ {
		v |= int(bytes[i]) << s
		s += 8
	}
	return v
}

// BigEndianToUint64 - build Int from slice bytes in Network (Motorola) order.
func BigEndianToInt(bytes []uint8) int64 {
	var v int64
	for i := 0; i < len(bytes); i++ {
		v <<= 8
		v |= int64(bytes[i])
	}
	return v
}
