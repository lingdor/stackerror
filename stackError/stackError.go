package stackError

import (
	"fmt"
	"os"
	"runtime/debug"
)

var Default_Stack bool

const POWERID_YES=-2

type stackError struct{
	stack []byte
	msg string
}
type StackError interface{
	Error() string
	PrintErr()
	LoadStack()
}
func NewType(msg string) StackError{
	val:= &stackError{
		msg:msg,
	}
	return val
}

func New(msg string,powerId int) StackError{
	val:= &stackError{
		msg:msg,
	}
	if Default_Stack||powerId==POWERID_YES || HasPower(powerId) {
		val.stack= debug.Stack()
	}
	return val
}

func NewFromError(err error,powerId int) StackError{
	if err==nil {
		return nil
	}
	val:= &stackError{
		msg:err.Error(),
	}
	if Default_Stack|| powerId == POWERID_YES || HasPower(powerId) {
		val.stack= debug.Stack()
	}
	return val
}

/**
print error message and stack
 */
func (this *stackError) PrintErr(){
	fmt.Fprintln( os.Stderr,"error message:",this.msg)
	if this.stack!=nil {
		fmt.Fprintln( os.Stderr, string(this.stack))
		return
	}
	fmt.Fprintln( os.Stderr, "printing-stack:")
	fmt.Fprintln( os.Stderr, string(debug.Stack()))
}

func (this *stackError) Error() string{
	return this.msg
}
func (this *stackError) LoadStack(){
	this.stack=debug.Stack()
}

func IsStackError(err error) bool{
	switch err.(type) {
	case StackError: return true
	case *stackError:return true
	}
	return false
}

func CheckStackError(err error) (StackError,bool){
	switch err.(type) {
	case *stackError: return err.(StackError),  true
	case StackError: return err.(StackError),  true
	}
	return nil,false
}

func ThrowExit(msg string){
	err:=New(msg,POWERID_YES)
	CheckExitError(err)
}