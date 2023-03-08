package services

import (
	"app/models"
	"crypto/md5"
	"errors"
	"fmt"
)

func AuthenticateUserForLogin(loginName, password string) (*models.User, error) {
	if len(password) == 0 || len(loginName) == 0 {
		return nil, errors.New("Error:用户或者密码为空")
	}
	data := []byte(password)
	has := md5.Sum(data)
	password = fmt.Sprintf("%x", has)
	v, err := GetUserByPhone(loginName) //数据库查询语句。自己写的
	fmt.Println(v)

	if err != nil {
		return nil, errors.New("Error:未找到该用户")

	} else if v.Password != password {
		return nil, errors.New("Error:密码错误")

	} else {
		return v, nil
	}
}

func GetUserByPhone(name string) (*models.User, error) {
	user := models.User{
		Name:     "fengshaomin",
		Password: "123",
		Phone:    "15110202919",
	}
	return &user, nil
}
