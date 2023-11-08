package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stephanie-cardoso/api-golang/handler"
)

func Initialize() {
	router := gin.Default()
	handler.InitializeHandler()

	router.GET("/opening", handler.GetOpening)
	router.POST("/opening", handler.NewOpening)
	router.DELETE("/opening", handler.DeleteOpening)
	router.PUT("/opening", handler.UpdateOpening)
	router.GET("/openings", handler.GetOpenings)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
