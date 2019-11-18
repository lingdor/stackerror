// Package stackError A have stack information with error library
package stackError

import (
	"bytes"
	"reflect"
	"runtime"
)

// StackMaxDeep Max deep with stack infromation
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

//String invoked Error() method
func (this *stackError) String() string {
	return this.Error()
}

//GetMsg get error message
func (this *stackError) GetMsg() string {
	return this.msg
}

//GetStacks get error stacks
func (this *stackError) GetStacks() []runtime.Frame {
	return this.stack
}

//GetChild Get error childInfo
func (this *stackError) GetChild() error {
	return this.child
}
