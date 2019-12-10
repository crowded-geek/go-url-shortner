package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "../go-hashids"
    "../mux"
)

var m map[string]string

func Root(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    setr := m[params["id"]]
    fmt.Println(setr)
    fmt.Println("Redirecting... ")
    http.Redirect(w, r, setr, 302)
}

func create(w http.ResponseWriter, r *http.Request){
	hd := hashids.NewData()
    h, _ := hashids.NewWithData(hd)
    now := time.Now()
	params := r.URL.Query()
    link := params.Get("longUrl")
    str, _ := h.Encode([]int{int(now.Unix())})
    m[str] = link
    fmt.Fprintf(w, "localhost:8282/%s", str)
    fmt.Println(str)
}

func handleRequests() {
    router := mux.NewRouter()
	m = make(map[string]string)
    router.HandleFunc("/{id}", Root).Methods("GET")
    router.HandleFunc("/create/", create).Methods("GET")
    log.Fatal(http.ListenAndServe(":8282", router))
}

func main() {
    handleRequests()
}