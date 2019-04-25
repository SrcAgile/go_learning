package string_prac

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringSplitCombine(t *testing.T) {
	a := "hello,world"
	split := strings.Split(a, ",")
	for _, val := range split {
		t.Log(val)
	}
	t.Logf("the type of a is %T", a)
	t.Log(strings.Join(split, "--"))
}

func TestStringTrans(t *testing.T) {
	s := "1000"
	if val, err := strconv.Atoi(s); err == nil {
		t.Log("err is", err)
		t.Log("val is", val)
	}

}
