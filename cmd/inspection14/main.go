package main

import (
	"example.com/inspection14/internal/archive"
	"example.com/inspection14/internal/audit"
	"example.com/inspection14/internal/catalog"
	"example.com/inspection14/internal/httpapi"
	"example.com/inspection14/internal/query"
	"example.com/inspection14/internal/store"
	"example.com/inspection14/internal/workflow14"
	"log"
	"net/http"
)

func main() {
	s, e := store.Open("inspection14.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	c := catalog.New(s)
	w := workflow14.New(s)
	q := query.New(s)
	_ = archive.New(s)
	_ = audit.New(s)
	log.Fatal(http.ListenAndServe(":8080", httpapi.New(c, w, q).Routes()))
}
