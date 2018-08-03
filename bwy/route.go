package bwy

import (
	"net/http"
	"regexp"
	"html/template"
	"www/bwy/config"
)
type www struct {
	match string
	handler func(http.ResponseWriter, *http.Request)
}

var route []www

func RouteAny(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	route = append(route, www{match:pattern,handler:handler})
}

func Match(w *http.ResponseWriter,r *http.Request) {
	for i:=0; i<len(route); i++ {
		regexps := regexp.MustCompile("("+route[i].match+")");
		matchs := regexps.FindSubmatch([]byte(r.URL.Path))
		if matchs != nil {
			route[i].handler(*w,r)
			return
		}
	}

	//w.WriteHeader(404)
	t, _ :=template.ParseFiles(config.CONFIG["init|homePath"]+"views/common/_404.html")
	t.Execute(*w, t)

}

