package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type BeeAlgorithm struct {
	graph *Graph
	sections []*Section
	workBees int
	bees int
}

func (b *BeeAlgorithm) Start() {
	b.generateSections()
	for iteration := 0; iteration < 1000; iteration++ {
		sections := b.getRandomSections()
		for _, section := range sections {
			workBees := b.workBees / section.colorsUsed
			usedVertexes := make(map[int]bool)
			for workBees > 0 {
				maxDegree := 0
				maxVertex := -1
				for i := 0; i < b.graph.numVertexes; i++ {
					sum := 0
					for j := 0; j < b.graph.numVertexes; j++ {
						sum += b.graph.adjMatrix[i][j]
					}
					if sum > maxDegree {
						if _, ok := usedVertexes[i]; !ok {
							maxDegree = sum
							maxVertex = i
						}
					}
				}
				if maxVertex == -1 {
					break
				}
				usedVertexes[maxVertex] = true
				for _, connectedVertex := range b.graph.getConnectedVertexes(maxVertex) {
					if workBees <= 0 {
						break
					}
					workBees--
					currentColor := section.coloring[connectedVertex]
					for i := 1; i < currentColor; i++ {
						section.coloring[connectedVertex] = i
						if !section.isValidColoring() {
							section.coloring[connectedVertex] = currentColor
						} else {
							break
						}
					}
				}
			}
			section.colorsUsed = section.getNumberOfColorsUsed()
		}
	}
	for _, v := range b.sections {
		if v != nil {
			fmt.Printf("%d ", v.colorsUsed)
		}
	}
}

func (b *BeeAlgorithm) getRandomSections() []*Section {
	var sections []*Section
	selected := make(map[int]bool)
	for i := 0; i < int(math.Min(float64(b.bees), float64(b.graph.numVertexes))); i++ {
		for {
			index := rand.Intn(len(b.sections))
			if b.sections[index] == nil {
				continue
			}
			if _, ok := selected[index]; !ok {
				selected[index] = true
				sections = append(sections, b.sections[index])
				break
			}
		}
	}
	return sections
}

func dfs(v, color int, visited []bool, section *Section) {
	visited[v] = true
	section.coloring[v] = color
	for _, i := range section.graph.getConnectedVertexes(v) {
		if !visited[i] {
			for j := 1; ; j++ {
				section.coloring[i] = j
				if section.isValidColoring() {
					dfs(i, j, visited, section)
					break
				} else {
					section.coloring[i] = 0
				}
			}
		}
	}
}

func (b *BeeAlgorithm) generateSections() {
	for i := 0; i < int(math.Min(float64(len(b.sections)), float64(b.graph.numVertexes))); i++ {
		visited := make([]bool, b.graph.numVertexes)
		section := &Section{
			graph:      b.graph,
			coloring:   make([]int, b.graph.numVertexes),
			colorsUsed: 0,
		}
		dfs(i, 1, visited, section)
		section.colorsUsed = section.getNumberOfColorsUsed()
		b.sections[i] = section
	}
}

func NewBeeAlgorithm(workBees, bees int) *BeeAlgorithm {
	rand.Seed(time.Now().Unix())
	graph := generateGraph(100)
	return &BeeAlgorithm{
		graph: graph,
		sections: make([]*Section, bees * 4),
		workBees: workBees,
		bees: bees,
	}
}
