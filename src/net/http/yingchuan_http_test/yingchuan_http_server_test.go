// Package yingchuan_http
// @Author chenyingchuan
// @Description //TODO
// @Date 16:20 2022/12/12
package yingchuan_http

import (
	"log"
	"net/http"
	"testing"
	"time"
)

// 地址进行重定向
func TestServerDemo1(t *testing.T) {
	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

type timeHandler struct {
	format string
}

func (t *timeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	tm := time.Now().Format(t.format)
	writer.Write([]byte("The time is: " + tm))
}

func timeHandler1(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	})
}

func timeHandler2(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}

func TestServerTimeHandler(t *testing.T) {
	mux := http.NewServeMux()

	th := &timeHandler{format: time.RFC1123}

	mux.Handle("/time", th)
	log.Println("Listening...")

	http.ListenAndServe(":3000", mux)
}

func TestServerTineHandler2(t *testing.T) {
	mux := http.NewServeMux()

	t1 := &timeHandler{format: time.RFC1123}
	mux.Handle("/rfc1123", t1)

	t2 := &timeHandler{
		format: time.RFC3339,
	}
	mux.Handle("/rfc3339", t2)

	log.Println("Listening...")

	http.ListenAndServe(":3000", mux)
}

func TestDefaultServeMux(t *testing.T) {
	var format string = time.RFC1123
	t1 := timeHandler1(format)
	http.Handle("/time", t1)
	log.Println("Listening...")

	//不写的话 就是默认的
	http.ListenAndServe(":3000", nil)
}
