/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-17 13:28:11
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

// Operations about Game
type SortController struct {
	baseController
}

/**
 * @description:
 * @param {*GameController} yxm --- 2020-12-16
 * @return {*}	json
 */
func (g *SortController) RotatingJigsaw() {
	arr := []int{5, 8, 9, 6, 4, 1, 3, 8, 10}
	returnData := util.BubbleSort(arr)
	u.Data["json"] = returnData
	u.ServeJSON()
}
