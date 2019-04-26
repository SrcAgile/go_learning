package mask

import (
	"fmt"
	"testing"
)

func TestMask(t *testing.T) {
	const (
		Readable = 1 << iota
		Writeable
		Executable
	)
	//create mask
	mask := Readable | Executable
	fmt.Println("readable:", (mask&Readable) == Readable)
	fmt.Println("executable:", (mask&Executable) == Executable)
	fmt.Println("writeable:", (mask&Writeable) == Writeable)

	mask = mask &^ Executable //clean mask
	fmt.Println("executable:", (mask&Executable) == Executable)
}
