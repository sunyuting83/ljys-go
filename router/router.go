package router

import (
	. "newapp/api/apis"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/index", Indexs)

	router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	return router
}
