package main

import (
	"flag"
	orm "newapp/database"
	leveldb "newapp/leveldb"
	"newapp/router"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	defer orm.Eloquent.Close()

	defer leveldb.Leveldb.Close()
	router := router.InitRouter()
	router.Run(strings.Join([]string{":", port}, ""))
}
