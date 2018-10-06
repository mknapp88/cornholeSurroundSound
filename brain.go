/*
Usage:
	-p="8080": port to serve on
	-d=".":    the directory of static files to host

Navigating to http://localhost:8080 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"fmt"
)

var logLevel = 1

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UP")
}

func main() {
  log.Printf("Starting brain...")

	port := flag.String("p", "8080", "port to serve on")

	flag.Parse()

	mux := http.NewServeMux()

	hostPort := net.JoinHostPort("0.0.0.0", *port)

	mux.HandleFunc("/health", health)
	log.Printf("Serving requests on HTTP port: %s\n", *port)

	server := http.Server{Handler: mux, Addr: hostPort}
	log.Fatal(server.ListenAndServe())

}

