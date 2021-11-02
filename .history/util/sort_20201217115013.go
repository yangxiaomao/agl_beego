/*
 * @Author: your name
 * @Date: 2020-10-31 02:16:11
 * @LastEditTime: 2020-12-17 11:50:12
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/util/logs.go
 */
package util

/**
 * @description:BubbleSort(冒泡排序)
 * @date		2020-12-17
 * @auther		yxm
 */

func BubbleSort(list []int, left, right int) {
	if right == 0 {
		return
	}
	for index, num := range list {
		if index < right && num > list[index+1] {
			utils.SwapGo(list, index, index+1)
		}
	}
	sort(list, left, right-1)
}
