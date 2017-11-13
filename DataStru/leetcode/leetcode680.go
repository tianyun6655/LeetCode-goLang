package leetcode


/**
680. Valid Palindrome II
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.
Example 1:
Input: "aba"
Output: True
Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'
 */

func ValidPalindrome(s string) bool {
	bytes:=[]byte(s)
	start := 0
	end :=len(bytes)-1
	for start<end{
	if bytes[start]!=bytes[end]{
		return isPalindrome(bytes,start+1,end)||isPalindrome(bytes,start,end-1)
	}
	start++
	end--
}
return true
}

func isPalindrome( s []byte,l,r int) bool{
	for l<r{
	if s[l]!=s[r]{
		return false
	}
	l++
	r--
}
return true
}