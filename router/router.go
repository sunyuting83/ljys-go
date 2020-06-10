package router

import (
	apis "newapp/api/apis"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.Use(CORSMiddleware())
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
	}

	return router
}

// CORSMiddleware cors middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
