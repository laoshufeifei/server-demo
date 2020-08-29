package websocket

import (
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var count int64

func atomicAddCount(delta int64) {
	n := atomic.AddInt64(&count, delta)
	if delta > 0 && n%100 == 0 {
		log.Printf("max number of connections: %v", n)
	}
}

// HandleWSEcho handle ws://HOST:PORT/echo
func HandleWSEcho(ctx *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("upgrade had error:", err)
		return
	}
	atomicAddCount(1)

	defer func() {
		ws.Close()
		atomicAddCount(-1)
	}()

	for {
		//读取ws中的数据
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read message had error:", err)
			break
		}

		err = ws.WriteMessage(messageType, message)
		// err = ws.WriteJSON(gin.H{
		//	"message": msg
		// })
		if err != nil {
			break
		}
	}
}
