package apis

import (
	leveldb "imovie/leveldb"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteCache Delete Cache
func DeleteCache(c *gin.Context) {
	var id string = c.DefaultQuery("id", "index")
	datas := gin.H{
		"status":  0,
		"message": "delete cache",
	}
	cache := leveldb.GetLevel("index")
	if cache != "leveldb: not found" {
		keys := []string{id}
		leveldb.DelLevel(keys)
	}
	c.JSON(http.StatusOK, datas)
}
