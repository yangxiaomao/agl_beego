/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-12 17:49:27
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	"beeapi/models/mymongo"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

// Operations about Rsa
type RsaController struct {
	baseController
}

type SignJson struct {
	Sign string
}

// User model definiton.
type User struct {
	Id       string
	Name     string
	Password string
	RegDate  time.Time
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

/**
 * @description:Mongodb数据测试操作（测试使用）
 * @param {*}	yxm---2020-11-27
 * @return {*}	终端输出string
 */

func (r *RsaController) MongoTest() {
	ticker := time.NewTicker(time.Second) // 每隔1s进行一次打印
	var num int64
	num = 1
	for {
		<-ticker.C

		a := []int{}
		dataArr := a[0:]
		dataArr = append(dataArr, 10)
		beego.Info(len(dataArr))
		num++
	}

	var code int
	var err error
	mConn := mymongo.Conn()
	defer mConn.Close()

	var user User
	user.Id = "2"
	user.Name = "yangxiaomao"
	user.Password = "123456"

	c := mConn.DB("db_shop").C("md5_01")
	err = c.Insert(&user)

	if err != nil {
		beego.Info(err.Error())
		if mgo.IsDup(err) {
			code = 200
		} else {
			code = 100
		}
	} else {
		code = 0
	}
	sign_map := make(map[string]int)
	sign_map["code"] = code
	r.Data["json"] = sign_map
	r.ServeJSON()

}
