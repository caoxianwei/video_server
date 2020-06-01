package main

import (
	"fmt"
	"net/http"
)

func Print1to20() int {
	res := 0
	for i := 0; i <= 20; i++ {
		res += 1
	}
	return res
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
