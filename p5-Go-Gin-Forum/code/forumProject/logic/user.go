package logic

import (
	"forumProject/dao/mysql"
	"forumProject/models"
	snowflake "forumProject/pkg/sonwflake"
)

func SignUp(p *models.ParamSignUp) (err error) {

	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2.生成UID
	var userID uint64
	if userID, err = snowflake.GetID(); err != nil {
		return err
	}
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}

	// 3.入库
	if err = mysql.InsertUser(user); err != nil {
		return err
	}

	return
}

func Login(p *models.ParamLogin) error {

	// 实例化user
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)

}
