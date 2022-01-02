package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ReportWrite 写实习报告
func ReportWrite(c *gin.Context) {
	var report model.Report
	_ = c.ShouldBind(&report)
	userid := c.GetString("userid")
	report.Title = c.PostForm("title")
	report.Desc = c.PostForm("desc")
	report.Content = c.PostForm("content")
	data, code := model.ReportWrite(userid, &report)
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

// ReportUpdate 修改实习报告
func ReportUpdate(c *gin.Context) {
	var data model.Report
	userid := c.GetString("userid")
	_ = c.ShouldBind(&data)
	data.Userid = c.PostForm("userid")
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	data.ID = uint(tempid)
	data.Title = c.PostForm("title")
	data.Desc = c.PostForm("desc")
	data.Content = c.PostForm("content")
	code := model.ReportUpdate(userid, &data)
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

// ReportDelete 删除报告
func ReportDelete(c *gin.Context) {
	var data model.Report
	_ = c.ShouldBind(&data)
	data.Userid = c.PostForm("userid")
	data.Class = c.PostForm("class")
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	data.ID = uint(tempid)
	userid := c.GetString("userid")
	code := model.ReportDelete(userid, &data)
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

// ReportCheckClass 查询对应老师班级实习报告
func ReportCheckClass(c *gin.Context) {
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

	data, total := model.ReportCheckClass(userid, pageSize, pageNum)
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

// ReportCheckDay 按日期查询实习报告
func ReportCheckDay(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	creatDay := c.GetHeader("creat_day")
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
	data, total := model.ReportCheckDay(userid, creatDay, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// ReportCheckuserid 查询对应学号的实习报告
func ReportCheckuserid(c *gin.Context) {
	userid := c.GetHeader("userid")
	data, code := model.ReportCheckuserid(userid)
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

// ReportCheckItself 查询自己的实习报告
func ReportCheckItself(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.ReportCheckItself(userid)
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
