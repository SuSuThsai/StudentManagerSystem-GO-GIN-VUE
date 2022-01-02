package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AssessmentWriteAll 导入班级所有考核成绩
func AssessmentWriteAll(c *gin.Context) {
	var assessment model.Assessment
	_ = c.ShouldBind(&assessment)
	userid := c.GetString("userid")
	data, code, total := model.AssessmentWriteAll(userid, &assessment)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// AssessmentUpdate 修改实习考核成绩
func AssessmentUpdate(c *gin.Context) {
	var assessment model.Assessment
	userid := c.GetString("userid")
	_ = c.ShouldBind(&assessment)
	assessment.Userid = c.PostForm("userid")
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	assessment.ID = uint(tempid)
	Ability, _ := strconv.Atoi(c.PostForm("working_ability"))
	assessment.WorkingAbility = Ability
	Communication, _ := strconv.Atoi(c.PostForm("communication"))
	assessment.Communication = Communication
	Attitude, _ := strconv.Atoi(c.PostForm("working_attitude"))
	assessment.WorkingAttitude = Attitude
	code := model.AssessmentUpdate(userid, &assessment)
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

// AssessmentDelete 老师删除评核
func AssessmentDelete(c *gin.Context) {
	var assessment model.Assessment
	_ = c.ShouldBind(&assessment)
	userid := c.GetString("userid")
	code := model.AssessmentDelete(userid, &assessment)
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

// AssessmentCheckClass 查询对应老师班级实习评核
func AssessmentCheckClass(c *gin.Context) {
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

	data, total := model.AssessmentCheckClass(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// AssessmentCheckAll 查询所有班级实习评核
func AssessmentCheckAll(c *gin.Context) {
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

	data, total := model.AssessmentCheckAll(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// AssessmentCheckItself 查询自己的实习考核成绩
func AssessmentCheckItself(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.AssessmentCheckItself(userid)
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

// AssessmentCheckuserid 查询自己的实习考核成绩
func AssessmentCheckuserid(c *gin.Context) {
	userid := c.GetHeader("userid")
	data, code := model.AssessmentCheckuserid(userid)
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
