package model

import (
	"YamadaKESHE/utils/errmsg"
	"fmt"
	"gorm.io/gorm"
)

type Gpa struct {
	gorm.Model
	Userid          string `json:"userid"`
	Username        string `json:"username"`
	Class           string `json:"class"`
	Role            int    `json:"role"`
	AcademicScore   int    `gorm:"type:int;DEFAULT:0"json:"academic_score"`
	InternshipScore int    `gorm:"type:int;DEFAULT:0"json:"internship_score"`
	TotalScore      int    `gorm:"type:int;DEFAULT:0"json:"total_score"`
}

// GpaWriteAll 导入班级所有考核成绩
func GpaWriteAll(userid string, data *Gpa) ([]Gpa, int, int64) {
	var user User
	var users []User
	var total int64
	db.Where("userid = ?", userid).First(&user)
	db.Where("class = ? and role = ?", user.Class, 2).Find(&users)
	if user.ID != 0 {
		for _, u := range users {
			var temp Gpa
			db.Where("userid = ?", u.Userid).First(&temp)
			if temp.ID == 0 {
				data.Userid = u.Userid
				data.Username = u.Username
				data.Class = u.Class
				data.Role = u.Role
				data.TotalScore = (data.AcademicScore*3)/4 + data.InternshipScore/4
				err = db.Create(&data).Error
				data.ID++
				if err != nil {
					return nil, errmsg.ERROR_POST_FAIL, 0
				}
			}
		}
		var gpa []Gpa
		db.Where("class = ? and role = ?", user.Class, 2).Find(&gpa)
		db.Model(&gpa).Where("class = ? and role = ?", user.Class, 2).Count(&total)
		return gpa, errmsg.SUCCESS, total
	}
	return nil, errmsg.ERROR_IT_IS_BEING_WIRRTRE, 0
}

// GpaUpdate 修改学习成绩
func GpaUpdate(userid string, data *Gpa) int {
	var user User
	var gpa Gpa
	var assessment Assessment
	db.Where("userid = ?", userid).First(&user)
	db.Where("userid = ?", data.Userid).First(&gpa)
	db.Where("userid = ?", gpa.Userid).First(&assessment)
	if gpa.ID != 0 && user.Class == gpa.Class {
		var maps = make(map[string]interface{})
		maps["academic_score"] = data.AcademicScore
		maps["internship_score"] = assessment.TotalScore
		maps["total_score"] = (data.AcademicScore)*3/4 + assessment.TotalScore/4
		err = db.Model(&gpa).Where("id = ? and class = ?", gpa.ID, user.Class).Updates(&maps).Error
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

// GpaDelete 老师删除评核
func GpaDelete(userid string, data *Gpa) int {
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

// GpaCheckClass 查询对应老师班级学业成绩
func GpaCheckClass(userid string, pageSize int, pageNum int) ([]Gpa, int64) {
	var gpa []Gpa
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Select("id,userid,username,class,role,academic_score,total_score,internship_score").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&gpa)
		db.Model(&gpa).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return gpa, total
	}

	err = db.Select("id,userid,username,class,role,total_score,academic_score,internship_score").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&gpa).Error
	db.Model(&gpa).Where("role = ?", 2).Count(&total)

	if err != nil {
		return gpa, 0
	}
	return gpa, total
}

// GpaCheckAll 查询所有班级学业成绩
func GpaCheckAll(userid string, pageSize int, pageNum int) ([]Gpa, int64) {
	var gpa []Gpa
	var user User
	var total int64
	if userid != "" {
		db.Where(" userid = ?", userid).First(&user)
		db.Select("id,userid,username,class,role,academic_score,internship_score").Where(
			"class = ? and role = ?", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&gpa)
		db.Model(&gpa).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return gpa, total
	}
	err = db.Select("id,userid,username,class,role,academic_score,internship_score").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&gpa).Error
	db.Model(&gpa).Where("role = ?", 2).Count(&total)
	if err != nil {
		return gpa, 0
	}
	return gpa, total
}

// GpaCheckItself 查询自己的学业成绩
func GpaCheckItself(userid string) ([]Gpa, int) {
	var gpa Gpa
	var gpas []Gpa
	if userid != "" {
		db.Where(" userid = ?", userid).First(&gpa)
		db.Where(" userid = ?", userid).Find(&gpas)
		if gpa.ID == 0 {
			return gpas, errmsg.ERROR_USER_NOT_EXIST
		}
		return gpas, errmsg.SUCCESS
	}
	return gpas, errmsg.ERROR_USERID_NULL
}

// GpaCheckuserid 查询自己的学业成绩
func GpaCheckuserid(userid string) ([]Gpa, int) {
	var gpa []Gpa
	if userid != "" {
		db.Where(" userid = ?", userid).First(&gpa)
		if len(gpa) == 0 {
			return gpa, errmsg.ERROR_USER_NOT_EXIST
		}
		return gpa, errmsg.SUCCESS
	}
	return gpa, errmsg.ERROR_USERID_NULL
}

// GpaCheckUpdate 更新对应老师班级学业成绩
func GpaCheckUpdate(userid string) ([]Gpa, int, int64) {
	var user User
	var users []Gpa
	var total int64
	db.Where("userid = ?", userid).First(&user)
	db.Where("class = ? and role = ?", user.Class, 2).Find(&users)
	if user.ID != 0 {
		for _, u := range users {
			var temp Gpa
			var assessment Assessment
			db.Where("userid = ?", u.Userid).First(&temp)
			db.Where("userid = ?", u.Userid).First(&assessment)
			if temp.ID != 0 {
				var maps = make(map[string]interface{})
				maps["internship_score"] = assessment.TotalScore
				fmt.Println("2", assessment.TotalScore)
				maps["total_score"] = (temp.AcademicScore)*3/4 + assessment.TotalScore/4
				err = db.Model(&temp).Where("id = ? and class = ?", u.ID, user.Class).Updates(&maps).Error
				if err != nil {
					return nil, errmsg.ERROR_UPDATE_FAIL, total
				}
			}
		}
		var gpa []Gpa
		db.Where("class = ? and role = ?", user.Class, 2).Find(&gpa)
		db.Model(&gpa).Where("class = ? and role = ?", user.Class, 2).Count(&total)
		return gpa, errmsg.SUCCESS, total
	}
	return nil, errmsg.ERROR_IT_IS_BEING_WIRRTRE, 0
}
