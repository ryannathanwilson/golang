// Run with `go run server.go`
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ex struct {
	Me  string
	You string
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		echo := r.URL.Query().Get("echo")
		fmt.Fprintf(w, echo)
	})

	http.HandleFunc("/rnw", func(w http.ResponseWriter, r *http.Request) {
		ob := ex{
			Me:  "ryan",
			You: "somebody",
		}
		response, err := json.MarshalIndent(ob, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(response))
	})

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
