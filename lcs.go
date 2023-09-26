package main

import "fmt"

func main() {
	const wordone string = "cgata"
	const wordtwo string = "tccata"

	var matrix [len(wordone) + 1][len(wordtwo) + 1]int

	for i := 0; i < len(wordone); i++ {
		for j := 0; j < len(wordtwo); j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 0
			} else {
				if wordone[i - 1] == wordtwo[j - 1] {
					matrix[i][j] = matrix[i-1][j-1] + 1
				} else {
					matrix[i][j] = max(matrix[i-1][j], matrix[i][j-1])
				}
			}
		}
	}

	fmt.Println("SIMPLE")
	for i := 0; i < len(wordone); i++ {
		for j := 0; j < len(wordtwo); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}

	// retorna a maior subsequencia comum

	maxValue := 4
	fmt.Println(maxValue)
	// tamanho da maior sequencia comum entre as strings
	i := len(wordone)
	j := len(wordtwo)
	lcs := make([]byte, maxValue)
	for (i > 0 && j > 0) {
		if wordone[i-1] == wordtwo[j-1] {
			lcs[maxValue-1] = wordone[i-1]
			i--
			j--
			maxValue--
		}else if matrix[i-1][j] > matrix[i][j-1] {
			i--
		}else {
			j--
		}
	}

	fmt.Println(string(lcs))
}
