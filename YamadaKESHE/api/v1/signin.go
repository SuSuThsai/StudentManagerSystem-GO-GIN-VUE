package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SignInCreat 签到
func SignInCreat(c *gin.Context) {
	id := c.GetString("userid")

	data, isValue, code := model.SignInCreat(id)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":   code,
			"data":     data,
			"is_value": isValue,
			"message":  errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":   code,
			"data":     data,
			"is_value": isValue,
			"message":  errmsg.GetErrMsg(code),
		})
	}
}

// SIgnInCheckClass 查询对应老师班级信息
func SIgnInCheckClass(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	userid := c.GetString("userid")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.SIgnInCheckClass(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// SIgnInCheckAll 查询所有班级信息
func SIgnInCheckAll(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	userid := c.Query("userid")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.SIgnInCheckAll(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// SIgnInCheckSignDay 按日期查询
func SIgnInCheckSignDay(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	signDay := c.GetHeader("sign_day")
	userid := c.GetString("userid")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.SIgnInCheckSignDay(userid, signDay, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// SignInCheckItself 查询自己是否签到
func SignInCheckItself(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.SIgnInCheckItself(userid)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}
