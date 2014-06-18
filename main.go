package main

import (
	"fmt"
	"net/http"
	"os"
)

var port string

func init() {
	port = os.Getenv("PORT")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")

		for _, env := range os.Environ() {
			fmt.Fprintf(w, "%v\n", env)
		}
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
