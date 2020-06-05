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
		api.GET("/getclass", apis.ClassLists)
		api.GET("/list", apis.GetLists)
		api.GET("/getmovie", apis.Movie)
		api.GET("/area", apis.Tags)
		api.GET("/director", apis.Tags)
		api.GET("/performer", apis.Tags)
	}

	return router
}
