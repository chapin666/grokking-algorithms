package chapter7

import (
	"fmt"
)

const (
	inf = 0x7FF0000000000000
)

func dijkstra() {

	// 一、init nodes
	graph := make(map[string]map[string]int)

	// node 'start '
	graph["start"] = make(map[string]int)
	// the edge(key) and weight(value) of node start.a
	graph["start"]["a"] = 6
	// the edge(key) and weight(value) of node start.b
	graph["start"]["b"] = 2

	// node 'a'
	graph["a"] = make(map[string]int)
	// the edge(key) and weight(value) of node a.fin
	graph["a"]["fin"] = 1

	// node 'b'
	graph["b"] = make(map[string]int)
	// the edge(key) and weight(value) of node b.a
	graph["b"]["a"] = 3
	// the edge(key) and weight(value) of node b.fin
	graph["b"]["fin"] = 5

	// node 'fin'(end)
	graph["fin"] = make(map[string]int)

	// 二、node costs
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = inf // end node

	// 三、parent node of childs
	parent := make(map[string]interface{})
	parent["a"] = "start"
	parent["b"] = "start"
	parent["fin"] = nil

	// 四、record processed element
	processed := []string{}

	// 五、implements
	minNode, minWeight := findLowestConstNode(costs, processed)
	for minNode != "" && minWeight != inf {
		// find const lowest node
		cost := costs[minNode]
		// get lowest node neighbors
		neighbors := graph[minNode]

		for n, _ := range neighbors {
			newCost := cost + neighbors[n]
			// if current node is closer
			if costs[n] > newCost {
				costs[n] = newCost
				parent[n] = minNode
			}
		}

		processed = append(processed, minNode)
		minNode, minWeight = findLowestConstNode(costs, processed)
	}

	fmt.Println(parent)

}

// findLowestConstNode return lowest value from costs
func findLowestConstNode(costs map[string]int, processed []string) (node string, weight int) {

	weight = inf
	node = ""

	for k, v := range costs {
		if v < weight && !contain(processed, k) {
			weight = v
			node = k
		}
	}

	return
}

// return  processed contain key
func contain(processd []string, key string) bool {
	for _, value := range processd {
		if value == key {
			return true
		}
	}
	return false
}
