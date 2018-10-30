package backend

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTodoData is the struct for the
type CreateTodoData struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed" binding:"exists"`
}

func createTodo(c *gin.Context) {
	var data CreateTodoData
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println(err)
	}
	todo := todoModel{Title: data.Title, Completed: data.Completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "id": todo.ID})
}

func fetchAllTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		_todos = append(_todos, transformedTodo{
			ID: item.ID, Title: item.Title, Completed: item.Completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "todos": _todos})
}
