package http

import (
	"fmt"

	"quickshare/helpers"

	"github.com/gin-gonic/gin"
)

const ConfigFile = "quickshare.json"

func StartHTTPDaemon() {
	router := gin.Default()

	router.GET("/pair", func(ctx *gin.Context) {
		ipAddress, ok := ctx.GetQuery("ip")
		if !ok {
			ctx.AbortWithStatus(422)
		}

		err := helpers.AddRecord(ipAddress, "pending")
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
	})

	router.GET("/accept", func(ctx *gin.Context) {
		ipAddress, ok := ctx.GetQuery("ip")
		if !ok {
			ctx.AbortWithStatus(422)
		}

		err := helpers.AddRecord(ipAddress, "paired")
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
	})

	router.POST("/send", func(ctx *gin.Context) {
	})

	router.Run()
}
