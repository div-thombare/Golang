package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//*****************GORILLA MUX
// func ArticleHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
// 	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
// 	srv := &http.Server{
// 		Handler:      r,
// 		Addr:         "127.0.0.1:8000",
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}
// 	log.Fatal(srv.ListenAndServe())
// }

// ******************QUERY BASED MATCHING
// func QueryHandler(w http.ResponseWriter, r *http.Request) {
// 	queryParams := r.URL.Query()

// 	id := queryParams.Get("id")
// 	category := queryParams.Get("Category")

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Got parameter id: %s\n", id)
// 	fmt.Fprintf(w, "Got parameter category: %s\n", category)
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/articles", QueryHandler).Queries("id", "{id}", "category", "{category}")

// 	srv := &http.Server{
// 		Handler:      r,
// 		Addr:         "127.0.0.1:8000",
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}
// 	log.Fatal(srv.ListenAndServe())
// }

// ****************HOST BASED MATCHING
func QueryHandler(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s!",
		queryParams["category"][0])

}
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
