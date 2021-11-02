/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-02 18:52:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	models "beeapi/models/mymysql"
	clientServices "beeapi/services/client"
	"encoding/json"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Operations about Users
type AdminMd5Controller struct {
	adminBaseController
}

/**
 * @description: 	根据Md5密文查询Md5原文（int类型原文，最大8位）
 * @param {*UserController} yxm --- 2020-10-29
 * @return {*}	json
 */
func (m *AdminMd5Controller) AdminSearchOriginalInt() {
	verifUser := m.verifAdminIdentity()
	verifUserData := make(map[string]interface{})
	if verifUser.code == 10001 {
		returnData := clientServices.ReturnJSONData(verifUser.code, verifUser.msg, verifUserData)
		m.Data["json"] = returnData
		m.ServeJSON()
		return
	}

	m.o = orm.NewOrm()
	md5int := []*models.Md5Int{}
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := m.Ctx.Input.RequestBody
	paramData := models.Md5Int{}
	beego.Info(jsonDatabytes)
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	beego.Info(reflect.TypeOf(jsonDatabytes))
	var err error
	err = json.Unmarshal(jsonDatabytes, &paramData)
	if err != nil {
		beego.Info("json.Unmarshal is err:" + err.Error())
	}
	beego.Info(paramData)

	m.o.QueryTable(new(models.Md5Int).TableName()).Filter("DenseString", paramData.DenseString).One(&md5int)
	beego.Info(md5int)
	m.Data["json"] = md5int

	m.ServeJSON()
}
