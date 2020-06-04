package models

import (
	orm "newapp/database"
	"time"
)

// MvMovie struct
type MvMovie struct {
	ID        int64     `json:"id" gorm:"primary_key;column:id"`
	Cid       int64     `json:"cid" gorm:"column:cid"`
	Title     string    `json:"title" gorm:"column:title"`
	Other     string    `json:"other" gorm:"column:other"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// Movie struct
type Movie struct {
	ID        int64          `json:"id" gorm:"primary_key:true;column:id"`
	Cid       int64          `json:"cid" gorm:"column:cid"`
	Title     string         `json:"title" gorm:"column:title"`
	Other     string         `json:"other" gorm:"column:other"`
	Director  []*MvDirector  `gorm:"many2many:mv_director_mv_movie;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:mv_director_id;jointable_foreignkey:mv_movie_id;" json:"director"`
	Performer []*MvPerformer `json:"performer" gorm:"many2many:mv_movie_mv_performer;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:mv_performer_id;jointable_foreignkey:mv_movie_id;"`
	Area      []*MvArea      `json:"area" gorm:"many2many:mv_area_mv_movie;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:mv_area_id;jointable_foreignkey:mv_movie_id;"`
}

// MvDirector struct
type MvDirector struct {
	ID    int64  `json:"id" gorm:"primary_key:true;column:id"`
	DName string `json:"d_name" gorm:"column:d_name"`
	Count int64  `json:"count" gorm:"column:count"`
}

// MvPerformer struct
type MvPerformer struct {
	ID    int64  `json:"id" gorm:"primary_key;column:id"`
	PName string `json:"p_name" gorm:"column:p_name"`
	Count int64  `json:"count" gorm:"column:count"`
}

// MvArea struct
type MvArea struct {
	ID    int64  `json:"id" gorm:"primary_key;column:id"`
	AName string `json:"a_name" gorm:"column:a_name"`
	Count int64  `json:"count" gorm:"column:count"`
}

// TableName change table name
func (Movie) TableName() string {
	return "mv_movie"
}

// Indexs 列表
func (index *MvMovie) Indexs(inid []int64) (indexs []MvMovie, err error) {
	if err = orm.Eloquent.
		Select("id, cid, title, other").
		Order("id desc").
		Limit(9).
		Find(&indexs, "cid in (?)", inid).Error; err != nil {
		return
	}
	return
}

// Classifys 列表
func (index *MvMovie) Classifys(id int64) (indexs []MvMovie, err error) {
	if err = orm.Eloquent.
		Select("id, cid, title, other").
		Order("id desc").
		Limit(9).
		Find(&indexs, "cid = ?", id).Error; err != nil {
		return
	}
	return
}

// Lists 列表
func (index *MvMovie) Lists(id, page int64) (indexs []MvMovie, err error) {
	p := makePage(page)
	if err = orm.Eloquent.
		Select("id, cid, title, other").
		Order("id desc").
		Limit(30).Offset(p).
		Find(&indexs, "cid = ?", id).Error; err != nil {
		return
	}
	return
}

// Movies data
func (movie *Movie) Movies(id int64) (movies Movie, err error) {
	orm.Eloquent.First(&movies, id)
	if err = orm.Eloquent.
		Model(&movies).
		Related(&movies.Director, "Director").
		Related(&movies.Performer, "Performer").
		Related(&movies.Area, "Area").
		First(&movies).Error; err != nil {
		return
	}
	// if err = orm.Eloquent.Preload("Director").First(&movies, id).Error; err != nil {
	// 	return
	// }
	return
}

// HotLists 列表
func (index *MvMovie) HotLists(inid int64) (indexs []MvMovie, err error) {
	if err = orm.Eloquent.
		Select("id, cid, title, other").
		Order("random()").
		Limit(3).
		Find(&indexs, "cid = ?", inid).Error; err != nil {
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
