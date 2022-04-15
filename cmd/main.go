package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":9927", "listen address")
	dir    = flag.String("dir", "../assets", "directory to serve")
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
