/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-18 21:41:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type GoodsSku struct {
	Id              int64
	GoodsId         int64
	GoodsSkuValue   string
	MarketPrice     float64
	SoldPrice       float64
	GoodsCostprice  float64
	Discount        float64
	CostShare       int64
	IsDiscount      int64
	GoodsDiscount   float64
	Stock           int64
	StockBeforeWarn int64
	SkuVideo        string
	SkuVideoicon    string
	SkuThunm        string
	SkuImage        string
	SkuArrt         string
	SkuDetail       string
	VirtualSales    int64
	OnSale          int64
	IsRecommend     int64
	StoreId         int64
	Sort            int64
	IsDefault       int64
	RateType        int64
	Created         int64
	Updated         int64
}

func (g *GoodsSku) TableName() string {
	return TableName("goods_sku")
}

// 添加商品SKU
// 2020-11-03
// yxm
func AddGoodsSku(g *GoodsSku, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()
	g.Created = time.Now().Unix()
	id, err = ormObj.Insert(g)
	return
}

// 更新商品
// 2020-11-03
// yxm
func UpdateGoodsSku(g *GoodsSku, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()

	if _, err = ormObj.Update(g); err == nil {
		id = g.Id
	}
	return
}

/**
 * @description: ORM查询，商品SKU单条查询
 * @param {*orm.Condition} yxm --- 2020-11-03
 * @return {*} obj *GoodsSku, err error
 */
func GetGoodsSkuInfo(cond *orm.Condition) (obj *GoodsSku, err error) {
	ormObj := orm.NewOrm()
	var goodssku GoodsSku
	ormObj.QueryTable(&goodssku).SetCond(cond).One(&goodssku)
	return &goodssku, nil
}
