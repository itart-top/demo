package pool

import (
	"testing"
	"fmt"
)

func TestPool_Schedule(t *testing.T) {
	pool:=New(128)
	for i:=0;i<1000; i++  {
		_i:=i
		pool.Schedule(func() {
			fmt.Println("task=", _i)
		})
	}
	fmt.Println("end")
}