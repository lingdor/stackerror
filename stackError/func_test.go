package stackError

import (
	"testing"
)

func TestCheckExitError(t *testing.T) {

	err:=New("stackErr")
	//err2 := errors.New("stackErr")
	//err.PrintErr()
	CheckExitError(err)

}
type tt struct{
	Attr string
}