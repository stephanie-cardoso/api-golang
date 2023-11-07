package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	router.GET("/opening", getOpening)
	router.POST("/opening", newOpening)
	router.DELETE("/opening", deleteOpening)
	router.PUT("/opening", updateOpening)
	router.GET("/openings", getOpenings)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get opening",
	})
}

func newOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "post opening",
	})
}

func deleteOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete opening",
	})
}

func updateOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "put opening",
	})
}

func getOpenings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get openings",
	})
}
