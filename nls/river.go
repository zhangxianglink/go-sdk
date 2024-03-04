package nls

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Connection(ws url.URL, headers http.Header, wait time.Duration) *websocket.Conn {
	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = wait
	c, _, err := websocket.DefaultDialer.Dial(ws.String(), headers)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return c
}

func Run(c *websocket.Conn, onMessage func(msg []byte), onClose func(err error)) chan struct{} {
	done := make(chan struct{})
	// 创建一个空结构体类型的通道，用于通知协程退出。
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				onClose(err)
				return
			}
			onMessage(message)
		}
	}()
	return done
}
