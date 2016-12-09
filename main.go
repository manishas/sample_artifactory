package main

import (
	"fmt"
	"net/http"
)

func welcome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Welcome to website !")
}

func main() {
	http.HandleFunc("/", welcome)
	http.Handle("/public", http.FileServer(http.Dir("public")))

	fmt.Printf("www started and listening on :9000")
	http.ListenAndServe(":9000", nil)
}
