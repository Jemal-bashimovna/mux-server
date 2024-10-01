package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Building Rest Api in GO")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "There are all comments")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Return a single comment with id=%s\n", id)
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Post a new comment")
	})

	mux.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "There are all comments")
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "There are all comments")
		case http.MethodPost:
			fmt.Fprintf(w, "post a new comment")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	if er := http.ListenAndServe("localhost:8081", mux); er != nil {
		fmt.Println(er.Error())
	}
}
