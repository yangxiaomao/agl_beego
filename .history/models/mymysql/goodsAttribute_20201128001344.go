/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-18 21:39:43
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type GoodsAttribute struct {
	Id              int64
	AttributeName   string
	GoodsCategoryId int64
	Optional        int64
	Sort            int64
	IsDelete        int64
	Created         int64
	Updated         int64
}

func (g *GoodsAttribute) TableName() string {
	return TableName("goods_attribute")
}

// 添加商品SKU
// 2020-11-03
// yxm
func AddGoodsAttribute(g *GoodsAttribute, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()
	g.Created = time.Now().Unix()
	id, err = ormObj.Insert(g)
	return
}

// 更新商品
// 2020-11-03
// yxm
func UpdateGoodsAttribute(g *GoodsAttribute, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()

	if _, err = ormObj.Update(g); err == nil {
		id = g.Id
	}
	return
}

/**
 * @description: ORM查询，商品属性单条查询
 * @param {*orm.Condition} yxm --- 2020-11-03
 * @return {*} obj *GoodsSku, err error
 */
func GetGoodsAttributeInfo(cond *orm.Condition) (obj *GoodsAttribute, err error) {
	ormObj := orm.NewOrm()
	var goodsattribute GoodsAttribute
	ormObj.QueryTable(&goodsattribute).SetCond(cond).One(&goodsattribute)
	return &goodsattribute, nil
}
