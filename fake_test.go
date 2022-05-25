package main

import (
	"encoding/json"

	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
To run this test without docker:
go test -v -run TestDummy

To run this test with docker:
docker-compose up -d
docker-compose exec go test -vet=off -v -run TestDummy

Some notes:
- The -vet=off is to avoid the cgo compiler error
- The -v is to see the output of the test

Todo:
Currently the dockerfile copies the go binary to an alpine image.
This image can't run go tests and will fail.

We could use the golang:alpine and run the tests there but the
resulting image will be much bigger (300MB instead of 7MB).

Ideally run the tests in one container then let the server run
in a lighter image.
*/

// Dummy test
func TestDummy(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/fake", nil)
	response := httptest.NewRecorder()

	middleware.Dummy(response, request)

	t.Run("returns Car data", func(t *testing.T) {
		var got string
		if json.Unmarshal(response.Body.Bytes(), &got) != nil {
			t.Errorf("Error unmarshalling resposne body")
		}
		want := "Hello world"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}