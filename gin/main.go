package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		resp := struct {
			Message string
		}{
			Message: "Hello World",
		}

		c.JSON(http.StatusOK, resp)
	})

	r.Run()
}
