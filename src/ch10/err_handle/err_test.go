// package errors
//
// // New returns an error that formats as the given text.
//
// // New 返回一个给定文本格式的错误。
// func New(text string) error {
//     return &errorString{text}
// }
//
// // errorString is a trivial implementation of error.
//
// // errorString 是 error 的一个琐碎的实现。
// type errorString struct {
//     s string
// }
//
// func (e *errorString) Error() string {
//     return e.s
// }
package testerr

import (
	"errors"
	"testing"
)

func checkEffect(n int) (int, error) {
	if n > 100 {
		return n, errors.New("n is too large")
	}
	if n < 10 {
		return n, errors.New("n is too small")
	}

	return n, nil
}

func TestErr(t *testing.T) {
	if _, err := checkEffect(1); err != nil {
		t.Error(err)
	}

}
