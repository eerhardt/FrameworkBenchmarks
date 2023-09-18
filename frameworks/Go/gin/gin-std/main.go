package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id         uint32    `json:"id"`
	Title      string    `json:"title"`
	DueBy      time.Time `json:"dueBy"`
	IsComplete bool      `json:"isComplete"`
}

var todos = []Todo{
	{Id: 1, Title: "Walk the dog"},
	{Id: 2, Title: "Do the dishes", DueBy: time.Now()},
	{Id: 3, Title: "Do the laundry", DueBy: time.Now().Add(time.Hour * 24)},
	{Id: 4, Title: "Clean the bathroom"},
	{Id: 5, Title: "Clean the car", DueBy: time.Now().Add(time.Hour * 48)},
}

// / Test 1: JSON serialization
func json(c *gin.Context) {
	c.JSON(200, todos)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/json", json)
	fmt.Println("Listening and serving HTTP")
	r.Run(":8080")
}
