package main

import (
	"fmt"
	"sort"
)

func main(){
	
	// nums := []int{2, 3, 10, 11}
	// target := 9
	// fmt.Println(twoSum(nums, target))

	// fmt.Println(productExceptSelf(nums))

	// s := "babad"
	// fmt.Println(longestPalindrome(s))

	// fmt.Println(isAnagram("abb", "bwb"))

	// a := []int{3,6,2}
	// fmt.Println(subSum(a, 9))

	// Create nodes
    // node1 := &ListNode{Val: 1}
    // node2 := &ListNode{Val: 2}
    // node3 := &ListNode{Val: 3}
    // node4 := &ListNode{Val: 4}
    // node5 := &ListNode{Val: 5}

    // Link nodes to create the list: 1 -> 2 -> 3 -> 4 -> 5 -> nil
    // node1.Next = node2
    // node2.Next = node3
    // node3.Next = node4
    // node4.Next = node5
    // node5.Next = nil

    // Reverse the linked list
    // reversedHead := reversedList(node1)
    // fmt.Println("Reversed linked list:")
    // printList(reversedHead)

	// Create the first sorted linked list: 1 -> 2 -> 4 -> nil
	// l1 := &ListNode{Val: 1}
	// l1.Next = &ListNode{Val: 2}
	// l1.Next.Next = &ListNode{Val: 4}
 
	// // Create the second sorted linked list: 1 -> 3 -> 4 -> nil
	// l2 := &ListNode{Val: 1}
	// l2.Next = &ListNode{Val: 3}
	// l2.Next.Next = &ListNode{Val: 4}

	// // Merge the two lists
    // mergedList := mergeTwoList(l1, l2)
    // fmt.Println("Merged List:")
    // printList(mergedList)

	// nums := []int{4,5,6,7,0,1,3,4}
	// // fmt.Println(search(nums, 4))

	// fmt.Println(kthLargest(nums, 3))

	// coins := []int{1, 2, 5}
	// amount := 11
	// result := coinChange(coins, amount)
	// fmt.Println(result) // Output: 3

	// nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	// result := lengthOfLIS(nums)
	// fmt.Println(result) // Output: 4

	// s := "leetcode"
	// wordDict := []string{"leet", "code"}
	// result := wordBreak(s, wordDict)
	// fmt.Println(result) // Output: true

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// Call level order traversal function
	// result := levelOrder(root)
	// Print the result
	// fmt.Println(result)

	p := root.Right.Left       
	q := root.Right.Right
	// Find the LCA
	lca := lowestCommonAncestor(root, p, q)
	// Print the result
	if lca != nil {
		fmt.Printf("The LCA of %d and %d is %d\n", p.Val, q.Val, lca.Val)
	} else {
		fmt.Println("LCA not found.")
	}




	
} 

//  Two Sum
// Problem: Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. 
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Input: nums = [2, 7, 11, 15], target = 9
// Output: [0, 1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func twoSum(a []int, target int) []int {
	numMap := make(map[int]int)

	for i, nums := range a {
		complement := target - nums
		if index, found := numMap[complement]; found{
			return []int{index, i}
		}
		numMap[nums] = i
	} 
	return nil
}

// Product of Array Except Self
/// Problem:Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]. 

func productExceptSelf(num []int) []int {
	answer := make([]int, len(num))
	n := len(num)
	
	leftProduct := 1
	for i := 0; i<n; i++ {
		answer[i] = leftProduct
		leftProduct *= num[i]
	}

	rightProduct := 1
	for i := n-1; i >= 0; i--{
		answer[i] *= rightProduct
		rightProduct *= num[i]
	}

	return answer
	
}

// Longest Palindromic Substring
// Problem: Given a string s, return the longest palindromic substring in s.
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.

func longestPalindrome(s string) string {
	if len(s) == 0{
		return ""
	}

	start, end := 0,0

	for i := 0 ; i < len(s); i++ {
		len1 := palindromeCheck(s, i, i)
		len2 := palindromeCheck(s, i, i+1)
		maxLen := max(len1,len2)

		if maxLen > end - start {
			start = i - (maxLen - 1) / 2
			end = i + maxLen / 2
		}
	}
	return s[start:end+1]
}

