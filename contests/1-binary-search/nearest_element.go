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

	pos := lowerBound(A, X)

	candidates := make([]int, 0)
	if pos < N {
		candidates = append(candidates, pos)
	}
	if pos > 0 {
		candidates = append(candidates, pos-1)
	}

	bestIdx := -1
	minDiff := -1

	for _, idx := range candidates {
		diff := abs(A[idx] - X)
		if bestIdx == -1 || diff < minDiff || (diff == minDiff && idx < bestIdx) {
			bestIdx = idx
			minDiff = diff
		}
	}

	fmt.Println(bestIdx)
}

func lowerBound(arr []int, x int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] >= x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*student@DESKTOP-MLRIE5E:~/dma$ python3 run_tests.py contests/1-binary-search/nearest_element.go contests/1-binary-search/nearest_element_tests.zip
Compiling with command: go build -o /tmp/tmpsj3es39f/solution_exec /home/student/dma/contests/1-binary-search/nearest_element.go
Extracted tests to /tmp/tmpsj3es39f/tests
Test test1 PASSED
Test test2 PASSED
Test test3 PASSED
Test test4 PASSED
Test test5 PASSED
All 5/5 tests passed!*/
