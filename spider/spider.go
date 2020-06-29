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
	var p int
	flag.IntVar(&p, "p", 5, "分页，默认为50")
	flag.Parse()
	list := makeList(p)
	for _, url := range list {
		fmt.Println(url)
		b, e := getData(url)
		if e {
			fmt.Println(b)
		}
	}
}

// makeList
func makeList(p int) []string {
	var rturl string = "http://cj.wlzy.tv/api/macs/vod/?ac=detail&pg="
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
