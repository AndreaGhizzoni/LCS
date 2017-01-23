package main

import "fmt"

const (
	P = "lasciaperderequesto"
	T = "questoeunesempio"

	/*
		P = "aaabbbccc"
		T = "zzzbbbttt"
	*/
)

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func lcs(d *[][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	} else {
		D := *d
		if D[i][j] == 0 {
			if P[i] == T[j] {
				D[i][j] = lcs(d, i-1, j-1) + 1
			} else {
				D[i][j] = max(lcs(d, i-1, j), lcs(d, i, j-1))
			}

		}
		return D[i][j]
	}
}

func printdiff(d *[][]int, i, j int) {
	D := *d
	if i >= 0 && j >= 0 && P[i] == T[j] {
		printdiff(d, i-1, j-1)
		fmt.Printf("%c", P[i])
	} else if j >= 0 && (i < 0 || D[i][j-1] >= D[i-1][j]) {
		printdiff(d, i, j-1)
		//fmt.Printf("%c", T[i])
	} else if i >= 0 && (j < 0 || D[i][j-1] < D[i-1][j]) {
		printdiff(d, i-1, j)
		//fmt.Printf("%c", P[i])
	}
}

func main() {
	lenP := len(P)
	lenT := len(T)
	D := make([][]int, lenP)
	for i := 0; i < lenP; i++ {
		D[i] = make([]int, lenT)
	}

	fmt.Println(lcs(&D, lenP-1, lenT-1))
	printdiff(&D, lenP-1, lenT-1)
	fmt.Println()
}
