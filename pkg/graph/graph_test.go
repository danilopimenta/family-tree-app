package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
}

func TestGraph_NodeExists(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	assert.False(t, g.NodeExists("a"))
	assert.False(t, g.NodeExists("b"))
	assert.False(t, g.NodeExists("c"))
	g.NewNode("a")
	assert.True(t, g.NodeExists("a"))
	assert.False(t, g.NodeExists("b"))
	assert.False(t, g.NodeExists("c"))
	g.NewNode("b")
	assert.True(t, g.NodeExists("a"))
	assert.True(t, g.NodeExists("b"))
	assert.False(t, g.NodeExists("c"))
	g.NewNode("c")
	assert.True(t, g.NodeExists("a"))
	assert.True(t, g.NodeExists("b"))
	assert.True(t, g.NodeExists("c"))
}

func TestGraph_NewNode(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	a := g.NewNode("a")
	assert.NotNil(t, a)
	assert.NotNil(t, a.edges)
	assert.Equal(t, 0, len(a.edges))
	assert.Equal(t, "a", a.Name)
}

func TestGraph_GetNode(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	a := g.GetNode("a")
	assert.NotNil(t, a)
	assert.NotNil(t, a.edges)
	assert.Equal(t, 0, len(a.edges))
	assert.Equal(t, "a", a.Name)
	a = g.GetNode("a")
	assert.NotNil(t, a)
	assert.NotNil(t, a.edges)
	assert.Equal(t, 0, len(a.edges))
	assert.Equal(t, "a", a.Name)
}

func TestGraph_AddRelational(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	g.AddRelational("a", "b")
	assert.Equal(t, 2, len(g.nodes))
	a := g.GetNode("a")
	assert.NotNil(t, a)
	assert.NotNil(t, a.edges)
	assert.Equal(t, 1, len(a.edges))
	assert.Equal(t, "a", a.Name)
	b := g.GetNode("b")
	assert.NotNil(t, b)
	assert.NotNil(t, b.edges)
	assert.Equal(t, 1, len(b.edges))
	assert.Equal(t, "b", b.Name)
}

func TestGraph_AddEdges(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	g.AddEdges("a", "b", 1)
	assert.Equal(t, 2, len(g.nodes))
	a := g.GetNode("a")
	assert.NotNil(t, a)
	assert.NotNil(t, a.edges)
	assert.Equal(t, 1, len(a.edges))
	assert.Equal(t, "a", a.Name)
	b := g.GetNode("b")
	assert.NotNil(t, b)
	assert.NotNil(t, b.edges)
	assert.Equal(t, 1, len(b.edges))
	assert.Equal(t, "b", b.Name)
}

func TestGraph_GraphCount(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	g.AddEdges("a", "b", 1)
	assert.Equal(t, 2, len(g.nodes))
	assert.Equal(t, 2, g.GraphCount())
}

func TestNode_GetEdges(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	g.AddEdges("a", "b", 1)
	assert.Equal(t, 2, len(g.nodes))
	assert.Equal(t, 2, g.GraphCount())
	aEdges := g.nodes["a"].GetEdges()
	assert.Equal(t, 1, len(aEdges))
	bEdges := g.nodes["b"].GetEdges()
	assert.Equal(t, 1, len(bEdges))
}

func TestNode_GetYourAncestors(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	g.AddEdges("a", "b", 1)
	assert.Equal(t, 2, len(g.nodes))
	assert.Equal(t, 2, g.GraphCount())
	aEdges := g.nodes["a"].GetEdges()
	assert.Equal(t, 1, len(aEdges))
	bEdges := g.nodes["b"].GetEdges()
	assert.Equal(t, 1, len(bEdges))
	aAncestors := g.nodes["a"].GetYourAncestors()
	assert.Equal(t, 0, len(aAncestors))
	bAncestors := g.nodes["b"].GetYourAncestors()
	assert.Equal(t, 1, len(bAncestors))
}

func TestGraph_GetNodes(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.NotNil(t, g.nodes)
	assert.Equal(t, 0, len(g.nodes))
	g.AddEdges("a", "b", 1)
	assert.Equal(t, 2, len(g.nodes))
	assert.Equal(t, 2, g.GraphCount())
	nodes := g.GetNodes()
	assert.Equal(t, 2, len(nodes))
}
