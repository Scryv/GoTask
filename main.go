package main

import (
	_ "errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct { //main task struct
	ID   string `json:"id"`
	Task string `json:"task"`
	Date string `json:"date"`
	Done bool   `json:"done"`
}

var tasks = []task{ //just some values to get it started will later do csv or sm
	{ID: "1", Task: "Clean room", Date: "02-04-26", Done: false},
	{ID: "2", Task: "Read book", Date: "01-05-26", Done: false},
	{ID: "3", Task: "Record", Date: "29-08-26", Done: false},
}

func getTodos(c *gin.Context) { //c = gincontext so gives gin JSON context
	c.IndentedJSON(http.StatusOK, tasks) //gets task and transform JSON
} //status okays the tasks
func main() {
	router := gin.Default()        //default gin server
	router.GET("/todos", getTodos) //on localhost/todos runs func getTodos
	router.Run("localhost:8080")   //server on 8080
}
