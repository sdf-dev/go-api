package data

import (
	"net/http"

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