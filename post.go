package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postTodos(context *gin.Context) {
	var newTodo todo

	if error := context.BindJSON(&newTodo); error != nil {
		fmt.Println(error)
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusOK, newTodo)
}
