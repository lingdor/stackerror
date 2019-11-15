package stackError

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

/**
Check error and print after exit
*/
func CheckExitError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "error message:", err.Error())
	os.Exit(1)
}

/**
check error and panic
*/
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
func Panic(errMsg string) {
	newErr := New(errMsg)
	panic(newErr)
}
func PanicError(err error) {
	newErr := ChildNew(err.Error(), err)
	panic(newErr)
}

func FuncNameToPkgName(funcName string) string {

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

func startWith(full string, partten string) bool {
	lenPartten := len(partten)
	if lenPartten > len(full) {
		return false
	}
	return full[0:lenPartten] == partten
}

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
