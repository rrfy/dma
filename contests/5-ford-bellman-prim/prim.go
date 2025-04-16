package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	u, v, w int
}

type MinHeap []Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].w < h[j].w }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Чтение N и M
	scanner.Scan()
	nm := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(nm[0])
	M, _ := strconv.Atoi(nm[1])

	// Построение списка смежности
	adj := make(map[int][]Edge)
	for i := 0; i < M; i++ {
		scanner.Scan()
		uvw := strings.Fields(scanner.Text())
		u, _ := strconv.Atoi(uvw[0])
		v, _ := strconv.Atoi(uvw[1])
		w, _ := strconv.Atoi(uvw[2])

		adj[u] = append(adj[u], Edge{u, v, w})
		adj[v] = append(adj[v], Edge{v, u, w})
	}

	// Алгоритм Прима
	visited := make(map[int]bool)
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	mstEdges := []Edge{}
	totalWeight := 0

	// Начинаем с вершины 1 (можно выбрать любую)
	if len(adj) == 0 && N > 1 {
		fmt.Println(-1)
		return
	}

	start := 1
	visited[start] = true
	for _, edge := range adj[start] {
		heap.Push(minHeap, edge)
	}

	for minHeap.Len() > 0 {
		smallestEdge := heap.Pop(minHeap).(Edge)

		if visited[smallestEdge.v] {
			continue
		}

		visited[smallestEdge.v] = true
		mstEdges = append(mstEdges, smallestEdge)
		totalWeight += smallestEdge.w

		for _, edge := range adj[smallestEdge.v] {
			if !visited[edge.v] {
				heap.Push(minHeap, edge)
			}
		}
	}

	// Проверка связности графа
	if len(visited) != N {
		fmt.Println(-1)
		return
	}

	// Вывод суммарного веса
	fmt.Println(totalWeight)

	// Сортировка ребер в порядке возрастания (по вершинам)
	sort.Slice(mstEdges, func(i, j int) bool {
		if mstEdges[i].u == mstEdges[j].u {
			return mstEdges[i].v < mstEdges[j].v
		}
		return mstEdges[i].u < mstEdges[j].u
	})

	// Формируем строку с ребрами
	var edgesStr strings.Builder
	for i, edge := range mstEdges {
		if i > 0 {
			edgesStr.WriteString(" ")
		}
		if edge.u < edge.v {
			edgesStr.WriteString(fmt.Sprintf("%d %d", edge.u, edge.v))
		} else {
			edgesStr.WriteString(fmt.Sprintf("%d %d", edge.v, edge.u))
		}
	}

	// Выводим все ребра в одной строке
	fmt.Println(edgesStr.String())
}
