package main

import (
	"os"
	"os/signal"
	"sync"
)

var (
	Version   string
	BuildTime string
)

func main() {
	//address := os.Args[1:2]

	var wg sync.WaitGroup
	wg.Add(1)
	var service = CreateService(":45000")
	service.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		service.Stop()
		wg.Done()
	}()
	wg.Wait()
}
