package main

var APPCONTENT = `package main
 
import (
	"net/http"
	"log"
)

const (
	STATIC = "/public/"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", DefaultHandler)

	mux.Handle(STATIC, http.StripPrefix(STATIC, http.FileServer(http.Dir("public"))))

	log.Panic(http.ListenAndServe(":8080", mux))
}

func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Server Running"))
}

`
