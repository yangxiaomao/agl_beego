/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-17 17:25:35
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import (
	"beeapi/util"
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
func (s *SortController) SortHandle() {
	arr := []int{5, 8, 9, 6, 4, 1, 3, 8, 10, 1000, 152, 190, 580, 680}
	typeSort, err := s.GetInt64("type")
	if err != nil {

	}
	var returnData []int
	if typeSort == 1 {
		returnData = util.BubbleSort(arr)
	} else if typeSort == 2 {
		returnData = util.BucketSort(arr)
	}

	s.Data["json"] = returnData
	s.ServeJSON()
}
