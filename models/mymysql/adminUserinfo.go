/*
 * @Author: your name
 * @Date: 2020-10-31 00:40:09
 * @LastEditTime: 2020-11-28 00:16:02
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/adminUserinfo.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminUserInfo struct {
	Id      int
	UserId  int64
	Sex     int
	IsVip   int
	Created int64
	Updated int64
}

func (au *AdminUserInfo) TableName() string {
	return TableName("admin_user_info")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddAdminUserInfo(au *AdminUserInfo, ormObj orm.Ormer) (id int64, err error) {
	au.Updated = time.Now().Unix()
	au.Created = time.Now().Unix()
	id, err = ormObj.Insert(au)
	return
}
