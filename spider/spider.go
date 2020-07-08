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
		fmt.Println(`配置文件参数不能为空，请使用 -c 配置文件路径`)
	} else {
		list := makeList(p)
		for _, url := range list {
			fmt.Println(url)
			b, e := getData(url)
			if e {
				fmt.Println(b)
				fmt.Println(getTopID("1"))
			}
		}
	}
}

// makeMovieData
// func makeMovieData(list []MovieList) {
// 	var movie models.MvMovie
// 	for _, item := range list {

// 	}
// }

// makeList
func makeList(p int) []string {
	var rturl string = "https://www.mhapi123.com/inc/api_mac10.php?ac=detail&pg="
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

// getTopID get top id
func getTopID(id string) (gid int64) {
	switch id {
	case "5":
		gid = 12
		break
	case "6":
		gid = 13
		break
	case "7":
		gid = 14
		break
	case "8":
		gid = 15
		break
	case "9":
		gid = 16
		break
	case "10":
		gid = 18
		break
	case "11":
		gid = 17
		break
	case "12":
		gid = 5
		break
	case "13":
		gid = 6
		break
	case "14":
		gid = 9
		break
	case "15":
		gid = 8
		break
	case "16":
		gid = 7
		break
	case "17":
		gid = 10
		break
	case "18":
		gid = 11
		break
	case "19":
		gid = 21
		break
	case "20":
		gid = 20
		break
	case "21":
		gid = 19
		break
	case "22":
		gid = 22
		break
	case "23":
		gid = 23
		break
	case "24":
		gid = 24
		break
	case "25":
		gid = 25
		break
	case "26":
		gid = 26
		break
	case "27":
		gid = 27
		break
	case "28":
		gid = 28
		break
	case "29":
		gid = 29
		break
	case "30":
		gid = 30
		break
	case "31":
		gid = 31
		break
	default:
		gid = 22
		break
	}
	return
}
