package apis

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	model "newapp/database/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Movie 列表数据
func Movie(c *gin.Context) {
	var id string = c.DefaultQuery("id", "1")
	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		intid = 100
	}
	b, m := MakeClassify()
	data := makeMovie(intid)
	hot := getHotLists(data.Cid)
	newd := makeMovieDatas(data.Other)

	c.JSON(http.StatusOK, gin.H{
		"status":    0,
		"menu":      b,
		"menumore":  m,
		"id":        data.ID,
		"title":     data.Title,
		"cid":       data.Cid,
		"director":  data.Director,
		"performer": data.Performer,
		"area":      data.Area,
		"img":       newd["img"],
		"reamarks":  newd["remarks"],
		"play_path": newd["play_path"],
		"languarge": newd["languarge"],
		"profiles":  newd["profiles"],
		"year":      newd["year"],
		"score":     newd["score"],
		"hotlist":   hot,
	})
}

// makeMovie make movie data
func makeMovie(id int64) model.Movie {
	var movie model.Movie
	s, err := movie.Movies(id)
	if err != nil {
		fmt.Println("err")
	}
	return s
}

// getHotLists get hot list
func getHotLists(cid int64) []MovieLs {
	var (
		hotlist model.MvMovie
		data    []MovieLs
	)
	d, err := hotlist.HotLists(cid)
	if err != nil {
		fmt.Println("err")
	}
	for _, item := range d {
		p := strTojson(item.Other)
		data = append(data, MovieLs{ID: item.ID, Title: item.Title, Img: p.Img, Score: p.Score, Remarks: p.Remarks})
	}
	return data
}

// makeMovieData fun
func makeMovieDatas(s string) map[string]interface{} {
	var (
		hlen int = 1
		plen int = 1
	)
	p := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		return p
	}
	hls := p["play_path"].(map[string]interface{})["hls"].([]interface{})
	pla := p["play_path"].(map[string]interface{})["player"].([]interface{})
	hlen = len(hls)
	plen = len(pla)
	p["play_path"].(map[string]interface{})["more"] = false
	p["play_path"].(map[string]interface{})["count"] = hlen
	if hlen >= 50 {
		p["play_path"].(map[string]interface{})["more"] = true
		p["play_path"].(map[string]interface{})["hls"] = makeTowArr(hls, hlen)
	}
	if plen >= 50 {
		p["play_path"].(map[string]interface{})["more"] = true
		p["play_path"].(map[string]interface{})["player"] = makeTowArr(pla, plen)
	}
	return p
}

// makeTowArr make tow arr
func makeTowArr(arr []interface{}, c int) []interface{} {
	// 先定义一个map数组。长度为总数/20(子数组的长度)，向上取整
	var ar []interface{} = make([]interface{}, int(math.Ceil(float64(c)/20.0)))
	var x int = -1 //定义循环原始数组起始index
	for i := 0; i < int(math.Ceil(float64(c)/20.0)); i++ {
		// 定义子map数组。长度子数组长度20
		var tmp []interface{} = make([]interface{}, 20)
		var u int = 20                              //定义子数组循环长度20
		if i == int(math.Ceil(float64(c)/20.0))-1 { //如果是最后一个子数组
			// 最后一个子数组长度 = 总长度 - 总长度/20向下取整 × 子数组循环长度20
			tmp = make([]interface{}, c-(int(math.Floor(float64(c)/20.0))*20))
			// 最后一个子数组循环长度 = 总长度 - 总长度/20向下取整 × 子数组循环长度20
			u = c - (int(math.Floor(float64(c)/20.0)) * 20)
		}
		for j := 0; j < u; j++ { //循环子数组长度20
			x++ // 原始数组起始值递增
			if x <= c-1 {
				tmp[j] = arr[x] // 给子数组赋值 j子数组下标 x原始数组下标
			}
		}
		ar[i] = tmp // 给父数组赋值 i父数组下标
	}
	return ar
}
