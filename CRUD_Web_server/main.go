package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// model
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// simulating database using in memory database
var users = []User{
	{ID: "1", Name: "Dibyajyoti Das", Gender: "Male", Age: 22},
	{ID: "2", Name: "Mishanraj Kalita", Gender: "Male", Age: 23},
}

func main() {
	r := gin.Default()

	//Routes
	r.GET("/", handleHome)
	r.GET("/users", getAllUsers)
	r.POST("/users", addUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":1020")
}

// handler Functions
// home route
func handleHome(ctx *gin.Context) {
	ctx.String(200, "use these routes - /users, /users/:id")
}

// get all books
func getAllUsers(ctx *gin.Context) {
	ctx.JSON(200, users)
}

// get single user
func getUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	for _, user := range users {
		if user.ID == userId {
			ctx.JSON(200, user)
			return
		}
	}

	ctx.String(400, "User not found...")
}

// create/Add a new user
func addUser(ctx *gin.Context) {
	var newUser User

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)
	ctx.JSON(http.StatusCreated, newUser)
}

// updating an user
func updateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	var updatedUser User

	if err := ctx.BindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == userId {
			users[i] = updatedUser
			ctx.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found..."})
}

// Deleting an User
func deleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	for i, user := range users {
		if user.ID == userId {
			users = append(users[:i], users[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"Message": "User Deleted Successfully..."})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found..."})
}

// check nil error
// func checkNilError(err error, ctx *gin.Context) {
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// }
