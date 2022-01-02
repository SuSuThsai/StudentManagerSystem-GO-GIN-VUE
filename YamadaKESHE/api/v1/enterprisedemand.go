package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DemandWrite 写企业需求
func DemandWrite(c *gin.Context) {
	var demand model.Demand
	_ = c.ShouldBind(&demand)
	userid := c.GetString("userid")
	demand.CompanyName = c.PostForm("company_name")
	demand.Station = c.PostForm("station")
	demand.PositionDescription = c.PostForm("position_description")
	demand.EndDay = c.PostForm("end_day")
	data, code := model.DemandWrite(userid, &demand)
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

// DemandUpdate 修改企业需求
func DemandUpdate(c *gin.Context) {
	var demand model.Demand
	userid := c.GetString("userid")
	_ = c.ShouldBind(&demand)
	demand.CompanyName = c.PostForm("company_name")
	demand.EndDay = c.PostForm("end_day")
	demand.Station = c.PostForm("station")
	demand.PositionDescription = c.PostForm("position_description")
	demand.Userid = c.PostForm("userid")
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	demand.ID = uint(tempid)
	code := model.DemandUpdate(userid, &demand)
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

// DemandDelete 老师删除企业需求
func DemandDelete(c *gin.Context) {
	var demand model.Demand
	_ = c.ShouldBind(&demand)
	userid := c.GetString("userid")
	demand.Userid = c.PostForm("userid")
	tempid, _ := strconv.Atoi(c.PostForm("id"))
	demand.ID = uint(tempid)
	code := model.DemandDelete(userid, &demand)
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

// DemandCheckClass 查询对应老师班级企业需求
func DemandCheckClass(c *gin.Context) {
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

	data, total := model.DemandCheckClass(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// DemandCheckAll 查询所有企业需求
func DemandCheckAll(c *gin.Context) {
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

	data, total := model.DemandCheckAll(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// DemandCheckDay 按日期查询企业需求
func DemandCheckDay(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	creatDay := c.Query("creat_day")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.DemandCheckDay(creatDay, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// DemandCheckItself 查询自己发表的企业需求
func DemandCheckItself(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.DemandCheckItself(userid)
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
