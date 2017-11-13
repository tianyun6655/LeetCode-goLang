package leetcode

var meno =make(map[int]int)
func ClimbStairs(n int) int{

	if n==1{
		 return 1
	 }
	if n==2{
		return 2
	}
	if meno[n]==0{
		meno[n]= ClimbStairs(n-1)+ClimbStairs(n-2)
	}
	return meno[n]
}

