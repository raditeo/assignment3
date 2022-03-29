package router

import (
	"assignment3/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	valueRouter := r.Group("/value")
	{
		valueRouter.GET("/", controllers.ReadJson)
	}

	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return r
}
