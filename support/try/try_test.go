package try

import (
	"fmt"
	"testing"
)

func TestTry_Catch(t *testing.T) {
	NewTry(func() {
		panic("foo")
	}).Catch(func(err interface{}) {
		fmt.Println(err)
	}).Finally()
}
