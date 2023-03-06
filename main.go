package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})

		user.GET("/:name/:action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
	}

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.String(200, id+page+name+message)
	})
	return router
}

func main() {
	r := setupRouter()
	r.Run(":8081")
}
