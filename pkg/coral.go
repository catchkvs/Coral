package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/user", userEcho)
	http.HandleFunc("/call", call.HandleUserCall)
	http.HandleFunc("/chat", chat.ChatHandler)
	http.HandleFunc("/admin", adminEcho)
	http.HandleFunc("/sms-handle", sms.SMSHandler)
	http.HandleFunc("/live-orders", restaurant.LiveOrderHandler)
	http.HandleFunc("/index-build", restaurant.MenuIndexBuildHandler)
	http.HandleFunc("/prefix-query", restaurant.MenuItemQueryHandler)
	http.HandleFunc("/get-current-session", chat.ChatHandler)
	http.HandleFunc("/ping", ping)
	go CleanupWorker()
	log.Fatal(http.ListenAndServe(":4040", nil))
}


func ping(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("healthy"))
}