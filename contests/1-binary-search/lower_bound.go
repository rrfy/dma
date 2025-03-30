package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	X, _ := strconv.Atoi(parts[1])

	scanner.Scan()
	arrLine := scanner.Text()
	arrParts := strings.Fields(arrLine)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(arrParts[i])
	}

	left := 0
	right := N

	for left < right {
		mid := (left + right) / 2
		if A[mid] >= X {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left < N && A[left] >= X {
		fmt.Println(left)
	} else {
		fmt.Println(-1)
	}
}

/* student@DESKTOP-MLRIE5E:~/dma$ python3 run_tests.py contests/1-binary-s
earch/lower_bound.go contests/1-binary-search/lower_bound_tests.zip
Compiling with command: go build -o /tmp/tmpgnihnily/solution_exec /home/student/dma/contests/1-binary-search/lower_bound.go
Extracted tests to /tmp/tmpgnihnily/tests
Test test1 PASSED
Test test2 PASSED
Test test3 PASSED
Test test4 PASSED
Test test5 PASSED
Test test6 PASSED
Test test7 PASSED
All 7/7 tests passed! */
