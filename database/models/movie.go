package models

import (
	orm "newapp/database"
	"time"
)

// MvMovie struct
type MvMovie struct {
	ID        int64     `json:"id" gorm:"primary_key, column:id"`
	Cid       int64     `json:"cid" gorm:"column:cid"`
	Title     string    `json:"title" gorm:"column:title"`
	Other     string    `json:"other" gorm:"column:other"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName change table name
// func (MvMovie) TableName() string {
// 	return "mv_movie"
// }

// Indexs 列表
func (index *MvMovie) Indexs(inid []int64) (indexs []MvMovie, err error) {
	if err = orm.Eloquent.Select("id, cid, title, other").Where("cid in (?)", inid).Order("id desc").Limit(9).Find(&indexs).Error; err != nil {
		return
	}
	return
}

// Classifys 列表
func (index *MvMovie) Classifys(id int64) (indexs []MvMovie, err error) {
	if err = orm.Eloquent.Select("id, cid, title, other").Where("cid = ?", id).Order("id desc").Limit(9).Find(&indexs).Error; err != nil {
		return
	}
	return
}

// Lists 列表
func (index *MvMovie) Lists(id, page int64) (indexs []MvMovie, err error) {
	p := makePage(page)
	if err = orm.Eloquent.Select("id, cid, title, other").Where("cid = ?", id).Order("id desc").Limit(30).Offset(p).Find(&indexs).Error; err != nil {
		return
	}
	return
}

// makePage make page
func makePage(p int64) int64 {
	p = p - 1
	if p <= 0 {
		p = 0
	}
	page := p * 30
	return page
}
