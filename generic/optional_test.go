package generic

import "testing"

func TestOptional(t *testing.T) {

	oe := Empty[int]()
	t.Logf("empty: %t %#v", oe.IsPresent(), oe)
	if oe.IsPresent() {
		t.Errorf("  incorrect empty: %t", oe.IsPresent())
	}
	ve, pe := oe.Get()
	t.Logf("empty get: %t %d", !pe, ve)
	if ve != 0 {
		t.Errorf("  ** incorrect empty 0 value: %d", ve)
	}
	ve = oe.OrElse(-1234)
	t.Logf("empty get: default %d", ve)
	if ve != -1234 {
		t.Errorf("  ** incorrect OrElse -1234 value: %d", ve)
	}
	op := Present[int](3)
	vp, pv := op.Get()
	t.Logf("present get: %t %d", pv, vp)
}
