package main

import (
	"encoding/json"
	"flag"
	"github.com/avast/retry-go"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"time"
)

var conn *websocket.Conn
var addr = flag.String("addr", ":8081", "http service address")
var wsAddr = flag.String("wsAddr", "localhost:3000", "ws service address")

func main() {
	flag.Parse()
	log.Printf("API running at port %s", *addr)
	http.HandleFunc("/", echo)

	err := retry.Do(func() error {
		c, err := connectWS()
		if err != nil {
			log.Println("attempting to connect to ws")
			return err
		}
		conn = c
		return nil
	})
	defer conn.Close()

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func connectWS() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: *wsAddr, Path: "/ws"}
	log.Printf("connecting ws at %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to WS")
	return c, nil
}

func echo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	enableCORS(&w, r)

	conn.WriteJSON(map[string]string{"sendTime": time.Now().Format(time.ANSIC)})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func enableCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
}
