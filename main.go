package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"
)
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Web coding by Xiang910")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
