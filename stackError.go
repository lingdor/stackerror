package stackError

import (
	"bytes"
	"reflect"
	"runtime"
)
const DEFAULT_STACK_MAX_DEEP=16
var StackMaxDeep=DEFAULT_STACK_MAX_DEEP

type stackError struct{
	stack []runtime.Frame
	msg string
	parent error
}

type StackError interface{
	error
	GetMsg() string
	GetStack() string
}

func New(msg string) error{
	return ChildNew(msg,nil)
}

func ChildNew(msg string, parent error) error{
	val:= &stackError{
		msg:msg,
	}
	val.stack= getOutCallers(StackMaxDeep)
	val.parent=parent
	return val
}

func (this *stackError) Error() string{
	var err error = this
	buffer:=bytes.NewBuffer([]byte{})
	for err != nil {
		tt:=reflect.TypeOf(err)
		buffer.WriteString(tt.String())
		buffer.WriteString(" : ")
		buffer.WriteString(this.msg)
		if tt==reflect.TypeOf(this) {
			buffer.WriteString("\n")
			stackErr:=err.(*stackError)
			buffer.Write(formatStackFrame(stackErr.stack))
			err=stackErr.parent
			continue
		}
		break
	}
	return buffer.String()
}

func Throw(msg string){
	err:=New(msg)
	CheckExitError(err)
}