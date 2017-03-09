package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("Service", t, func() {
		var port = 45001
		var endpoint = fmt.Sprintf(":%d", port)
		var service = CreateService(endpoint)
		defer service.Stop()
		service.Start()

		Convey("returns Hello World", func() {
			var client = &http.Client{}
			var url = fmt.Sprintf("http://localhost:%d/", port)

			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				panic(err)
			}

			response, err := client.Do(request)
			if err != nil {
				panic(err)
			}

			So(response.StatusCode, ShouldEqual, http.StatusOK)
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)

			So(string(body), ShouldEqual, "Hello, World!")
		})
	})
}
