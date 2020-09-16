package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"log"
	"scaling-ws/internal/ws"
	"time"
)

var addr = flag.String("addr", ":3000", "api service address")
var isNatsPublisher = flag.Bool("natsPublisher", true, "is nats publisher")

const subject = "com.scaling-ws.updates"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer ec.Close()
	defer nc.Close()

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		ws.ServeWs(hub, c, nc, subject)
	})

	r.POST("/ping", func(c *gin.Context) {
		// publish to nats
		if *isNatsPublisher {
			message := map[string]string{"sendTime": time.Now().Format(time.ANSIC)}
			if err := ec.Publish(subject, message); err != nil {
				log.Fatal(err)
			}
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run(*addr)
}
