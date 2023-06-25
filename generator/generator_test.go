package generator

import (
	"fmt"
	"testing"
)

func TestUUIDGenerator(t *testing.T) {
	fmt.Println(UUIDGenerator.Generate())
	fmt.Println(UUIDGenerator.Generate())
	fmt.Println(UUIDGenerator.Generate())
}

func TestRandomGenerator(t *testing.T) {
	fmt.Println(RandomGenerator.Generate())
	fmt.Println(RandomGenerator.Generate())
	fmt.Println(RandomGenerator.Generate())
}
