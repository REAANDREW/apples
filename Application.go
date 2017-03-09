package main

import (
	"os"
	"sync"
)

type application struct {
	service *HTTPService
}

func (instance application) Run(c chan os.Signal) {
	var wg sync.WaitGroup
	wg.Add(1)
	instance.service.Start()
	go func() {
		<-c
		instance.service.Stop()
		wg.Done()
	}()
	wg.Wait()
}

func CreateApplication(address string) application {
	return application{
		service: CreateHelloWorldHTTPService(address),
	}
}
