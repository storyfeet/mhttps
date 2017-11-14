package main

import (
	"fmt"
	"net/http"
	"strings"
)

type IHandler struct{}

type SHandler struct{}

func (ih IHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostdom := strings.Split(r.Host, ":")[0]
	http.Redirect(w, r, "https://"+hostdom+":8181"+r.URL.Path, 302)
}

func (sh SHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I hope you feel secure now you are here")
}
