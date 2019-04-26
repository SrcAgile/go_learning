package interPracPackage

import (
	"fmt"
	"testing"
)

type code string

type programmer interface {
	writeHelloWorld() code
}

type goProgrammer struct {
}

func (goer *goProgrammer) writeHelloWorld() code {
	return "fmt.println helloworld"
}

type javaProgrammer struct {
}

func (jaer *javaProgrammer) writeHelloWorld() code {
	return "system.out.println helloworld "
}

func sayHelloWorld(p programmer) {
	fmt.Printf("the type is %T\n", p)
	fmt.Println(p.writeHelloWorld())
}

func TestInterface(t *testing.T) {
	goer := goProgrammer{}
	javer := javaProgrammer{}
	sayHelloWorld(&goer)
	sayHelloWorld(&javer)
}
