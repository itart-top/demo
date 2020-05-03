package pipeline

import (
	"net"
	"bufio"
	"os"
	"strconv"
)

func NetworkSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listener.Close()
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSink(writer, in)
	}()
}
func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		r := ReaderSource(conn, -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}
func CreateNetworkPipeline(
	fileName string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	var sortAddr [] string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := ReaderSource(bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000+i)
		NetworkSink(addr, InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	return nil
	var sortResults [] <-chan int
	for _, addr := range sortAddr {
		sortResults = append(sortResults, NetworkSource(addr))
	}
	return MergeN(sortResults...)
}
