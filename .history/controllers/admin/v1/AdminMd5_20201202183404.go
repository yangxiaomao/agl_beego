/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-02 18:28:46
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

type Md5DenseString struct {
	Id           int64
	Name         string
	ParentId     int64
	Image        string
	Sort         int64
	Level        int64
	Status       int64
	IsNavigation int64
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
	md5 := []*models.Md5{}
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := g.Ctx.Input.RequestBody
	paramData := GoodsCate{}
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	beego.Info(reflect.TypeOf(jsonDatabytes))
	err = json.Unmarshal(jsonDatabytes, &paramData)
	if err != nil {
		beego.Info("json.Unmarshal is err:" + err.Error())
	}
	beego.Info(paramData)

	m.o.QueryTable(new(models.Md5).TableName()).Filter("Id", u.user_id).One(&md5)
	m.Data["json"] = md5

	m.ServeJSON()
}
