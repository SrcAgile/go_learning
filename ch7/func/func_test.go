package func_test

import (
	"fmt"
	"testing"
	"time"
)

type objfunc func(args1 []int) int

func timeSpend(inner objfunc) objfunc {
	return func(args1 []int) int {
		start := time.Now()
		ret := inner(args1)
		fmt.Println("Time spend: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(args []int) int {
	time.Sleep(time.Second * 1)
	return args[1]
}
func TestFunc(t *testing.T) {
	newfunc := timeSpend(slowFunc)
	newfunc([]int{1, 2, 3, 4})
}
