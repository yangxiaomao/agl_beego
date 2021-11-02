/*
 * @Author: your name
 * @Date: 2020-10-29 04:15:24
 * @LastEditTime: 2020-10-29 04:15:32
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/signdetail.go
 */
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/**
 * @description: 签到详情页模型
 * @param {*}
 * @return {*}
 */

type SignDetail struct {
	Id       int
	UserId   int64
	GeteType int
	TaskNum  float64
	Days     int64
	Created  int64
	Updated  int64
}

func (s *SignDetail) TableName() string {
	return TableName("sign_detail")
}

// 添加用户签到明细记录
func AddSignDetail(m *SignDetail, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}
