package main

import (
	"fmt"
	"net/http"
	// "github.com/gin-gonic/gin"
)

//using Gin Framework
// func main() {
// 	//setting up the router
// 	r := gin.Default()

// 	//Defining a simple route
// 	r.GET("/", func(ctx *gin.Context) {

// 		//displaying as a JSON
// 		// ctx.JSON(http.StatusOK, gin.H{"Message": "Hello World! This is my first Simple Web server..."})

// 		//displaying directly on the page
// 		ctx.String(200, "Hello World! This is my first Simple Web server...")
// 	})

// 	//starting the server on port
// 	r.Run(":1000")
// }

// using http
func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on port http://localhost:1000")

	//listening on server
	http.ListenAndServe(":1000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! This is my first Simple Web server...")
}
