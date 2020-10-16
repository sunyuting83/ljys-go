package models

import orm "imovie/database"

// SearchKey movie data
type SearchKey struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// TableName change table name
func (SearchKey) TableName() string {
	return "mv_movie"
}

// SearchKey 列表
func (search *MvMovie) SearchKey(key string) (searchs []SearchKey, err error) {
	if err = orm.Eloquent.
		Where("title LIKE ?", "%"+key+"%").
		Or("entitle LIKE ?", "%"+key+"%").
		Select("id, title").
		Order("id desc").
		Limit(7).
		Find(&searchs).Error; err != nil {
		return
	}
	return
}

// Search 列表
func (search *MvMovie) Search(key string, page int64) (searchs []MvMovie, err error) {
	p := makePage(page)
	if err = orm.Eloquent.
		Where("title LIKE ?", "%"+key+"%").
		Or("entitle LIKE ?", "%"+key+"%").
		Select("id, title, other").
		Order("id desc").
		Limit(15).
		Offset(p).
		Find(&searchs).Error; err != nil {
		return
	}
	return
}

// SearchHot 列表
func (search *MvMovie) SearchHot() (searchs []SearchKey, err error) {
	if err = orm.Eloquent.
		Select("id, title").
		Order("random()").
		Limit(10).
		Find(&searchs).Error; err != nil {
		return
	}
	return
}
