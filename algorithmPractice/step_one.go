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
