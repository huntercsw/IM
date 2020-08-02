package main

import "github.com/gin-gonic/gin"

func RouterInit(r *gin.Engine) {
	r.GET("/ws/upgradeToWs", UpgradeHttpToWs)
}
