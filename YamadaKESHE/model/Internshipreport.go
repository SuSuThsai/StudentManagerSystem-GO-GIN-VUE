package model

import (
	"YamadaKESHE/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type Report struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100);not null"json:"title"`
	Desc      string `gorm:"type:varchar(200)"json:"desc"`
	Content   string `gorm:"type:longtext"json:"content"`
	CreatDay  string `gorm:"type:varchar(30);not nul;"json:"creat_day"`
	UpdateDay string `gorm:"type:varchar(30)"json:"update_day"`
	Userid    string `json:"userid"`
	Username  string `json:"username"`
	Class     string `json:"class"`
	Role      int    `json:"role"`
}

// ReportWrite 写实习报告
func ReportWrite(userid string, data *Report) (Report, int) {
	var user User
	db.Where("userid = ?", userid).First(&user)
	db.Where("title = ?", data.Title).Where("userid = ?", user.Userid).First(&data)
	if data.ID == 0 {
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
	return *data, errmsg.ERROR_REPORT_IS_BEING_WIRRTRE
}

// ReportUpdate 修改实习报告
func ReportUpdate(userid string, data *Report) int {
	var report Report
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		var maps = make(map[string]interface{})
		maps["title"] = data.Title
		maps["desc"] = data.Desc
		maps["content"] = data.Content
		// 15:04:05 精确到秒
		maps["update_day"] = time.Now().Format("2006-01-02")
		err = db.Model(&report).Where("id = ? and userid = ?", data.ID, user.Userid).Updates(&maps).Error
		if err != nil {
			return errmsg.ERROR_UPDATE_FAIL
		}
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_USERID_NULL
}

// ReportDelete 老师删除报告
func ReportDelete(userid string, data *Report) int {
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 && user.Class == data.Class {
		err = db.Where("id = ? and class = ?", data.ID, data.Class).Delete(&data).Error
		if err != nil {
			return errmsg.ERROR_DELETE_FAIL
		}
		return errmsg.SUCCESS
	}
	if user.Class != data.Class {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.ERROR_USERID_NULL
}

// ReportCheckClass 查询对应老师班级实习报告
func ReportCheckClass(userid string, pageSize int, pageNum int) ([]Report, int64) {
	var report []Report
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&report)
		db.Model(&report).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return report, total
	}

	err = db.Select("id,userid,username,class,title,role,creat_day,update_day").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&report).Error
	db.Model(&report).Where("role = ?", 2).Count(&total)

	if err != nil {
		return report, 0
	}
	return report, total
}

//// ReportCheckAll 查询所有班级实习报告
//func ReportCheckAll(userid string,pageSize int, pageNum int) ([]Report,int64) {
//	var report []Report
//	var user User
//	var total int64
//	if userid!="" {
//		db.Where(" userid = ?",userid).First(&user)
//		db.Select("id,userid,username,class,title,creat_day,update_day").Where(
//			"class = ? and role = ?",user.Class,2,
//		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&report)
//		db.Model(&report).Where(
//			"class = ? and role = ?", user.Class,2,
//		).Count(&total)
//		return report, total
//	}
//	err = db.Select("id,userid,username,class,title,creat_day,update_day").Where("role = ?",2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&report).Error
//	db.Model(&report).Where("role = ?",2).Count(&total)
//	if err != nil {
//		return report, 0
//	}
//	return report, total
//}

// ReportCheckDay 按日期查询实习报告
func ReportCheckDay(userid string, creatDay string, pageSize int, pageNum int) ([]Report, int64) {
	var report []Report
	var user User
	var total int64
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		err = db.Where("creat_day = ? and role = ? and class = ? ", creatDay, 2, user.Class).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&report).Error
		db.Model(&report).Where("creat_day = ? and role = ? and class = ? ", creatDay, 2, user.Class).Count(&total)
		if err != nil {
			return report, 0
		}
		return report, total
	} else {
		return report, 0
	}
}

// ReportCheckItself 查询自己的实习报告
func ReportCheckItself(userid string) ([]Report, int) {
	var report Report
	var reports []Report
	if userid != "" {
		db.Where(" userid = ?", userid).First(&report)
		db.Where(" userid = ?", userid).Find(&reports)
		if report.ID == 0 {
			return reports, errmsg.ERROR_USER_NOT_EXIST
		}
		return reports, errmsg.SUCCESS
	}
	return reports, errmsg.ERROR_USERID_NULL
}

// ReportCheckuserid 查询对应学号的实习报告
func ReportCheckuserid(userid string) (Report, int) {
	var report Report
	if userid != "" {
		db.Where(" userid = ?", userid).First(&report)
		if report.ID == 0 {
			return report, errmsg.ERROR_USER_NOT_EXIST
		}
		return report, errmsg.SUCCESS
	}
	return report, errmsg.ERROR_USERID_NULL
}
