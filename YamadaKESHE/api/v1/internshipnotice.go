package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// NoticeWrite 写实习通知
func NoticeWrite(c *gin.Context) {
	var notice model.Notice
	_ = c.ShouldBind(&notice)
	userid := c.GetString("userid")
	notice.Title = c.PostForm("title")
	notice.Text = c.PostForm("text")
	data, code := model.NoticeWrite(userid, &notice)
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

// NoticeUpdate 修改实习通知
func NoticeUpdate(c *gin.Context) {
	var notice model.Notice
	userid := c.GetString("userid")
	_ = c.ShouldBind(&notice)
	notice.Title = c.PostForm("title")
	notice.Text = c.PostForm("text")
	notice.Userid = c.PostForm("userid")
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	notice.ID = uint(tempid)
	code := model.NoticeUpdate(userid, &notice)
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

// NoticeDelete 老师删除实习通知
func NoticeDelete(c *gin.Context) {
	var notice model.Notice
	_ = c.ShouldBind(&notice)
	userid := c.GetString("userid")
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	notice.ID = uint(tempid)
	notice.Userid = c.PostForm("userid")
	code := model.NoticeDelete(userid, &notice)
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

// NoticeCheckClass 查询对应老师班级实习通知
func NoticeCheckClass(c *gin.Context) {
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

	data, total := model.NoticeCheckClass(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//// NoticeCheckAll 查询所有班级实习通知
//func NoticeCheckAll(c *gin.Context){
//	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
//	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
//	userid := c.Query("userid")
//	switch {
//	case pageSize >= 100:
//		pageSize = 100
//	case pageSize <= 0:
//		pageSize = 10
//	}
//
//	if pageNum == 0 {
//		pageNum = 1
//	}
//
//	data, total := model.NoticeCheckAll(userid, pageSize, pageNum)
//	code := errmsg.SUCCESS
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"total":   total,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// NoticeCheckDay 按日期查询实习通知
func NoticeCheckDay(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	creatDay := c.Query("creat_day")
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
	data, total := model.NoticeCheckDay(userid, creatDay, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// NoticeCheckItselfTeacher 查询教师自己的实习通知
func NoticeCheckItselfTeacher(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.NoticeCheckItselfTeacher(userid)
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
