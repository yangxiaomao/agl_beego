/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-12-02 18:43:13
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

func (m *Md5) TableName() string {
	return TableName("md5")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AdMd5(m *Md5, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}
