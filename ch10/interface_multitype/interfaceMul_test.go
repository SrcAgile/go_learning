package empty_interface

import (
	"fmt"
	"testing"
)

//switch v:= in.type
func DoSomething(in interface{}) {
	switch v := in.(type) {
	case int:
		fmt.Println("i am int and val is", v)
	case string:
		fmt.Println("i am strign and val is", v)
	default:
		fmt.Println("i donot know who i am")
	}

}

func TestEmptyInterface(t *testing.T) {
	var a int
	a = 10
	DoSomething(a)

	var b string
	b = "HelloWorld"
	DoSomething(b)
}
