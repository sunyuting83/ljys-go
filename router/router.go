package router

import (
	. "newapp/api/apis"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/index", Indexs)

	return router
}
