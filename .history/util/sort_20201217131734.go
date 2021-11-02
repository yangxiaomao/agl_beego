/*
 * @Author: your name
 * @Date: 2020-10-31 02:16:11
 * @LastEditTime: 2020-12-17 13:17:34
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

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

/**
 * @description:变量交换
 * @date 		2020-12-17
 * @auther 		yxm
 */

func Swap(list []int, i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

/**
 * @description:go特有变量交换
 * @date 		2020-12-17
 * @auther 		yxm
 */

func SwapGo(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}

/***
 * go变量高阶交换(不推荐，一般不好理解)
 */
func SwapGoAdvanced(list []int, i, j int) {
	list[i] = list[i] + list[j]
	list[j] = list[i] - list[j]
	list[i] = list[i] - list[j]
}
