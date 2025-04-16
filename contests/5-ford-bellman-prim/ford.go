package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	u, v, w int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])
	A, _ := strconv.Atoi(parts[2])
	B, _ := strconv.Atoi(parts[3])

	edges := make([]Edge, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		w, _ := strconv.Atoi(parts[2])
		edges[i] = Edge{u, v, w}
	}

	dist := make([]int, N+1)
	parent := make([]int, N+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[A] = 0

	// Релаксация рёбер
	for i := 1; i < N; i++ {
		for _, edge := range edges {
			if dist[edge.u] != math.MaxInt32 && dist[edge.u]+edge.w < dist[edge.v] {
				dist[edge.v] = dist[edge.u] + edge.w
				parent[edge.v] = edge.u
			}
		}
	}

	// Проверка на отрицательные циклы
	for _, edge := range edges {
		if dist[edge.u] != math.MaxInt32 && dist[edge.u]+edge.w < dist[edge.v] {
			fmt.Println(-1)
			return
		}
	}

	if dist[B] == math.MaxInt32 {
		fmt.Println(-1)
		return
	}

	// Восстановление пути
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
	if dist[B] == -1 {
		fmt.Println(dist[B])
	} else {
		fmt.Println(dist[B])
		fmt.Println(strings.Trim(fmt.Sprint(path), "[]"))
	}

}
