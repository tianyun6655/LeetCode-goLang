package dynamicprogramming

import "math"

func MinPathSum(grid [][]int) int {


	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if i==0&&j!=0{
				grid[i][j]+= grid[i][j-1]
			}else if j==0 && i!=0{
				grid[i][j]+=grid[i-1][j]
			}else if i!=0&&j!=0{
				grid[i][j]+= int(math.Min(float64(grid[i-1][j]),float64(grid[i][j-1])))
			}

		}
	}

	return grid[len(grid)-1][len(grid[len(grid)-1])-1]
}
