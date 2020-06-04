package fu

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	fmt.Println(Foo())
	// Output: 22
}

func ExampleFoo(t *testing.T) {
	fmt.Println(Foo())
	// Output: 22
}
