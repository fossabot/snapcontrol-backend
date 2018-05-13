package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
)

func TestEchoHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		Run(setupEchoEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "Hello World", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestHelloWorldWithoutName(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := echoHelloHandler()

	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World", rec.Body.String())
	}
}

func TestHelloWorldWithName(t *testing.T) {
	rec, c := setupEchoContextForHelloWorld()

	h := echoHelloHandler()

	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World, Ulsmar!", rec.Body.String())
	}
}

func setupEchoContextForHelloWorld() (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", strings.NewReader("Ulsmar"))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}
