package recovertest

import (
	"errors"
	"testing"
)

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
			t.Log("recover from the error")

		}
	}()

	panic(errors.New("Unknown err happend there"))
}
