/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-11-28 00:21:56
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

// Operations about Game
type GameController struct {
	baseController
}

/**
 * @description: 	3D旋转拼图小游戏
 * @param {*GameController} yxm --- 2020-12-16
 * @return {*}	json
 */
func (g *GameController) RotatingJigsaw() {
	verifUser := u.verifUserIdentity()
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

/**
 * @description: 用户签到
 * @param {*}	 yxm --- 2020-10-29
 * @return {*}	 json
 */

func (u *UserController) UserSignIn() {
	verifUserData := make(map[string]interface{})
	verifUser := u.verifUserIdentity()
	if verifUser.code == 10001 {
		returnData := clientServices.ReturnJSONData(verifUser.code, verifUser.msg, verifUserData)
		u.Data["json"] = returnData
		u.ServeJSON()
		return
	}
	serviceData, _ := clientServices.UserSignInService(u.user_id)
	u.Data["json"] = serviceData
	u.ServeJSON()
}
