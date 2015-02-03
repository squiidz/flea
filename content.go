package main

var APPCONTENT = `package main
 
import (
	"net/http"
	"log"
)

const (
	STATIC = "/public/"
	TEMPLATE = STATIC + "template"
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
var INDEX = `<!DOCTYPE HTML>
<html>
	<head>
		<meta charset="utf-8" />
		<title></title>
		
		<link rel="stylesheet" type="text/css" href="/public/css/style.css">
		<script type="text/javascript" src="/public/js/script.js"></script>
	</head>
	<body>
		<p>Server Work</p>
	</body>
</html>
`
