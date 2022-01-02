package model

import (
	"YamadaKESHE/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type Rtsl struct {
	gorm.Model
	Lng      string `gorm:"type:varchar(30)"json:"lng"`
	Lat      string `gorm:"type:varchar(30)"json:"lat"`
	LDay     string `gorm:"type:varchar(30)"json:"l_day"`
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Class    string `json:"class"`
	Role     int    `json:"role"`
}

var nums =make(map[string]bool)
// RtslInCreat 录入实时位置
func RtslInCreat(userid string, lng string, lat string) int {
	var rtsl Rtsl
	var user User
	Day := time.Now().Format("2006-01-02")
	db.Where("userid = ?", userid).First(&rtsl)
	db.Where("userid = ?", userid).First(&user)
	if rtsl.ID == 0&&user.ID!=0&&!nums[userid] {
		nums[userid]=true
		rtsl.Userid = user.Userid
		rtsl.Username = user.Username
		rtsl.Class = user.Class
		rtsl.LDay = Day
		rtsl.Role = user.Role
		rtsl.Lng = lng
		rtsl.Lat = lat
		err = db.Create(&rtsl).Error
		if err != nil {
			return errmsg.ERROR_UPDATE_FAIL
		}
		return errmsg.SUCCESS
	}else {
		err = db.Model(&rtsl).Updates(map[string]interface{}{"lng": lng, "lat": lat}).Error
		if err != nil {
			return errmsg.ERROR_UPDATE_FAIL
		}
	}
	return errmsg.SUCCESS
}

// RstlCheck 实时定位查询
func RstlCheck(userid string) ([]Rtsl, int) {
	var rtsl []Rtsl
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		err = db.Where("class=? AND l_day=?", user.Class, time.Now().Format("2006-01-02")).Find(&rtsl).Error
		if err != nil {
			return rtsl, errmsg.ERROR_Check_FAIL
		}
		return rtsl, errmsg.SUCCESS
	}
	return rtsl, errmsg.ERROR_USERID_NULL
}
