/*
 * @Author: your name
 * @Date: 2020-10-29 04:13:44
 * @LastEditTime: 2020-11-18 21:36:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/models/soybeandetail.go
 */
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type SoybeanDetail struct {
	Id            int64
	UserId        int64
	OptSoybean    float64
	BeforeSoybean float64
	AfterSoybean  float64
	OptType       int
	Source        int
	OrderId       int
	Created       int64
	Updated       int64
}

func (s *SoybeanDetail) TableName() string {
	return TableName("soybean_detail")
}

// 添加毛豆明细记录
func AddSoybeanDetail(m *SoybeanDetail, ormObj orm.Ormer) (id int64, err error) {
	m.Updated = time.Now().Unix()
	m.Created = time.Now().Unix()
	id, err = ormObj.Insert(m)
	return
}
