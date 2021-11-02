package services

import (
	models "beeapi/models/mymysql"
	"beeapi/util"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
)

/**
 * @description: 用户注册服务
 * @param {*models.User} yxm --- 2020-10-29
 * @return {*} response map[string]interface{}, err error
 */
func UserRegisterService(obj *models.User) (response map[string]interface{}, err error) {
	o := orm.NewOrm()
	returnData := make(map[string]interface{})
	userRes := make(map[string]interface{})
	u1 := uuid.NewV4().String()
	useruuid := util.Md5(u1)
	obj.Uuid = useruuid
	obj.Email = "353125014@qq.com"
	obj.LastIp = util.GetLocalIP()
	obj.LastTime = time.Now().Unix()
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
	// 事务处理过程
	if userid, err := models.AddUser(obj, o); err != nil {
		beego.Error("失败！" + err.Error())
		returnData = ReturnJSONData(10001, "失败！"+err.Error(), userRes)
	} else {
		//用户基本信息添加
		_, err = addUserInfo(userid, o)
		//添加用户资金信息
		_, err = addUserAmount(userid, o)
		userRes, err = UserInfoRedis(userid, useruuid)
		if err != nil {
			returnData = ReturnJSONData(10001, "失败！"+err.Error(), userRes)
		} else {
			returnData = ReturnJSONData(200, "成功", userRes)
		}

	}
	response = returnData
	return response, nil
}

/**
 * @description: 用户基本信息入redis
 * @param {*} yxm --- 2020-10-29
 * @return {*} response map[string]interface{}, err error
 */

func UserInfoRedis(userid int64, useruuid string) (response map[string]interface{}, err error) {
	returnData := make(map[string]interface{})
	randint := rand.Intn(10000000000000)
	//生成token
	tokenstring := strconv.Itoa(randint)
	usertoken := util.Md5(tokenstring)

	returnData["user_uuid"] = useruuid
	returnData["user_token"] = usertoken
	returnData["user_id"] = userid
	redisPool, err := util.GetRedisConnection(2)
	if err != nil {
		return
	}
	defer redisPool.Close() //函数运行结束 ，把连接放回连接池

	jsonDataSeq, err := json.Marshal(returnData)
	if err != nil {
		return
	}
	redisKey := "login:token:uuid:" + useruuid

	_, error := redisPool.Do("Set", redisKey, jsonDataSeq)
	if error != nil {
		return
	}
	redisPool.Close() //关闭连接池
	return returnData, nil
}

/**
 * @description: 用户基本信息入库
 * @param {*} yxm --- 2020-10-29
 * @return {*} response int64, err error
 */

func addUserInfo(userid int64, ormObj orm.Ormer) (response int64, err error) {
	UserInfo := models.UserInfo{}
	UserInfo.UserId = userid
	UserInfo.Sex = 0
	UserInfo.IsVip = 0
	userinfoid, err := models.AddUserInfo(&UserInfo, ormObj)
	if err != nil {
		return
	}
	response = userinfoid
	return response, nil
}

/**
 * @description: 用户资金信息入库
 * @param {*} yxm --- 2020-10-29
 * @return {*} response int64, err error
 */

func addUserAmount(userid int64, ormObj orm.Ormer) (response int64, err error) {
	UserAmount := models.UserAmount{}
	UserAmount.UserId = userid
	UserAmount.Soybean = 0
	UserAmount.Integral = 0
	useramountid, err := models.AddUserAmount(&UserAmount, ormObj)
	if err != nil {
		return
	}
	response = useramountid
	return response, nil
}

/**
 * @description: 用户登录服务
 * @param {*} yxm --- 2020-10-29
 * @return {*} response int64, err error
 */
func UserLoginService(username string, password string) (response map[string]interface{}, err error) {
	userParam := make(map[string]interface{})
	returnData := make(map[string]interface{})
	userRes := make(map[string]interface{})
	// //自定义条件查询
	// cond := orm.NewCondition()
	// cond1 := cond.And("Username", username).And("Password", util.Md5(password))
	// userinfo, err := models.GetUserInfo(cond1)
	userParam["field"] = "u.id,u.username,u.uuid,ui.sex,ua.soybean"
	userParam["where"] = "u.username = ? and u.password = ?"
	userParam["username"] = username
	userParam["password"] = util.Md5(password)
	userObj, _ := models.GetUserJoinInfo(userParam)
	if userObj.Id == 0 {
		returnData = ReturnJSONData(10001, "失败！", returnData)
		return returnData, nil
	}

	userRes, err = UserInfoRedis(userObj.Id, userObj.Uuid)
	if err != nil {
		return
	}
	userSoybean := fmt.Sprintf("%.2f", userObj.Soybean)
	beego.Info(userSoybean)
	returnData["user_id"] = userObj.Id
	returnData["user_name"] = userObj.Username
	returnData["user_uuid"] = userObj.Uuid
	returnData["user_token"] = userRes["user_token"]
	returnData["user_sex"] = userObj.Sex
	returnData["soybean"] = userSoybean
	response = returnData
	return response, nil

}
