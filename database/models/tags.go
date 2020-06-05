package models

import (
	orm "newapp/database"
	"strconv"
)

// TagLs 列表
func (tag *MvPerformer) TagLs(id string, page int64) (tags MvPerformer, err error) {
	// orm.Eloquent.Debug().First(&tags, id)
	// if err = orm.Eloquent.
	// 	Debug().
	// 	Model(&tags).
	// 	Association("Movie").
	// 	Find(&tags.Movie).Error; err != nil {
	// 	return
	// }
	// orm.Eloquent.LogMode(true)
	// if err = orm.Eloquent.
	// 	Debug().
	// 	Preload("Movie").
	// 	First(&tags, id).Error; err != nil {
	// 	return
	// }
	p := makePageS(page)
	orm.Eloquent.First(&tags, id)
	sql := `SELECT * FROM "mv_movie" INNER JOIN "mv_movie_mv_performer" ON "mv_movie_mv_performer"."mv_movie_id" = "mv_movie"."id" WHERE ("mv_movie_mv_performer"."mv_performer_id" IN ("` + id + `")) ` + ` ORDER BY "mv_movie"."id" DESC Limit 30 OFFSET ` + p
	if err = orm.Eloquent.Raw(sql).Scan(&tags.Movie).Error; err != nil {
		return
	}
	// fmt.Println(&tags)
	return
}

// makePage make page
func makePageS(p int64) string {
	p = p - 1
	if p <= 0 {
		p = 0
	}
	page := p * 30
	ipage := strconv.FormatInt(page, 10)
	return ipage
}
