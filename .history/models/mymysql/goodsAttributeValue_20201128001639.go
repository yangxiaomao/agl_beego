/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-28 00:16:39
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type GoodsAttributeValue struct {
	Id              int64
	AttributeName   string
	GoodsCategoryId int64
	Optional        int64
	Sort            int64
	IsDelete        int64
	Created         int64
	Updated         int64
}

func (g *GoodsAttributeValue) TableName() string {
	return TableName("goods_attribute_value")
}

// 添加商品属性值
// 2020-11-03
// yxm
func AddGoodsAttributeValue(g *GoodsAttributeValue, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()
	g.Created = time.Now().Unix()
	id, err = ormObj.Insert(g)
	return
}

// 更新商品属性值
// 2020-11-03
// yxm
func UpdateGoodsAttributeValue(g *GoodsAttributeValue, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()

	if _, err = ormObj.Update(g); err == nil {
		id = g.Id
	}
	return
}

/**
 * @description: ORM查询，商品属性值单条查询
 * @param {*orm.Condition} yxm --- 2020-11-03
 * @return {*} obj *GoodsSku, err error
 */
func GetGoodsAttributeValueInfo(cond *orm.Condition) (obj *GoodsAttributeValue, err error) {
	ormObj := orm.NewOrm()
	var goodsattributevalue GoodsAttributeValue
	ormObj.QueryTable(&goodsattributevalue).SetCond(cond).One(&goodsattributevalue)
	return &goodsattributevalue, nil
}
