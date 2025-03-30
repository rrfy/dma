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
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	arrLine := scanner.Text()
	arrParts := strings.Fields(arrLine)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(arrParts[i])
	}

	left := 0
	right := N - 1

	for left < right {
		mid := (left + right) / 2
		if A[mid] < A[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}

/*student@DESKTOP-MLRIE5E:~/dma$ python3 run_tests.py contests/1-binary-search/peak.go contests/1-binary-search/peak_tests.zip
Compiling with command: go build -o /tmp/tmp21ra2p6d/solution_exec /home/student/dma/contests/1-binary-search/peak.go
Extracted tests to /tmp/tmp21ra2p6d/tests
Test test1 PASSED
Test test2 PASSED
Test test3 PASSED
Test test4 PASSED
All 4/4 tests passed!*/
