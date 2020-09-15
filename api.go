package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"time"
)

var conn *websocket.Conn

func main() {
	log.Println("API running at port 8081")
	http.HandleFunc("/", echo)

	c, err := connectWS()
	if err != nil {
		log.Fatal("unable to connect to ws")
	}
	conn = c
	defer conn.Close()

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func connectWS() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
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
