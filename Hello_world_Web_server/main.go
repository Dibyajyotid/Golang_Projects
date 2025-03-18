package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//setting up the router
	r := gin.Default()

	//Defining a simple route
	r.GET("/", func(ctx *gin.Context) {

		//displaying as a JSON
		// ctx.JSON(http.StatusOK, gin.H{"Message": "Hello World! This is my first Simple Web server..."})

		//displaying directly on the page
		ctx.String(200, "Hello World! This is my first Simple Web server...")
	})

	//starting the server on port
	r.Run(":1000")
}
