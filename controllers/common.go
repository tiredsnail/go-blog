package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type RetJson struct {
	Status		string			`json:"status"`
	Msg			string			`json:"msg"`
	Data		interface{}		`json:"data"`
}

func LayoutType() []map[string]string {
	data := []map[string]string{
		{"name":"GoLang","url":"golang"},
		{"name":"Mysql","url":"mysql"},
		{"name":"PHP","url":"php"},
		{"name":"TrashCan","url":"trashcan"},
	}
	return data
}



func (ret *RetJson) Success(w http.ResponseWriter) {
	retjson, _ := json.Marshal(ret)
	fmt.Fprintf(w, string(retjson) )
	return
}

func Error404(w http.ResponseWriter,r *http.Request) {

}
