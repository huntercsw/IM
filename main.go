package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
)

func init() {
	InitLogger()
}

func main() {
	r := gin.Default()
	RouterInit(r)

	go func() {
		if err := r.Run("0.0.0.0:9999"); err != nil {
			Logger.Error("run gin service error:", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	log.Println(s)
}
