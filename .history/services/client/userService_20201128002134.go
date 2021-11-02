package services

import (
	models "beeapi/models/mymysql"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

//封装json数据结构返回
func ReturnJSONData(code int16, msg string, data map[string]interface{}) (response map[string]interface{}) {
	userObj := make(map[string]interface{})
	userObj["code"] = code
	userObj["msg"] = msg
	userObj["data"] = data
	response = userObj
	return response
}

/**
 * @description: 用户签到服务
 * @param {*}	yxm --- 2020-10-30
 * @return {*} response map[string]interface{}, err error
 */
func UserSignInService(userId int64) (response map[string]interface{}, err error) {
	//定义创建时间
	var createat int64 = 0
	//定义已签到天数
	var days int64 = 0
	//定义签到后天数
	var signdays int64 = 1
	//定义赠送毛豆数量
	var soybeansum float64 = 0.50

	//用户签到
	o := orm.NewOrm()
	signdetail := models.SignDetail{}
	o.QueryTable(new(models.SignDetail).TableName()).Filter("UserId", userId).OrderBy("-id").One(&signdetail)

	if signdetail.Id == 0 {
		createat = 0
		days = 0
	} else {
		createat = signdetail.Created
		days = signdetail.Days
	}
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")

	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().UTC()
	//获取昨天零点时间戳
	yesterdaytime := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, loc).Unix()
	//获取当天零点时间戳
	dayzerotime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).Unix()
	//获取当天23:59:59秒
	dayaftertime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, loc).Unix()
	//如果用户签到记录存在，并且创建时间在今天，则说明今天已签到
	if signdetail.Id > 0 && createat >= dayzerotime && createat <= dayaftertime {
		response = ReturnJSONData(10001, "今天已签到", make(map[string]interface{}))
		return response, err
	}
	//检测用户昨天是否签到
	if signdetail.Id > 0 && createat >= yesterdaytime && createat < dayzerotime {
		//判断用户签到是否是连续7天，如果是，归零
		if days != 7 {
			signdays = days + 1
		}
	}
	//查询用户资金表，查看用户当前毛豆数量
	useramount := models.UserAmount{}
	o.QueryTable(new(models.UserAmount).TableName()).Filter("UserId", userId).One(&useramount)
	if useramount.Id == 0 {
		response = ReturnJSONData(10001, "账户异常", make(map[string]interface{}))
		return response, err
	}
	o.Begin()
	//延迟调用处理事务
	defer func() {
		recovered := recover()
		if recovered != nil {
			o.Rollback()
			beego.Info(recovered)
		}
		if err == nil {
			o.Commit()
		} else {
			o.Rollback()
		}
	}()

	//计算之后毛豆数量
	beforesoybeansum := decimal.NewFromFloat(useramount.Soybean)
	//签到送毛豆数量
	signsoybeansum := decimal.NewFromFloat(soybeansum)

	aftersoybeansum := beforesoybeansum.Add(signsoybeansum) //beforesoybeansum + signsoybeansum
	afterSoybeanFloat64, _ := aftersoybeansum.Float64()
	//用户毛豆变更记录
	soybeandetail := models.SoybeanDetail{}
	soybeandetail.UserId = userId
	soybeandetail.OptSoybean = soybeansum
	soybeandetail.BeforeSoybean = useramount.Soybean
	soybeandetail.AfterSoybean = afterSoybeanFloat64
	soybeandetail.OptType = 1
	soybeandetail.Source = 1
	soybeandetail.OrderId = 0
	// beego.Info(reflect.TypeOf(aftersoybeansum))
	soybeanId, err := models.AddSoybeanDetail(&soybeandetail, o)
	if err != nil {
		return
	}
	//更新用户资金表，更新用户毛豆数量
	useramount.Soybean = afterSoybeanFloat64
	amountId, err := o.Update(&useramount)
	if err != nil {
		return
	}
	//用户签到记录
	newSignDetail := models.SignDetail{}
	newSignDetail.UserId = userId
	newSignDetail.GeteType = 1
	newSignDetail.TaskNum = soybeansum
	newSignDetail.Days = signdays
	newSignId, err := models.AddSignDetail(&newSignDetail, o)
	if err != nil {
		return
	}

	if soybeanId > 0 && amountId > 0 && newSignId > 0 {
		successData := make(map[string]interface{})

		var surplus int64 = 0
		if signdays == 7 {
			surplus = signdays
		} else {
			surplus = 7 - signdays
		}

		successData["sgin_days"] = signdays
		successData["surplus"] = surplus
		successData["tips"] = "恭喜您-签到成功，获得" + strconv.FormatFloat(soybeansum, 'g', -1, 64) + "克毛豆"
		response = ReturnJSONData(200, "成功", successData)
	} else {
		response = ReturnJSONData(10001, "失败", make(map[string]interface{}))
	}
	return response, err

}
