package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Service", t, func() {
		var port = 8000
		cmd := exec.Command("./something-continuous", "-a", fmt.Sprintf(":%d", port))
		cmd.Start()
		defer cmd.Process.Kill()
		time.Sleep(10 * time.Millisecond)
		AssertHealthCheck(port)
	})
}

func AssertHealthCheck(port int) {
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
}
