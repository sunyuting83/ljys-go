package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetDBPath get data file path
func GetDBPath(t string, d bool) string {
	var (
		p   string
		dir string
	)
	if d {
		dir = "/home/sun/Works/go/src/newapp"
	} else {
		path, err := os.Executable()
		if err != nil {
		}
		dir = filepath.Dir(path)
	}

	p = "movie.sqlite"
	if t == "level" {
		p = "Cache"
	}
	if t == "spider" {
		p = "IgnoreDB"
	}
	return strings.Join([]string{dir, p}, "/")
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
