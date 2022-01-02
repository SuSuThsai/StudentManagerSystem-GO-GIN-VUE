package model

import (
	"YamadaKESHE/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type SignIn struct {
	gorm.Model
	Userid   string `gorm:"type:varchar(20);not null"json:"userid"`
	Username string `gorm:"type:varchar(20);not null"json:"username"`
	Class    string `gorm:"type:varchar(20);not null"json:"class"`
	Role     int    `gorm:"type:int;DEFAULT:2"json:"role"`
	SignDay  string `gorm:"type:varchar(40);not null"json:"sign_day"`
	IsValue  bool   `gorm:"type:bool;DEFAULT:false"json:"is_value" label:"状态"`
}

// SignInCreat 签到
func SignInCreat(userid string) ([]SignIn, bool, int) {
	var sigin SignIn
	var sigins []SignIn
	var user User
	SignDay := time.Now().Format("2006-01-02")
	db.Where("userid = ?", userid).First(&user)
	db.Where("userid", userid).First(&sigin)
	db.Where("isValue = ? and sign_day = ?", sigin.IsValue, SignDay).First(&sigin)
	if sigin.ID == 0 {
		sigin.Userid = userid
		sigin.Username = user.Username
		sigin.Class = user.Class
		sigin.SignDay = SignDay
		sigin.Role = user.Role
		sigin.IsValue = true
		db.Create(&sigin)
		db.Where("userid = ?", userid).Find(&sigins)
		return sigins, sigin.IsValue, errmsg.SUCCESS
	} else if sigin.ID != 0 && sigin.IsValue == false {
		var maps = make(map[string]interface{})
		maps["is_value"] = true
		err = db.Model(&sigin).Where("id = ? and userid = ?", sigin.ID, user.Userid).Updates(&maps).Error
		if err != nil {
			return sigins, sigin.IsValue, errmsg.ERROR_UPDATE_FAIL
		}
		db.Where("userid = ?", userid).Find(&sigins)
		return sigins, sigin.IsValue, errmsg.SUCCESS
	}
	return nil, sigin.IsValue, errmsg.ERROR_USER_IS_SIGNIN
}

// SIgnInCheckClass 按班级查询
func SIgnInCheckClass(userid string, pageSize int, pageNum int) ([]SignIn, int64) {
	var signin []SignIn
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Select("id,userid,username,class,is_value,sign_day,created_at").Where(
			"class = ? and role = ? and sign_day = ?", user.Class, 2, time.Now().Format("2006-01-02"),
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&signin)
		db.Model(&signin).Where(
			"class = ? and role = ? and sign_day = ?", user.Class, 2, time.Now().Format("2006-01-02"),
		).Count(&total)
		return signin, total
	}
	err = db.Select("id,userid,username,class,sign_day,is_value,created_at").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&signin).Error
	db.Model(&signin).Where("role = ?", 2).Count(&total)

	if err != nil {
		return signin, 0
	}
	return signin, total
}

// SIgnInCheckSignDay 按日期查询
func SIgnInCheckSignDay(userid string, signDay string, pageSize int, pageNum int) ([]SignIn, int64) {
	var signin []SignIn
	var user User
	var total int64
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		err = db.Where("sign_day = ? and role = ? and class = ?", signDay, 2, user.Class).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&signin).Error
		db.Model(&signin).Where("sign_day = ? and role = ? and class = ?", signDay, 2, user.Class).Count(&total)
		if err != nil {
			return signin, 0
		}
		return signin, total
	} else {
		return signin, 0
	}

}

// SIgnInCheckAll 查询所有班级的签到信息
func SIgnInCheckAll(userid string, pageSize int, pageNum int) ([]SignIn, int64) {
	var signin []SignIn
	var user User
	var total int64
	if userid != "" {
		db.Where(" userid = ?", userid).First(&user)
		db.Select("id,userid,username,class,is_value,sign_day,created_at").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&signin)
		db.Model(&signin).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return signin, total
	}
	err = db.Select("id,userid,username,class,sign_day,is_value,created_at").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&signin).Error
	db.Model(&signin).Where("role = ?", 2).Count(&total)

	if err != nil {
		return signin, 0
	}
	return signin, total
}

// SIgnInCheckItself 查询自己是否签到
func SIgnInCheckItself(userid string) ([]SignIn, int) {
	var signins []SignIn
	var signin SignIn
	if userid != "" {
		db.Where(" userid = ?", userid).First(&signin)
		if signin.ID == 0 {
			return nil, errmsg.ERROR_USER_NOT_EXIST
		}
		db.Where(" userid = ?", userid).Find(&signins)
		return signins, errmsg.SUCCESS
	}
	return signins, errmsg.ERROR_USERID_NULL
}

// SignInCreatNew 定期更新考勤表
func SignInCreatNew(timex time.Time) {
	var users []User
	db.Where("role = ?", 2).Find(&users)
	time1 := timex.Format("2006-01-02")

	for _, user := range users {
		var temp = SignIn{}
		db.Where("sign_day = ? and userid = ?", time1, user.Userid).First(&temp)
		if temp.ID == 0 {
			temp.Userid = user.Userid
			temp.Username = user.Username
			temp.Class = user.Class
			temp.Role = user.Role
			temp.SignDay = time1
			db.Create(&temp)
		}
	}
}

// ReFlashSignIn 定时
func ReFlashSignIn() {
	temp := time.Tick(1 * 24 * time.Hour)
	SignInCreatNew(time.Now())
	//SignInCreatNew(time.Now().Add(1*24*time.Hour))
	for {
		select {
		case <-temp:
			SignInCreatNew(time.Now().Add(1 * 24 * time.Hour))
		}
	}
}
