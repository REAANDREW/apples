package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	sut "github.com/reaandrew/something-continuous"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloWorldHTTPService(t *testing.T) {
	var address = ":51000"

	var service = sut.CreateHelloWorldHTTPService(address)

	Convey("HTTP Service Factory", t, func() {
		Convey("returns a new service", func() {
			service.Start()
			defer service.Stop()

			var client = http.Client{}
			var url = fmt.Sprintf("http://localhost%s/", address)
			request, err := http.NewRequest("GET", url, nil)
			So(err, ShouldBeNil)
			response, err := client.Do(request)
			So(err, ShouldBeNil)
			defer response.Body.Close()

			data, err := ioutil.ReadAll(response.Body)
			So(err, ShouldBeNil)
			So(string(data), ShouldEqual, "Hello, World!")
		})
	})
}
