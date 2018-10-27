package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func echoHandle(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	r.ParseForm()
	echo := r.FormValue("echo")
	fmt.Print(echo)
	w.Write([]byte(echo))
}

func jsonParse(body []byte) map[string]string {
	var j interface{}
	json.Unmarshal(body, &j)
	result := make(map[string]string, 0)
	convertJson := j.(map[string]interface{})
	for k, v := range convertJson {
		switch ty := v.(type) {
		case string:
			result[k] = ty
		case int:
			result[k] = strconv.Itoa(ty)
		case float64:
			result[k] = strconv.FormatFloat(ty, 'f', 1, 32)
		case map[string]interface{}:
			fmt.Printf("....")
		}
	}
	return result
}
func PostHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read post body error : %s", err.Error())
		w.Write([]byte("error"))
		return
	}
	var jsonstr interface{}
	err = json.Unmarshal(body, &jsonstr)
	if err != nil {
		fmt.Printf("json parse error : %s", err.Error())
		w.Write([]byte("json parse error"))
	}
	mp := jsonParse(body)

	var rsp string
	for k, v := range mp {
		rsp += fmt.Sprintf("[%s] = [%s]\n", k, v)
	}
	w.Write([]byte(rsp))
}

func httpServ() {
	http.HandleFunc("/echo", echoHandle)
	http.HandleFunc("/postjson", PostHandle)
	http.ListenAndServe("localhost:7890", nil)
}

func main() {
	httpServ()
}
