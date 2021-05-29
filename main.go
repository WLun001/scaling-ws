package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

var addr = flag.String("addr", ":3000", "api service address")

const subject = "com.scaling-ws.updates"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	defer ec.Close()
	defer nc.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	r.POST("/ping", func(c *gin.Context) {
		// publish to nats
		message := map[string]string{"sendTime": time.Now().Format(time.ANSIC)}
		if err := ec.Publish(subject, message); err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run(*addr)
}
