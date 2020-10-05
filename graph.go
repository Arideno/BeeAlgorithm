package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Graph struct {
	adjMatrix [][]int
	numVertexes int
}

func (g *Graph) getDegrees() []int {
	var degrees []int
	for _, v := range g.adjMatrix {
		degree := 0
		for _, k := range v {
			degree += k
		}
		degrees = append(degrees, degree)
	}
	return degrees
}

func (g *Graph) printDegrees() {
	for _, v := range g.adjMatrix {
		degree := 0
		for _, k := range v {
			degree += k
		}
		fmt.Println(degree)
	}
}

func (g *Graph) getConnectedVertexes(vertex int) []int {
	var vertexes []int
	for i, v := range g.adjMatrix[vertex] {
		if v == 1 {
			vertexes = append(vertexes, i)
		}
	}
	return vertexes
}

func randomRange(min, max int) int {
	return rand.Intn(max + 1 - min) + min
}

func generateGraph(size int) *Graph {
	graph := &Graph{
		numVertexes: size,
	}
	graph.adjMatrix = make([][]int, size)
	for i := 0; i < size; i++ {
		graph.adjMatrix[i] = make([]int, size)
	}

	for vertex := 0; vertex < size; vertex++ {
		connections := graph.adjMatrix[vertex]
		currentDegree := 0
		for _, v := range connections {
			currentDegree += v
		}
		degree := int(math.Min(float64(randomRange(1, 20)-currentDegree), float64(size-vertex-1)))
		for i := 0; i < degree; i++ {
			notConnected := true
			tries := 0
			var newVertex int
			for notConnected && tries < size {
				newVertex = randomRange(vertex + 1, size - 1)
				tries++
				newVertexDegree := 0
				for _, v := range graph.adjMatrix[newVertex] {
					newVertexDegree += v
				}
				if connections[newVertex] == 0 && newVertexDegree <= 20 {
					notConnected = false
					graph.adjMatrix[vertex][newVertex] = 1
					graph.adjMatrix[newVertex][vertex] = 1
				}
			}
		}
	}
	return graph
}
