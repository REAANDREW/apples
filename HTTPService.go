package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
)

type HTTPService struct {
	engine   *gin.Engine
	stopFunc func()
	endpoint string
}

func (instance *HTTPService) Start() {
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

func (instance *HTTPService) Stop() {
	instance.stopFunc()
	time.Sleep(1)
}

func CreateHTTPService(endpoint string, configureRouting func(router *gin.Engine)) *HTTPService {
	router := gin.Default()

	configureRouting(router)

	router.GET("/meta/health", func(ctx *gin.Context) {
		fmt.Println("returning status ok")
		ctx.Status(200)
	})

	return &HTTPService{
		engine:   router,
		endpoint: endpoint,
	}
}
