// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multi

import (
	"testing"

	"gonum.org/v1/gonum/graph"
)

// Tests Issue #27
func TestWeightedEdgeOvercounting(t *testing.T) {
	g := generateDummyWeightedGraph()

	if neigh := graph.NodesOf(g.From(int64(2))); len(neigh) != 2 {
		t.Errorf("Node 2 has incorrect number of neighbors got neighbors %v (count %d), expected 2 neighbors {0,1}", neigh, len(neigh))
	}
}

func generateDummyWeightedGraph() *WeightedDirectedGraph {
	nodes := [4]struct{ srcID, targetID int }{
		{2, 1},
		{1, 0},
		{2, 0},
		{0, 2},
	}

	g := NewWeightedDirectedGraph()

	for i, n := range nodes {
		g.SetWeightedLine(WeightedLine{F: Node(n.srcID), T: Node(n.targetID), W: 1, UID: int64(i)})
	}

	return g
}

// Test for issue #123 https://github.com/gonum/graph/issues/123
func TestIssue123WeightedDirectedGraph(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	g := NewWeightedDirectedGraph()

	n0 := g.NewNode()
	g.AddNode(n0)

	n1 := g.NewNode()
	g.AddNode(n1)

	g.RemoveNode(n0.ID())

	n2 := g.NewNode()
	g.AddNode(n2)
}
