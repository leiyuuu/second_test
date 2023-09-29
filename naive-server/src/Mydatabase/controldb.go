package Mydatabase

import (
	"errors"

	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
	"gorm.io/gorm"
)

func Check_in(access_token interface{}) int32 { //仅返回错误类型，分数直接在服务器代码处用随机数生成  分数范围是1-100
	var find_person Person
	result := DB.Where("access_token = ?", access_token).First(&find_person)
	//  				里面也可以传strcut
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //token无效
		return 1
	}
	if find_person.Has_checkin { //已经签到过了
		return 2
	}
	if !find_person.Has_signin { //没登录怎么能签到呢
		return 3
	}
	find_person.Has_checkin = true
	DB.Save(&find_person)
	//改变为已经签到---------
	return 0
}

func Sign_in(username, password interface{}) (string, int32) { //登录
	var find_person Person
	result := DB.Where("username = ?", username.(string)).First(&find_person)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //用户不存在
		return "", 1
	}
	if find_person.Password != password.(string) { //密码错误
		return "", 2
	}
	if find_person.Has_signin { //已经登录
		return "", 3
	}
	find_person.Has_signin = true
	DB.Save(&find_person)
	//更改对象登录情况
	return find_person.Access_token, 0
}

func Sign_up(username, password interface{}) (string, int32) { //注册

	result := DB.Where("username = ?", username).First(&Person{}) //用户名重复
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", 1
	}
	if password == "" { //密码空
		return "", 2
	}
	if username == "" { //用户名空
		return "", 3
	}
	access_token := Gen_a_token(username, password)
	new_person := Person{
		Username:     username.(string),
		Password:     password.(string),
		Access_token: access_token,
	}
	DB.Create(&new_person) //注册进入数据库
	if Auto_sign_in {
		Sign_in(username, password)
	}
	return access_token, 0
}

func Remove_all_data() { //调试用，删库跑路
	loggenerator.Info("DELETE EVERYTHING IN DATABASE")
	DB.Where("can_delete = ?", true).Delete(&Person{})

}
