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
	A, _ := strconv.Atoi(parts[1])
	B, _ := strconv.Atoi(parts[2])

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
	parent := make([]int, N+1)
	queue := []int{A}
	visited[A] = true
	found := false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == B {
			found = true
			break
		}

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = node
				queue = append(queue, neighbor)
			}
		}
	}

	if !found {
		fmt.Println(-1)
		return
	}

	path := []int{}
	current := B
	for current != A {
		path = append(path, current)
		current = parent[current]
	}
	path = append(path, A)

	// Реверсируем путь
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	fmt.Printf("%d\n%s\n", len(path)-1, strings.Trim(fmt.Sprint(path), "[]"))
}
