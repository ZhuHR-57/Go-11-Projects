package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"forumProject/models"
	"forumProject/settings"
)

func CheckUserExist(username string) (err error) {

	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func InsertUser(user *models.User) (err error) {

	// 对密码加密
	password := encryptPassword(user.Password)

	// 插入
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)

	return
}

func Login(user *models.User) (err error) {

	oldPassword := user.Password

	sqlStr := `select username,password from user where username = ?`
	err = db.Get(user, sqlStr, user.UserName)
	// 一般不会判断不存在，因为不能让用户知道
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		// 数据库错误
		return err
	}

	waitProvePassword := encryptPassword(oldPassword)
	if waitProvePassword != user.Password {
		return errors.New("密码错误")
	}
	return
}

func encryptPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(settings.Conf.Salt))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}
