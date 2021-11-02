/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-18 21:36:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
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

type UserJoinInfo struct {
	Id         int64
	Uuid       string
	Username   string
	Password   string
	Email      string
	LoginCount int64
	LastTime   int64
	LastIp     string
	State      int64
	Sex        int64
	Soybean    float64
	Created    int64
	Updated    int64
}

func (m *User) TableName() string {
	return TableName("user")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(m *User, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}

// UpdateUser update User into database and returns id on success
func UpdateUser(m *User, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.Id
	}
	return
}

/**
 * @description: ORM查询，用户表基本查询
 * @param {*orm.Condition} yxm --- 2020-10-30
 * @return {*} obj *User, err error
 */
func GetUserInfo(cond *orm.Condition) (obj *User, err error) {
	ormObj := orm.NewOrm()
	var user User
	ormObj.QueryTable(&user).SetCond(cond).One(&user)
	return &user, nil
}

/**
 * @description: 构造查询，用户表关联用户基本信息表查询
 * @param {*} yxm --- 2020-10-30
 * @return {*} obj []User, err error
 */
func GetUserJoinInfo(userParam map[string]interface{}) (obj *UserJoinInfo, err error) {
	var userJoinInfo UserJoinInfo

	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select(userParam["field"].(string)).
		From("tb_user u").
		LeftJoin("tb_user_info ui").On("u.id = ui.user_id").
		LeftJoin("tb_user_amount ua").On("u.id = ua.user_id").
		Where(userParam["where"].(string)).
		// OrderBy("u.id").Desc().
		Limit(1)

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, userParam["username"].(string), userParam["password"].(string)).QueryRow(&userJoinInfo)
	return &userJoinInfo, nil
}
