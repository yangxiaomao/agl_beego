/*
 * @Author: your name
 * @Date: 2020-11-28 00:17:14
 * @LastEditTime: 2020-11-28 00:17:15
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/models/mymysql/userinfo.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id      int
	UserId  int64
	Sex     int
	IsVip   int
	Created int64
	Updated int64
}

func (m *UserInfo) TableName() string {
	return TableName("user_info")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUserInfo(m *UserInfo, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}
