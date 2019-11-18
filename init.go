package stackError

import (
	"runtime"
)

var pkgName string = "github.com/lingdor/stackError"

func init() {
	pc := make([]uintptr, 1)
	runtime.Callers(1, pc)
	frm := runtime.CallersFrames(pc)
	frmval, ok := frm.Next()
	if ok {
		newName := funcNameToPkgName(frmval.Function)
		if newName != "" {
			pkgName = newName
		}
	}
}
