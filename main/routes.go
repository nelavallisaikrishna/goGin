package main

import test "oms/controller"

func InitializeRoutes() {
	router.GET("/", test.TestFuncs)
	router.POST("/", test.TestPost)
	router.GET("/query", test.TestQuery)
	router.GET("/query/:name/:age", test.TestParam)
}