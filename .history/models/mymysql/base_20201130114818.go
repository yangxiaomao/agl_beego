/*
 * @Author: your name
 * @Date: 2020-11-18 21:37:18
 * @LastEditTime: 2020-11-28 00:16:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/models/base.go
 */
package mymysql

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {

		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Config), new(AdminUser), new(AdminUserInfo), new(Goods), new(GoodsAttribute), new(GoodsAttributeValue), new(GoodsCategory), new(User), new(UserInfo),
		new(UserAmount), new(SoybeanDetail), new(SignDetail), new(GoodsSku), new(Md5))
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
