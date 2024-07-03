package routing

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sdf-dev/go-api/Data"
)

func Routing() {
	fmt.Println("Hello, World!")
	router := gin.Default()
	router.GET("/programmers", data.GetProgrammers)
	router.POST("/programmers", data.PostProgrammer)
	router.GET("/programmers/:id", data.ReturnProgrammerByID) 
	router.Run("localhost:8080")
}