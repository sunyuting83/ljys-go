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
		data, name := getTags(id, ipage, path)
		if len(data) <= 0 {
			data = make([]MovieLs, 0)
		}
		datas = gin.H{
			"status":   0,
			"menu":     b,
			"menumore": m,
			"movie":    data,
			"c_name":   name,
		}
		if len(data) > 0 {
			leveldb.SetLevel(tname, jsonToStr(datas), 86400000)
		}
	} else {
		datas = strToJsons(cache)
	}
	c.JSON(http.StatusOK, datas)
}

// makeMovieList make movie list for index
func getTags(id string, page int64, path string) (data []MovieLs, name string) {
	var (
		a   model.MvArea
		dir model.MvDirector
		p   model.MvPerformer
		err error
	)
	switch path {
	case "area":
		a, err = a.TagALs(id, page)
		name = a.AName
		data = getMovieDDatag(a.Movie)
	case "director":
		dir, err = dir.TagDLs(id, page)
		name = dir.DName
		data = getMovieDDatag(dir.Movie)
	case "performer":
		p, err = p.TagPLs(id, page)
		name = p.PName
		data = getMovieDDatag(p.Movie)
	default:
		p, err = p.TagPLs(id, page)
		name = p.PName
		data = getMovieDDatag(p.Movie)
		break
	}
	if err != nil {
		return data, ""
	}
	// fmt.Println(data)

	return data, name
}

// getMovieDData get Movie d data
func getMovieDDatag(m []*model.MvMovie) []MovieLs {
	var data []MovieLs
	for _, item := range m {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data
}
