package apis

import (
	"fmt"
	model "imovie/database/models"
	leveldb "imovie/leveldb"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ClassLists 列表数据
func ClassLists(c *gin.Context) {
	var id string = c.DefaultQuery("id", "1")
	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("err")
	}
	var datas gin.H
	cname := strings.Join([]string{"class", id}, ":")
	cache := leveldb.GetLevel(cname)
	if cache == "leveldb: not found" {
		b, m := MakeClassify()
		data := makeClassList(intid)
		datas = gin.H{
			"status":    0,
			"menu":      b,
			"menumore":  m,
			"movielist": data,
		}
		if len(data) > 0 {
			leveldb.SetLevel(cname, jsonToStr(datas), 86400000)
		}
	} else {
		datas = strToJsons(cache)
	}
	c.JSON(http.StatusOK, datas)
}

// makeMovieList make movie list for index
func makeClassList(id int64) []Classlist {
	var d []Classlist
	s, err := classify.GetClass(id)
	if err != nil {
		fmt.Println("err")
	}
	for _, item := range s {
		movie := makeClassData(item.ID)
		d = append(d, Classlist{ID: item.ID, CName: item.CName, Movie: movie})
	}
	return d
}

// makeMovieData make movie data
func makeClassData(i int64) []MovieLs {
	var (
		index model.MvMovie
		data  []MovieLs
	)
	d, err := index.Classifys(i)

	if len(d) <= 0 {
		data = make([]MovieLs, 0)
	}
	if err != nil {
		return data
	}
	for _, item := range d {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data
}
