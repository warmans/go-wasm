package main

import (
	"log"
	"net/http"
	"flag"
	"net"
	"fmt"
)

var (
	wasmFile  = flag.String("wasm-file", "example.wasm", "The file to serve as WASM")
	publicDir = flag.String("public", "cmd/server/public", "Location of public HTML dir")
	port      = flag.String("port", "", "Port to bind to")
)

func init() {
	flag.Parse()
}

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, *wasmFile)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*publicDir)))
	mux.HandleFunc("/main.wasm", wasmHandler)

	if *port == "" {
		*port = ":0"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		panic(err)
	}
	log.Printf("Running example %s on :%d\n", *wasmFile, listener.Addr().(*net.TCPAddr).Port)

	// and the main HTTP server
	log.Fatal(http.Serve(listener, mux))
}
