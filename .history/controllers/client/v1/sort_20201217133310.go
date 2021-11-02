/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-17 13:33:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	"beeapi/util"
)

// Operations about Game
type SortController struct {
	baseController
}

/**
 * @description:	冒泡排序BubbleSort
 * @param {*GameController} yxm --- 2020-12-16
 * @return {*}	json
 */
func (s *SortController) BubbleSort() {
	arr := []int{5, 8, 9, 6, 4, 1, 3, 8, 10}
	returnData := util.BubbleSort(arr)
	s.Data["json"] = returnData
	s.ServeJSON()
}
