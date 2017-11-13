package dynamicprogramming

import (
	"fmt"
	"math"
)

func IntegerBreak(n int) int {
	max:=math.MinInt64
	for i:=1;i<int(math.Sqrt(float64(n)));i++{
		times:=n/i
		if times==1{
			if max<times*n%i{
				max=times*n%i
			}
			continue
		}
		left:=n%i
		last:=i+left

		if last<i*left{
			last = i*left
		}

		timespow:=pow(i,times-1)

		fmt.Println("timespow",timespow)

		result :=timespow*last
		fmt.Println(i," ",result)
		if max<result{
			max=result
		}

	}
	return max
}

func pow (number,pow int) int{
	result:=number

	for i:=1;i<pow;i++{
		result*=number
	}
	return result
}