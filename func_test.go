package stackError

import (
	"errors"
	"testing"
)

func TestCheckPanic(t *testing.T) {

	defer func() {
		if err := recover(); err == nil {
			t.Error()
		}
	}()

	err := errors.New("TestCheckPanic")
	CheckPanic(err)
	PanicError(err)
	Panic("TestCheckPanic")
}
