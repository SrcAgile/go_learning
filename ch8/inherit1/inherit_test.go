package inherit_test

import (
	"fmt"
	"testing"
)

type pet struct {
}

func (p *pet) Speak() {
	fmt.Println("...")
}
func (p *pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *pet
}

//extends a class by combine
func (d *Dog) Speak() {
	//fmt.Println("....")
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	// p.Speak()
	// fmt.Println(" ",host)
	d.p.SpeakTo(host)
}

func TestInherit(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("wangwangwang")
}
