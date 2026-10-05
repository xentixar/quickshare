package http

import (
	"fmt"

	"quickshare/helpers"

	"github.com/gin-gonic/gin"
)

const port = "8082"

func InitServer(files map[string]string) {
	helpers.GenerateQRCode("http://" + helpers.GetIPAddress() + ":" + port)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		_, _ = ctx.Writer.Write([]byte(generateLinks(files)))
	})

	router.GET("/download", func(ctx *gin.Context) {
		path, exists := ctx.GetQuery("path")
		if !exists {
			ctx.AbortWithStatus(404)
		}

		if file := files[path]; file != "" {
			ctx.FileAttachment(path, file)
			return
		}

		ctx.AbortWithStatus(403)
	})

	_ = router.Run(":" + port)
}

func generateLinks(files map[string]string) string {
	html := ""
	for abs, file := range files {
		html = html + fmt.Sprintf("<a href='/download?path=%v'>%v</a><br>", abs, file)
	}

	return html
}
