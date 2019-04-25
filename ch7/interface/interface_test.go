package inter_test

import (
	"fmt"
	"testing"
)

type say interface {
	sayhello(int) int
}

type chinese struct {
	language string
}
type foreign struct {
	language string
}

func (ch *chinese) sayhello(n int) int {
	fmt.Println("你好")
	return 0
}

func (en *foreign) sayhello(n int) int {
	fmt.Println("'s up'")
	return 0
}

// func someoneSayhello(sayInterface *say,n int){
//   sayInterface.sayhello(n)
// }

func TestInterface(t *testing.T) {
	// var interCsay say := new(chinese)
	// var interFsay say := new(foreign)
	cn := chinese{language: "zh"}
	en := foreign{language: "en"}
	var interCsay say
	interCsay = &cn
	var interFsay say
	interFsay = &en

	cn.sayhello(1)
	interCsay.sayhello(1)
	interFsay.sayhello(1)

}
