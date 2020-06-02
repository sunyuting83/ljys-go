package main

import (
	orm "newapp/database"
	"newapp/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8456")
}
