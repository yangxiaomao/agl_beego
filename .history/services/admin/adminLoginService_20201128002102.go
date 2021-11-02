package services

import (
	models "beeapi/models/mymysql"
	"beeapi/util"
	"encoding/json"
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
func AdminRegisterService(obj *models.AdminUser) (response map[string]interface{}, err error) {
	o := orm.NewOrm()
	returnData := make(map[string]interface{})
	userRes := make(map[string]interface{})
	u1 := uuid.NewV4().String()
	adminuuid := util.Md5(u1)
	obj.Uuid = adminuuid
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
	if userid, err := models.AddAdminUser(obj, o); err != nil {
		beego.Error("失败！" + err.Error())
		returnData = ReturnJSONData(10001, "失败！"+err.Error(), userRes)
	} else {
		//用户基本信息添加
		_, err = addAdminInfo(userid, o)
		userRes, err = AdminUserInfoRedis(userid, adminuuid)
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

func AdminUserInfoRedis(userid int64, adminuuid string) (response map[string]interface{}, err error) {
	returnData := make(map[string]interface{})
	randint := rand.Intn(10000000000000)
	//生成token
	tokenstring := strconv.Itoa(randint)
	admintoken := util.Md5(tokenstring)

	returnData["admin_uuid"] = adminuuid
	returnData["admin_token"] = admintoken
	returnData["admin_id"] = userid
	redisPool, err := util.GetRedisConnection(2)
	if err != nil {
		return
	}
	defer redisPool.Close() //函数运行结束 ，把连接放回连接池

	jsonDataSeq, err := json.Marshal(returnData)
	if err != nil {
		return
	}
	redisKey := "login:admin:uuid:" + adminuuid

	_, error := redisPool.Do("Set", redisKey, jsonDataSeq)
	if error != nil {
		return
	}
	redisPool.Close() //关闭连接池
	return returnData, nil
}

/**
 * @description: 管理员用户基本信息入库
 * @param {*} yxm --- 2020-10-30
 * @return {*} response int64, err error
 */

func addAdminInfo(userid int64, ormObj orm.Ormer) (response int64, err error) {
	AdminInfo := models.AdminUserInfo{}
	AdminInfo.UserId = userid
	AdminInfo.Sex = 0
	AdminInfo.IsVip = 0
	adminInfoid, err := models.AddAdminUserInfo(&AdminInfo, ormObj)
	if err != nil {
		return
	}
	response = adminInfoid
	return response, nil
}

/**
 * @description: 用户登录服务
 * @param {*} yxm --- 2020-10-29
 * @return {*} response int64, err error
 */
func AdminLoginService(username string, password string) (response map[string]interface{}, err error) {
	userParam := make(map[string]interface{})
	returnData := make(map[string]interface{})
	userRes := make(map[string]interface{})
	// //自定义条件查询
	// cond := orm.NewCondition()
	// cond1 := cond.And("Username", username).And("Password", util.Md5(password))
	// userinfo, err := models.GetUserInfo(cond1)
	userParam["field"] = "u.id,u.username,u.uuid"
	userParam["where"] = "u.username = ? and u.password = ?"
	userParam["username"] = username
	userParam["password"] = util.Md5(password)
	userObj, _ := models.GetAdminUserJoinInfo(userParam)
	if userObj.Id == 0 {
		returnData = ReturnJSONData(10001, "失败！", returnData)
		return returnData, nil
	}

	userRes, err = AdminUserInfoRedis(userObj.Id, userObj.Uuid)
	if err != nil {
		return
	}
	returnData["admin_id"] = userObj.Id
	returnData["admin_name"] = userObj.Username
	returnData["admin_uuid"] = userObj.Uuid
	returnData["admin_token"] = userRes["admin_token"]
	response = returnData
	return response, nil

}

//封装json数据结构返回
func ReturnJSONData(code int16, msg string, data map[string]interface{}) (response map[string]interface{}) {
	userObj := make(map[string]interface{})
	userObj["code"] = code
	userObj["msg"] = msg
	userObj["data"] = data
	response = userObj
	return response
}
