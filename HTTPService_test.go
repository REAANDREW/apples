package main_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	sut "github.com/reaandrew/something-continuous"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHTTPService(t *testing.T) {
	var address = ":51000"
	var service = sut.CreateHTTPService(address, func(router *gin.Engine) {})
	Convey("HTTP Service Factory", t, func() {
		Convey("returns a new service", func() {
			So(service, ShouldNotBeNil)
		})
	})

	Convey("HTTP Service Start Stop", t, func() {
		service.Start()

		var url = fmt.Sprintf("http://localhost%s/meta/health", address)

		client := http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		So(err, ShouldBeNil)
		response, err := client.Do(request)
		So(err, ShouldBeNil)
		So(response.StatusCode, ShouldEqual, http.StatusOK)

		service.Stop()
		response, err = client.Do(request)

		So(response, ShouldBeNil)
		So(err, ShouldNotBeNil)

	})
}
