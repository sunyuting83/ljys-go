package apis

import (
	"fmt"
	"net/http"
	model "newapp/database/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ClassLists 列表数据
func ClassLists(c *gin.Context) {
	var id string = c.DefaultQuery("id", "1")
	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("err")
	}
	b, m := MakeClassify()
	data := makeClassList(intid)

	c.JSON(http.StatusOK, gin.H{
		"status":    0,
		"menu":      b,
		"menumore":  m,
		"movielist": data,
	})
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
	if err != nil {
		fmt.Println("err")
	}
	for _, item := range d {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data
}
