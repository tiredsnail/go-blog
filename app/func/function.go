package _func

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
func (ret *RetJson) Success(w http.ResponseWriter) {
	retjson, _ := json.Marshal(ret)
	fmt.Fprintf(w, string(retjson) )
	return
}




