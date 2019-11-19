// package  stackerror A have stack information with error library
package stackerror

import (
	"runtime"
)

var pkgName string = "github.com/lingdor/stackerror"

func init() {
	pc := make([]uintptr, 1)
	runtime.Callers(1, pc)
	frm := runtime.CallersFrames(pc)
	frmval, _ := frm.Next()
	//if ok {
	newName := funcNameToPkgName(frmval.Function)
	if newName != "" {
		pkgName = newName
	}
	//}
}
