package apis

import (
	"net/http"
	model "newapp/database/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSearchKey 列表数据
func GetSearchKey(c *gin.Context) {
	var (
		key  string = c.DefaultQuery("key", "向西")
		list model.MvMovie
	)
	data, err := list.SearchKey(key)
	if err != nil {
		data = make([]model.SearchKey, 0)
	}
	if len(data) <= 0 {
		data = make([]model.SearchKey, 0)
	}

	c.JSON(http.StatusOK, data)
}

// GetSearch 列表数据
func GetSearch(c *gin.Context) {
	var (
		key  string = c.DefaultQuery("word", "向西")
		page string = c.DefaultQuery("page", "1")
		list model.MvMovie
		data []MovieLs
	)
	ipage, err := strconv.ParseInt(page, 10, 64)
	d, err := list.Search(key, ipage)
	if err != nil {
		data = make([]MovieLs, 0)
	}
	if len(data) <= 0 {
		data = make([]MovieLs, 0)
	}
	for _, item := range d {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}

	c.JSON(http.StatusOK, data)
}
