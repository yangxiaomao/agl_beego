/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-28 00:16:31
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package mymysql

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Goods struct {
	Id              int64
	GoodsTitle      string
	GoodsTag        string
	OneCateId       int64
	TwoCateId       int64
	ThreeCateId     int64
	GoodsCateName   string
	Unit            string
	TrueSales       int64
	GoodsWeight     int64
	VirtualSales    int64
	StartSaleTime   int64
	EndSaleTime     int64
	TotalStock      int64
	GoodsState      int64
	Easybin         int64
	CreateId        int64
	CreateName      string
	AuditStatus     int64
	AuditId         int64
	AuditTime       int64
	IsAppoint       int64
	RefuseReason    string
	StoreId         int64
	StoreName       string
	GoodsType       int64
	IsCoupons       int64
	IsHaveGift      int64
	GoodsVat        int64
	IsOverRefund    int64
	IsAllowActivity int64
	PayType         int64
	Sorts           int64
	IsRecommend     int64
	IsNow           int64
	ServerGuarantee string
	IsPost          int64
	FullAmount      float64
	FullPiece       int64
	NotPostalArea   string
	IsGift          int64
	IsQtian         int64
	Created         int64
	Updated         int64
}

func (g *Goods) TableName() string {
	return TableName("goods")
}

// 添加商品
// 2020-11-03
// yxm
func AddGoods(g *Goods, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()
	g.Created = time.Now().Unix()
	id, err = ormObj.Insert(g)
	return
}

// 更新商品
// 2020-11-03
// yxm
func UpdateGoods(g *Goods, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()

	if _, err = ormObj.Update(g); err == nil {
		id = g.Id
	}
	return
}

/**
 * @description: ORM查询，商品单条查询
 * @param {*orm.Condition} yxm --- 2020-11-03
 * @return {*} obj *GoodsCategory, err error
 */
func GetGoodsInfo(cond *orm.Condition) (obj *Goods, err error) {
	ormObj := orm.NewOrm()
	var goods Goods
	ormObj.QueryTable(&goods).SetCond(cond).One(&goods)
	return &goods, nil
}

/**
 * @description: ORM查询，商品分类列表查询
 * @param {*} yxm --- 2020-10-30
 * @return {*} obj []User, err error
 */
func GetGoodsList(cond *orm.Condition) (obj []*Goods, err error) {
	ormObj := orm.NewOrm()
	var goodsmodel Goods
	var goods []*Goods

	ormObj.QueryTable(&goodsmodel).SetCond(cond).All(&goods)
	return goods, nil
}
