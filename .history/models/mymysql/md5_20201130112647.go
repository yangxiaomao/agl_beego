/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-30 11:26:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Md5 struct {
	Id         int64
	Uuid       string
	Username   string
	Password   string
	Email      string
	LoginCount int64
	LastTime   int64
	LastIp     string
	State      int64
	Created    int64
	Updated    int64
}

func (m *Md5) TableName() string {
	return TableName("md5")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(m *Md5, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}
