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

// GetLists 列表数据
func GetLists(c *gin.Context) {
	var cid string = c.DefaultQuery("cid", "1")
	var page string = c.DefaultQuery("page", "1")
	icid, err := strconv.ParseInt(cid, 10, 64)
	ipage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		fmt.Println("err")
	}

	var datas gin.H
	lname := strings.Join([]string{"list", cid, page}, ":")
	cache := leveldb.GetLevel(lname)
	if cache == "leveldb: not found" {
		b, m := MakeClassify()
		cname, err := classify.GetOneClass(icid)
		if err != nil {
			fmt.Println("err")
		}
		data := makeList(icid, ipage)
		if len(data) <= 0 {
			data = make([]MovieLs, 0)
		}
		datas = gin.H{
			"status":   0,
			"menu":     b,
			"menumore": m,
			"ctitle":   cname.CName,
			"movies":   data,
		}
		if len(data) > 0 {
			leveldb.SetLevel(lname, jsonToStr(datas), 86400000)
		}
	} else {
		datas = strToJsons(cache)
	}
	c.JSON(http.StatusOK, datas)
}

// makeMovieList make movie list for index
func makeList(id, page int64) []MovieLs {
	var (
		list model.MvMovie
		data []MovieLs
	)
	d, err := list.Lists(id, page)
	if err != nil {
		return data
	}
	for _, item := range d {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data
}
