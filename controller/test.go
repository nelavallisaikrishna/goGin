package controller

import (
	"fmt"
	"io/ioutil"
	"oms/db"

	"github.com/gin-gonic/gin"
)


func TestFuncs(c *gin.Context) {
	query := make(map[string]interface{})
	comm_top, err := db.ReadData(query)
	if err != nil {
				fmt.Println("------", err)
	}
	fmt.Println(comm_top)
	 c.JSON(200, gin.H{"message": "Hello, World!"})
}

func TestPost(c *gin.Context) {
	body := c.Request.Body
	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------bidy----", data)
	 c.JSON(200, gin.H{"message": "Hello, POST!"})
}

func TestQuery(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	 c.JSON(200, gin.H{"name": name, "age": age})
}

func TestParam(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	 c.JSON(200, gin.H{"name": name, "age": age})
}