package main

import (
	"fmt"
	"net"
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

/*
//LRU缓存淘汰算法(最近最少使用)
type LRUCache struct {
	Capacity int
	Size     int
	Head     *DLinkNode
	Tail     *DLinkNode
	Mp       map[int]*DLinkNode
}

type DLinkNode struct {
	Key   int
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
	if cur, ok := this.Mp[key]; ok {
		//就一个元素的时候不需要移动位置
		if cur.Next == this.Tail {
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
		left.Next = cur
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
	if this.Size == this.Capacity {
		if this.Size == 1 {

			p = this.Head
			//删除最开始的节点的位置信息，因为容量已经满了
			delete(this.Mp, p.Next.Key)

			//和前面一个节点建立双向联系
			p.Next = node
			node.Prev = p
			//和后面一个节点建立双向联系
			node.Next = q
			q.Prev = node

		} else {

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
	} else {
		var p *DLinkNode
		q := this.Tail
		if this.Size == 0 {
			p = this.Head
		} else {
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

func TestLRUCache(str []string, nums [][]int) {
	ret := make([]string, 0)

	for i, v := range str {

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
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := (*ListNode)(nil)
	q := (*ListNode)(nil)
	prev := (*ListNode)(nil)
	cur := (*ListNode)(nil)
	if head.Val < x {
		p = head
		q = head.Next
	} else {
		p.Next = head
		q = head
	}
	prev = head
	cur = head.Next
	for cur != nil {
		if cur.Val > x {
			prev = cur
			cur = cur.Next
		}
		prev.Next = cur.Next
		//插入
		if cur != q {
			cur.Next = q
		}
		if p != cur {
			p.Next = cur
		}

		p = cur
		cur = prev.Next
	}

	return head

	//prev := &ListNode{}
	//next := &ListNode{}
	//l := prev
	//r := next
	//cur := head
	//for cur != nil{
	//	p := cur
	//	cur = cur.Next
	//	p.Next = nil
	//	if p.Val < x{
	//		prev.Next = p
	//		prev = prev.Next
	//	}else{
	//		next.Next = p
	//		next = next.Next
	//	}
	//}
	//prev.Next = r.Next
	//return l.Next
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 0
	cur := head
	ret := (*ListNode)(nil)
	prev := (*ListNode)(nil)
	back := (*ListNode)(nil)
	for cur != nil {
		count++
		if left > 1 && count == left-1 {
			prev = cur
		}
		if count == left {
			ret = cur
		}
		if count == right {
			back = cur.Next
			cur.Next = nil
			break
		}
		cur = cur.Next
	}
	first := ret
	tmp := reverseList(ret)
	if prev == nil {
		first.Next = back
		return tmp
	}
	prev.Next = tmp
	first.Next = back

	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}
	fast := head.Next.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next == head {
		return head
	}
	mp := make(map[*ListNode]int)
	fast := head.Next.Next
	slow := head
	//mp[fast] = 1
	mp[slow] = 1
	for fast != nil && fast.Next != nil {
		if fast == slow {
			if _, ok := mp[slow]; ok {
				return slow
			}
		}
		slow = slow.Next
		fast = fast.Next.Next
		if _, ok := mp[fast]; ok {
			return fast
		}
		if _, ok := mp[slow]; ok {
			return slow
		}
		//mp[fast] = 1
		mp[slow] = 1
	}
	return nil
	//if !bCycle{
	//	return nil
	//}
	//ret := (*ListNode)(nil)
	//cur := head
	//for cur != nil{
	//	if _, ok := mp[cur];ok{
	//		return cur
	//	}
	//	mp[cur]++
	//	cur = cur.Next
	//}
	//return ret
}

//实现一个栈
type MyStack struct {
	First  *LinkListNode
	Second *LinkListNode
	Size   int
}

type LinkListNode struct {
	Head *DataNode
	Tail *DataNode
}

type DataNode struct {
	Val  int
	Next *DataNode
	Prev *DataNode
}

func Constructor() MyStack {
	ret := MyStack{}

	first := &LinkListNode{}
	nodeHead := &DataNode{}
	nodetail := &DataNode{}
	first.Head = nodeHead
	first.Tail = nodetail

	nodeHead.Next = nodetail
	nodetail.Prev = nodeHead
	ret.First = first

	//第二个双向链表记录插入位置的前后节点位置
	ret.Second = &LinkListNode{}

	ret.Size = 0

	return ret
}

func (this *MyStack) Push(x int) {
	p := (*DataNode)(nil)
	q := (*DataNode)(nil)
	if this.Size == 0 {
		p = this.First.Head
		q = this.First.Tail
	} else {
		p = this.Second.Head
		q = this.Second.Tail
		//node := &DataNode{}
		//node.Val = x
		//p.Next = node
		//node.Prev = p
		//
		//node.Next = q
		//q.Prev = node
		//
		//
		////已经有层级关系了
		//this.Second.Head = node
		//this.Second.Tail = q
	}
	node := &DataNode{}
	node.Val = x
	p.Next = node
	node.Prev = p
	node.Next = q
	q.Prev = node
	//已经有层级关系了
	this.Second.Head = node
	this.Second.Tail = q
	this.Size++
}

func (this *MyStack) Pop() int {
	ret := this.Second.Head.Val

	p := this.Second.Head.Prev
	q := this.Second.Tail

	p.Next = q
	q.Prev = p

	this.Second.Head = p

	this.Size--
	return ret
}

func (this *MyStack) Top() int {
	return this.Second.Head.Val
}

func (this *MyStack) Empty() bool {
	if this.Size == 0 {
		return true
	}
	return false
}

func TestMyStack() {
	nums := []string{"MyStack", "push", "push", "top", "pop", "empty"}
	numsList := [][]int{{}, {1}, {2}, {}, {}, {}}

	nums = []string{"MyStack", "push", "push", "pop", "top"}
	numsList = [][]int{{}, {1}, {2}, {}, {}}

	ret := make([]string, 0)
	var obj MyStack
	for i, v := range nums {
		//if i == 2{
		//	fmt.Println("11")
		//}

		switch v {
		case "MyStack":
			obj = Constructor()
			ret = append(ret, "null")
		case "push":
			obj.Push(numsList[i][0])
			ret = append(ret, "null")
		case "top":
			x := obj.Top()
			ret = append(ret, strconv.Itoa(x))
		case "pop":
			x := obj.Pop()
			ret = append(ret, strconv.Itoa(x))
		case "empty":
			b := obj.Empty()
			if b {
				ret = append(ret, "true")
			} else {
				ret = append(ret, "false")
			}
		default:
			fmt.Println("err")
		}
	}
	fmt.Println(ret)
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums []int

func preorderTraversal(root *TreeNode) []int {
	nums = make([]int, 0)
	if root == nil {
		return nums
	}
	preorderTraversalBack(root)
	return nums
}

func preorderTraversalBack(node *TreeNode) {
	if node == nil {
		return
	}
	nums = append(nums, node.Val)
	preorderTraversalBack(node.Left)
	preorderTraversalBack(node.Right)
}

func postorderTraversal(root *TreeNode) []int {
	nums = make([]int, 0)
	if root == nil {
		return nums
	}
	postorderTraversalBack(root)
	return nums
}

func postorderTraversalBack(node *TreeNode) {
	if node == nil {
		return
	}
	postorderTraversalBack(node.Left)
	postorderTraversalBack(node.Right)
	nums = append(nums, node.Val)
}

func singleNumbers(nums []int) []int {
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	or := 1
	for sum^or != 0 {
		or <<= 1
	}
	a := 0
	b := 0
	for _, v := range nums {
		if v^or == 1 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

func main() {

	{
		root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
		//fmt.Println(preorderTraversal(root))
		fmt.Println(postorderTraversal(root))
		return

		TestMyStack()
		return

		reverseBetweenList := []int{1, 2, 3, 4, 5}
		reverseBetweenLeft := 2
		reverseBetweenRight := 4

		reverseBetweenHead := &ListNode{reverseBetweenList[0], nil}

		cur := reverseBetweenHead
		for i := 1; i < len(reverseBetweenList); i++ {
			tmp := &ListNode{reverseBetweenList[i], nil}
			cur.Next = tmp
			cur = cur.Next
		}

		fmt.Println(reverseBetween(reverseBetweenHead, reverseBetweenLeft, reverseBetweenRight))
		return

		partitionkList := []int{1, 4, 3, 2, 5, 2}

		partitionHead := &ListNode{partitionkList[0], nil}

		cur = partitionHead
		for i := 1; i < len(partitionkList); i++ {
			tmp := &ListNode{partitionkList[i], nil}
			cur.Next = tmp
			cur = cur.Next
		}

		partitionk := 3
		partition(partitionHead, partitionk)
		return

		//TestLRUCacheStr := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
		//TestLRUCacheMp := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
		//
		//TestLRUCacheStr = []string{"LRUCache", "put", "get", "put", "get", "get"}
		//TestLRUCacheMp = [][]int{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}}
		//
		//TestLRUCacheStr = []string{"LRUCache", "put", "put", "get", "put", "put", "get"}
		//TestLRUCacheMp = [][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}}
		//
		//TestLRUCache(TestLRUCacheStr, TestLRUCacheMp)
		//return

		list, err := net.Listen("tcp", ":5001")
		if err != nil {
			return
		}
		fmt.Println(list)

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
