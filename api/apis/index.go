package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "newapp/database/models"
	leveldb "newapp/leveldb"

	"github.com/gin-gonic/gin"
)

var classify model.MvClassify

// Indexs 列表数据
func Indexs(c *gin.Context) {
	var (
		datas  gin.H
		swiper []Swiper
		note   []string
	)
	sw := leveldb.GetLevel("swiper")
	if sw == "leveldb: not found" {
		swiper, note = MakeSwiperData()
	} else {
		nnn := leveldb.GetLevel("note")
		note = NoteToJsons(nnn)
		swiper = SwiperTojson(sw)
	}
	cache := leveldb.GetLevel("index")
	if cache == "leveldb: not found" {
		b, m := MakeClassify()
		data := makeMovieList(b)
		if len(data) <= 0 {
			data = make([]Movielist, 0)
		}
		datas = gin.H{
			"status":    0,
			"menu":      b,
			"menumore":  m,
			"movielist": data,
			"swiper":    swiper,
			"notice":    note,
		}
		if len(data) > 0 {
			leveldb.SetLevel("index", jsonToStr(datas), 86400000)
		}
	} else {
		datas = strToJsons(cache)
	}

	c.JSON(http.StatusOK, datas)
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

// strToJsons fun
func strToJsons(s string) gin.H {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return result
	}
	return result
}

// jsonToStr fun
func jsonToStr(d gin.H) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return result
	}
	return result
}

// NoteToJsons fun
func NoteToJsons(s string) (result []string) {
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		return
	}
	return
}

// SwiperTojson fun
func SwiperTojson(s string) (p []Swiper) {
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		// fmt.Println(err.Error())
		return
	}
	return
}
