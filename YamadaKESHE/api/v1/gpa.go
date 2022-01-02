package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GpaWriteAll 导入班级所有考核成绩
func GpaWriteAll(c *gin.Context) {
	var gpa model.Gpa
	_ = c.ShouldBind(&gpa)
	userid := c.GetString("userid")
	data, code, total := model.GpaWriteAll(userid, &gpa)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// GpaUpdate 修改学习成绩
func GpaUpdate(c *gin.Context) {
	var gpa model.Gpa
	userid := c.GetString("userid")
	_ = c.ShouldBind(&gpa)
	gpa.Userid = c.PostForm("userid")
	AcademicScore, _ := strconv.Atoi(c.PostForm("academic_score"))
	gpa.AcademicScore = AcademicScore
	code := model.GpaUpdate(userid, &gpa)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// GpaDelete 老师删除评核
func GpaDelete(c *gin.Context) {
	var gpa model.Gpa
	_ = c.ShouldBind(&gpa)
	userid := c.GetString("userid")
	code := model.GpaDelete(userid, &gpa)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// GpaCheckClass 查询对应老师班级学业成绩
func GpaCheckClass(c *gin.Context) {
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

	data, total := model.GpaCheckClass(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GpaCheckAll 查询所有班级学业成绩
func GpaCheckAll(c *gin.Context) {
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

	data, total := model.GpaCheckAll(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GpaCheckItself 查询自己的学业成绩
func GpaCheckItself(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.GpaCheckItself(userid)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// GpaCheckuserid 查询对于学号的学业成绩
func GpaCheckuserid(c *gin.Context) {
	userid := c.GetHeader("userid")
	data, code := model.GpaCheckuserid(userid)
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

// GpaCheckUpdate 更新班级所有考核成绩
func GpaCheckUpdate(c *gin.Context) {
	var gpa model.Gpa
	_ = c.ShouldBind(&gpa)
	userid := c.GetString("userid")
	data, code, total := model.GpaCheckUpdate(userid)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	}
}
