package dyamicProgramming

var A int = 0
var meno =make(map[int]int)
func Fib (n int) int{
	A++
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
    if meno[n]==0 {
		meno[n ]=Fib(n-1)+Fib(n-2)
	}
	return meno[n]
}


func FibFromBottom(n int) int{
	meno[0] =0
	meno[1] =1
	for i:=2;i<=n;i++{
		meno[i] = meno[i-1]+meno[i-2]
	}

	return meno[n]
}