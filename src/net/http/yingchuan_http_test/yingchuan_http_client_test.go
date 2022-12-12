// Package yingchuan_http
// @Author chenyingchuan
// @Description //ref: https://studygolang.com/articles/9467
// @Date 14:28 2022/12/12
package yingchuan_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestHttpClientGet(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Println(err)
		return
	}

	//关闭 尽量不要这么写
	defer resp.Body.Close()

	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}
	fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)

	fmt.Printf("resp Proto %s\n", resp.Proto)

	fmt.Printf("resp content length %d\n", resp.ContentLength)

	fmt.Printf("resp transfer encoding %v\n", resp.TransferEncoding)

	fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)

	fmt.Println(reflect.TypeOf(resp.Body)) // *http.gzipReader

	buf := bytes.NewBuffer(make([]byte, 0, 512))

	length, _ := buf.ReadFrom(resp.Body)

	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	fmt.Println(string(buf.Bytes()))
}

func TestHttpClentPost(t *testing.T) {
	client := http.Client{}
	req, err := http.NewRequest("POST", "http://www.maimaiche.com/loginRegister/login.do", strings.NewReader("mobile=xxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp.Header.Get("Content-Type")) // application/json;charset=UTF-8
	type Result struct {
		Msg    string
		Status string
		Obj    string
	}
	res := &Result{}
	json.Unmarshal(body, res)

	if res.Status == "1" {
		fmt.Println(res.Msg)
	} else {
		fmt.Println("login err")
	}

	fmt.Println(res)
}

func TestHttpClientPost2(t *testing.T) {
	resp, err := http.Post("http://www.maimaiche.com/loginRegister/login.do",
		"application/x-www-form-urlencoded",
		strings.NewReader("mobile=xxxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestHttpClientPostForm(t *testing.T) {
	postParam := url.Values{
		"mobile":      {"xxxxxx"},
		"isRemberPwd": {"1"},
	}
	resp, err := http.PostForm("http://www.maimaiche.com/loginRegister/login.do", postParam)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
