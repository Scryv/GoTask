package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
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

func getIdTask(id string) (*task, int, error) { //outputs a struct and err and int for clarification for deletion
	for i, t := range tasks { //task loop index and value
		if t.ID == id { //if value == id that u got taken in
			return &tasks[i], i, nil //will return the line with that id
		}
	}
	return nil, -1, errors.New("No task with that id") //error no task -1 isnt a thigns so auto err
}
func getTheIdOfTask(c *gin.Context) {
	id := c.Param("id")           //sais id is the id param from /:id
	task, _, err := getIdTask(id) //uses task id
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Task with this id is not found"})
		return //if there is error it will status not found and custom message
	}
	c.IndentedJSON(http.StatusOK, task) //if everything good status okay and sends task

}
func deleteTask(c *gin.Context) {
	id := c.Param("id")            //takes id param again
	_, index, err := getIdTask(id) //since i dont need the book i just need to know if error or nah
	if err != nil {                //error handling
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Task with this id doesnt exist"})
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...) //deletes the json piece con to id
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "The task has been deleted succesfully"})

}
func getTodos(c *gin.Context) { //c = gincontext so gives gin JSON context
	c.IndentedJSON(http.StatusOK, tasks) //gets task and transform JSON
} //status okays the tasks
func main() {
	router := gin.Default()                        //default gin server
	router.GET("/todos", getTodos)                 //on localhost/todos runs func getTodos
	router.GET("/todos/:id", getTheIdOfTask)       //gets task by id
	router.DELETE("/todos/delete/:id", deleteTask) //deletes task
	router.Run("localhost:8080")                   //server on 8080
}
