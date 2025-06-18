/*
Алгоритм Дейкстры поиск минимального веса по взвешенному графу.
Грокаем Алгоритмы глава 9.
*/
package deikstra

import (
	"math"
)

type Node map[string]int
type Graph map[string]Node
type Costs map[string]int
type Processed map[string]bool
type Parents map[string]string

func Deikstra(graph Graph, start, end string) ([]string, int) {
	costs := make(Costs)
	parents := make(Parents)
	processed := make(Processed)

	// Все затраты устанавливаются как бесконечность
	for node := range graph {
		costs[node] = math.MaxInt32
	}

	// Начальная точка равна 0
	costs[start] = 0

	// Основной цикл
	node := findLowestCostNode(costs, processed)
	for node != "" {
		neighbors := graph[node]
		for nextNode, edgeWeight := range neighbors {
			newCost := costs[node] + edgeWeight
			if costs[nextNode] > newCost {
				costs[nextNode] = newCost
				parents[nextNode] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode(costs, processed)
	}

	// Построение пути
	path := []string{end}
	current := end
	for current != start {
		current = parents[current]
		path = append(path, current)
	}

	// Обратная сортировка пути
	reversedPath := make([]string, len(path))
	for i, j := 0, len(path)-1; i <= j; i, j = i+1, j-1 {
		reversedPath[i], reversedPath[j] = path[j], path[i]
	}

	return reversedPath, costs[end]
}

func findLowestCostNode(costs Costs, processed Processed) string {
	lowestCost := math.MaxInt32
	lowestCostNode := ""
	for node, cost := range costs {
		if cost < lowestCost && !processed[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}
