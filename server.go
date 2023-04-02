package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hotbrandon/go-gin-tutorial/controller"
	"github.com/hotbrandon/go-gin-tutorial/service"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)
)

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "alive")
}

func parseProcedure(c *gin.Context) {
	procedureName := c.Param("procedure")

	c.String(200, procedureName)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", index)

	router.GET("/parser/p/:procedure", parseProcedure)

	router.GET("/videos", func(ctx *gin.Context) {
		videos := videoController.GetAll()
		ctx.JSON(200, videos)
	})

	router.POST("/videos", func(ctx *gin.Context) {
		video := videoController.Save(ctx)
		ctx.JSON(200, video)
	})

	router.GET("/health", healthCheck)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}
