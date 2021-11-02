/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-10-31 21:13:39
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	"encoding/base64"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Rsa
type AdminRsaController struct {
	adminBaseController
}

type AdminSignJson struct {
	Sign string
}

/**
 * @description: rsa加密接口（测试使用）
 * @param {*}	yxm---2020-10-31
 * @return {*}	终端输出json
 */

func (r *AdminRsaController) RsaEncrypting() {
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := r.Ctx.Input.RequestBody
	// beego.Info(reflect.TypeOf(r.GetString("")))
	data, _ := r.AdminRsaEncrypt(jsonDatabytes)
	rsaData := make(map[string]interface{})
	rsaData["sign"] = base64.StdEncoding.EncodeToString(data)
	r.Data["json"] = rsaData
	r.ServeJSON()
}

/**
 * @description: rsa解密接口（测试使用）
 * @param {*}	yxm---2020-10-31
 * @return {*}	终端输出json
 */

func (r *AdminRsaController) RsaDecrypting() {
	var signData AdminSignJson
	beego.Info(signData)
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := r.Ctx.Input.RequestBody
	err := json.Unmarshal(jsonDatabytes, &signData)
	if err != nil {
		beego.Info("json.Unmarshal is err:" + err.Error())
	}
	signBytes, err := base64.StdEncoding.DecodeString(signData.Sign)

	if err != nil {
		beego.Info("Rsa解密失败!,失败原因：" + err.Error())
	}
	origData, _ := r.AdminRsaDecrypt(signBytes)
	sign_map := make(map[string]string)
	err = json.Unmarshal(origData, &sign_map)
	if err != nil {
		beego.Info("Rsa解密失败!,失败原因：" + err.Error())
	}
	r.Data["json"] = sign_map
	r.ServeJSON()

}
