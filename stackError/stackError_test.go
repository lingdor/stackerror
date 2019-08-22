package stackError

import (
	"errors"
	"testing"
)

type ii interface {
	XX()string
}

func TestNew(t *testing.T){

	err :=New("test error")
	err.PrintErr()
}

func TestIsStackError(t *testing.T){

	err :=New("test error")
	isStack:=IsStackError(err)
	if !isStack {
		t.Fail()
		return
	}
	rawErr:=errors.New("raw error")
	isStack=IsStackError(rawErr)
	if isStack {
		t.Fail()
		return
	}
}