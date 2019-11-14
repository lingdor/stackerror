package stackError

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"testing"
)

type ii interface {
	XX()string
}

func TestNew(t *testing.T){

	err := New("test error")
	fmt.Println(err.Error())
}

func BenchmarkA(b *testing.B){
	for i:=0;i<b.N;i++{
		pc := make([]uintptr, 16)
		runtime.Callers(2,pc)
		frm :=runtime.CallersFrames(pc)
		frmval,_:= frm.Next()
		fmt.Println(frmval.Function)
	}

}

func BenchmarkB(b *testing.B){
	for i:=0;i<b.N;i++{
		debug.Stack()
	}

}

func ff(str string){

	str =string(debug.Stack())
	fmt.Println(str)
}

func TestXX(t *testing.T){
	ff("123")
	frames:=getOutCallers(30)
	for _,frame:=range frames {
		reflect.ValueOf(frame.Func)


		fmt.Println(frame.Function)
	}

}