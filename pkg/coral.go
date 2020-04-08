package main

import (
	"flag"
	"github.com/catchkvs/Coral/pkg/handler"
	"github.com/catchkvs/Coral/pkg/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the server...")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/session", handler.Handle)
	http.HandleFunc("/", ping)
	http.Handle("/metrics", promhttp.Handler())
	go server.CleanupWorker()
	log.Fatal(http.ListenAndServe(":4040", nil))
}


func ping(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("healthy"))
}