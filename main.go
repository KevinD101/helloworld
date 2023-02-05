package main

import "fmt"

func QuickSort(nums []int, left, right int){
	mid := PartSort(nums, left,right)
	if left >= right{
		return
	}
	QuickSort(nums, left, mid - 1)
	QuickSort(nums, mid + 1, right)
}

func PartSort(nums []int, left, right int)int{
	mid := left
	for left < right{
		for left < right && nums[mid] <= nums[right]{
			right--
		}
		for left < right && nums[mid] > nums[left]{
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[mid] = nums[mid], nums[left]
	return left
}

func missingNumberBack(nums []int, left, right int)int{
	for left < right{
		if left == right -1{
			if left != nums[left] {
				return left
			}
			if right != nums[right]{
				return right
			}
			return right + 1
		}
		mid := (left + right) >> 1
		if mid == nums[mid] && (mid + 1) != nums[mid + 1]{
			return mid + 1
		}else if mid == nums[mid] && (mid + 1) == nums[mid + 1]{
			left = mid
		}else if mid != nums[mid] && (mid - 1) == nums[mid - 1]{
			return mid
		}else {
			return missingNumberBack(nums,left, mid - 1)
		}
	}
	if left == nums[left]{
		return left + 1
	}
	return left
}

func missingNumber(nums []int) int {
	QuickSort(nums, 0, len(nums) - 1)

	return  missingNumberBack(nums,0, len(nums) - 1)

	//for i := 0; i < len(nums); i++{
	//	if i != nums[i] {
	//		return i
	//	}
	//}
	//return len(nums)


	//mp := make(map[int]int)
	//for i := 0; i < len(nums) + 1;i++{
	//	mp[i] = 0
	//}
	//for i := 0; i < len(nums);i++{
	//	mp[nums[i]]++
	//}
	//for i, v := range mp{
	//	if v == 0{
	//		return i
	//	}
	//}
	//return 0
}


func main() {

	missingNumberList := []int{3,0,1}
	//missingNumberList = []int{9,6,4,2,3,5,7,0,1}
	missingNumberList = []int{0,1}
	fmt.Println(missingNumber(missingNumberList))
	return
	{

		fmt.Println("hello world")
	}

}
