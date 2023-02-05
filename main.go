package main

import "fmt"

func QuickSort(nums []int, left, right int){
	mid := PartSort(nums, left,right)
	QuickSort(nums, left, mid - 1)
	QuickSort(nums, mid + 1, right)

}

func PartSort(nums []int, left, right int)int{

	key := nums[left]

	for left < right{
		if left < right && nums[mi]


	}


	return left
}

func missingNumber(nums []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums) + 1;i++{
		mp[i] = 0
	}
	for i := 0; i < len(nums);i++{
		mp[nums[i]]++
	}
	for i, v := range mp{
		if v == 0{
			return i
		}
	}
	return 0
}


func main() {

	missingNumberList := []int{3,0,1}
	missingNumberList = []int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(missingNumberList))
	return

	fmt.Println("hello world")
}
