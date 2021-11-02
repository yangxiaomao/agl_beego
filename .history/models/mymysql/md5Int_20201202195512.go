/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-12-02 19:46:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Md5Int struct {
	Id             int64
	OriginalString string
	DenseString    string
	Created        int64
	Updated        int64
}

func (m *Md5Int) TableName() string {
	return TableName("md5_int")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AdMd5Int(m *Md5Int, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}

// UpdateUser update User into database and returns id on success
func UpdateMd5Int(m *orm.Params, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.Id
	}
	return
}
