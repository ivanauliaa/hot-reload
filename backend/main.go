package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World Ehe"))
	})

	const address = "0.0.0.0:9000"
	fmt.Println("Listening at", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
