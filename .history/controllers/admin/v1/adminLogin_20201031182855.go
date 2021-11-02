/*
 * @Author: your name
 * @Date: 2020-10-31 04:21:56
 * @LastEditTime: 2020-10-31 18:28:56
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/admin/v1/adminLogin.go
 */
package controllers

import (
	"beeapi/models"
	adminServices "beeapi/services/admin"
	clientServices "beeapi/services/client"
	"beeapi/util"
	"encoding/json"
)

// Operations about AdminLogin
type AdminLoginController struct {
	adminBaseController
}

// @Title AdminUserRegister
// @Description 后台管理员用户注册接口
// @Param paasid path string true "The paasid name"
// @Param field query string true "field"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @router /paas/:paasid [get]

func (u *AdminLoginController) AdminUserRegister() {
	var obj models.AdminUser
	obj.Username = u.GetString("username")
	obj.Password = util.Md5(u.GetString("password"))

	userRes, _ := adminServices.AdminRegisterService(&obj)
	u.Data["json"] = userRes
	u.ServeJSON()
}

func (u *AdminLoginController) AdminUserLogin() {
	var err error
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := u.Ctx.Input.RequestBody
	paramMap := make(map[string]string)
	err = json.Unmarshal(jsonDatabytes, &paramMap)
	if err != nil {
		clientServices.ReturnJSONData(10001, "参数错误！", make(map[string]interface{}))
	}

	userRes, _ := adminServices.AdminLoginService(paramMap["username"], paramMap["password"])
	u.Data["json"] = userRes
	u.ServeJSON()
}