func palindromeCheck(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Valid Anagram
// Problem: Given two strings s and t, return true if t is an anagram of s, and false otherwise

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}

	// fmt.Println(count)
	for _, char := range t {
		count[char]--
		// fmt.Println(count)
		if count[char] < 0 {
			return false
		}
	}	

	return true
}

// Subarray Sum Equals K
// Problem: Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.

func subSum(a []int, k int) int {
	count := 0
	sum := 0
	subMap := make(map[int]int)
	subMap[0] = 1

	for _, num := range a{
		// fmt.Println(num)
		sum += num
		if _, found := subMap[sum-k]; found{
			count += subMap[sum-k]
		}

		subMap[sum]++
	}
	return count
}

// Reverse Linked List
// Problem: Given the head of a singly linked list, reverse the list and return the new head.

type ListNode struct {
    Val  int
    Next *ListNode
}

// Function to print the linked list
func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Print(current.Val, " -> ") // Print the value of the current node
        current = current.Next          // Move to the next node
    }
    fmt.Println("nil") // Indicate the end of the list
}

func reversedList (head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// Merge Two Sorted Lists
// Problem: You are given the heads of two sorted linked lists, list1 and list2. 
// Merge the two lists into one sorted linked list, and return the head of the merged list.

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

// Search in Rotated Sorted Array
// Problem Description: Given an integer array nums that is sorted in ascending order but then rotated at some unknown pivot,
// and a target integer target, find the index of the target. If the target is not present, return -1.

func search(num []int, target int) int {
	left, right := 0, len(num)-1

	for left <= right {
		mid := left + (right - left)/2

		if num[mid] == target{
			return mid
		}

		if num[left] <= num[mid]{
			if num[left] <= target && target < num[mid]{
				right = mid -1
			} else {
				left  = mid +1
			}
		} else {
			if num[right] >= target && target <= num[mid]{
				left = mid +1
			} else {
				right = mid -1
			}
		}
	}
	return -1
}

// Kth Largest Element in an Array
// Problem Description: Find the kth largest element in an unsorted array. 
// You need to do this with a time complexity better than O(n log n), such as using quickselect.

func kthLargest (a []int, t int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
    return a[t - 1]
}

func coinChange(coins []int, amount int) int {
    // Initialize the DP array with a large value
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }
	
    dp[0] = 0 // 0 coins needed to make 0 amount
	// fmt.Println(dp)

    // Fill the DP array
    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            dp[i] = min(dp[i], dp[i-coin]+1)
        }
    }

    // If dp[amount] is still the initialized large value, return -1
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

// Helper function to find the minimum of two numbers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0{
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] =1
	}

	maxLis := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLis = max(maxLis, dp[i])
	}
	return maxLis
}

func wordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]bool)
    for _, word := range wordDict {
        wordSet[word] = true
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true

    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }

    return dp[len(s)]
}

// Problem:
// Given a binary tree, return the level order traversal of its nodes' values. 
// Each level of the tree should be represented as a separate list in the output.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	// Initialize a queue to store nodes
	queue := []*TreeNode{root}

	// While there are nodes to process in the queue
	for len(queue) > 0 {
		// Track the number of nodes at the current level
		levelSize := len(queue)
		// To store values of the current level
		level := []int{}

		// Process all nodes in the current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]       // Get the front node
			queue = queue[1:]      // Dequeue
			level = append(level, node.Val) // Add the node's value to the level

			// Add the node's children to the queue (if they exist)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Append the level's values to the result
		result = append(result, level)
	}

	return result
}

// Problem Description
// Given a binary search tree, find the lowest common ancestor of two given nodes, p and q.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

// Problem Description
// You are given a number of courses and a list of prerequisite pairs. Each prerequisite pair indicates that one course must be completed before another.
// The task is to determine if it is possible to finish all courses given the prerequisites.

