package generic

/*

  File:    optional.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
    Description: methods for an "optional" value.
  Helpful for SQL or JSON situations to handle missing values.
*/

// Optional - wrapper for any (interface{}
type Optional[T any] struct {
	present bool
	value   T
}

// Present - create an Optional with value.
func Present[T any](value T) Optional[T] {
	return Optional[T]{present: true, value: value}
}

// Empty - create an Optional with NO value
func Empty[T any]() Optional[T] {
	return Optional[T]{present: false}
}

// IsPresent - true is a value is set
func (o *Optional[T]) IsPresent() bool {
	return o.present
}

// Get the current value and true if present
// not present gives zero value and false
func (o *Optional[T]) Get() (T, bool) {
	return o.value, o.present
}

// OrElse gives the current value if present
//
//	or gives the supplied default value
func (o *Optional[T]) OrElse(value T) T {
	if o.present {
		return o.value
	}
	return value
}
