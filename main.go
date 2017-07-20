package main

import (
	"sync"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var wg sync.WaitGroup
	kingpin.Parse()
	service := CreateHelloWorldHTTPService(":8000")
	service.Start()

	wg.Add(1)
	wg.Wait()
}
