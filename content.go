package main

var APPCONTENT = `package main
 
import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", DefaultHandler)

	log.Panic(http.ListenAndServe(":8080", mux))
}

func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Server Running"))
}

`
