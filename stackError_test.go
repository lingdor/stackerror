package stackError

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {

	err := New("test error")
	if err == nil || err.GetMsg() != "test error" || len(err.GetStacks()) < 1 {
		t.Fail()
	}
	parentErr := NewParent("parent msg", err)
	if parentErr.GetChild() != err {
		t.Fail()
	}
	parentStr := parentErr.Error()
	if parentStr == parentErr.GetMsg() || strings.Index(parentStr, err.GetMsg()) == -1 {
		t.Fail()
	}

}
