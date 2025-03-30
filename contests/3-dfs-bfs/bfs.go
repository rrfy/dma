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
	queue := []int{1}
	visited[1] = true
	result := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}

/*python3 run_tests.py contests/3-dfs-bfs/bfs.go contests/3-dfs-bfs/bfs_tests.zip
Compiling with command: go build -o /tmp/tmpovqot4ps/solution_exec /home/student/dma/contests/3-dfs-bfs/bfs.go
Extracted tests to /tmp/tmpovqot4ps/tests
Test test1 PASSED
Test test2 PASSED
Test test3 PASSED
Test test4 PASSED
All 4/4 tests passed!*/
