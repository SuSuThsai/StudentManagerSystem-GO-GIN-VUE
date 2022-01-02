package model

import (
	"YamadaKESHE/utils/errmsg"
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	WorkingAbility  int `gorm:"type:int;DEFAULT:0"json:"working_ability"`
	Communication   int `gorm:"type:int;DEFAULT:0"json:"communication"`
	WorkingAttitude int `gorm:"type:int;DEFAULT:0"json:"working_attitude"`
	TotalScore      int `gorm:"type:int;DEFAULT:0"json:"total_score"`
	//term		    string	  `gorm:"type:varchar(20)"json:"term"`
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Class    string `json:"class"`
	Role     int    `json:"role"`
}

// AssessmentWriteAll 导入班级所有考核成绩
func AssessmentWriteAll(userid string, data *Assessment) ([]Assessment, int, int64) {
	var user User
	var users []User
	var total int64
	db.Where("userid = ?", userid).First(&user)
	db.Where("class = ? and role = ?", user.Class, 2).Find(&users)
	if user.ID != 0 {
		for _, u := range users {
			var temp Assessment
			db.Where("userid = ?", u.Userid).First(&temp)
			if temp.ID == 0 {
				data.Userid = u.Userid
				data.Username = u.Username
				data.Class = u.Class
				data.Role = u.Role
				data.TotalScore = data.WorkingAbility/2 + data.Communication/4 + data.WorkingAttitude/4
				err = db.Create(&data).Error
				data.ID++
				if err != nil {
					return nil, errmsg.ERROR_POST_FAIL, 0
				}
			}
		}
		var assessment []Assessment
		db.Where("class = ? and role = ?", user.Class, 2).Find(&assessment)
		db.Model(&assessment).Where("class = ? and role = ?", user.Class, 2).Count(&total)
		return assessment, errmsg.SUCCESS, total
	}
	return nil, errmsg.ERROR_IT_IS_BEING_WIRRTRE, 0
}

// AssessmentUpdate 修改实习考核成绩
func AssessmentUpdate(userid string, data *Assessment) int {
	var user User
	var assessment Assessment
	db.Where("userid = ?", userid).First(&user)
	db.Where("userid = ?", data.Userid).First(&data)
	if data.ID != 0 && user.Class == data.Class {
		var maps = make(map[string]interface{})
		maps["working_ability"] = data.WorkingAbility
		maps["communication"] = data.Communication
		maps["working_attitude"] = data.WorkingAttitude
		maps["TotalScore"] = data.WorkingAbility/2 + data.Communication/4 + data.WorkingAttitude/4
		err = db.Model(&assessment).Where("id = ? and class = ?", data.ID, user.Class).Updates(&maps).Error
		if err != nil {
			return errmsg.ERROR_UPDATE_FAIL
		}
		return errmsg.SUCCESS
	}
	if user.Class != data.Class && data.Class != "" {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.ERROR_USERID_NULL
}

// AssessmentDelete 老师删除评核
func AssessmentDelete(userid string, data *Assessment) int {
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 && user.Class == data.Class {
		err = db.Where("id = ? and class = ?", data.ID, user.Class).Delete(&data).Error
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

// AssessmentCheckClass 查询对应老师班级实习评核
func AssessmentCheckClass(userid string, pageSize int, pageNum int) ([]Assessment, int64) {
	var assessment []Assessment
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Select("id,userid,username,class,role,working_ability,communication,working_attitude").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&assessment)
		db.Model(&assessment).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return assessment, total
	}

	err = db.Select("id,userid,username,class,role,working_ability,communication,working_attitude").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&assessment).Error
	db.Model(&assessment).Where("role = ?", 2).Count(&total)

	if err != nil {
		return assessment, 0
	}
	return assessment, total
}

// AssessmentCheckAll 查询所有班级实习评核
func AssessmentCheckAll(userid string, pageSize int, pageNum int) ([]Assessment, int64) {
	var assessment []Assessment
	var user User
	var total int64
	if userid != "" {
		db.Where(" userid = ?", userid).First(&user)
		db.Select("id,userid,username,class,role,working_ability,communication,working_attitude").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&assessment)
		db.Model(&assessment).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return assessment, total
	}
	err = db.Select("id,userid,username,class,role,working_ability,communication,working_attitude").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&assessment).Error
	db.Model(&assessment).Where("role = ?", 2).Count(&total)
	if err != nil {
		return assessment, 0
	}
	return assessment, total
}

// AssessmentCheckItself 查询自己的实习考核成绩
func AssessmentCheckItself(userid string) ([]Assessment, int) {
	var assessment Assessment
	var assessments []Assessment
	if userid != "" {
		db.Where(" userid = ?", userid).First(&assessment)
		db.Where(" userid = ?", userid).First(&assessments)
		if assessment.ID == 0 {
			return assessments, errmsg.ERROR_USER_NOT_EXIST
		}
		return assessments, errmsg.SUCCESS
	}
	return assessments, errmsg.ERROR_USERID_NULL
}

// AssessmentCheckuserid 查询自己的实习考核成绩
func AssessmentCheckuserid(userid string) ([]Assessment, int) {
	var assessment []Assessment
	if userid != "" {
		db.Where(" userid = ?", userid).First(&assessment)
		if len(assessment) == 0 {
			return assessment, errmsg.ERROR_USER_NOT_EXIST
		}
		return assessment, errmsg.SUCCESS
	}
	return assessment, errmsg.ERROR_USERID_NULL
}
