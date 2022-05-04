package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
	fmt.Println("Please Visit http://localhost:8888/")
	http.ListenAndServe(":8888", nil)
	cmd := exec.Command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}