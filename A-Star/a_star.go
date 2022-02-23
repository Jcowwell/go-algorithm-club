package astar

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

type Node[E comparable] struct {
	vertex   E       // The graph vertex.
	cost     float32 // The actual cost between the start vertex and this vertex.
	estimate float32 // Estimated (heuristic) cost betweent this vertex and the target vertex.
}

func (n Node[E]) init(vertex E, cost, estimate float32) {
	n.vertex = vertex
	n.cost = cost
	n.estimate = estimate
}

func (n Node[E]) hashValue() uintptr {
	return Hash(n.vertex)

}

func less[E comparable, n Node[E]](lhs, rhs Node[E]) bool {
	return lhs.cost+lhs.estimate < rhs.cost+lhs.estimate
}

func equal[E comparable, n Node[E]](lhs, rhs Node[E]) bool {
	return lhs.vertex == rhs.vertex
}

type Edge[E comparable, V Node[E]] struct {
	Vertex V       // The edge vertex
	cost   float32 // The edge's cost.
	target float32 // The target vertex.
}

type Graph[E comparable, V Node[E], WE Edge[E, Node[E]]] struct {
	Vertex V
	Edge   WE
	Edges  func(V) []WE // Lists all edges going out from a vertex.
}

type AStar[E comparable, N Node[E], WE Edge[E, Node[E]], G Graph[E, N, WE]] struct {
	graph     G                           // The graph to search on.
	heuristic func(N, N) float32          // The heuristic cost function that estimates the cost between two vertices.
	open      func()                      // Open list of nodes to expand. HashedHeap
	closed    map[N]string                // Set of vertices already expanded.
	costs     map[N]float32               // Actual vertex cost for vertices we already encountered (refered to as `g` on the literature).
	parents   map[N]N                     // Store the previous node for each expanded node to recreate the path.
	init      func(G, func(N, N) float32) // Initializes `AStar` with a graph and a heuristic cost function.
	path      func(N) []N                 // Finds an optimal path between `source` and `target`. - Precondition: both `source` and `target` belong to `graph`.
	expand    func(N, N)
	cost      func(N) float32
	buildPath func(N, N) []N
	cleanup   func()
}
