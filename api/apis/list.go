package apis

import (
	"fmt"
	"net/http"
	model "newapp/database/models"
	"strconv"

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
	b, m := MakeClassify()
	cname, err := classify.GetOneClass(icid)
	data := makeList(icid, ipage)

	c.JSON(http.StatusOK, gin.H{
		"status":   0,
		"menu":     b,
		"menumore": m,
		"ctitle":   cname.CName,
		"movies":   data,
	})
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
