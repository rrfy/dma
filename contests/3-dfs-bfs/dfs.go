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

	adj := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		k, _ := strconv.Atoi(parts[0])
		adj[i] = make([]int, k)
		for j := 0; j < k; j++ {
			adj[i][j], _ = strconv.Atoi(parts[j+1])
		}
	}

	visited := make([]bool, N+1)
	stack := []int{1}
	result := []int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true
			result = append(result, node)

			for i := len(adj[node]) - 1; i >= 0; i-- {
				neighbor := adj[node][i]
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}

/*python3 run_tests.py contests/3-dfs-bfs/dfs.go contests/3-dfs-bfs/dfs_tests.zip
Compiling with command: go build -o /tmp/tmpdzdk0t6i/solution_exec /home/student/dma/contests/3-dfs-bfs/dfs.go
Extracted tests to /tmp/tmpdzdk0t6i/tests
Test test1 PASSED
Test test2 PASSED
Test test3 PASSED
Test test4 PASSED
All 4/4 tests passed!*/
