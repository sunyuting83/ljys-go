package models

import (
	orm "newapp/database"
)

// MvClassify struct
type MvClassify struct {
	ID    int64  `json:"id" gorm:"primary_key, column:id"`
	TopID int64  `json:"top_id" gorm:"column:top_id"`
	CName string `json:"c_name" gorm:"column:c_name"`
}

// BigClass 列表
func (classify *MvClassify) BigClass() (classifys []MvClassify, err error) {
	if err = orm.Eloquent.Select("id, top_id, c_name").Find(&classifys).Error; err != nil {
		return
	}
	return
}

// SmallClass 列表
func (classify *MvClassify) SmallClass(id int64) (classifys []MvClassify, err error) {
	if err = orm.Eloquent.Where("top_id = ?", id).Limit(4).Select("id, top_id, c_name").Find(&classifys).Error; err != nil {
		return
	}
	return
}
