package main

import (
	"itart-demo/pipeline"
	"os"
	"fmt"
	"bufio"
)

func main() {
	fileName:="small.in"
	size:=64
	file, err :=os.Create(fileName)
	if err!=nil{
		panic(file)
	}
	defer file.Close()
	p:=pipeline.RandomSource(size)
	writer:= bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()
	// reader
	file, err =os.Open(fileName)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	for v:=range p {
		fmt.Println(v)
	}
}
