package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func UpgradeHttpToWs(c *gin.Context) {
	conn, err := (&websocket.Upgrader{
		// disable cross-site check
		CheckOrigin: func(r *http.Request) bool {
			log.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Upgrade http to webSocket error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorCode": 5001,
			"Data": fmt.Sprintf("Upgrade http to webSocket error: %f", err),
		})
		return
	}

	log.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	//clientManager.Register <- client
}
