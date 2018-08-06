package main

import (
	"log"
	"net/http"
	"flag"
)

var (
	wasmFile = flag.String("wasm-file", "example.wasm", "The file to serve as WASM")
	publicDir = flag.String("public", "cmd/server/public", "Location of public HTML dir")
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

	log.Printf("Running example %s on :3000", *wasmFile)
	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
