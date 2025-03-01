package main

import (
	"etextbook_server/auth"
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", TestHandler)
	http.HandleFunc("/signup", auth.SignUpHandle)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
