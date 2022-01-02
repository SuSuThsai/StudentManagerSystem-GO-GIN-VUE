package model

import (
	"YamadaKESHE/utils"
	"YamadaKESHE/utils/errmsg"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Userid   string `gorm:"type:varchar(12);primary_key;not null"json:"userid" validate:"required,min=10,max=12" label:"账号"`
	Username string `gorm:"type:varchar(20);not null"json:"username" validate:"required,min=2,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null"json:"password"validate:"required,min=6,max=20" label:"密码"`
	Class    string `gorm:"type:varchar(20);not null" json:"class" validate:"required,min=2,max=12" label:"班级"`
	Role     int    `gorm:"type:int;DEFAULT:2"json:"role"validate:"required,gte=2" label:"角色"`
}

type Manager struct {
	Username string `gorm:"type:varchar(20);not null"json:"username" validate:"required,min=2,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null"json:"password" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:0";json:"role" label:"SuperRoot"`
}

// CheckUser 查询用户是否存在
func CheckUser(userid string) (code int) {
	var users User
	db.Select("id").Where("userid = ?", userid).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERID_USED //1001
	}
	return errmsg.SUCCESS
}

// CheckUpUser 更新查询
func CheckUpUser(userid string) (code int) {
	var user User
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		return errmsg.SUCCESS
	} else if user.ID > 0 {
		return errmsg.ERROR_USERID_USED
	} else if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// CreatUser 创建用户
func CreatUser(data *User) (code int) {
	//data.Password = ScryptPW(data.Password)
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

// GetUser 查询用户
func GetUser(userid string) (User, int) {
	var user User
	err = db.Limit(1).Where("userid = ?", userid).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(userid string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var user User
	var total int64
	db.Where("userid = ?", userid).First(&user)
	if user.ID != 0 {
		db.Select("id,userid,username,class,role,created_at").Where(
			"class = ? and role =? ", user.Class, 2,
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"class = ? and role = ?", user.Class, 2,
		).Count(&total)
		return users, total
	}
	err = db.Select("id,userid,username,class,role,created_at").Where("role = ?", 2).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	db.Model(&users).Where("role = ?", 2).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

// DeleteUser 删除用户
func DeleteUser(userid string) int {
	var user User
	fmt.Println("删除用户：", userid)
	err = db.Where("userid = ?", userid).Delete(&user).Error
	//对不存在id没有进行处理
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 编辑用户
func EditUser(userid string, role int) int {
	var user User
	var maps = make(map[string]interface{})
	maps["role"] = role
	err = db.Model(&user).Where("userid = ?", userid).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ChangePassword 修改密码
func ChangePassword(userid string, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password

	err = db.Select("password").Where("userid = ?", userid).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeCreate 密码加密 密码加密方法有：bcrypt，scrypt，加salt hash
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPW(u.Password)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPW(u.Password)
	return nil
}

func ScryptPW(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{255, 22, 18, 33, 99, 66, 25, 11}
	HashPW, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	//HashPW, err :=bcrypt.GenerateFromPassword([]byte(password),KeyLen)

	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw
}

// ValidateLogin 后台登陆验证
func ValidateLogin(userid string, password string) (User, int) {
	var user User
	//var PasswordErr error

	db.Where("userid = ?", userid).First(&user)

	//PasswordErr=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPW(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	//if PasswordErr != nil {
	//	return user,errmsg.ERROR_PASSWORD_WRONG
	//}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NOT_RIGHT
	}
	return user, errmsg.SUCCESS
}

// CheckLoginFront 前台登录
func CheckLoginFront(userid string, password string) (User, int) {
	var user User
	//var PasswordErr error

	db.Where("userid = ?", userid).First(&user)

	//PasswordErr=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPW(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	//if PasswordErr != nil {
	//	return user,errmsg.ERROR_PASSWORD_WRONG
	//}
	return user, errmsg.SUCCESS
}

func CheckLoginLocal(userid string, password string) (Manager, int) {
	nameLocal := utils.LocalUser
	passwordLocal := utils.LocalPassword
	roleLocal := utils.LocalRole
	if nameLocal == userid && passwordLocal == password {
		return Manager{Username: nameLocal, Password: passwordLocal, Role: roleLocal}, errmsg.SUCCESS
	} else if nameLocal != userid {
		return Manager{}, errmsg.ERROR_USER_NOT_EXIST
	}

	return Manager{}, errmsg.ERROR_PASSWORD_WRONG
}
