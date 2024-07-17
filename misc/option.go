package misc

import "fmt"

/*

  File:    option.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: methods for an "optional" value. pre generic.
*/

type Option struct {
	present bool
	value   interface{} // - default is nil
}

func NewOption(value interface{}) *Option {
	return &Option{present: true, value: value}
}

func (o *Option) IsPresent() bool {
	return o.present
}

func (o *Option) Set(value interface{}) {
	o.value = value
	o.present = true
}

func (o *Option) Get() (interface{}, bool) {
	return o.value, o.present
}

func (o *Option) OrElse(value interface{}) (interface{}, bool) {
	if o.present {
		return o.value, o.present
	}
	return value, false
}

func (o *Option) String() string {
	return fmt.Sprintf("%t, %T : %#v", o.present, o.value, o.value)
}
