package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RtslInCreat 录入实时位置
func RtslInCreat(c *gin.Context) {
	userid := c.GetString("userid")
	lng := c.PostForm("lng")
	lat := c.PostForm("lat")
	fmt.Println("1", lng, lat)
	code := model.RtslInCreat(userid, lng, lat)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    lng,
			"lat":     lat,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// RstlCheck 实时定位查询
func RstlCheck(c *gin.Context) {
	userid := c.GetString("userid")
	data, code := model.RstlCheck(userid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
