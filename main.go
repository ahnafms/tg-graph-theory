package main

import "fmt"

type matrix [][]int

func main() {
	graph := make(matrix, 6)

	for i := range graph {
		graph[i] = make([]int, 6)
	}

	graph.addEdge(0, 1)
	graph.addEdge(1, 0)
	graph.addEdge(1, 3)
	graph.addEdge(1, 4)
	graph.addEdge(2, 0)
	graph.addEdge(2, 5)
	graph.addEdge(3, 1)
	graph.addEdge(4, 1)
	graph.addEdge(5, 2)

	graph.bfs(0)
}

func (m matrix) bfs(s int) {
	visited := make([]int, len(m))
	queue := []int{}

	for i := range visited {
		visited[i] = 0
	}

	visited[s] = 1
	queue = append(queue, s)
	fmt.Printf("Visited : %d\n", s)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for i, connected := range m[node] {
			if connected == 1 && visited[i] == 0 {
				visited[i] = 1
				fmt.Printf("Visited : %d\n", i)
				queue = append(queue, i)
			}
		}
	}
}

func (m matrix) addEdge(u int, v int) {
	m[u][v] = 1
	m[v][u] = 1
}
