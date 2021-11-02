/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-12-03 16:46:17
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Md5String00 struct {
	Id             int64
	OriginalString string
	DenseString    string
	SearchCount    int64
	Created        int64
	Updated        int64
}

func (m *Md5String00) TableName() string {
	return TableName("md5_string")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddMd5String(m *Md5String00, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}

// UpdateUser update User into database and returns id on success
func UpdateMd5String(m *Md5String00, ormObj orm.Ormer) (id int64, err error) {
	var num int64
	if num, err = ormObj.Update(m); err == nil {
		id = num
	}
	return
}
