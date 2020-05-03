package main

import (
	"itart-demo/pipeline"
	"fmt"
)

func main() {
	p:=pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(1,3,4,6)),
		pipeline.InMemSort(pipeline.ArraySource(9,8,7,5)))
	for v := range p {
		fmt.Println(v)
	}
	d:=1/2
	fmt.Println("",d)
}
