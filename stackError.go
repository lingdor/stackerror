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
	GetStacks() []runtime.Frame
	GetChild() error
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
	for i := 0; err != nil; i++ {
		if i > 0 {
			buffer.WriteString("\n")
		}
		tt := reflect.TypeOf(err)
		buffer.WriteString(tt.String())
		buffer.WriteString(" : ")
		childStackError, isStack := err.(StackError)
		if isStack {
			buffer.WriteString(childStackError.GetMsg())
			buffer.WriteString("\n")
			buffer.Write(formatStackFrame(childStackError.GetStacks()))
			err = childStackError.GetChild()
			continue
		} else {
			buffer.WriteString(err.Error())
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

//get error stacks
func (this *stackError) GetStacks() []runtime.Frame {
	return this.stack
}

//Get error childInfo
func (this *stackError) GetChild() error {
	return this.child
}
