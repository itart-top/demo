package main

import "itart-demo/pipeline"

func main() {
	p:=pipeline.CreateNetworkPipeline("large.in", 800000, 4)
	writeToFile(p, "large.out" )
	printFile("large.out")
}
