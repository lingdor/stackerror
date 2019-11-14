package stackError

import (
	"fmt"
	"testing"
)

type ii interface {
	XX()string
}

func TestNew(t *testing.T){

	err := New("test error")
	fmt.Println(err.Error())
}