package graph

// Edge is a struct that represents a relationship between two nodes
type Edge struct {
	weight         int
	RelationalType string

	// Case is parental relationship, Node1 is parent, Node2 is child
	Node1 *Node
	Node2 *Node
}

// Node is a struct that represents a person
type Node struct {
	Name  string
	edges []*Edge
}

// Graph is a struct that represents a family tree
type Graph struct {
	nodes map[string]*Node
}

func NewGraph() *Graph {
	var nodes = make(map[string]*Node)
	return &Graph{
		nodes: nodes,
	}
}

func NewEdge(node1 *Node, node2 *Node, weight int, relationalType string) *Edge {
	return &Edge{
		weight,
		relationalType,
		node1,
		node2,
	}
}

func (g *Graph) NodeExists(name string) bool {
	_, ok := g.nodes[name]
	return ok
}

func (g *Graph) NewNode(name string) *Node {
	newNode := &Node{
		name,
		[]*Edge{},
	}
	g.nodes[name] = newNode
	return newNode
}

func (g *Graph) GetNode(name string) *Node {
	if !g.NodeExists(name) {
		return g.NewNode(name)
	}
	return g.nodes[name]
}

func (g *Graph) AddEdges(src string, dest string, weight int) *Node {
	nNode := g.GetNode(dest)
	sNode := g.GetNode(src)
	newEdge := NewEdge(sNode, nNode, weight, "path")
	sNode.edges = append(sNode.edges, newEdge)
	nNode.edges = append(nNode.edges, newEdge)
	return nNode
}

func (g *Graph) AddRelational(parent string, child string) *Node {
	nNode := g.GetNode(child)
	sNode := g.GetNode(parent)
	newEdge := NewEdge(sNode, nNode, 0, "parental")
	sNode.edges = append(sNode.edges, newEdge)
	nNode.edges = append(nNode.edges, newEdge)
	return nNode
}

func (g *Graph) GraphCount() int {
	return len(g.nodes)
}

func (g *Graph) GetNodes() []Node {
	var nodes []Node
	for _, n := range g.nodes {
		nodes = append(nodes, *n)
	}
	return nodes
}

func (n *Node) GetYourAncestors() []*Node {
	return getAncestors(n, n.Name)
}

func (n *Node) GetEdges() []*Edge {
	return n.edges
}

// Todo: make concurrent and more efficient
// getAncestors returns a list of ancestors for a given person in a graph
func getAncestors(person *Node, first string) []*Node {
	ancestors := make([]*Node, 0)
	for _, pEdge := range person.edges {
		parent := pEdge.Node1
		if parent.Name != person.Name {
			ancestors = append(ancestors, getAncestors(parent, first)...)
		}
	}
	if person.Name != first {
		ancestors = append(ancestors, person)
	}
	return ancestors
}

func GraphTest() *Graph {
	g := NewGraph()
	g.AddRelational("Sonny", "Mike")
	g.AddRelational("Mike", "Dunny")
	g.AddRelational("Mike", "Bruce")
	g.AddRelational("Martin", "Phoebe")
	g.AddRelational("Phoebe", "Bruce")
	g.AddRelational("Phoebe", "Dunny")
	g.AddRelational("Anastasia", "Phoebe")
	g.AddRelational("Anastasia", "Ursula")
	g.AddRelational("Martin", "Ursula")
	g.AddRelational("Ursula", "Jacqueline")
	g.AddRelational("Eric", "Jacqueline")
	g.AddRelational("Ellen", "Eric")
	g.AddRelational("Oprah", "Eric")
	g.AddRelational("Eric", "Melody")
	g.AddRelational("Arial", "Melody")
	return g
}
