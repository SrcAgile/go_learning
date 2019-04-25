package func2_prac

import (
	"fmt"
	"testing"
)

func deferTest() {
	defer func() {
		fmt.Println("screw it.i am lazy guy")
	}()
	fmt.Println("---welcome---")

	panic("human err")

}

func TestFunc2(t *testing.T) {
	deferTest()
}
