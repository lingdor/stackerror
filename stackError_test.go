// Package stackerror A have stack information with error library
package stackerror

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	rawError := New("test error")

	stackErr := NewParent("parent msg", rawError)
	if stackErr.GetChild() != rawError || stackErr.GetMsg() != "parent msg" || len(stackErr.GetStacks()) < 1 {
		t.Fail()
	}

	formatStr := fmt.Sprint(stackErr)
	if formatStr == "" {
		t.Fail()
	}
	//rawError := errors.New("newerror")
	//stackParent := NewParent("parentStack", rawError)

}
