/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-02 18:28:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	models "beeapi/models/mymysql"
	clientServices "beeapi/services/client"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Operations about Users
type AdminMd5Controller struct {
	adminBaseController
}

/**
 * @description: 	根据Md5密文查询Md5原文
 * @param {*UserController} yxm --- 2020-10-29
 * @return {*}	json
 */
func (u *AdminMd5Controller) UserPersonalCenter() {
	verifUser := u.verifAdminIdentity()
	verifUserData := make(map[string]interface{})
	if verifUser.code == 10001 {
		returnData := clientServices.ReturnJSONData(verifUser.code, verifUser.msg, verifUserData)
		u.Data["json"] = returnData
		u.ServeJSON()
		return
	}
	beego.Info(verifUser)
	u.o = orm.NewOrm()
	users := []*models.User{}

	u.o.QueryTable(new(models.User).TableName()).Filter("Id", u.user_id).One(&users)
	u.Data["json"] = users

	u.ServeJSON()
}
