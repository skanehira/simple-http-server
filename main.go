package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("p", "8080", "port")
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("http server listen: " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
