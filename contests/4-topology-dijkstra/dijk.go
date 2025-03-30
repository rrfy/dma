package main

import (
	"bufio"
	"fmt"
	"math"
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

	adj := make([][][2]int, N+1)
	for i := 1; i <= N; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		k, _ := strconv.Atoi(parts[0])
		adj[i] = make([][2]int, k)
		for j := 0; j < k; j++ {
			adj[i][j][0], _ = strconv.Atoi(parts[1+2*j])
			adj[i][j][1], _ = strconv.Atoi(parts[1+2*j+1])
		}
	}

	dist := make([]int, N+1)
	parent := make([]int, N+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[A] = 0

	visited := make([]bool, N+1)

	for {
		minDist := math.MaxInt32
		u := -1

		for i := 1; i <= N; i++ {
			if !visited[i] && dist[i] < minDist {
				minDist = dist[i]
				u = i
			}
		}

		if u == -1 || u == B {
			break
		}

		visited[u] = true

		for _, edge := range adj[u] {
			v, w := edge[0], edge[1]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				parent[v] = u
			}
		}
	}

	if dist[B] == math.MaxInt32 {
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

	fmt.Println(dist[B])
	fmt.Println(strings.Trim(fmt.Sprint(path), "[]"))
}

/*python3 run_tests.py contests/4-topology-dijkstra/dijk.go contests/4-topology-dijkstra/dijkstra_tests.zip
Compiling with command: go build -o /tmp/tmp07u0xjuk/solution_exec /home/student/dma/contests/4-topology-dijkstra/dijk.go
Extracted tests to /tmp/tmp07u0xjuk/tests
Test test1 PASSED
Test test2 PASSED
Test test3 PASSED
Test test4 PASSED
Test test5 PASSED
All 5/5 tests passed!*/
