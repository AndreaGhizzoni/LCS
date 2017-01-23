package main

import (
	"flag"
	"fmt"
)

// TODO using built in max function
func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// Function that calculates the Longest Common Subsequence between P and T.
// d : is a pointer to a slice of slice of dimension: len(P)xlen(T)
// p : pointer to string to calculate the lcs
// t : pointer to string to calculate the lcs
// i : index that goes from len(P)-1 to -1
// j : index that goes from len(T)-1 to -1
func lcs(d *[][]int, p, t *string, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	} else {
		D := *d
		if D[i][j] == 0 {
			if (*p)[i] == (*t)[j] {
				D[i][j] = lcs(d, p, t, i-1, j-1) + 1
			} else {
				D[i][j] = max(lcs(d, p, t, i-1, j), lcs(d, p, t, i, j-1))
			}

		}
		return D[i][j]
	}
}

// Function that prints the longest common subsequence between P and T.
// In order to work, this function needs the pre-elaborated matrix of values (d)
// from lcs(), so call lcs() first and then printLCS()
// d : is a pointer to a slice of slice.
// p : pointer to string to calculate the lcs
// t : pointer to string to calculate the lcs
// i : index that goes from len(P)-1 to -1
// j : index that goes from len(T)-1 to -1
func printLCS(d *[][]int, p, t *string, i, j int) {
	D := *d
	if i >= 0 && j >= 0 && (*p)[i] == (*t)[j] {
		printLCS(d, p, t, i-1, j-1)
		fmt.Printf("%c", (*p)[i])
	} else if j > 0 && (i == 0 || D[i][j-1] >= D[i-1][j]) {
		printLCS(d, p, t, i, j-1)
		//fmt.Printf("%c", T[i])
	} else if i > 0 && (j == 0 || D[i][j-1] < D[i-1][j]) {
		printLCS(d, p, t, i-1, j)
		//fmt.Printf("%c", P[i])
	}
}

func main() {
	P := flag.String("s1", "", "first string")
	T := flag.String("s2", "", "second string")

	flag.Parse()

	lenP := len((*P))
	lenT := len((*T))
	D := make([][]int, lenP)
	for i := 0; i < lenP; i++ {
		D[i] = make([]int, lenT)
	}

	fmt.Println(lcs(&D, P, T, lenP-1, lenT-1))
	printLCS(&D, P, T, lenP-1, lenT-1)
	fmt.Println()
}
