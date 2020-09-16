package main

import (
	"flag"
	"github.com/avast/retry-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"scaling-ws/internal/ws"
	"time"
)

var addr = flag.String("addr", ":3000", "api service address")
var wsAddr = flag.String("wsAddr", "localhost:4000", "ws service address")
var skipWs = flag.Bool("skipWs", false, "skip ws setup")

var conn *websocket.Conn

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	var serverName string
	if *addr == ":3000" {
		serverName = "server 1"
	} else {
		serverName = "server 2"
	}
	hub := ws.NewHub(serverName)
	go hub.Run()

	if !*skipWs {
		err := retry.Do(func() error {
			c, err := connectWS()
			if err != nil {
				log.Println("attempting to connect to ws")
				return err
			}
			conn = c
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
	}

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		ws.ServeWs(hub, c, nil)
	})

	r.POST("/ping", func(c *gin.Context) {
		// send to ws
		if !*skipWs {
			conn.WriteJSON(map[string]string{"sendTime": time.Now().Format(time.ANSIC)})
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run(*addr)
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
