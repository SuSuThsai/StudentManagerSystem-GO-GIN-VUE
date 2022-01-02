package model

import (
	"YamadaKESHE/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type Notice struct {
	gorm.Model
	Title     string `gorm:"type:varchar(40)"json:"title"`
	Text      string `gorm:"type:varchar(120)"json:"text"`
	CreatDay  string `gorm:"type:varchar(30);not nul;"json:"creat_day"`
	UpdateDay string `gorm:"type:varchar(30)"json:"update_day"`
	Userid    string `json:"userid"`
	Username  string `json:"username"`
	Class     string `json:"class"`
	Role      int    `json:"role"`
}

// NoticeWrite 写实习通知
func NoticeWrite(userid string, data *Notice) (Notice, int) {
	var user User
	db.Where("userid = ?", userid).First(&user)
	data.Userid = user.Userid
	data.Username = user.Username
	data.Class = user.Class
	data.Role = user.Role
	data.CreatDay = time.Now().Format("2006-01-02")
	data.UpdateDay = "无"
	err = db.Create(&data).Error
	if err != nil {
		return *data, errmsg.ERROR_POST_FAIL
	}
	return *data, errmsg.SUCCESS
}

// NoticeUpdate 修改实习通知
func NoticeUpdate(userid string, data *Notice) int {
	var notice Notice
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 && data.Userid == user.Userid {
		var maps = make(map[string]interface{})
		maps["title"] = data.Title
		maps["text"] = data.Text
		// 15:04:05 可以精确到秒
		maps["update_day"] = time.Now().Format("2006-01-02")
		err = db.Model(&notice).Where("id = ? and userid = ?", data.ID, user.Userid).Updates(&maps).Error
		if err != nil {
			return errmsg.ERROR_UPDATE_FAIL
		}
		return errmsg.SUCCESS
	}
	if data.Userid != user.Userid {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.ERROR_USERID_NULL
}

// NoticeDelete 老师删除实习通知
func NoticeDelete(userid string, data *Notice) int {
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 && data.Userid == user.Userid {
		err = db.Where("id = ? and userid = ?", data.ID, user.Userid).Delete(&data).Error
		if err != nil {
			return errmsg.ERROR_DELETE_FAIL
		}
		return errmsg.SUCCESS
	}
	if data.Userid != user.Userid {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.ERROR_USERID_NULL
}

// NoticeCheckClass 查询对应老师班级实习通知
func NoticeCheckClass(userid string, pageSize int, pageNum int) ([]Notice, int64) {
	var notice []Notice
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Select("id,userid,username,class,title,text,role,creat_day,update_day").Where(
			"class = ? and role = ?", user.Class, 1,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notice)
		db.Model(&notice).Where(
			"class = ? and role = ?", user.Class, 1,
		).Count(&total)
		return notice, total
	}

	err = db.Select("id,userid,username,class,title,text,role,creat_day,update_day").Where("role = ?", 1).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notice).Error
	db.Model(&notice).Where("role = ?", 1).Count(&total)

	if err != nil {
		return notice, 0
	}
	return notice, total
}

// NoticeCheckAll 查询所有班级实习通知
func NoticeCheckAll(userid string, pageSize int, pageNum int) ([]Notice, int64) {
	var notice []Notice
	var user User
	var total int64
	if userid != "" {
		db.Where(" userid = ?", userid).First(&user)
		db.Select("id,userid,username,class,text,role,creat_day,update_day").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notice)
		db.Model(&notice).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return notice, total
	}
	err = db.Select("id,userid,username,class,text,role,creat_day,update_day").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notice).Error
	db.Model(&notice).Where("role = ?", 2).Count(&total)
	if err != nil {
		return notice, 0
	}
	return notice, total
}

// NoticeCheckDay 按日期查询实习通知
func NoticeCheckDay(userid string, creatDay string, pageSize int, pageNum int) ([]Notice, int64) {
	var notice []Notice
	var user User
	var total int64
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		err = db.Where("creat_day = ? and role = ? and class = ? ", creatDay, 1, user.Class).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notice).Error
		db.Model(&notice).Where("creat_day = ? and role = ? and class = ? ", creatDay, 1, user.Class).Count(&total)
		if err != nil {
			return notice, 0
		}
		return notice, total
	} else {
		return notice, 0
	}
}

// NoticeCheckItselfTeacher 查询教师自己的实习通知
func NoticeCheckItselfTeacher(userid string) (Notice, int) {
	var notice Notice
	if userid != "" {
		db.Where(" userid = ?", userid).First(&notice)
		if notice.ID == 0 {
			return notice, errmsg.ERROR_USER_NOT_EXIST
		}
		return notice, errmsg.SUCCESS
	}
	return notice, errmsg.ERROR_USERID_NULL
}
