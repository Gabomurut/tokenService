package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type tokens struct {
	First  string `json:"first,omitempty"`
	Second string `json:"second,omitempty"`
	Third  string `json:"third,omitempty"`
}

func main() {
	router := gin.Default()
	router.POST("/service", getStrings)

	err := router.Run()
	if err != nil {
		return
	}
}


func getStrings(c *gin.Context) {

	var tokens tokens
	 c.BindJSON(&tokens)

	firstToken := tokens.First
	secondToken := tokens.Second
	thirdToken := tokens.Third

	response := firstToken + secondToken + thirdToken


	c.IndentedJSON(http.StatusOK, response)
}
