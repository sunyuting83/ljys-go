package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "newapp/database/models"

	"github.com/gin-gonic/gin"
)

var classify model.MvClassify

// Indexs 列表数据
func Indexs(c *gin.Context) {

	b, m := MakeClassify()
	data := makeMovieList(b)

	c.JSON(http.StatusOK, gin.H{
		"status":    0,
		"menu":      b,
		"menumore":  m,
		"movielist": data,
	})
}

// makeMovieList make movie list for index
func makeMovieList(s []model.MvClassify) []Movielist {
	var d []Movielist
	for _, item := range s {
		smallclass, err := classify.SmallClass(item.ID)
		if err != nil {
			fmt.Println("err")
		}
		small, inid := makeSmallFun(smallclass)
		movie := makeMovieData(inid)
		ls := Movielist{
			ID:         item.ID,
			CName:      item.CName,
			Smallclass: small,
			Movie:      movie,
		}
		d = append(d, ls)
	}
	return d
}

// makeSmallFun make small fun
func makeSmallFun(s []model.MvClassify) ([]MovieSmall, []int64) {
	var sm []MovieSmall
	inid := []int64{}
	for _, item := range s {
		sm = append(sm, MovieSmall{item.ID, item.CName})
		inid = append(inid, item.ID)
	}
	return sm, inid
}

// makeMovieData make movie data
func makeMovieData(i []int64) []MovieLs {
	var (
		index model.MvMovie
		data  []MovieLs
	)
	d, err := index.Indexs(i)
	if err != nil {
		fmt.Println("err")
	}
	for _, item := range d {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data
}

// strTojson fun
func strTojson(s string) MovieLs {
	var p MovieLs
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		// fmt.Println(err.Error())
		return p
	}
	return p
}
