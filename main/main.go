package main

import (
	"fmt"
	"oms/config"
	"oms/db"
	"os"

	// test "oms/controller"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
		fmt.Println("Loading " + os.Getenv("env") + "config...")
		config.SetConfigVars()
		fmt.Println("Connecting to databases...")
		port := ":" + os.Getenv("port")
		db.Connect()
		router = gin.Default()
		InitializeRoutes()
		fmt.Println(" Server running at port : ", port)
		router.Run(port)
}