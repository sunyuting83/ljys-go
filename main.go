package main

import (
	orm "newapp/database"
	leveldb "newapp/leveldb"
	"newapp/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	defer orm.Eloquent.Close()

	defer leveldb.Leveldb.Close()
	router := router.InitRouter()
	router.Run(":8456")
}
