package models

import (
	"test/mysql"
)

// user Model
type User struct {
	ID       uint   `gorm:"AUTO_INCREMENT"`
	Account  string `json:"account" binding:"required" gorm:"size:18;not null,unique; Comment:账号"`
	Password string `json:"password" binding:"required" gorm:"size:18;not null; Comment:密码"`
	Name     string `json:"name" gorm:"size:255; Comment:姓名"`
	Sex      string `json:"sex" gorm:"size:3; Comment:性别"`
	Age      byte   `json:"age" gorm:"default:18; Comment:年龄"`
	Card     string `json:"card" gorm:"size:11; Comment:卡号"`
	Iphone   string `json:"iphone" gorm:"size:11; Comment:电话号码"`
}

/*
	这个Model的增删改查操作都放在这里
*/
// CreateUserInfo 创建UserInfo

//创建用户
func CreateUser(userInfo *User) (err error) {
	//mysql.DB.Where("account = ?", userInfo.Account).First(&userInfo)
	err = mysql.DB.Create(&userInfo).Error
	return
}

//查询用户列表
func GetUserList() (userList []*User, err error) {
	if err = mysql.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

//查询单个用户详细信息
func GetUserInfo(id string) (userInfo *User, err error) {
	userInfo = new(User)
	if err = mysql.DB.Debug().Where("id=?", id).First(userInfo).Error; err != nil {
		return nil, err
	}
	return
}

//更新用户信息
func UpdateAVisitInfo(userInfo *User) (err error) {
	err = mysql.DB.Save(userInfo).Error
	return
}

//删除用户
func DeleteUser(id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&User{}).Error
	return
}
