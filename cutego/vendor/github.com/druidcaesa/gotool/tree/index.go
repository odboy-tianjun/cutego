package tree

import "sort"

// Tree uniformly defines the data structure of the menu tree, you can also customize other fields to add
type Tree struct {
	Title           string      `json:"title"`
	Data            interface{} `json:"data"`
	Leaf            bool        `json:"leaf"`
	Selected        bool        `json:"checked"`
	PartialSelected bool        `json:"partiallySelected"`
	Children        []Tree      `json:"children"`
}

// INode Other structures of  want to generate a menu tree and directly implement this interface
type INode interface {
	GetTitle() string
	GetId() int
	GetPid() int
	GetData() interface{}
	IsRoot() bool
}
type INodes []INode

func (nodes INodes) Len() int {
	return len(nodes)
}
func (nodes INodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}
func (nodes INodes) Less(i, j int) bool {
	return nodes[i].GetId() < nodes[j].GetId()
}

// GenerateTree After the  custom structure implements the INode interface, call this method to generate the tree structure
// nodes need to generate tree nodes
// selectedNode selected node after spanning tree
// The tree structure object after menuTrees is successfully generated
func (Tree) GenerateTree(nodes, selectedNodes []INode) (trees []Tree) {
	trees = []Tree{}
	// Define the top-level root and child nodes
	var roots, childs []INode
	for _, v := range nodes {
		if v.IsRoot() {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &Tree{
			Title: v.GetTitle(),
			Data:  v.GetData(),
		}
		// Before recursion, confirm the selected state of childTree according to the parent node
		childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children)
		// recursive
		recursiveTree(childTree, childs, selectedNodes)
		// After recursion, confirm the selected state of childTree according to the child node
		if !childTree.Selected {
			childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children)
		}
		// After recursion, confirm the half-selected state of childTree according to the child nodes
		childTree.PartialSelected = nodePartialSelected(childTree.Children)
		// After recursion, confirm whether it is a leaf node according to the subsection
		childTree.Leaf = len(childTree.Children) == 0
		trees = append(trees, *childTree)
	}
	return
}

// recursiveTree recursively spanning tree structure
// tree recursive tree object
// nodes recursive nodes
// selectedNodes selected nodes
func recursiveTree(tree *Tree, nodes, selectedNodes []INode) {
	data := tree.Data.(INode)

	for _, v := range nodes {
		if v.IsRoot() {
			// If the current node is the top-level root node, skip
			continue
		}
		if data.GetId() == v.GetPid() {
			childTree := &Tree{
				Title: v.GetTitle(),
				Data:  v.GetData(),
			}
			childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children) || tree.Selected
			recursiveTree(childTree, nodes, selectedNodes)

			if !childTree.Selected {
				childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children)
			}
			childTree.PartialSelected = nodePartialSelected(childTree.Children)
			childTree.Leaf = len(childTree.Children) == 0
			tree.Children = append(tree.Children, *childTree)
		}
	}
}

// FindRelationNode queries all parent nodes of nodes in nodes in allTree
// nodes to query the child node array of the parent node
// allTree all nodes array
func (Tree) FindRelationNode(nodes, allNodes []INode) (respNodes []INode) {
	nodeMap := make(map[int]INode)
	for _, v := range nodes {
		recursiveFindRelationNode(nodeMap, allNodes, v, 0)
	}

	for _, v := range nodeMap {
		respNodes = append(respNodes, v)
	}
	sort.Sort(INodes(respNodes))
	return
}

// recursiveFindRelationNode recursively query related parent and child nodes
// nodeMap query results are collected in the map
// allNodes all nodes
// node recursive node
// t Recursive search type: 0 finds parent and child nodes; 1 only finds parent nodes; 2 only finds child nodes
func recursiveFindRelationNode(nodeMap map[int]INode, allNodes []INode, node INode, t int) {
	nodeMap[node.GetId()] = node
	for _, v := range allNodes {
		if _, ok := nodeMap[v.GetId()]; ok {
			continue
		}
		// Find the parent node
		if t == 0 || t == 1 {
			if node.GetPid() == v.GetId() {
				nodeMap[v.GetId()] = v
				if v.IsRoot() {
					// When it is the top-level root node, no recursion
					continue
				}
				recursiveFindRelationNode(nodeMap, allNodes, v, 1)
			}
		}
		// Find child nodes
		if t == 0 || t == 2 {
			if node.GetId() == v.GetId() {
				nodeMap[v.GetId()] = v
				recursiveFindRelationNode(nodeMap, allNodes, v, 2)
			}
		}
	}
}

// nodeSelected determines the selected state of the node
// node judges the node
func nodeSelected(node INode, selectedNodes []INode, children []Tree) bool {
	for _, v := range selectedNodes {
		if node.GetId() == v.GetId() {
			// 1. If the current node exists in the selected node array
			return true
		}
	}

	if len(children) == 0 {
		// 2. Precondition 1 is not met, and there are no child nodes
		return false
	}
	selectedNum := 0
	for _, v := range children {
		if v.Selected {
			selectedNum++
		}
	}
	if selectedNum == len(children) {
		// The preconditions 1, 2 are not met, and all the child nodes are selected
		return true
	}
	return false
}

// nodePartialSelected judges the half-selected state of the node
func nodePartialSelected(trees []Tree) bool {
	selectedNum := 0
	for _, v := range trees {
		if v.Selected {
			selectedNum++
		}
	}
	if selectedNum == len(trees) || selectedNum == 0 {
		// All child nodes are selected, or none of them are selected
		return false
	}
	return true
}
