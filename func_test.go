package stackError

import (
	"fmt"
	"testing"
)

func TestCheckExitError(t *testing.T) {

	err:= New("stackErr")
	//err2 := errors.New("stackErr")
	//err.PrintErr()
	CheckExitError(err)

}
type tt struct{
	Attr string
}

func TestDebugStack(t *testing.T){

	fmt.Println(1)


}
func TestDebugCaller(t *testing.T){

	fmt.Println(2)



}