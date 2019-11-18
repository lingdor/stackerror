package stackError

import (
	"bytes"
	"reflect"
	"runtime"
)

var StackMaxDeep = 16
var MaxStackSize = 1024

type stackError struct {
	stack []runtime.Frame
	msg   string
	child error
}

type StackError interface {
	error
	GetMsg() string
}

func New(msg string) StackError {
	return NewParent(msg, nil)
}

func NewParent(msg string, child error) StackError {
	val := &stackError{
		msg: msg,
	}
	val.stack = getOutCallers(StackMaxDeep)
	val.child = child
	return val
}

func (this *stackError) Error() string {
	var err error = this
	buffer := bytes.Buffer{}
	for err != nil {
		tt := reflect.TypeOf(err)
		buffer.WriteString(tt.String())
		buffer.WriteString(" : ")
		buffer.WriteString(this.msg)
		if tt == reflect.TypeOf(this) {
			buffer.WriteString("\n")
			stackErr := err.(*stackError)
			buffer.Write(formatStackFrame(stackErr.stack))
			err = stackErr.child
			continue
		}
		break
	}
	return buffer.String()
}

func (this *stackError) String() string {
	return this.Error()
}

//get error message
func (this *stackError) GetMsg() string {
	return this.msg
}
