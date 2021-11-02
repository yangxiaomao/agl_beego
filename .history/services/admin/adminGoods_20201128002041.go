/*
 * @Author: your name
 * @Date: 2020-11-01 01:55:10
 * @LastEditTime: 2020-11-28 00:20:41
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/services/admin/adminGoodsCatgoryService.go
 */
package services

import (
	models "beeapi/models/mymysql"
	"encoding/json"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//获取商品列表数据返回
func GetGoodsList() (response map[string]interface{}, err error) {
	returnData := make(map[string]interface{})
	cond := orm.NewCondition()
	cond1 := cond.And("status", 1)
	goodsCategoryList, _ := models.GetGoodsList(cond1)
	if len(goodsCategoryList) == 0 {
		returnData = ReturnJSONData(10001, "失败！", returnData)
		return returnData, nil
	}
	returnData["catgory_list"] = goodsCategoryList
	response = returnData

	return response, nil
}

//添加商品
func AddGoods(jsonDatabytes []uint8) (response map[string]interface{}, err error) {
	// goods := models.Goods{}
	// //r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	// beego.Info(reflect.TypeOf(jsonDatabytes))
	// err = json.Unmarshal(jsonDatabytes, &goods)
	// if err != nil {
	// 	beego.Info("json.Unmarshal is err:" + err.Error())
	// }
	// beego.Info(goods)
	// returnData := make(map[string]interface{})
	// //延迟调用处理异常
	// defer func() {
	// 	recovered := recover()
	// 	if recovered != nil {
	// 		response = ReturnJSONData(10001, "失败", returnData)
	// 	}
	// }()
	// ormObj := orm.NewOrm()
	// categoryId, err := models.AddGoods(&goods, ormObj)
	// if err != nil {
	// 	return
	// }
	// returnData["id"] = categoryId
	// response = ReturnJSONData(200, "成功", returnData)
	return response, nil
}

//修改商品分类
func EditGoods(jsonDatabytes []uint8) (response map[string]interface{}, err error) {
	paramData := GoodsCate{}
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	beego.Info(reflect.TypeOf(jsonDatabytes))
	err = json.Unmarshal(jsonDatabytes, &paramData)
	if err != nil {
		beego.Info("json.Unmarshal is err:" + err.Error())
	}
	beego.Info(paramData)
	returnData := make(map[string]interface{})
	//延迟调用处理异常
	defer func() {
		recovered := recover()
		if recovered != nil {
			beego.Info(recovered)
			response = ReturnJSONData(10001, "失败", returnData)
		}
	}()
	ormObj := orm.NewOrm()
	goodscategory := models.GoodsCategory{}
	goodscategory.Id = paramData.Id
	goodscategory.Name = paramData.Name
	goodscategory.ParentId = paramData.ParentId
	goodscategory.Image = paramData.Image
	goodscategory.Sort = paramData.Sort
	goodscategory.Level = paramData.Level
	goodscategory.Status = paramData.Status
	goodscategory.IsNavigation = paramData.IsNavigation
	categoryId, err := models.UpdateGoodsCategory(&goodscategory, ormObj)
	if err != nil {
		return
	}
	returnData["id"] = categoryId
	response = ReturnJSONData(200, "成功", returnData)
	return response, nil
}
