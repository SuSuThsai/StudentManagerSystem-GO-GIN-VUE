package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// WorkLogWrite 写工作日志
func WorkLogWrite(c *gin.Context) {
	var worklog model.Log
	_ = c.ShouldBind(&worklog)
	userid := c.GetString("userid")
	_ = c.ShouldBind(&worklog)
	worklog.Userid = c.PostForm("userid")
	worklog.Title = c.PostForm("title")
	worklog.TeamMember = c.PostForm("team_member")
	worklog.Goal = c.PostForm("goal")
	worklog.Results = c.PostForm("results")
	data, code := model.WorkLogWrite(userid, &worklog)
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

// WorkLogUpdate 修改工作日志
func WorkLogUpdate(c *gin.Context) {
	var worklog model.Log
	userid := c.GetString("userid")
	_ = c.ShouldBind(&worklog)
	id, _ := strconv.Atoi(c.PostForm("id"))
	worklog.ID = uint(id)
	worklog.Userid = c.PostForm("userid")
	worklog.Title = c.PostForm("title")
	worklog.TeamMember = c.PostForm("team_member")
	worklog.Goal = c.PostForm("goal")
	worklog.Results = c.PostForm("results")
	code := model.WorkLogUpdate(userid, &worklog)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    worklog,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"data":    worklog,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

//// WorkLogDelete 删除工作日志
//func WorkLogDelete(c *gin.Context){
//	var worklog model.Log
//	_=c.ShouldBind(&worklog)
//	code := model.WorkLogDelete(&worklog)
//	if code==errmsg.SUCCESS {
//		c.JSON(http.StatusOK,gin.H{
//			"status":code,
//			"message":errmsg.GetErrMsg(code),
//		})
//	}else {
//		c.JSON(http.StatusOK,gin.H{
//			"status":code,
//			"message":errmsg.GetErrMsg(code),
//		})
//	}
//}

// WorkLogCheckClass 查询对应老师班级实习报告
func WorkLogCheckClass(c *gin.Context) {
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

	data, total := model.WorkLogCheckClass(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//// ReportCheckAll 查询所有班级实习报告
//func ReportCheckAll(c *gin.Context){
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
//	data, total := model.reportcheckAll(userid, pageSize, pageNum)
//	code := errmsg.SUCCESS
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"total":   total,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// WorkLogCheckDay 按日期查询实习报告
func WorkLogCheckDay(c *gin.Context) {
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
	data, total := model.WorkLogCheckDay(userid, creatDay, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// WorkLogCheckItself 查询自己的实习报告
func WorkLogCheckItself(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.WorkLogCheckItself(userid)
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

// WorkLogCheckuserid 查询对应学号的工作日志
func WorkLogCheckuserid(c *gin.Context) {
	userid := c.GetHeader("userid")
	data, code := model.WorkLogCheckuserid(userid)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": data, "status": code,

			"message": errmsg.GetErrMsg(code),
		})
	}
}
