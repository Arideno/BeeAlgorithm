package main

type Section struct {
	graph *Graph
	coloring []int
	colorsUsed int
}

func (s *Section) getNumberOfColorsUsed() int {
	max := 0
	for _, v := range s.coloring {
		if v > max {
			max = v
		}
	}
	return max
}

func (s *Section) isValidColoring() bool {
	for i, color := range s.coloring {
		connectedVertexes := s.graph.getConnectedVertexes(i)
		for _, vertex := range connectedVertexes {
			if s.coloring[vertex] == color && color != 0 {
				return false
			}
		}
	}
	return true
}
