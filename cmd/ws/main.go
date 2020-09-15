package main

import (
	"flag"
	"log"
	"net/http"
	"scaling-ws/internal/ws"
)

var addr = flag.String("addr", ":3000", "ws service address")

func main() {
	flag.Parse()
	log.Printf("WS running at port %s", *addr)
	var serverName string
	if *addr == ":3000" {
		serverName = "server 1"
	} else {
		serverName = "server 2"
	}
	hub := ws.NewHub(serverName)
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
