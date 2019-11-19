package stackerror

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
}

func TestPanicError(t *testing.T) {

	defer func() {
		if err := recover(); err == nil {
			t.Error()
		}
	}()

	err := errors.New("TestCheckPanic")
	PanicError(err)
}

func TestPanic(t *testing.T) {

	defer func() {
		if err := recover(); err == nil {
			t.Error()
		}
	}()
	Panic("TestCheckPanic")
}
