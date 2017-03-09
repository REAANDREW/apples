package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
)

type Service struct {
	engine   *gin.Engine
	stopFunc func()
	endpoint string
}

func (instance *Service) Start() {
	var server = manners.NewWithServer(&http.Server{
		Addr:    instance.endpoint,
		Handler: instance.engine,
	})

	instance.stopFunc = func() {
		server.Close()
	}

	go server.ListenAndServe()
	time.Sleep(1)
}

func (instance *Service) Stop() {
	instance.stopFunc()
	time.Sleep(1)
}

func CreateService(endpoint string) *Service {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	})
	router.GET("/meta/health", func(ctx *gin.Context) {
		fmt.Println("returning status ok")
		ctx.Status(200)
	})

	return &Service{
		engine:   router,
		endpoint: endpoint,
	}
}
