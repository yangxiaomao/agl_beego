/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-17 15:53:14
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	"beeapi/util"

	"github.com/astaxie/beego"
)

// Operations about Sort
type SortController struct {
	baseController
}

/**
 * @description:	冒泡排序BubbleSort
 * @param {*SortController} yxm --- 2020-12-16
 * @return {*}	json
 */
func (s *SortController) BubbleSortHandle() {
	arr := []int{5, 8, 9, 6, 4, 1, 3, 8, 10, 1000, 152, 190, 580, 680}
	beego.Info(arr)
	returnData := util.BubbleSort(arr)
	s.Data["json"] = returnData
	s.ServeJSON()
}
