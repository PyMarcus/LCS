package main

import "fmt"

func max(value1, value2 int) int{
	if value1 > value2{
		return value1
	}
	return value2
}

func main() {
	const yi = "cgata"
	const xj = "tccata"
	var table [len(yi) + 1][len(xj) + 1]int
	
	for i := 0; i< len(yi); i++{
		for j := 0; j < len(xj); j++{
			if i == 0 || j == 0{
				table[i][j] = 0
			}else{
				if yi[i - 1] == xj[j - 1]{
					table[i][j] = table[i-1][j-1] + 1
				}else{
					table[i][j] = max(table[i - 1][j], table[i][j - 1])
				}
			}
		}
	}
	
	for i := 0; i< len(yi); i++{
		for j := 0; j < len(xj); j++{
			fmt.Print(table[i][j])	
		}
		fmt.Println()
	}
}
