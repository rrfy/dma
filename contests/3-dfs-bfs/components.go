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
	count := 0

	for i := 1; i <= N; i++ {
		if !visited[i] {
			bfs(adj, visited, i)
			count++
		}
	}

	fmt.Println(count)
}

func bfs(adj [][]int, visited []bool, start int) {
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
