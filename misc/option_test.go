package misc

import "testing"

func TestOption(t *testing.T) {

	o := &Option{}
	if o.IsPresent() {
		t.Errorf("Empty, but reported NOT Empty")
	}
	o.Set("abc")
	if !o.IsPresent() {
		t.Errorf("Not Empty, but reported Empty")
	}
	t.Logf("abc: %s", o)
	v, _ := o.Get()
	if v != "abc" {
		t.Errorf("Get, expected \"abc\", got %#v", v)
	}

	var names []string = []string{
		"ralph",
		"sam",
		"sue",
	}
	on := NewOption(names)
	gnames, _ := on.Get()
	t.Logf("names: %#v", gnames)

}
