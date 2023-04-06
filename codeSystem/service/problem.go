package service

import (
	"codeSystem/define"
	"codeSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Summary 问题列表
// @Tags 公共方法
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json"{"code":200,"data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultPageSize))
	if err != nil {
		log.Println("GetProblemList", err)
		return
	}
	page = (page - 1) * size
	keyword := c.Query("keyword")
	//categoryIdentity := "a"

	data := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword)
	err = tx.Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("GetProblemList", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
