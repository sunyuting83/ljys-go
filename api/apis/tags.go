package apis

import (
	"fmt"
	"net/http"
	model "newapp/database/models"
	leveldb "newapp/leveldb"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Tags 列表数据
func Tags(c *gin.Context) {
	path := c.FullPath()[5:len(c.FullPath())]
	var id string = c.DefaultQuery("id", "1")
	var page string = c.DefaultQuery("page", "1")
	ipage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		fmt.Println("err")
	}

	var datas gin.H
	tname := strings.Join([]string{"tag", id, page}, ":")
	cache := leveldb.GetLevel(tname)
	if cache == "leveldb: not found" {
		b, m := MakeClassify()
		data, name := getTags(id, ipage)
		datas = gin.H{
			"status":        0,
			"menu":          b,
			"menumore":      m,
			"matched_route": path,
			"movie":         data,
			"c_name":        name,
		}
		// leveldb.SetLevel(tname, jsonToStr(datas), 86400000)
	} else {
		datas = strToJsons(cache)
	}
	c.JSON(http.StatusOK, datas)
}

// makeMovieList make movie list for index
func getTags(id string, page int64) (data []MovieLs, name string) {
	var (
		tags model.MvPerformer
	)
	d, err := tags.TagLs(id, page)
	if err != nil {
		return data, ""
	}

	for _, item := range d.Movie {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data, d.PName
}
