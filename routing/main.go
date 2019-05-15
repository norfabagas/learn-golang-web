package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main()  {
    r := mux.NewRouter()
    r.HandleFunc("/books/{title}/page/{page}", func (w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book : %s on page %s\n", title, page)
    })

    r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
        fmt.Fprintf(w, "Home page")
    })

    http.ListenAndServe(":8080", r)
}
