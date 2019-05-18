package controllers

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

const (
	URL_ROOT = "http://localhost:8889"
)

// APIGET -
func APIGET(t *testing.T, path string) *httpexpect.Object {
	return httpexpect.New(t, URL_ROOT).
		GET("/api"+path).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", 0).
		ContainsKey("message").
		ValueEqual("message", "成功").
		ContainsKey("data").
		Value("data").
		Object()
}

// APIGETList -
func APIGETList(t *testing.T, path string) *httpexpect.Array {
	return APIGET(t, path).
		ContainsKey("list").
		Value("list").
		Array()
}
