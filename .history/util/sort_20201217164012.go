/*
 * @Author: your name
 * @Date: 2020-12-17 13:31:53
 * @LastEditTime: 2020-12-17 16:40:12
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/util/sort.go
 */
package util

/**
 * @description:BubbleSort(冒泡排序)
 * @date		2020-12-17
 * @auther		yxm
 */

func BubbleSort(arr []int) []int {
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
 * @description:BucketSort(桶排序)
 * @date		2020-12-17
 * @auther		yxm
 */

func BubbleSort(arr []int) []int {
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
