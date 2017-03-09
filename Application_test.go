package main_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	sut "github.com/reaandrew/something-continuous"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("Service", t, func() {
		var port = 8000
		var endpoint = fmt.Sprintf(":%d", port)

		var c = make(chan os.Signal)
		defer func() {
			c <- os.Interrupt
		}()

		go func() {
			sut.CreateApplication(endpoint).Run(c)
		}()
		time.Sleep(10 * time.Millisecond)

		Convey("returns Healthy", func() {
			AssertHealthCheck(port)
		})
	})
}
