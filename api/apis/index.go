package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "newapp/database/models"

	"github.com/gin-gonic/gin"
)

// Movielist index movie list data
type Movielist struct {
	ID         int64        `json:"id"`
	CName      string       `json:"c_name"`
	Smallclass []MovieSmall `json:"smallclass"`
	Movie      []MovieLs    `json:"movie"`
}

// MovieSmall small classify
type MovieSmall struct {
	ID    int64  `json:"id"`
	CName string `json:"c_name"`
}

// MovieLs movie data
type MovieLs struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Img     string `json:"img"`
	Remarks string `json:"remarks"`
	Score   string `json:"score"`
}

var (
	classify model.MvClassify
	// index model.MvMovie
)

// Indexs 列表数据
func Indexs(c *gin.Context) {

	allclass, err := classify.BigClass()
	b, m := makeClassify(allclass)
	data := makeMovieList(b)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    0,
		"menu":      b,
		"menumore":  m,
		"movielist": data,
	})
}

// makeClassify make classify list
func makeClassify(c []model.MvClassify) ([]model.MvClassify, []model.MvClassify) {
	var (
		b []model.MvClassify
		s []model.MvClassify
	)
	for _, item := range c {
		if item.TopID == 0 {
			b = append(b, item)
		} else {
			s = append(s, item)
		}
	}
	return b, s
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
		p     MovieLs
	)
	d, err := index.Indexs(i)
	if err != nil {
		fmt.Println("err")
	}
	for _, item := range d {
		var other []byte = []byte(item.Other)
		if err := json.Unmarshal(other, &p); err != nil {
			// fmt.Println(err.Error())
			fmt.Println("err")
		}
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data
}
