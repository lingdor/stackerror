package stackError

import (
	"fmt"
	"os"
	"runtime/debug"
)

type StackError struct{
	stack []byte
	msg string
}

func New(msg string) *StackError{
	return &StackError{
		stack:debug.Stack(),
		msg:msg,
	}
}

func (this *StackError) PrintErr(){
	fmt.Fprintln( os.Stderr,"error:",this.msg)
	fmt.Fprintln( os.Stderr, string(this.stack))
}
func (this *StackError) Error() string{
	return this.msg
}

func IsStackError(err error) bool{
	switch err.(type) {
	case *StackError: return true
	}
	return false
}