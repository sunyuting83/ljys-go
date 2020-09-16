package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// OldData Old Data
type OldData struct {
	Code int         `json:"code"`
	List []MovieList `json:"list"`
}

// MovieList Movie List
type MovieList struct {
	VodID       string `json:"vod_id"`
	VodName     string `json:"vod_name"`
	TypeID      string `json:"type_id"`
	VodEn       string `json:"vod_en"`
	VodPic      string `json:"vod_pic"`
	VodPlayFrom string `json:"vod_play_from"`
	VodDownFrom string `json:"vod_down_from"`
	VodPlayurl  string `json:"vod_play_url"`
	VodContent  string `json:"vod_content"`
	VodYear     string `json:"vod_year"`
	VodArea     string `json:"vod_area"`
	VodLang     string `json:"vod_lang"`
	TypeScore   string `json:"vod_score"`
	VodDuration string `json:"vod_duration"`
	VodRemarks  string `json:"vod_remarks"`
}

// ConfigFile Config File
type ConfigFile struct {
	URI  string       `json:"uri"`
	ZYFL bool         `json:"zyfl"`
	ZYID string       `json:"zyid"`
	List []ConfigList `json:"list"`
	AREA []ConfigArea `json:"area"`
}

// ConfigList Config List
type ConfigList struct {
	ID  string `json:"id"`
	SID int64  `json:"sid"`
}

// ConfigArea config area
type ConfigArea struct {
	NAME string `json:"name"`
	ID   int64  `json:"id"`
}

// PlayerList player list
type PlayerList struct {
	HLS    []PlayerContent `json:"hls"`
	Player []PlayerContent `json:"player"`
}

// PlayerContent player content
type PlayerContent struct {
	NAME string `json:"name"`
	Path string `json:"path"`
}

// main
func main() {
	var (
		p int
		c string
	)
	flag.IntVar(&p, "p", 5, "分页，默认为5")
	flag.StringVar(&c, "c", "", "分页，默认为5")
	flag.Parse()
	if len(c) <= 0 {
		fmt.Println("\x1B[31m配置文件参数不能为空，请使用 -c 配置文件路径\x1B[0m")
	} else {
		data, err := ioutil.ReadFile(c)
		if err != nil {
			fmt.Println("\x1B[31m错误：配置文件路径错误\x1B[0m")
			return
		}
		config := ConfigTojson(data)
		// 先读取配置文件 传入url获取到列表
		list := makeList(p, config.URI)
		for _, url := range list {
			fmt.Println(url)
			b, e := getData(url)
			if e {
				// fmt.Println(b)
				MakeData(b, config.List, config.ZYFL, config.ZYID, config.AREA)
			}
		}
	}
}

// ConfigTojson fun
func ConfigTojson(s []byte) (p ConfigFile) {
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		// fmt.Println(err.Error())
		return
	}
	return
}

// makeMovieData
// func makeMovieData(list []MovieList) {
// 	var movie models.MvMovie
// 	for _, item := range list {

// 	}
// }

// makeList
func makeList(p int, url string) []string {
	var rturl string = url
	var listurls []string
	if p <= 0 {
		p = 1
	}
	for i := p; i >= 1; i-- {
		si := strconv.Itoa(i)
		urllist := strings.Join([]string{rturl, si}, "")
		listurls = append(listurls, urllist)
	}
	// console.log(listurls);
	return listurls
}

// getData
func getData(u string) ([]MovieList, bool) {
	var result OldData
	client := &http.Client{}                   // 新建一个http请求
	req, err := http.NewRequest("GET", u, nil) //创建请求头
	if err != nil {
		return result.List, false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36") // 加入模拟终端
	// 发起请求
	res, err := client.Do(req)
	if err != nil {
		return result.List, false
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return result.List, false
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result.List, false
	}
	errs := json.Unmarshal([]byte(string(body)), &result)
	if errs != nil {
		fmt.Println(errs)
		return result.List, false
	}
	return result.List, true
}

// MakeData make data
func MakeData(b []MovieList, l []ConfigList, z bool, id string, areas []ConfigArea) {
	for _, item := range b {
		classifyid := getTopID(item.TypeID, l, z, item.VodArea, id, areas) //传入id对应获取到分类id
		fmt.Println(classifyid)
		// fmt.Println(item)
		player := makePlayer(item.VodPlayurl)
		fmt.Println(player)
		// VodPlayurl $分割文字与播放地址 #分割多集
	}
	return
}

// getTopID get top id
func getTopID(id string, l []ConfigList, z bool, area string, i string, areas []ConfigArea) (gid int64) {
	if z {
		for _, item := range l {
			if id == i {
				return getFID(area, areas)
			} else if item.ID == id {
				return item.SID
			}
		}
	} else {
		for _, item := range l {
			if item.ID == id {
				return item.SID
			}
		}
	}
	return
}

// getFID get f id
func getFID(area string, areas []ConfigArea) (gid int64) {
	for _, item := range areas {
		if area == item.NAME {
			return item.ID
		}
		return areas[0].ID
	}
	return
}

// Make Player
func makePlayer(p string) (player PlayerList) {
	var (
		data PlayerList
	)
	if strings.Contains(p, "$") {
		list := strings.Split(p, "#")
		for _, item := range list {
			namep := strings.Split(item, "$")
			blen := len(namep[1])
			b := strings.LastIndex(namep[1], ".")
			types := namep[1][b+1 : blen]
			if types == "m3u8" {
				data.HLS = append(data.HLS, PlayerContent{NAME: namep[0], Path: namep[1]})
			} else {
				data.Player = append(data.Player, PlayerContent{NAME: namep[0], Path: namep[1]})
			}
		}
	}
	return data
}
