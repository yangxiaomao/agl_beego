/*
 * @Author: your name
 * @Date: 2020-12-17 13:31:53
 * @LastEditTime: 2020-12-17 16:52:33
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

func BucketSort(list []int) []int {
	max := max(list)
	min := min(list)
	base := 0
	if min < 0 {
		base = -min
	} else {
		base = min
	}
	max = (max + base) / 10
	min = (min + base) / 10
	bucket := make([][]int, max-min+1)
	var result []int
	for _, value := range list {
		i := (int)((value + base) / 10)
		bucket[i] = append(bucket[i], value)
	}

	for _, value := range bucket {
		if len(value) > 0 {
			quicksort.Sort(value, 0, len(value)-1)
		}
	}

	for _, value := range bucket {
		if len(value) > 0 {
			for _, v := range value {
				result = append(result, v)
			}
		}
	}
	return result
}

func max(list []int) int {
	max := list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}

func min(list []int) int {
	min := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	return min
}

/**
 * @description:QuickSort(快速排序)
 * @date		2020-12-17
 * @auther		yxm
 */

func QuickSort(list []int, left, right int) {
	if right < left {
		return
	}
	flag := list[left]
	start := left
	end := right
	for {
		if start == end {
			break
		}
		for list[end] >= flag && end > start {
			end--
		}
		for list[start] <= flag && end > start {
			start++
		}
		if end > start {
			utils.SwapGo(list, start, end)
		}
	}
	utils.SwapGo(list, left, start)
	Sort(list, left, start-1)
	Sort(list, start+1, right)
}
