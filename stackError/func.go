package stackError

import (
	"fmt"
	"os"
	"runtime/debug"
)

/**
print stack and exit
 */
func CheckExitError (err error) {
	if err==nil {
		return
	}
	if stackErr,isTrue:= CheckStackError(err);isTrue {
		stackErr.PrintErr()
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr,"error message:",err.Error())
	fmt.Fprintln(os.Stderr,"stack:",string(debug.Stack()))
	os.Exit(1)
}
func Panic(errMsg string){
	newErr:=New(errMsg,POWERID_YES)
	panic(newErr)
}
func PanicError(err error){
	newErr:=NewFromError(err,POWERID_YES)
	panic(newErr)
}