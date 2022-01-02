package model

import (
	"YamadaKESHE/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type Demand struct {
	gorm.Model
	CompanyName         string `gorm:"type:varchar(20);not null"json:"company_name"`
	Station             string `gorm:"type:varchar(20);not nul"json:"station"`
	Record              string `gorm:"type:varchar(20);"json:"record"`
	CreatDay            string `gorm:"type:varchar(40);"json:"creat_day"`
	EndDay              string `gorm:"type:varchar(40);"json:"end_day"`
	UpdateDay           string `gorm:"type:varchar(30)"json:"update_day"`
	PositionDescription string `gorm:"type:varchar(110);"json:"position_description"`
	DemandProfession    string `gorm:"type:varchar(20);"json:"demand_profession"`
	OtherText           string `gorm:"type:varchar(110);"json:"other_text"`
	Userid              string `json:"userid"`
	Username            string `json:"username"`
	Class               string `json:"class"`
	Role                int    `json:"role"`
}

// DemandWrite 写企业需求
func DemandWrite(userid string, data *Demand) (Demand, int) {
	var user User
	//var demand Demand
	db.Where("userid = ?", userid).First(&user)
	db.Where("station = ? and company_name = ? and userid = ?", data.Station, data.CompanyName, userid).First(&data)
	//db.Where("company_name = ? and userid",data.Station,user.Userid).First(&demand)
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
	//if demand.ID==0 {
	//	return *data,errmsg.ERROR_USER_NOT_CHARGE_FOR_IT
	//}
	return *data, errmsg.ERROR_STATION_IS_BEING_WIRRTRE
}

// DemandUpdate 修改企业需求
func DemandUpdate(userid string, data *Demand) int {
	var demand Demand
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 && data.Userid == user.Userid {
		var maps = make(map[string]interface{})
		maps["company_name"] = data.CompanyName
		maps["station"] = data.Station
		maps["record"] = data.Record
		maps["end_day"] = data.EndDay
		maps["position_description"] = data.PositionDescription
		maps["demand_profession"] = data.DemandProfession
		maps["other_text"] = data.OtherText
		// 15:04:05 精确到秒
		maps["update_day"] = time.Now().Format("2006-01-02")
		err = db.Model(&demand).Where("id = ? and userid = ?", data.ID, user.Userid).Updates(&maps).Error
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

// DemandDelete 老师删除企业需求
func DemandDelete(userid string, data *Demand) int {
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

// DemandCheckClass 查询对应老师班级企业需求
func DemandCheckClass(userid string, pageSize int, pageNum int) ([]Demand, int64) {
	var demand []Demand
	var user User
	var total int64
	db.Where(" userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Where(
			"class = ? and role = ?", user.Class, 1,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&demand)
		db.Model(&demand).Where(
			"class = ? and role = ?", user.Class, 1,
		).Count(&total)
		return demand, total
	}

	err = db.Where("role = ?", 1).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&demand).Error
	db.Model(&demand).Where("role = ?", 1).Count(&total)

	if err != nil {
		return demand, 0
	}
	return demand, total
}

// DemandCheckAll 查询所有企业需求
func DemandCheckAll(userid string, pageSize int, pageNum int) ([]Demand, int64) {
	var demand []Demand
	var user User
	var total int64
	if userid != "" {
		db.Where(" userid = ?", userid).First(&user)
		db.Where(
			"class = ? and role = ?", user.Class, 1,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&demand)
		db.Model(&demand).Where(
			"class = ? and role = ?", user.Class, 1,
		).Count(&total)
		return demand, total
	}
	err = db.Where("role = ?", 1).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&demand).Error
	db.Model(&demand).Where("role = ?", 1).Count(&total)
	if err != nil {
		return demand, 0
	}
	return demand, total
}

// DemandCheckDay 按日期查询企业需求
func DemandCheckDay(creatDay string, pageSize int, pageNum int) ([]Demand, int64) {
	var demand []Demand
	var total int64
	err = db.Where("creat_day = ? and role = ?", creatDay, 1).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&demand).Error
	db.Model(&demand).Where("creat_day = ? and role = ?", creatDay, 1).Count(&total)
	if err != nil {
		return demand, 0
	}
	return demand, total
}

// DemandCheckItself 查询自己发表的企业需求
func DemandCheckItself(userid string) (Demand, int) {
	var demand Demand
	if userid != "" {
		db.Where(" userid = ?", userid).First(&demand)
		if demand.ID == 0 {
			return demand, errmsg.ERROR_USER_NOT_EXIST
		}
		return demand, errmsg.SUCCESS
	}
	return demand, errmsg.ERROR_USERID_NULL
}
