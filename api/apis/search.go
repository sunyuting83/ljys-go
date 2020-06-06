package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	model "newapp/database/models"
	leveldb "newapp/leveldb"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// HotKeys s
type HotKeys []HotKey

//Len()
func (s HotKeys) Len() int {
	return len(s)
}

// Less 排序
func (s HotKeys) Less(i, j int) bool {
	return s[i].Click > s[j].Click
}

//Swap()
func (s HotKeys) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

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
			SaveHotKey(key)
		}
	} else {
		data = SearchToJsons(cache)
		SaveHotKey(key)
	}

	c.JSON(http.StatusOK, data)
}

// GetSearchHot fun
func GetSearchHot(c *gin.Context) {
	var (
		datas  gin.H
		list   model.MvMovie
		data   []model.SearchKey
		err    error
		hotkey HotKeys
	)
	data, err = list.SearchHot()
	if err != nil {
		data = make([]model.SearchKey, 0)
	}
	if len(data) <= 0 {
		data = make([]model.SearchKey, 0)
	}
	hotkey = GetHotKey()
	if len(hotkey) <= 0 {
		hotkey = make(HotKeys, 0)
	}
	datas = gin.H{
		"status":  0,
		"hotkey":  hotkey,
		"hotlist": data,
	}
	c.JSON(http.StatusOK, datas)
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

// GetHotKey (hotkey map[]in)
func GetHotKey() HotKeys {
	var hotkey HotKeys
	cache := leveldb.GetLevel("hotkey")
	if cache == "leveldb: not found" {
		return hotkey
	}
	hotkey = HotToJsons(cache)
	if len(hotkey) >= 10 {
		hotkey = hotkey[0:10]
	}
	return hotkey
}

// HotToJsons fun
func HotToJsons(s string) HotKeys {
	var result HotKeys
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return result
	}
	return result
}

// HotToString fun
func HotToString(d HotKeys) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return result
	}
	return result
}

// SaveHotKey fun
func SaveHotKey(key string) {
	var (
		hotkey HotKeys
	)
	cache := leveldb.GetLevel("hotkey")
	if cache == "leveldb: not found" {
		hotkey = append(hotkey, HotKey{
			Key:   key,
			Click: 1,
		})
		leveldb.SetLevel("hotkey", HotToString(hotkey), -1)
	} else {
		hotkey = HotToJsons(cache)
		h := ChecKey(hotkey, key)
		if h {
			for i, item := range hotkey {
				if item.Key == key {
					fmt.Println(hotkey[i])
					hotkey[i].Click = item.Click + 1
					break
				}
			}
		} else {
			hotkey = append(hotkey, HotKey{
				Key:   key,
				Click: 1,
			})
		}
		sort.Sort(hotkey)
		leveldb.SetLevel("hotkey", HotToString(hotkey), -1)
	}
	return
}

// ChecKey check key
func ChecKey(arr HotKeys, k string) bool {
	for _, item := range arr {
		if item.Key == k {
			return true
		}
	}
	return false
}
