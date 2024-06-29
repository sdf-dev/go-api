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
	router.Run("localhost:8080")
}