package v1

import (
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"YamadaKESHE/utils/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var code int
	var validCode int
	data.Userid = c.PostForm("userid")
	data.Username = c.PostForm("username")
	data.Class = c.PostForm("class")
	data.Password = c.PostForm("password")
	data.Role = 2
	msg, validCode = validate.Validate(&data)
	if validCode != errmsg.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  validCode,
			"message": msg,
		})
		c.Abort()
		return
	}
	code = model.CheckUser(data.Userid)
	if code == errmsg.SUCCESS {
		model.CreatUser(&data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else if code == errmsg.ERROR_USERID_USED {
		code = errmsg.ERROR_USERID_USED
		c.JSON(http.StatusConflict, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusGone, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
		code = errmsg.ERROR
	}
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	data, code := model.GetUser(id)
	var maps = make(map[string]interface{})
	maps["userid"] = data.Userid
	maps["username"] = data.Username
	maps["class"] = data.Class
	maps["role"] = data.Role
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    maps,
		"total":   1,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
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

	data, total := model.GetUsers(userid, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	userid := c.PostForm("userid")
	role, _ := strconv.Atoi(c.PostForm("role"))
	code := model.CheckUpUser(userid)
	if code == errmsg.SUCCESS {
		code = model.EditUser(userid, role)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("删除用户：", id)
	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id := c.Param("id")
	_ = c.ShouldBind(&data)

	code := model.ChangePassword(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
