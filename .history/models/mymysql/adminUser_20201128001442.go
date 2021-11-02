/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-18 21:38:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminUser struct {
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

func (au *AdminUser) TableName() string {
	return TableName("admin_user")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddAdminUser(au *AdminUser, ormObj orm.Ormer) (id int64, err error) {
	au.Updated = time.Now().Unix()
	au.Created = time.Now().Unix()
	id, err = ormObj.Insert(au)
	return
}

// UpdateUser update User into database and returns id on success
func UpdateAdminUser(au *AdminUser, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(au); err == nil {
		id = au.Id
	}
	return
}

/**
 * @description: ORM查询，用户表基本查询
 * @param {*orm.Condition} yxm --- 2020-10-30
 * @return {*} obj *User, err error
 */
func GetAdminUserInfo(cond *orm.Condition) (obj *AdminUser, err error) {
	ormObj := orm.NewOrm()
	var adminUser AdminUser
	ormObj.QueryTable(&adminUser).SetCond(cond).One(&adminUser)
	return &adminUser, nil
}

/**
 * @description: 构造查询，用户表关联用户基本信息表查询
 * @param {*} yxm --- 2020-10-30
 * @return {*} obj []User, err error
 */
func GetAdminUserJoinInfo(userParam map[string]interface{}) (obj *AdminUser, err error) {
	var adminUser AdminUser

	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select(userParam["field"].(string)).
		From("tb_admin_user u").
		LeftJoin("tb_admin_user_info ui").On("u.id = ui.user_id").
		Where(userParam["where"].(string)).
		// OrderBy("u.id").Desc().
		Limit(1)

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, userParam["username"].(string), userParam["password"].(string)).QueryRow(&adminUser)
	return &adminUser, nil
}
