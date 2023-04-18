package ints

import (
	"github.com/go-fires/framework/support/strs"
	"math/rand"
	"strconv"
	"strings"
)

// InArray returns true if val is in arr.
func InArray(val int, arr []int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

// Contains returns true if val is in arr.
func Contains(arr []int, val int) bool {
	return InArray(val, arr)
}

// Itoa converts an int to a string.
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// Split splits a string into an array of ints.
func Split(s, sep string) []int {
	split := strings.Split(s, sep)
	arr := make([]int, len(split))

	for i, v := range split {
		arr[i] = strs.Atoi(v).Val()
	}

	return arr
}

// Random returns a random int between min and max.
func Random(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandomString returns a random string of length.
func RandomString(length int) string {
	var letters = []rune("0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// Range returns an array of ints between min and max.
func Range(min, max int) []int {
	arr := make([]int, max-min+1)

	for i := min; i <= max; i++ {
		arr[i-min] = i
	}

	return arr
}

// Min returns the minimum of ints.
func Min(v ...int) int {
	min := v[0]

	for _, v := range v {
		if v < min {
			min = v
		}
	}

	return min
}

// Max returns the maximum of ints.
func Max(v ...int) int {
	max := v[0]

	for _, v := range v {
		if v > max {
			max = v
		}
	}

	return max
}

// Sum returns the sum of ints.
func Sum(v ...int) int {
	sum := 0

	for _, v := range v {
		sum += v
	}

	return sum
}

// SumSlice returns the sum of a slice of ints.
func SumSlice(v []int) int {
	return Sum(v...)
}

// Between returns true if val is between min and max.
func Between(val, min, max int) bool {
	min, max = Min(min, max), Max(min, max)

	return val >= min && val <= max
}

// Sort alias for SortAsc.
func Sort(slice []int) []int {
	return SortAsc(slice)
}

// SortAsc sorts a slice of ints in ascending order.
func SortAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

// SortDesc sorts a slice of ints in descending order.
func SortDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}
