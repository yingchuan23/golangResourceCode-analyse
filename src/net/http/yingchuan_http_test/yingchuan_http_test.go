// Package yingchuan_http_test
// @Author chenyingchuan
// @Description //TODO
// @Date 09:21 2022/12/12
package yingchuan_http_test

import (
	"fmt"
	"net/http"
	"testing"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi there,I love %s!", request.URL.Path[1:])
}

func TestHttpDemo1(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
