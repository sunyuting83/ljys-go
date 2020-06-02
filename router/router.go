package router

import (
	apis "newapp/api/apis"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/index", apis.Indexs)
	}

	return router
}
