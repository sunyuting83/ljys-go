package router

import (
	apis "imovie/api/apis"
	utils "imovie/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.Use(utils.CORSMiddleware())
	{
		api.GET("/index", apis.Indexs)
		api.GET("/getclass", apis.ClassLists)
		api.GET("/list", apis.GetLists)
		api.GET("/getmovie", apis.Movie)
		api.GET("/area", apis.Tags)
		api.GET("/director", apis.Tags)
		api.GET("/performer", apis.Tags)
		api.GET("/getkey", apis.GetSearchKey)
		api.GET("/search", apis.GetSearch)
		api.GET("/gethot", apis.GetSearchHot)
		api.GET("/180nm3ysib84", apis.DeleteCache)
	}

	return router
}
