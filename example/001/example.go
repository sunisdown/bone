package main

import (
	"log"
	"net/http"

	"github.com/go-zoo/bone"
)

func main() {
	// New mux instance
	mux := bone.New()
	// Custom 404
	mux.NotFound(Handler404)
	// Handle with any http method, Handle takes http.Handler as argument.
	mux.Handle("/index", http.HandlerFunc(homeHandler))
	mux.Handle("/index/:var/info/:test", http.HandlerFunc(varHandler))
	// Get, Post etc... takes http.HandlerFunc as argument.
	mux.Post("/home", http.HandlerFunc(homeHandler))
	mux.Get("/home/:var", http.HandlerFunc(varHandler))

	mux.Get("/:any", http.HandlerFunc(homeHandler))

	// Start Listening
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("WELCOME HOME"))
}

func varHandler(rw http.ResponseWriter, req *http.Request) {
	varr := bone.GetValue("var")
	test := bone.GetValue("test")
	log.Println("VAR = ", varr)
	log.Println("TEST = ", test)

	rw.Write([]byte(varr + " " + test))
}

func Handler404(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("These are not the droids you're looking for ..."))
}
