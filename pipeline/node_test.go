package pipeline

import (
	"testing"
	"fmt"
)

func TestArraySource(t *testing.T) {
	p:= ArraySource(1,2,3,3,4)
	for d:=range p {
		fmt.Println(d)
	}
}