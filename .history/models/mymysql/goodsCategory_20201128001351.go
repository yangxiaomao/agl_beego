/*
 * @Author: your name
 * @Date: 2020-10-30 01:54:21
 * @LastEditTime: 2020-11-18 21:40:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/user.go
 */
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type GoodsCategory struct {
	Id           int64
	Name         string
	ParentId     int64
	Image        string
	Sort         int64
	Level        int64
	Status       int64
	IsNavigation int64
	Created      int64
	Updated      int64
}

type GoodsCategoryList struct {
	Id       int64
	Name     string
	ParentId int64
	Child    []*GoodsCategoryList `orm:"-"`
}

func (g *GoodsCategory) TableName() string {
	return TableName("goods_category")
}

// 添加商品分类
// 2020-10-31
// yxm
func AddGoodsCategory(g *GoodsCategory, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()
	g.Created = time.Now().Unix()
	id, err = ormObj.Insert(g)
	return
}

// 更新商品分类
// 2020-10-31
// yxm
func UpdateGoodsCategory(g *GoodsCategory, ormObj orm.Ormer) (id int64, err error) {
	g.Updated = time.Now().Unix()

	if _, err = ormObj.Update(g, "Name", "ParentId", "Image", "Sort", "Level", "Status", "IsNavigation", "Updated"); err == nil {
		id = g.Id
	}
	return
}

/**
 * @description: ORM查询，商品分类单条查询
 * @param {*orm.Condition} yxm --- 2020-10-31
 * @return {*} obj *GoodsCategory, err error
 */
func GetGoodsCategoryInfo(cond *orm.Condition) (obj *GoodsCategory, err error) {
	ormObj := orm.NewOrm()
	var goodscategory GoodsCategory
	ormObj.QueryTable(&goodscategory).SetCond(cond).One(&goodscategory)
	return &goodscategory, nil
}

/**
 * @description: ORM查询，商品分类列表查询
 * @param {*} yxm --- 2020-10-30
 * @return {*} obj []User, err error
 */
func GetGoodsCategoryList(cond *orm.Condition) (obj []*GoodsCategoryList, err error) {
	ormObj := orm.NewOrm()
	var goodscategorymodel GoodsCategory
	var goodscategory []*GoodsCategory

	ormObj.QueryTable(&goodscategorymodel).SetCond(cond).All(&goodscategory, "id", "name", "parent_id")
	var categorylist []*GoodsCategoryList
	var TreeAnode []*GoodsCategoryList
	//如果取出的列表有数据，则做树状排序
	if len(goodscategory) > 0 {

		for _, v := range goodscategory {
			var goodscategoryinfo GoodsCategoryList
			goodscategoryinfo.Id = v.Id
			goodscategoryinfo.Name = v.Name
			goodscategoryinfo.ParentId = v.ParentId
			categorylist = append(categorylist, &goodscategoryinfo)

		}

		Anode := categorylist[0]      //父节点
		makeTree(categorylist, Anode) //调用生成tree
		TreeAnode = append(TreeAnode, Anode)
	}

	return TreeAnode, nil
}

func makeTree(Allnode []*GoodsCategoryList, node *GoodsCategoryList) { //参数为父节点，添加父节点的子节点指针切片
	childs, _ := haveChild(Allnode, node) //判断节点是否有子节点并返回
	if childs != nil {

		node.Child = append(node.Child, childs[0:]...) //添加子节点
		for _, v := range childs {                     //查询子节点的子节点，并添加到子节点
			_, has := haveChild(Allnode, v)
			if has {
				makeTree(Allnode, v) //递归添加节点
			}
		}
	}
}

func haveChild(Allnode []*GoodsCategoryList, node *GoodsCategoryList) (childs []*GoodsCategoryList, yes bool) {
	for _, v := range Allnode {
		if v.ParentId == node.Id {
			childs = append(childs, v)
		}
	}
	if childs != nil {
		yes = true
	}
	return
}
