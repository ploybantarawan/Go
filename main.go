package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main(){

	// two sum
	// s := []int{11, 15, 7, 2}
	// fmt.Println(twoSum(s, 9))

	// Palindrome Number
	// fmt.Println(isPalindrome(101))

	// Palindrome String
	// fmt.Println(isWordPalindrome("ovo"))

	// Reversed String
	// s := []byte{'h', 'e', 'l', 'l', 'o'}
	// fmt.Println(string(reversedString(s)))

	// Reverse Integer
	// fmt.Println(reversedInt(1534236469))

	// Valid Parentheses
	// fmt.Println(isValid("([)]"))

	// climbStairs
	// fmt.Println(stair(5))

	//max Profits
	// prices := []int{7, 1, 5, 3, 9, 4}
	// fmt.Println(maxProfits(prices))
	
	// anagram
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    // fmt.Println(groupAnagrams(strs)) 

	// Maximum Subarray
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 8}
    // fmt.Println(maxSubArray(nums))

	// Container With Most Water
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 9}
    // fmt.Println(maxArea(height))

	s := "abcabcbb"
	fmt.Println(longestS(s))
	
} 

// Two sum
// Description: Given an array of integers, return the indices of the two numbers such that they add up to a specific target.

func twoSum(nums []int, target int) []int {
    numMap := map[int]int{}
    for i, num := range nums {
        complement := target - num
        if j, found := numMap[complement]; found {
			fmt.Println(j)
            return []int{j, i}
        }
        numMap[num] = i
		fmt.Println(numMap)
    }
    return nil
}

// Palindrome Number
// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same forward and backward.

func isPalindrome(x int) bool{
	if x < 0{
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed

}

// Palindrome String

func isWordPalindrome(s string) bool{

	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for left < right{
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	
	return true
}

// Reverse a String
// Write a function that reverses a string. The input string is given as an array of characters s.

func reversedString(s []byte) []byte{
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

// Reverse Integer:
// Description: Given a 32-bit signed integer, reverse its digits. If the reversed integer overflows, return 0.
func reversedInt(x int) int{
	reversed := 0
	for x!= 0{
		pop := x % 10
		x /= 10

		//check overflow
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && pop > 7){
			return 0 // Overflow for positive numbers
		}
		if reversed < math.MinInt32/10 || (reversed == math.MinInt32/10 && pop < -8) {
            return 0 // Overflow for negative numbers
        }

		reversed = reversed*10 + pop
	}
	return reversed
}

// Valid Parentheses:
// Description: Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

func isValid(s string) bool{

	// Create a stack to hold opening parentheses
	stack := []rune{}

	// Create a map for matching parentheses
	matchParentheses := map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}

	// Iterate through the string
	for _, char := range s {
		if match, exists := matchParentheses[char]; exists{
			if len(stack) > 0 && stack[len(stack)-1] == match{
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

// Climbing Stairs
// Description: You are climbing a staircase. It takes n steps to reach the top. Each time you can either climb 1 or 2 steps. 
// In how many distinct ways can you climb to the top?

func stair(n int) int {
	// if only 1 stair = 1 step needed
	if n == 1{
		return 1
	}

	// based case 
	stair1 := 1
	stair2 := 2

	for i := 3; i<= n; i++ {
		current := stair1 + stair2
		stair1 = stair2
		stair2 = current
	} 

	return stair2
}

// Best Time to Buy and Sell Stock
// Description: You want to maximize your profit by choosing a single day to buy one stock 
// and choosing a different day in the future to sell that stock. return maximum profits

func maxProfits(price []int) int {
	if len(price) == 0 {
		return 0
	}

	minPrice := price[0]
	maxProfits := 0

	for i := 1; i < len(price); i++ {
		if price[i] < minPrice{
			minPrice = price[i]
		} else {
			profits := price[i] - minPrice
			if profits > maxProfits {
				maxProfits = profits
			}
		}

	}
	return maxProfits
}

// Group Anagrams
// Description: Given an array of strings, group the anagrams together.

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := SortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	
	result := [][]string{} 
	
	for _, group := range anagramMap {
		// fmt.Println(group)
		result = append(result, group)
	}

	return result
	
}

func SortString(s string) string {
    strSlice := strings.Split(s,"")
	sort.Strings(strSlice)
	return strings.Join(strSlice, "")
}

// Problem: Maximum Subarray
// Given an integer array nums, find the contiguous subarray (containing at least one number) 
// that has the largest sum and return its sum.

func maxSubArray(nums []int) int {
	// Kadane's Algorithm

	current := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		current = findMax(nums[i], nums[i] + current)

		if current > maxSum {
			maxSum = current
		}
	}
	return maxSum
}

func findMax(a,b int) int {
	if a > b {
		return a
	}
	return b
}

// Container With Most Water
// Description: Given n non-negative integers representing the height of lines on a chart, 
// find two lines that together with the x-axis form a container that holds the most water.
// The amount of water the container can hold is the minimum height between the two lines multiplied by the distance between them.

func maxArea(height [] int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		area := min(height[left],height[right]) * (right - left)

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right]{
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.

func longestS(s string) int {
	//Sliding Window
	charMap := make(map[byte]int)

	left, maxLength := 0, 0

	for right := 0; right < len(s); right++{
		// if match, update left pointer
		if index, found := charMap[s[right]]; index >= left && found {
			left = index + 1
		}
		// if not match, update char in map
		charMap[s[right]] = right

		currentLength := right - left +1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}