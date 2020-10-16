package apis

import (
	"encoding/json"
	"fmt"
	leveldb "imovie/leveldb"
)

// MakeSwiperData make swiper data
func MakeSwiperData() ([]Swiper, []string) {
	var swiper []Swiper
	swiper = append(swiper, Swiper{ID: "33571", Title: "2019科幻动作《终结者：黑暗命运》BD720P&BD1080P.英语中英双字", Img: "https://ae01.alicdn.com/kf/U9849b6d178b843bbada26bdec412f073g.jpg"}, Swiper{ID: "33340", Title: "2019高分韩剧《流浪者/浪客行》16集全.HD1080P.国韩双语中字", Img: "https://ae01.alicdn.com/kf/Uc9d893747c024afa960296414ef7476ev.jpg"}, Swiper{ID: "40845", Title: "2019科幻冒险《星球大战9：天行者崛起》BD720P&BD1080P.国英双语中英双字", Img: "https://ae01.alicdn.com/kf/Uef16413b321244e4a99519450c8138d6A.jpg"}, Swiper{ID: "41959", Title: "2020韩剧《王国第二季》6集全.HD1080P.韩语中字", Img: "https://ae01.alicdn.com/kf/U50a32afe33fe4ef8a727314dddc40f89Q.jpg"}, Swiper{ID: "30720", Title: "2019爱情犯罪《少年的你》HD1080P.国语中字", Img: "https://ae01.alicdn.com/kf/U49e5e2811edb4d08a07b63b6ec6081fdG.jpg"})
	sByte, errError := json.Marshal(swiper)
	if errError != nil {
		fmt.Println("err")
	}
	leveldb.SetLevel("swiper", string(sByte), -1)

	note := []string{}
	note = append(note, "后端GO版本全新上线")
	note = append(note, "千万不要相信视频中的广告")
	nByte, errnError := json.Marshal(note)
	if errnError != nil {
		fmt.Println("err")
	}
	leveldb.SetLevel("note", string(nByte), -1)
	return swiper, note
}
