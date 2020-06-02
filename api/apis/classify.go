package apis

import (
	model "newapp/database/models"
)

// MakeClassify make classify list
func MakeClassify() ([]model.MvClassify, []model.MvClassify) {
	var classify model.MvClassify
	allclass, err := classify.BigClass()

	var (
		b []model.MvClassify
		s []model.MvClassify
	)
	if err != nil {
		return b, s
	}
	for _, item := range allclass {
		if item.TopID == 0 {
			b = append(b, item)
		} else {
			s = append(s, item)
		}
	}
	return b, s
}
