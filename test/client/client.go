package main

// https://github.com/eranyanay/1m-go-websockets/blob/master/client.go

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"server-demo/utils"
	"time"

	"github.com/gorilla/websocket"
)

var (
	url         = flag.String("url", "ws://172.17.0.1:9090/echo", "test url")
	connections = flag.Int("conn", 10, "number of websocket connections")
)

func main() {
	utils.IncreaseResourcesLimit()

	flag.Usage = func() {
		io.WriteString(os.Stderr,
			"Usage: ./client -url=ws://172.17.0.1:9090/echo -conn=10\n",
		)
		flag.PrintDefaults()
	}
	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("Connecting to %s", *url)
	var conns []*websocket.Conn
	for i := 0; i < *connections; i++ {
		c, _, err := websocket.DefaultDialer.Dial(*url, nil)
		if err != nil {
			log.Println("Failed to connect", i, err)
			break
		}

		conns = append(conns, c)
	}

	connectedCount := len(conns)
	log.Printf("Finished initializing %d connections", connectedCount)
	if connectedCount == 0 {
		return
	}

	needsExit := false
	for !needsExit {
		select {
		case <-interrupt:
			needsExit = true
			log.Println("interrupt")
		}

		// time.Sleep(time.Millisecond * 1)
		for i := 0; i < connectedCount && !needsExit; i++ {
			conn := conns[i]

			err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second*5))
			if err != nil {
				fmt.Printf("Failed to receive pong: %v", err)
			}

			msg := fmt.Sprintf("Hello from conn %v", i)
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}

	for i := 0; i < connectedCount; i++ {
		conn := conns[i]
		closeMsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
		conn.WriteControl(websocket.CloseMessage, closeMsg, time.Now().Add(time.Second))
		conn.Close()
	}
}
