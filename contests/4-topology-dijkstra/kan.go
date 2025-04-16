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
	inDegree := make([]int, N+1)

	for i := 1; i <= N; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		k, _ := strconv.Atoi(parts[0])
		adj[i] = make([]int, k)
		for j := 0; j < k; j++ {
			adj[i][j], _ = strconv.Atoi(parts[j+1])
			inDegree[adj[i][j]]++
		}
	}

	queue := []int{}
	for i := 1; i <= N; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}
	count := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		count++

		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if count != N {
		fmt.Println(-1)
	} else {
		fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
	}
}
