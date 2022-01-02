package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// LeavePost 提交请假申请
func LeavePost(c *gin.Context) {
	var leave model.Leave
	_ = c.ShouldBind(&leave)
	userid := c.GetString("userid")
	leave.Reason = c.PostForm("reason")
	Lastfor, _ := strconv.Atoi(c.PostForm("last_for"))
	leave.LastFor = Lastfor
	data, code := model.LeavePost(userid, &leave)
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

// LeaveCheckClass 查看班级所有请假信息
func LeaveCheckClass(c *gin.Context) {
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

	data, total := model.LeaveCheckClass(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// LeaveCheckAll 查询所有班级的请假信息
func LeaveCheckAll(c *gin.Context) {
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

	data, total := model.LeaveCheckAll(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// LeaveCheckSignDay 按日期查询所有班级请假信息
func LeaveCheckSignDay(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	signDay := c.Query("sign_day")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.LeaveCheckSignDay(signDay, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// LeaveCheckItself 查询自己的请假信息
func LeaveCheckItself(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.LeaveCheckItself(userid)
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

// LeaveUpdateStudent 学生修改请假信息
func LeaveUpdateStudent(c *gin.Context) {
	var data model.Leave
	userid := c.GetString("userid")
	data.Userid = c.PostForm("userid")
	_ = c.ShouldBind(&data)
	data.Reason = c.PostForm("reason")
	Lastfor, _ := strconv.Atoi(c.PostForm("last_for"))
	data.LastFor = Lastfor
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	data.ID = uint(tempid)
	code := model.LeaveUpdateStudent(userid, &data)
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

// LeaveUpdateTeacher 老师批改请假信息
func LeaveUpdateTeacher(c *gin.Context) {
	var data model.Leave
	_ = c.ShouldBind(&data)
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	data.ID = uint(tempid)
	tempStatus, _ := strconv.Atoi(c.PostForm("status"))
	data.Status = tempStatus
	userid := c.GetString("userid")
	code := model.LeaveUpdateTeacher(userid, &data)
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

// LeaveDelete 删除请假信息
func LeaveDelete(c *gin.Context) {
	var data model.Leave
	_ = c.ShouldBind(&data)
	code := model.LeaveDelete(&data)
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
