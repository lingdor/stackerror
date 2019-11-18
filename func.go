package stackError

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//CheckExitError Check error and print after exit
func CheckExitError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "error message:", err.Error())
	os.Exit(1)
}

//CheckPanic check error and panic
func CheckPanic(err error) {
	if err == nil {
		return
	}
	switch err.(type) {
	case *stackError:
		panic(err)
	default:
		panic(New(err.Error()))
	}
}

//Panic make a stackError panic()
func Panic(errMsg string) {
	newErr := New(errMsg)
	panic(newErr)
}

//PanicError make a stacError panic() with error
func PanicError(err error) {
	newErr := NewParent(err.Error(), err)
	panic(newErr)
}

//Get PkgName from *Func.Function
func funcNameToPkgName(funcName string) string {

	index := strings.LastIndex(funcName, "/")
	if index == -1 {
		return ""
	}
	index = strings.Index(funcName[index:], ".") + index
	if index == -1 {
		return ""
	}
	pkgName = funcName[0:index]
	return pkgName
}

//condition string start with
func startWith(full string, partten string) bool {
	lenPartten := len(partten)
	if lenPartten > len(full) {
		return false
	}
	return full[0:lenPartten] == partten
}

//Get callers of current
func getOutCallers(maxDeep int) []runtime.Frame {
	pc := make([]uintptr, maxDeep)
	result := make([]runtime.Frame, 0, maxDeep)
	runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc)
	skipMod := true
	for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
		if skipMod {
			if startWith(frame.Function, pkgName) {
				continue
			}
			skipMod = false
		}
		result = append(result, frame)
	}
	return result
}

//stack Frame to []byte message
func formatStackFrame(frames []runtime.Frame) []byte {
	buffer := bytes.Buffer{}
	for i, frame := range frames {
		if i != 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString("  at ")
		buffer.WriteString(frame.Function)
		buffer.WriteString("( ")
		buffer.WriteString(frame.File)
		buffer.WriteString(":")
		buffer.WriteString(strconv.Itoa(frame.Line))
		buffer.WriteString(" )")
	}
	return buffer.Bytes()

}
