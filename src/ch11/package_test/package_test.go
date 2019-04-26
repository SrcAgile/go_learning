package testpkg

import (
	packageName "ch11/pkg1"
	"fmt"
	"testing"
)

//不用管里面的包叫什么名字，直接就能换成新名字

//the import "ch11/pkg1" command means finding package in the file which belong
//to this direcotory
//即寻找当前目录下所有go文件中的package文件，go文件叫什么名字都无所谓
//而且同一目录下只能存在一个package，ch11/pkg1下的所有go文件中的package要一致
func init() {
	fmt.Println("the testpkg is loading...")
}

func TestPkg(t *testing.T) {
	packageName.Spreading()
}
