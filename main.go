package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	requests()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func aboutMe(res http.ResponseWriter, r *http.Request) {

	who := "shamimenamul"

	fmt.Println(res, "My name is")
	fmt.Println(who)
	io.WriteString(res, "<h1>This is my website</h1>")
}

func requests() {
	http.HandleFunc("/aboutme", aboutMe)
}
