package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type UserInfo struct {
	Userid string
	Passwd string
}

type UserInfo2 struct {
	Userid string `json:"userid"`
	Passwd string `json:"passwd"`
}

func main() {
	var f interface{}
	jsonstr := []byte(`{"name":"ita","age":24,"fav":["game","ball"],"friend":[{"age":25,"name":"abc"}]}`)
	json.Unmarshal(jsonstr, &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is ", reflect.TypeOf(vv), vv)
		case float64:
			fmt.Println(k, "is ", reflect.TypeOf(vv), vv)
		case []interface{}:
			fmt.Println(k, "is an :", reflect.TypeOf(vv))
			for i, u := range vv {
				fmt.Println(i, u, reflect.TypeOf(u))
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	itachi := UserInfo{"itachi", "123456"}
	b, _ := json.Marshal(itachi)
	fmt.Println(b) //{"Userid":"itachi","Passwd":"123456"}

	itachi2 := UserInfo2{"itachi", "123456"}
	b, _ = json.Marshal(itachi2)
	fmt.Println(b) //{"userid":"itachi","passwd":"123456"}
}
