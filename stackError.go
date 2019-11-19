package stackerror

import (
	"bytes"
	"reflect"
	"runtime"
)

// StackMaxDeep Max deep with stack information
var StackMaxDeep = 16

// MaxStackSize Max byte size of formating stack information
var MaxStackSize = 1024

type stackError struct {
	stack []runtime.Frame
	msg   string
	child error
}

// StackError stackError package information interface
type StackError interface {
	error
	GetMsg() string
	GetStacks() []runtime.Frame
	GetChild() error
}

// New Get a StackError object with msg
func New(msg string) StackError {
	return NewParent(msg, nil)
}

// NewParent Get a parent stackError object with msg and child
func NewParent(msg string, child error) StackError {
	val := &stackError{
		msg: msg,
	}
	val.stack = getOutCallers(StackMaxDeep)
	val.child = child
	return val
}

//Error format stackError information with string
func (my *stackError) Error() string {
	var err error = my
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

//String invoked Error() method
func (my *stackError) String() string {
	return my.Error()
}

//GetMsg get error message
func (my *stackError) GetMsg() string {
	return my.msg
}

//GetStacks get error stacks
func (my *stackError) GetStacks() []runtime.Frame {
	return my.stack
}

//GetChild Get error childInfo
func (my *stackError) GetChild() error {
	return my.child
}
