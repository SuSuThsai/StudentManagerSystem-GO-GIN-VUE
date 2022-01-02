package model

import (
	"YamadaKESHE/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type Log struct {
	gorm.Model
	Title      string `gorm:"type:varchar(80);not null"json:"title"`
	Manager    string `gorm:"type:varchar(20);not null"json:"manager"`
	TeamMember string `gorm:"type:varchar(80)"json:"team_member"`
	Goal       string `gorm:"type:varchar(40)"json:"goal"`
	Results    string `gorm:"type:longtext"json:"results"`
	CreatDay   string `gorm:"type:varchar(30);not nul;"json:"creat_day"`
	UpdateDay  string `gorm:"type:varchar(30)"json:"update_day"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
	Class      string `json:"class"`
	Role       int    `json:"role"`
}

// WorkLogWrite 写工作日志
func WorkLogWrite(userid string, data *Log) (Log, int) {
	var user User
	db.Where("userid = ?", userid).First(&user)
	db.Where("title = ?", data.Title).Where("userid = ?", user.Userid).First(&data)
	if data.ID == 0 {
		data.Title = user.Username + time.Now().Format("2006-01-02") + "的工作日志"
		data.Manager = user.Username
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
	return *data, errmsg.ERROR_WORKLOG_IS_BEING_WIRRTRE
}

// WorkLogUpdate 修改工作日志
func WorkLogUpdate(userid string, data *Log) int {
	var worklog Log
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 && data.Userid == user.Userid {
		var maps = make(map[string]interface{})
		maps["team_member"] = data.TeamMember
		maps["goal"] = data.Goal
		maps["results"] = data.Results
		//15:04:05 精确到秒
		maps["update_day"] = time.Now().Format("2006-01-02")
		err = db.Model(&worklog).Where("id = ? and userid = ?", data.ID, user.Userid).Updates(&maps).Error
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

//// WorkLogDelete 删除工作日志
//func WorkLogDelete(data *Log) int{
//	var worklog Log
//	db.Where("id = ?",data.ID).First(&worklog)
//	err = db.Where("id = ?", data.ID).Delete(&data).Error
//	if err != nil {
//		return errmsg.ERROR_DELETE_FAIL
//	}
//	return errmsg.SUCCESS
//}

// WorkLogCheckClass 查询对应老师班级工作日志
func WorkLogCheckClass(userid string, pageSize int, pageNum int) ([]Log, int64) {
	var report []Log
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Select("id,userid,username,class,title,manager,results,goal,team_member,creat_day,update_day").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&report)
		db.Model(&report).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return report, total
	}

	err = db.Select("id,userid,username,class,title,manager,results,goal,team_member,creat_day,update_day").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&report).Error
	db.Model(&report).Where("role = ?", 2).Count(&total)

	if err != nil {
		return report, 0
	}
	return report, total
}

//// WorkLogCheckAll 查询所有班级工作日志
//func WorkLogCheckAll(userid string,pageSize int, pageNum int) ([]Log,int64) {
//	var worklog []Log
//	var user User
//	var total int64
//	if userid!="" {
//		db.Where(" userid = ?",userid).First(&user)
//		db.Select("id,userid,username,class,title,manager,results,goal,team_member,creat_day,update_day").Where(
//			"class = ? and role = ?",user.Class,2,
//		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&worklog)
//		db.Model(&worklog).Where(
//			"class = ? and role = ?", user.Class,2,
//		).Count(&total)
//		return worklog, total
//	}
//	err = db.Select("id,userid,username,class,title,manager,results,goal,team_member,creat_day,update_day").Where("role = ?",2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&worklog).Error
//	db.Model(&worklog).Where("role = ?",2).Count(&total)
//	if err != nil {
//		return worklog, 0
//	}
//	return worklog, total
//}

// WorkLogCheckDay 按日期查询工作日志
func WorkLogCheckDay(userid string, creatDay string, pageSize int, pageNum int) ([]Log, int64) {
	var worklog []Log
	var user User
	var total int64
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		err = db.Where("creat_day = ? and role = ? and class = ? ", creatDay, 2, user.Class).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&worklog).Error
		db.Model(&worklog).Where("creat_day = ? and role = ? and class = ? ", creatDay, 2, user.Class).Count(&total)
		if err != nil {
			return worklog, 0
		}
		return worklog, total
	} else {
		return worklog, 0
	}
}

// WorkLogCheckItself 查询自己的工作日志
func WorkLogCheckItself(userid string) ([]Log, int) {
	var worklog Log
	var worklogs []Log
	if userid != "" {
		db.Where(" userid = ?", userid).First(&worklog)
		db.Where(" userid = ?", userid).Find(&worklogs)
		if worklog.ID == 0 {
			return worklogs, errmsg.ERROR_USER_NOT_EXIST
		}
		return worklogs, errmsg.SUCCESS
	}
	return worklogs, errmsg.ERROR_USERID_NULL
}

// WorkLogCheckuserid 查询对应学号的工作日志
func WorkLogCheckuserid(userid string) ([]Log, int) {
	var worklog []Log
	if userid != "" {
		db.Where(" userid = ?", userid).Find(&worklog)
		if len(worklog) == 0 {
			return worklog, errmsg.ERROR_USER_NOT_EXIST
		}
		return worklog, errmsg.SUCCESS
	}
	return worklog, errmsg.ERROR_USERID_NULL
}
