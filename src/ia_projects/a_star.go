package main

import (
	"fmt"
	"sort"
)

type Node struct {
	name         string
	neighbours   map[*Node]int
	heuristic    map[*Node]int
	hasSemaphore map[*Node]bool
	previous     *Node
	cost         int
}

func showFullPath(n *Node) {
	if n.previous != nil {
		showFullPath(n.previous)
		fmt.Println(n)
	} else {
		fmt.Println(n)
	}
}

func (n Node) getNeighbours() (keys []*Node) {
	keys = make([]*Node, len(n.neighbours))
	i := 0
	for key := range n.neighbours {
		keys[i] = key
		i++
	}
	return
}

func (n Node) String() string {
	return fmt.Sprintf("%s: %d", n.name, n.cost)
}

type AStar struct {
	initial *Node
	destiny *Node
	border  []*Node
}

func (a AStar) exec() *Node {
	a.border = make([]*Node, 0, 100)
	a.border = append(a.border, a.initial)

	for {
		if a.border[0] == a.destiny {
			break
		}
		temp := a.border[0]
		for _, n := range temp.getNeighbours() {
			n.cost = temp.neighbours[n] + temp.heuristic[n] + temp.cost
			n.previous = temp
			a.border = append(a.border, n)
		}
		a.border = a.border[1:]
		sort.Sort(ByCost(a.border))
	}
	return a.border[0]
}

type ByCost []*Node

func (a ByCost) Len() int           { return len(a) }
func (a ByCost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCost) Less(i, j int) bool { return a[i].cost < a[j].cost }

func main() {
	nodeA := &Node{
		name:         "A",
		neighbours:   make(map[*Node]int),
		heuristic:    make(map[*Node]int),
		hasSemaphore: make(map[*Node]bool),
	}
	nodeB := &Node{
		name:         "B",
		neighbours:   make(map[*Node]int),
		heuristic:    make(map[*Node]int),
		hasSemaphore: make(map[*Node]bool),
	}
	nodeC := &Node{
		name:         "C",
		neighbours:   make(map[*Node]int),
		heuristic:    make(map[*Node]int),
		hasSemaphore: make(map[*Node]bool),
	}
	nodeD := &Node{
		name:         "D",
		neighbours:   make(map[*Node]int),
		heuristic:    make(map[*Node]int),
		hasSemaphore: make(map[*Node]bool),
	}

	nodeA.neighbours[nodeB] = 2
	nodeA.heuristic[nodeB] = 4

	nodeA.neighbours[nodeC] = 3
	nodeA.heuristic[nodeC] = 1

	nodeB.neighbours[nodeD] = 3
	nodeB.heuristic[nodeD] = 2

	nodeC.neighbours[nodeD] = 4
	nodeC.heuristic[nodeD] = 4

	as := &AStar{initial: nodeA, destiny: nodeD}
	showFullPath(as.exec())
}
