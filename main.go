package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
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
