package _const

import "net/http"

var REQUEST_URI string

func Request(r *http.Request) {
	REQUEST_URI = r.URL.Path
}