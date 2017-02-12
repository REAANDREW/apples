package main

import (
	"errors"
	"os"
	"os/signal"
	"sync"

	"github.com/fvbock/endless"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	Version   string
	BuildTime string
)

var (
	ErrCannotGetHostInfo = errors.New("Cannot get host information")
)

type Service struct {
	engine   *gin.Engine
	stopFunc func()
	endpoint string
}

func (instance *Service) Start() {
	var server = endless.NewServer(instance.endpoint, instance.engine)

	instance.stopFunc = func() {
		server.EndlessListener.Close()
	}

	go server.ListenAndServe()
}

func (instance *Service) Stop() {
	instance.stopFunc()
}

func CreateService(endpoint string) *Service {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	})

	return &Service{
		engine:   router,
		endpoint: endpoint,
	}
}

func main() {
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
