package main

import (
	"bufio"
	"fmt"
	"os"
)

var w = bufio.NewWriter(os.Stdout)

func main() {
	var n, m int
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()
	fmt.Fscan(r, &n, &m)
	a := make([]int, m)
	check := make([]int, n+1)
	permutation(a, check, 0, 0, n, m)
}

func permutation(a, check []int, pos, prev, n, m int) {
	if pos == m {
		for i := 0; i < m; i++ {
			fmt.Fprintf(w, "%d ", a[i])
		}
		fmt.Fprintln(w)
	} else {
		for i := prev; i < n && a[0] == 0; i++ {
			if check[i] == 0 {
				a[pos] = i
				check[i] = 1
				permutation(a, check, pos+1, i, n, m)
				check[i] = 0
			}
		}
	}
}
