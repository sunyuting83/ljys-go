package apis

import (
	"encoding/json"
	"net/http"
	model "newapp/database/models"
	leveldb "newapp/leveldb"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetSearchKey 列表数据
func GetSearchKey(c *gin.Context) {
	var (
		key  string = c.DefaultQuery("key", "向西")
		list model.MvMovie
		data []model.SearchKey
		err  error
	)
	sname := strings.Join([]string{"searchkey", key}, ":")
	cache := leveldb.GetLevel(sname)
	if cache == "leveldb: not found" {
		data, err = list.SearchKey(key)
		if err != nil {
			data = make([]model.SearchKey, 0)
		}
		if len(data) <= 0 {
			data = make([]model.SearchKey, 0)
		} else {
			leveldb.SetLevel(sname, SearchKeyToStr(data), 21600000)
		}
	} else {
		data = SearchKeyToJsons(cache)
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
	if err != nil {
		ipage = 1
	}
	sname := strings.Join([]string{"search", key, page}, ":")
	cache := leveldb.GetLevel(sname)
	if cache == "leveldb: not found" {
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
		if len(data) > 0 {
			leveldb.SetLevel(sname, SearchToStr(data), 21600000)
		}
	} else {
		data = SearchToJsons(cache)
	}

	c.JSON(http.StatusOK, data)
}

// SearchKeyToJsons fun
func SearchKeyToJsons(s string) []model.SearchKey {
	var result []model.SearchKey
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return result
	}
	return result
}

// SearchKeyToStr fun
func SearchKeyToStr(d []model.SearchKey) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return result
	}
	return result
}

// SearchToJsons fun
func SearchToJsons(s string) []MovieLs {
	var result []MovieLs
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return result
	}
	return result
}

// SearchToStr fun
func SearchToStr(d []MovieLs) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return result
	}
	return result
}
