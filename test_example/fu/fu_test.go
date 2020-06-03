package fu

import (
	"testing"
)

func TestFoo(t *testing.T) {
	got := Foo()
	if got != 1 {
		t.Error("Abs(-1); want 1", got)
	}
}
