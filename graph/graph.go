package graph

type Graph struct {
	vertices int
	edges []Edge
}

type Edge struct {
	from int
	to int
}
