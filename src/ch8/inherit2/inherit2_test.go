package inherit2_test

import (
	"fmt"
	"testing"
)

type human struct {
}

func (h *human) selfcheck() {
	fmt.Println("human self checking...")
}

func (h *human) attr() {
	h.selfcheck()
	fmt.Println("i am belong to Human")
}

type man struct {
	human
}

func (m *man) selfcheck() {
	fmt.Println("man self checking...")
}

func TestInherit(t *testing.T) {
	var m *man = new(man)
	m.attr()
}
