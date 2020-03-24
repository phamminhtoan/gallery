package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello there, Wellcome to my page!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an mail to <a href=\"mailto:minhtoan.jc@gmail.com\">minhtoan.jc@gmail.com</a>.")
	} else {
		fmt.Fprint(w, r.URL.Path)
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
