/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-11-18 05:40:40
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	adminServices "beeapi/services/admin"
)

// Operations about GoodsGategoryController
type AdminGoodsController struct {
	adminBaseController
}

/**
 * @description: 商品列表
 * @param {*}	yxm---2020-11-03
 * @return {*}	终端输出json
 */

func (g *AdminGoodsController) GoodsList() {
	verifAdminData := make(map[string]interface{})
	verifAdmin := g.verifAdminIdentity()
	if verifAdmin.code == 10001 {
		returnData := adminServices.ReturnJSONData(verifAdmin.code, verifAdmin.msg, verifAdminData)
		g.Data["json"] = returnData
		g.ServeJSON()
		return
	}
	//获取商品列表
	serviceData, _ := adminServices.GetGoodsList()
	g.Data["json"] = serviceData
	g.ServeJSON()
}

/**
 * @description: 商品添加
 * @param {*}  yxm---2020-11-03
 * @return {*}
 */
func (g *AdminGoodsController) AddGoods() {
	verifAdminData := make(map[string]interface{})
	verifAdmin := g.verifAdminIdentity()
	if verifAdmin.code == 10001 {
		returnData := adminServices.ReturnJSONData(verifAdmin.code, verifAdmin.msg, verifAdminData)
		g.Data["json"] = returnData
		g.ServeJSON()
		return
	}
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := g.Ctx.Input.RequestBody
	serviceData, _ := adminServices.AddGoods(jsonDatabytes)
	g.Data["json"] = serviceData
	g.ServeJSON()
}

/**
 * @description: 商品修改
 * @param {*}  yxm---2020-11-03
 * @return {*}
 */
func (g *AdminGoodsController) UpdateGoods() {
	verifAdminData := make(map[string]interface{})
	verifAdmin := g.verifAdminIdentity()
	if verifAdmin.code == 10001 {
		returnData := adminServices.ReturnJSONData(verifAdmin.code, verifAdmin.msg, verifAdminData)
		g.Data["json"] = returnData
		g.ServeJSON()
		return
	}
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := g.Ctx.Input.RequestBody
	//获取商品分类列表
	serviceData, _ := adminServices.EditGoods(jsonDatabytes)
	g.Data["json"] = serviceData
	g.ServeJSON()
}
