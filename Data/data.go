package data

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type programmer struct {
	ID int
	Name string
	Language string
	Frontend bool
	Backend bool
	Senior bool
	Junior bool
}
	var programmers = []programmer {
		{ID: 1, Name: "John", Language: "Go", Frontend: false, Backend: true, Senior: true, Junior: false},
		{ID: 2, Name: "Steven", Language: "PHP", Frontend: true, Backend: true, Senior: false, Junior: false},
		{ID: 3, Name: "Thomas", Language: "JavaScript", Frontend: true, Backend: true, Senior: false, Junior: false},
		{ID: 4, Name: "Sarah", Language: "TypeScript", Frontend: false, Backend: true, Senior: true, Junior: false},
		{ID: 5, Name: "Gary", Language: "Java", Frontend: false, Backend: true, Senior: false, Junior: true},
		{ID: 6, Name: "Chris", Language: "C", Frontend: false, Backend: true, Senior: true, Junior: false},
	}


func GetProgrammers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, programmers)
}


func PostProgrammer(c *gin.Context) {
	var newProgrammer programmer

	if err := c.BindJSON(&newProgrammer); err != nil {
		return
	}

	programmers = append(programmers, newProgrammer)
	c.IndentedJSON(http.StatusCreated, newProgrammer)
}

func ReturnProgrammerByID(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("error during type conversion", err)
	}

	for _, a := range programmers {
		if a.ID == i {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "programmer not found!"})
}

func DeleteProgrammerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("error during type conversion", err)
	}

	for _, a := range programmers {
		if a.ID == i {
			// remember index 0
			programmers = append(programmers[:i-1], programmers[i:]...)
			return
		}
	}
}

func UpdateProgrammerByID(c *gin.Context) {
	var newProgrammer programmer
	id := c.Params.ByName("id")
	i, error := strconv.Atoi(id)

	if error != nil {
		fmt.Println("error during type conversion", error)
	}

	if err := c.BindJSON(&newProgrammer); err != nil {
		return
	}

	for _, a := range programmers {
		if a.ID == i {
			programmers[i-1] = newProgrammer
		}
	}
}