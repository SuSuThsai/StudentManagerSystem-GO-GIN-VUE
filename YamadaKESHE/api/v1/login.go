package v1

import (
	"YamadaKESHE/middleware"
	"YamadaKESHE/model"
	"YamadaKESHE/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login 后台登录
func Login(c *gin.Context) {
	var data model.User
	data.Userid = c.PostForm("userid")
	data.Password = c.PostForm("password")

	var code int
	var token string

	data, code = model.ValidateLogin(data.Userid, data.Password)
	if code == errmsg.SUCCESS {
		setToken(c, data)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"userid":  data.Userid,
			"data":    data.Username,
			"token":   token,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var formData model.User
	formData.Userid = c.PostForm("userid")
	formData.Password = c.PostForm("password")

	var code int
	formData, code = model.CheckLoginFront(formData.Userid, formData.Password)
	if formData.ID != 0 {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"id":      formData.ID,
			"userid":  formData.Userid,
			"data":    formData.Username,
			"class":   formData.Class,
			"role":    formData.Role,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// LoginLocalManager LocalManager
func LoginLocalManager(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var code int
	formData, code := model.CheckLoginLocal(username, password)
	if code == errmsg.SUCCESS {
		setTokenLocal(c, formData)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":   code,
			"username": formData.Username,
			"role":     formData.Role,
			"message":  errmsg.GetErrMsg(code),
		})
	}
}

//token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.Claims{
		Userid: user.Userid,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "Yamada",
		},
	}
	tokenString, err := j.CreatToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   tokenString,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"id":      user.ID,
		"userid":  user.Userid,
		"data":    user.Username,
		"class":   user.Class,
		"role":    user.Role,
		"token":   tokenString,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})
	return
}

//token生成函数
func setTokenLocal(c *gin.Context, user model.Manager) {
	j := middleware.NewJWT()
	claims := middleware.Claims{
		Userid: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "Yamada",
		},
	}
	tokenString, err := j.CreatToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   tokenString,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    user.Username,
		"role":    user.Role,
		"token":   tokenString,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})
	return
}
