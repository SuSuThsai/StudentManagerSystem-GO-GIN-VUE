package model

import (
	"YamadaKESHE/utils/errmsg"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Leave struct {
	gorm.Model
	Userid   string `gorm:"type:varchar(20);not null"json:"userid"`
	Username string `gorm:"type:varchar(20);not null"json:"username"`
	Class    string `gorm:"type:varchar(20);not null"json:"class"`
	Role     int    `gorm:"type:int;DEFAULT:2"json:"role"`
	Reason   string `gorm:"type:longtext;DEFAULT:null"json:"reason"`
	StartDay string `gorm:"type:varchar(40);DEFAULT:0"json:"start_day"`
	LastFor  int    `gorm:"type:int;DEFAULT:1"json:"last_for"`
	EndDay   string `gorm:"type:varchar(40);DEFAULT:1"json:"end_day"`
	Status   int    `gorm:"type:int;DEFAULT:1"json:"status"`
}

func LeavePost(userid string, leave *Leave) (Leave, int) {
	var user User
	db.Where("userid = ?", userid).First(&user)
	time1 := time.Now()
	db.Where("start_day = ?", time1.Format("2006-01-02")).Where("userid = ?", user.Userid).First(&leave)
	fmt.Println("1", leave.LastFor)
	if leave.ID == 0 {
		leave.Userid = user.Userid
		leave.Username = user.Username
		leave.Class = user.Class
		leave.Role = user.Role
		leave.StartDay = time1.Format("2006-01-02")
		leave.EndDay = time1.Add(time.Duration(leave.LastFor) * 24 * time.Hour).Format("2006-01-02")
		fmt.Println("11", time1.Add(time.Duration(leave.LastFor)*24*time.Hour))
		leave.Status = 1
		err = db.Create(&leave).Error
		if err != nil {
			return *leave, errmsg.ERROR_POST_FAIL
		}
		return *leave, errmsg.SUCCESS
	}
	return *leave, errmsg.ERROR_USER_IS_LEAVING
}

// LeaveCheckClass 查看班级所有请假信息
func LeaveCheckClass(userid string, pageSize int, pageNum int) ([]Leave, int64) {
	var leave []Leave
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Select("id,userid,username,class,reason,start_day,end_day,last_for,status").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&leave)
		db.Model(&leave).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return leave, total
	}
	err = db.Select("id,userid,username,class,reason,start_day,end_day,last_for,status").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&leave).Error
	db.Model(&leave).Where("role = ?", 2).Count(&total)

	if err != nil {
		return leave, 0
	}
	return leave, total
}

// LeaveCheckAll 查询所有班级的请假信息
func LeaveCheckAll(userid string, pageSize int, pageNum int) ([]Leave, int64) {
	var leave []Leave
	var user User
	var total int64
	if userid != "" {
		db.Where(" userid = ?", userid).First(&user)
		db.Select("id,userid,username,class,reason,start_day,end_day,last_for,status").Where(
			"role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&leave)
		db.Model(&leave).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return leave, total
	}
	err = db.Select("id,userid,username,class,reason,start_day,end_day,last_for,status").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&leave).Error
	db.Model(&leave).Where("role = ?", 2).Count(&total)

	if err != nil {
		return leave, 0
	}
	return leave, total
}

// LeaveCheckSignDay 按日期查询所有班级请假信息
func LeaveCheckSignDay(startDay string, pageSize int, pageNum int) ([]Leave, int64) {
	var leave []Leave
	var total int64
	err = db.Where("start_day = ? and role = ?", startDay, 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&leave).Error
	db.Model(&leave).Where("start_day = ? and role = ?", startDay, 2).Count(&total)
	if err != nil {
		return leave, 0
	}
	return leave, total
}

// LeaveCheckItself 查询自己的请假信息
func LeaveCheckItself(userid string) ([]Leave, int) {
	var leave Leave
	var leaves []Leave
	if userid != "" {
		db.Where(" userid = ?", userid).First(&leave)
		db.Where("userid = ?", userid).Find(&leaves)
		if leave.ID == 0 {
			return leaves, errmsg.ERROR_USER_NOT_EXIST
		}
		return leaves, errmsg.SUCCESS
	}
	return leaves, errmsg.ERROR_USERID_NULL
}

// LeaveUpdateStudent 学生修改请假信息
func LeaveUpdateStudent(userid string, data *Leave) int {
	var leave Leave
	var old Leave
	time1 := time.Now()
	db.Model(&leave).Where("id = ? and userid = ?", data.ID, userid).First(&old)
	if old.ID != 0 {
		var maps = make(map[string]interface{})
		maps["reason"] = data.Reason
		maps["start_day"] = time1.Format("2006-01-02")
		maps["last_for"] = data.LastFor
		maps["end_day"] = time1.Add(time.Duration(data.LastFor) * 24 * time.Hour).Format("2006-01-02")
		err = db.Model(&leave).Where("id = ? and userid = ?", data.ID, userid).Updates(&maps).Error
		if err != nil {
			return errmsg.ERROR_UPDATE_FAIL
		}
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_USERID_NULL
}

// LeaveUpdateTeacher 老师批改请假信息
func LeaveUpdateTeacher(userid string, data *Leave) int {
	var leave Leave
	var user User
	db.Where("userid = ?", userid).First(&user)
	db.Where("id = ?", data.ID).First(&leave)
	if user.Class == leave.Class {
		var maps = make(map[string]interface{})
		maps["status"] = data.Status
		err = db.Model(&leave).Where("id = ?", data.ID).Updates(&maps).Error
		if err != nil {
			return errmsg.ERROR_TOEXAMINE_FAIL
		}
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_USER_NOT_RIGHT
}

// LeaveDelete 删除请假信息
func LeaveDelete(data *Leave) int {
	var leave Leave
	db.Where("id = ?", data.ID).First(&leave)
	if leave.Status == 2 || leave.Status == 1 {
		err = db.Where("id = ?", data.ID).Delete(&leave).Error
		if err != nil {
			return errmsg.ERROR_DELETE_FAIL
		}
		return errmsg.SUCCESS
	}
	return errmsg.ERROR
}
