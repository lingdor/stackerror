package stackError

import (
	"fmt"
	"os"
	"runtime/debug"
)

type stackError struct{
	stack []byte
	msg string
}

func New(msg string) *stackError{
	return &stackError{
		stack:debug.Stack(),
		msg:msg,
	}

}

func (this *stackError) PrintErr(){
	fmt.Fprintln( os.Stderr,"error:",this.msg)
	fmt.Fprintln( os.Stderr, string(this.stack))
}