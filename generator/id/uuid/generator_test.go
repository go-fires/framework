package uuid

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	fmt.Println(Generate())
	fmt.Println(Generate())
	fmt.Println(Generate())
}
