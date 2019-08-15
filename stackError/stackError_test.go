package stackError

import "testing"

func TestNew(t *testing.T){

	err :=New("test error")
	err.PrintErr()

}


