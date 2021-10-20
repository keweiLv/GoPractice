package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	s := "leetcode"
	word := []string{"leet", "code"}
	res := wordBreak(s, word)
	fmt.Println(res)
}

//单词拆分
func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= l; i++ {
		for j := i - 1; j >= 0; j-- {
			suffix := s[j:i]
			if wordMap[suffix] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[l]
}

//旋转字符串
func rotateString(A string, B string) bool {
	if A == "" && B == "" {
		return true
	}
	for i := 0; i < len(A); i++ {
		if A[i:]+A[:i] == B {
			return true
		}
	}
	return false
}

// 最后一个单词长度
func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}

// 不同的二叉搜索树
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

// 最小操作数使元素相等
func minMoves(nums []int) (ans int) {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return
}
