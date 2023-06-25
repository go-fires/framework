package ints

type Int int

// InArray returns true if the int is in the array.
func (i Int) InArray(arr []int) bool {
	return InArray(int(i), arr)
}

// Val returns the int value.
func (i Int) Val() int {
	return int(i)
}

// Itoa converts an int to a string.
func (i Int) Itoa() string {
	return Itoa(int(i))
}

// String alias for Itoa.
func (i Int) String() string {
	return Itoa(int(i))
}

// Bytes return the bytes of the string.
func (i Int) Bytes() []byte {
	return []byte(i.String())
}

// Between returns true if the int is between min and max.
func (i Int) Between(min, max int) bool {
	return Between(int(i), min, max)
}
