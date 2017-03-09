package main

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestService(t *testing.T) {
	Convey("Service Factory", t, func() {
		Convey("returns a new service", func() {
			So(CreateService(":45000"), ShouldNotBeNil)
		})
	})

	FocusConvey("Service Start Stop", t, func() {
		var address = ":51000"
		var service = CreateService(address)
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
