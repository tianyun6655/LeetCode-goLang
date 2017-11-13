package leetcode

import "fmt"

/**
Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.

Example 1:
Input: "()"
Output: True
Example 2:
Input: "(*)"
Output: True
Example 3:
Input: "(*))"
Output: True

 */

func CheckValidString(s string) bool {
	//( = 40,*=42,)=41
     bytes:=[]byte(s)

	 startCount:=0
	 disMatchCount:=0
	for leftIndex:=0;leftIndex<len(bytes);{
BREAK:
		rightIndex :=leftIndex+1
        if bytes[leftIndex]==42{
			startCount++
			leftIndex++
			continue
		}else if bytes[leftIndex]==0{
			leftIndex++
			continue
		}else if bytes[leftIndex]==41{
			leftIndex++
			disMatchCount++
			continue
		}

		for rightIndex<len(bytes){
			if bytes[rightIndex]==41{
				bytes[leftIndex]=0
                bytes[rightIndex]=0
				leftIndex++
				goto BREAK
			}
			rightIndex++
		}
		    leftIndex++
		    disMatchCount++

	}
	fmt.Println(startCount)
	fmt.Println(disMatchCount)
	if disMatchCount==0&&startCount==0{
		return  true
	}else if disMatchCount==0 && startCount!=0{
			return  true
	}else {
		if startCount>=disMatchCount{
			return true
		}
		return false
	}

}