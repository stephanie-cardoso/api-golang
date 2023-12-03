package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/stephanie-cardoso/api-golang/docs"
	"github.com/stephanie-cardoso/api-golang/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initialize() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	handler.InitializeHandler()

	router.GET("/opening", handler.GetOpening)
	router.POST("/opening", handler.CreateOpening)
	router.DELETE("/opening", handler.DeleteOpening)
	router.PUT("/opening", handler.UpdateOpening)
	router.GET("/openings", handler.GetOpenings)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
