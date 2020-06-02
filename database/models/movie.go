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
func (index *MvMovie) Indexs() (indexs []MvMovie, err error) {
	if err = orm.Eloquent.Select("id, cid, title, other, created_at").Order("id desc").Limit(3).Find(&indexs).Error; err != nil {
		return
	}
	return
}
