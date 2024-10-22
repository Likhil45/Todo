package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getItem(context *gin.Context) {
	item := context.Param("item")
	todo, error := getItems(item)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To do not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)

}

func getItems(item string) (*todo, error) {

	for i, t := range todos {
		if t.Item == item {
			return &todos[i], nil
		}
	}
	return nil, errors.New("item not found")

}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func toggleTodo(context *gin.Context) {
	id := context.Param("id")
	todo, error := getTodoById(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	todo.Status = !todo.Status
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, error := getTodoById(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
