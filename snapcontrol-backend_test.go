package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestEchoHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		Run(setupEchoEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, "Hello World", r.Body.String())
		assert.Equal(t, http.StatusOK, r.Code)
	})
}