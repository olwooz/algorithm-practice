package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	s := make([]int, n)
	pos := make([]int, m)

	for i := 0; i < n; i++ {
		s[i] = nextInt()
	}
	sort.Ints(s)

	for pos[0] < n {
		for i := 0; i < m; i++ {
			fmt.Fprintf(w, "%d ", s[pos[i]])
		}
		fmt.Fprintln(w)
		pos[m-1]++
		for i := m - 2; i >= 0; i-- {
			if pos[i+1] == n {
				pos[i]++
			}
		}
		//find first occurrence of n
		startPos := 0
		for i := 1; i < m; i++ {
			if pos[i] == n {
				startPos = i
				break
			}
		}
		if startPos > 0 {
			for i := startPos; i < m; i++ {
				pos[i] = pos[startPos-1]
			}
		}
	}
}

func nextInt() int {
	scanner.Scan()
	r := 0
	for _, c := range scanner.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}
