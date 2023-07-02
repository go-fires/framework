package generator

import (
	"fmt"
	"testing"
)

func TestUUIDGenerator(t *testing.T) {
	fmt.Println(UUID.Generate())
	fmt.Println(UUID.Generate())
	fmt.Println(UUID.Generate())
}

func TestRandomGenerator(t *testing.T) {
	fmt.Println(Random.Generate())
	fmt.Println(Random.Generate())
	fmt.Println(Random.Generate())
}
