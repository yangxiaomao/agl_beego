/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-11-03 17:38:29
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
type RsaController struct {
	baseController
}

type SignJson struct {
	Sign string
}

/**
 * @description: rsa加密接口（测试使用）
 * @param {*}	yxm---2020-10-29
 * @return {*}	终端输出string
 */

func (r *RsaController) RsaEncrypting() {
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := r.Ctx.Input.RequestBody
	// beego.Info(reflect.TypeOf(r.GetString("")))
	data, _ := r.RsaEncrypt(jsonDatabytes)
	rsaData := make(map[string]interface{})
	rsaData["sign"] = base64.StdEncoding.EncodeToString(data)
	r.Data["json"] = rsaData
	r.ServeJSON()
}

/**
 * @description: rsa解密接口（测试使用）
 * @param {*}	yxm---2020-10-29
 * @return {*}	终端输出string
 */

func (r *RsaController) RsaDecrypting() {
	var signData SignJson
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
	origData, _ := r.RsaDecrypt(signBytes)

	sign_map := make(map[string]string)
	err = json.Unmarshal(origData, &sign_map)
	if err != nil {
		beego.Info("Rsa解密失败!,失败原因：" + err.Error())
	}
	r.Data["json"] = sign_map
	r.ServeJSON()

}
