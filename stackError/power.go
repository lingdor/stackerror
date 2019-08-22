package stackError

import (
	"sync"
)

var powers map[int] bool

var powerIndex int
var powerLock sync.Mutex

func GetPowerKey() int{
	powerLock.Lock()
	defer powerLock.Unlock()
	powerIndex++
	return powerIndex
}

func HasPower(index int) bool{
	if index == -1 {
		return false
	}
	if val,has :=powers[index];has {
		return val
	}
	return false
}

func SetPower(isOpen bool,index int){
	powers[index]=isOpen
}