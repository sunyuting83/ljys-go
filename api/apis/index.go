package apis

import (
	"fmt"
	"net/http"

	model "newapp/database/models"

	"github.com/gin-gonic/gin"
)

// Indexs 列表数据
func Indexs(c *gin.Context) {
	var (
		classify model.MvClassify
		// index model.MvMovie
	)
	allclass, err := classify.BigClass()
	b, m := makeClassify(allclass)
	for _, item := range b {
		smallclass, err := classify.SmallClass(item.ID)
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println(smallclass)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   0,
		"menu":     b,
		"menumore": m,
	})
}

// makeClassify make classify list
func makeClassify(c []model.MvClassify) ([]model.MvClassify, []model.MvClassify) {
	var (
		b []model.MvClassify
		s []model.MvClassify
	)
	for _, item := range c {
		if item.TopID == 0 {
			b = append(b, item)
		} else {
			s = append(s, item)
		}
	}
	return b, s
}
