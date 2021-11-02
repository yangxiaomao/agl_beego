/*
 * @Author: your name
 * @Date: 2020-11-18 21:36:31
 * @LastEditTime: 2020-11-18 21:36:31
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/models/useramount.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type UserAmount struct {
	Id       int
	UserId   int64
	Soybean  float64
	Integral float64
	Created  int64
	Updated  int64
}

func (m *UserAmount) TableName() string {
	return TableName("user_amount")
}

/*
 * 添加用户资金信息
 * yxm
 * 2020-10-28
 */
func AddUserAmount(m *UserAmount, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}
