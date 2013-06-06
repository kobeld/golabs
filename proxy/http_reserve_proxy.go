package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	target, err := url.Parse("http://kobeld2.local:5280/http-bind")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(rw, req)

	err = http.ListenAndServe(":8888", proxy)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
