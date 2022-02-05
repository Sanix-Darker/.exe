package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Code struct {
	Name  string `json:"name"`
	Lang  string `json:"lang"`
	Code  string `json:"version"`
	Args  string `json:"args"`
	Stdin string `json:"stdin"`
}

func main() {

	fmt.Println("[-] .exe started at :4321")

	router := gin.Default()
	router.POST("/execute", executeHandler)

	router.Run("localhost:4321")

}

func executeHandler(c *gin.Context) {
	var code Code

	if err := c.BindJSON(&code); err != nil {
		return
	}

	//c.IndentedJSON(http.StatusCreated, newAlbum)
}
