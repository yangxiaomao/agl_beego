/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-10 20:00:19
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	models "beeapi/models/mymysql"
	clientServices "beeapi/services/client"
	"beeapi/util"
	"encoding/json"
)

// Operations about UserLogin
type LoginController struct {
	baseController
}

/**
 * @description: 用户注册接口
 * @param {*}	yxm---2020-10-29
 * @return {*}	json
 */

func (u *LoginController) UserRegister() {
	var obj models.User
	obj.Username = u.GetString("username")
	obj.Password = util.Md5(u.GetString("password"))

	userRes, _ := clientServices.UserRegisterService(&obj)
	u.Data["json"] = userRes
	u.ServeJSON()
}

/**
 * @description: 用户登录接口（手机号）
 * @param {*}	yxm---2020-10-29
 * @return {*}	json
 */

func (u *LoginController) UserLogin() {
	var err error
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := u.Ctx.Input.RequestBody
	paramMap := make(map[string]string)
	err = json.Unmarshal(jsonDatabytes, &paramMap)
	if err != nil {
		clientServices.ReturnJSONData(10001, "参数错误！", make(map[string]interface{}))
	}

	userRes, _ := clientServices.UserLoginService(paramMap["username"], paramMap["password"])
	u.Data["json"] = userRes
	u.ServeJSON()
}

// Get implemented Get() method for AppController.
func (this *LoginController) GetHome() {
	this.TplName = "welcome.html"
}
