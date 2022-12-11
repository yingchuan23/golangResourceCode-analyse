// Package yingchuan_json_test
// @Author chenyingchuan
// @Description //TODO
// @Date 15:52 2022/12/11
package yingchuan_json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Wife  string   `json:"-"`          //不将该字段做序列化输出
	Hobby []string `json:",omitempty"` //该字段做序列化输出，如果没有则不做序列化输出
	//Hobby []string `json:"hobby,omitempty"`  //区别 hobby字段会小写
}

func TestJsonPerson1(t *testing.T) {
	var one = Person{
		Name: "11",
		Age:  12,
		Wife: "12",
	}
	var two = Person{
		Name:  "12",
		Age:   12,
		Wife:  "11",
		Hobby: []string{"1", "2"},
	}
	res1, err := json.Marshal(one)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res1))
	res2, err := json.Marshal(two)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res2))
}

func TestJsonPerson2(t *testing.T) {
	var three Person
	stringOfJson := "{\"name\":\"ethan\",\"age\":12}"
	json.Unmarshal([]byte(stringOfJson), &three)
	fmt.Println(three)
}

func TestJsonPerson3(t *testing.T) {
	var mapJson = map[string]interface{}{"name": "小王", "address": "earth", "Wife": "May"}
	b, err := json.Marshal(mapJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	//反序列化 JSON =》 MAP
	var recMap map[string]interface{}
	var jsonString string = string(b)
	json.Unmarshal([]byte(jsonString), &recMap)
	fmt.Println(recMap)
	for key, value := range recMap {
		fmt.Println(key, "=>", value)
	}
}
