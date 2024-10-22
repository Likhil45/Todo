package main

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	Id     string `json:"id"`
	Item   string `json:"item"`
	Status bool   `json:"status"`
}

var todos = []todo{
	{Id: "1", Item: "clean room", Status: false},
	{Id: "2", Item: "Play football", Status: true},
	{Id: "3", Item: "Play Cricket", Status: false},
}

func main() {
	router := gin.Default() //Create server
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodo)
	router.GET("/todos/item/:item", getItem)
	router.POST("/todos", postTodos)
	router.Run("localhost:9080") //Run server
}
