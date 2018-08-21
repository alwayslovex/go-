package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func echoHandle(w http.ResponseWriter,r *http.Request){
	_,err := ioutil.ReadAll(r.Body)
	if err != nil{
		return
	}
	r.ParseForm()
	echo := r.FormValue("echo")
	fmt.Print(echo)
	w.Write([]byte(echo))
}
func httpServ(){
	http.HandleFunc("/echo",echoHandle)
	http.ListenAndServe("localhost:7890",nil)
}

func main()  {
	httpServ()
}
