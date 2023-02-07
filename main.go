package main

import (
	"fmt"
	"strconv"
)

func QuickSort(nums []int, left, right int) {
	mid := PartSort(nums, left, right)
	if left >= right {
		return
	}
	QuickSort(nums, left, mid-1)
	QuickSort(nums, mid+1, right)
}

func PartSort(nums []int, left, right int) int {
	mid := left
	for left < right {
		for left < right && nums[mid] <= nums[right] {
			right--
		}
		for left < right && nums[mid] > nums[left] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[mid] = nums[mid], nums[left]
	return left
}

func missingNumberBack(nums []int, left, right int) int {
	for left < right {
		if left == right-1 {
			if left != nums[left] {
				return left
			}
			if right != nums[right] {
				return right
			}
			return right + 1
		}
		mid := (left + right) >> 1
		if mid == nums[mid] && (mid+1) != nums[mid+1] {
			return mid + 1
		} else if mid == nums[mid] && (mid+1) == nums[mid+1] {
			left = mid
		} else if mid != nums[mid] && (mid-1) == nums[mid-1] {
			return mid
		} else {
			return missingNumberBack(nums, left, mid-1)
		}
	}
	if left == nums[left] {
		return left + 1
	}
	return left
}

func missingNumber(nums []int) int {
	QuickSort(nums, 0, len(nums)-1)

	return missingNumberBack(nums, 0, len(nums)-1)

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

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if left == right-1 {
			if nums[right] > nums[left] {
				return right
			}
			break
		}
		mid := (left + right) >> 1
		if mid > 0 && nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nums := make([]int, 0)
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	if k == len(nums) || k == 0 {
		return head
	} else if k > len(nums) {
		k = k % len(nums)
	}
	tmp1 := nums[len(nums)-k:]
	tmp2 := nums[:len(nums)-k]
	p := &ListNode{}
	p.Val = tmp1[0]
	ret := p
	for i, v := range tmp1 {
		if i == 0 {
			continue
		}
		next := &ListNode{}
		next.Val = v
		p.Next = next
		p = next
	}
	for _, v := range tmp2 {
		next := &ListNode{}
		next.Val = v
		p.Next = next
		p = next
	}
	//return ret.Next
	return ret
}

func ShowValue(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	realHead := &ListNode{}
	p := head
	q := head.Next
	cur := realHead
	count := 0
	for q != nil {
		if p.Val != q.Val {
			if count == 0 {
				p.Next = nil //去除连接关系
				cur.Next = p
				cur = p
			}
			p = q
			q = q.Next
			count = 0
		} else {
			count++
			q = q.Next
		}
	}
	//最后一个节点
	if p != nil && count == 0 {
		cur.Next = p
	}
	return realHead.Next
}

func containsNearbyDuplicate(nums []int, k int) bool {

	//mp := make(map[int]int)
	//for i, v := range nums{
	//	if pos, ok := mp[v]; ok && i - pos <= k{
	//		return true
	//	}
	//	mp[v] = i
	//}
	//return false

	////1.暴力解法 超时
	//for i := 0; i < len(nums); i++ {
	//	for j := 0; j < len(nums); j++ {
	//		if j == i {
	//			continue
	//		}
	//		var tmp float64 = float64(i - j)
	//		if math.Abs(tmp) <= float64(k) && nums[i] == nums[j] {
	//			return true
	//		}
	//	}
	//
	//}
	//return false

	//2.滑动窗口

	mp := make(map[int]int)
	for i, v := range nums {
		if i > k {
			delete(mp, nums[i-k-1])
		}
		if _, ok := mp[v]; ok {
			return true
		}
		mp[v] = 1
	}
	return false
}

func findMaxAverage(nums []int, k int) float64 {
	var average float64 = 0
	sum := 0
	if len(nums) <= k {
		for _, v := range nums {
			sum += v
		}
		average = float64(sum) / float64(k)
		return average
	}
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	tmp := sum
	for i := k; i < len(nums); i++ {
		tmp -= nums[i-k]
		tmp += nums[i]
		if tmp > sum {
			sum = tmp
		}
	}
	average = float64(sum) / float64(k)
	return average

	//暴力解法 超出时间限制
	//var average float64 =  0
	//sum := 0
	//if len(nums) <= k{
	//	for _, v := range nums{
	//		sum += v
	//	}
	//	average = float64(sum)/ float64(k)
	//	return average
	//}
	//for i := 0; i < k; i++{
	//	sum += nums[i]
	//}
	//for i := 1; i < len(nums) - k + 1; i++{
	//	tmp := nums[i]
	//	for j := i + 1; j < k + i; j++{
	//		tmp += nums[j]
	//	}
	//	if tmp > sum{
	//		sum = tmp
	//	}
	//}
	//average = float64(sum)/ float64(k)
	//return average
}

func findRepeatedDnaSequences(s string) []string {
	k := 10
	ret := make([]string, 0)
	mp := make(map[string]int)
	if len(s) <= k {
		return ret
	}
	for i := 0; i < len(s)-k+1; i++ {
		tmpStr := string(s[i:(i + k)])
		mp[tmpStr]++
	}
	for key, v := range mp {
		if v > 1 {
			ret = append(ret, key)
		}
	}
	return ret
}

//LRU缓存淘汰算法(最近最少使用)
type LRUCache struct {
	Capacity int
	Size     int
	Head     *DLinkNode
	Tail     *DLinkNode
	Mp       map[int]*DLinkNode
}




type DLinkNode struct {
	Key int
	Value int
	Next  *DLinkNode
	Prev  *DLinkNode
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{}
	ret.Mp = make(map[int]*DLinkNode)
	head := &DLinkNode{}
	tai := &DLinkNode{}
	ret.Head = head
	ret.Tail = tai
	ret.Head.Next = tai
	ret.Tail.Prev = head
	ret.Capacity = capacity
	ret.Size = 0
	return ret
}

func (this *LRUCache) Get(key int) int {
	if cur, ok := this.Mp[key];ok{
		//就一个元素的时候不需要移动位置
		if cur.Next == this.Tail{
			return this.Mp[key].Value
		}

		//先去掉原始的连接关系
		p := cur.Prev
		q := cur.Next
		p.Next = q
		q.Prev = p

		//再建立新的连接关系
		left := this.Tail.Prev
		right := this.Tail
		left.Next= cur
		cur.Prev = left
		cur.Next = right
		right.Prev = cur

		return this.Mp[key].Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.Mp[key].Value = value
		return
	}

	node := &DLinkNode{}
	node.Key = key
	node.Value = value
	this.Mp[key] = node

	var p *DLinkNode
	q := this.Tail
	if this.Size == this.Capacity{
		if this.Size == 1{

			p = this.Head
			//删除最开始的节点的位置信息，因为容量已经满了
			delete(this.Mp, p.Next.Key)

			//和前面一个节点建立双向联系
			p.Next = node
			node.Prev = p
			//和后面一个节点建立双向联系
			node.Next = q
			q.Prev = node


		}else{

			//去掉最开始的首节点
			pos := this.Head.Next
			pos.Next.Prev = this.Head
			this.Head.Next = pos.Next

			//删除最开始的节点的位置信息，因为容量已经满了
			delete(this.Mp, pos.Key)

			p = this.Tail.Prev

			//和前面一个节点建立双向联系
			p.Next = node
			node.Prev = p
			//和后面一个节点建立双向联系
			node.Next = q
			q.Prev = node
		}
	}else{
		var p *DLinkNode
		q := this.Tail
		if this.Size == 0{
			p = this.Head
		}else{
			p = this.Tail.Prev
		}

		//和前面一个节点建立双向联系
		p.Next = node
		node.Prev = p
		//和后面一个节点建立双向联系
		node.Next = q
		q.Prev = node

		this.Size++
	}

}


var obj LRUCache
func TestLRUCache(str []string, nums [][]int){
	ret := make([]string, 0)

	for i, v := range str{

		switch v {
		case "LRUCache":
			obj = Constructor(nums[i][0])
			ret = append(ret, "null")
		case "put":
			obj.Put(nums[i][0], nums[i][1])
			ret = append(ret, "null")
		case "get":
			tmp := strconv.Itoa(obj.Get(nums[i][0]))
			ret = append(ret, tmp)
		default:
			fmt.Println("unknown err")
		}
	}
	fmt.Println(ret)
}




func main() {
	TestLRUCacheStr := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	TestLRUCacheMp := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}

	TestLRUCacheStr = []string{"LRUCache","put","get","put","get","get"}
	TestLRUCacheMp = [][]int{   {1},     {2, 1}, {2}, {3,2}, {2},  {3}}

	TestLRUCacheStr = []string{"LRUCache","put","put","get","put","put","get"}
	TestLRUCacheMp = [][]int{   {2},    {2, 1}, {2,2}, {2}, {1,1},  {4,1}, {2}}

	TestLRUCache(TestLRUCacheStr, TestLRUCacheMp)
	return

	{
		findRepeatedDnaSequencesStr := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
		findRepeatedDnaSequencesStr = "AAAAAAAAAAA"
		fmt.Println(findRepeatedDnaSequences(findRepeatedDnaSequencesStr))
		return

		findMaxAverageList := []int{1, 12, -5, -6, 50, 3}
		findMaxAverageListK := 4
		//findMaxAverageList = []int{5}
		//findMaxAverageListK = 1
		//findMaxAverageList = []int{0,1,1,3,3}
		//findMaxAverageListK = 4

		fmt.Println(findMaxAverage(findMaxAverageList, findMaxAverageListK))
		return

		containsNearbyDuplicateList := []int{1, 2, 3, 1}
		containsNearbyDuplicateK := 3
		//containsNearbyDuplicateList = []int{1, 0, 1, 1}
		//containsNearbyDuplicateK = 1
		containsNearbyDuplicateList = []int{1, 2, 3, 1, 2, 3}
		containsNearbyDuplicateK = 2
		//containsNearbyDuplicateList = []int{4, 1, 2, 3,1,5}
		//containsNearbyDuplicateK = 3

		fmt.Println(containsNearbyDuplicate(containsNearbyDuplicateList, containsNearbyDuplicateK))
		return

		deleteDuplicatesData1 := &ListNode{1, nil}
		deleteDuplicatesData2 := &ListNode{2, nil}
		deleteDuplicatesData3 := &ListNode{3, nil}
		deleteDuplicatesData4 := &ListNode{3, nil}
		deleteDuplicatesData5 := &ListNode{4, nil}
		deleteDuplicatesData6 := &ListNode{4, nil}
		deleteDuplicatesData7 := &ListNode{5, nil}

		deleteDuplicatesData8 := &ListNode{5, nil}

		deleteDuplicatesData1.Next = deleteDuplicatesData2
		deleteDuplicatesData2.Next = deleteDuplicatesData3
		deleteDuplicatesData3.Next = deleteDuplicatesData4
		deleteDuplicatesData4.Next = deleteDuplicatesData5
		deleteDuplicatesData5.Next = deleteDuplicatesData6
		deleteDuplicatesData6.Next = deleteDuplicatesData7

		deleteDuplicatesData7.Next = deleteDuplicatesData8

		ret := deleteDuplicates(deleteDuplicatesData1)
		ShowValue(ret)
		return

		rotateRightHead := &ListNode{}
		p := rotateRightHead
		size := 5
		size = 2
		for i := 0; i < size; i++ {
			next := &ListNode{}
			next.Val = (i + 1)
			//next.Val = i
			p.Next = next
			p = next
		}
		rotateRightHeadK := 4
		rotateRightHeadK = 5
		ret = rotateRight(rotateRightHead.Next, rotateRightHeadK)
		ShowValue(ret)
		return

		findPeakElementList := []int{1, 2, 3, 1}
		//findPeakElementList = []int{1,2,1,3,5,6,4}
		findPeakElementList = []int{1, 2, 3}
		fmt.Println(findPeakElement(findPeakElementList))
		return

		missingNumberList := []int{3, 0, 1}
		//missingNumberList = []int{9,6,4,2,3,5,7,0,1}
		missingNumberList = []int{0, 1}
		fmt.Println(missingNumber(missingNumberList))
		return
		fmt.Println("hello world")
	}

}
