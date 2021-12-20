package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"

	"github.com/awolden/adventofcode2021/helpers"
)

type Node struct {
	data   *int
	lnode  *Node
	rnode  *Node
	parent *Node
}

func (node *Node) AddNode(nodeType string, newNode *Node) {
	switch nodeType {
	case "left":
		node.lnode = newNode
	case "right":
		node.rnode = newNode
	}
}

func (node *Node) ToString() string {
	parentData := ""
	if node.parent == nil {
		parentData = "nil"
	} else {
		parentData = nilOrNil(node.parent.data)
	}
	return fmt.Sprintf("(%p)(data: %s: leftNode: %s: rightNode: %s parent: %s)",
		node, nilOrNil(node.data), nilOrNil(node.lnode.data), nilOrNil(node.rnode.data), parentData,
	)
}

func nilOrNil(i *int) string {
	if i == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", *i)
}

func main() {
	input := readInput()
	// printTree(os.Stdout, &input[0], 0, 'M')

	tree := &input[0]
	for i := 1; i < len(input); i++ {
		tree = addTrees(tree, &input[i])
		reduce(tree)
	}
	// fmt.Println("-------- Final ----------")
	fmt.Println("Result:", formalPrintTree(tree))
	fmt.Println("magnitude:", getMagnitude(tree))
	fmt.Println("----------------------------")
	newInput := readInput()
	getLargestMagnitudes(&newInput)

}

func getLargestMagnitudes(input *[]Node) {
	highest := 0
	for i := range *input {
		for j := range *input {
			if i == j {
				continue
			}
			newTree := makeFreshTree(i, j)
			reduce(newTree)
			magnitude := getMagnitude(newTree)
			//fmt.Println("testing:", formalPrintTree(&tree), "+", formalPrintTree(&tree2), ":", magnitude, formalPrintTree(newTree))
			if magnitude > highest {
				highest = magnitude
				//fmt.Println("New Highest:", magnitude, formalPrintTree(&tree), "+", formalPrintTree(&tree2))
			}
		}
	}
	fmt.Println("Highest", highest)
}

func makeFreshTree(i int, j int) *Node {
	input := readInput()
	return addTrees(&input[i], &input[j])
}

func reduce(tree *Node) {

	// fmt.Println("-------- INITIAL ----------")
	// fmt.Println(formalPrintTree(tree))
	// fmt.Println("----------------------------")
	//fmt.Println("Reducing:", formalPrintTree(tree))
	for treeNeedsExplode(tree, 0) || treeNeedsSplit(tree) {
		healTree(tree)
		if treeNeedsExplode(tree, 0) {
			explodeTreeSimply(tree)
			//fmt.Println("-----------Exploded-----------------")
			//fmt.Printf("exploded:%s\n", formalPrintTree(tree))
			//fmt.Println("----------------------------")
			continue
		}
		if treeNeedsSplit(tree) {
			splitTreeSimply(tree)
			//fmt.Println("-----------Split-----------------")
			//fmt.Printf("splitted:%s\n", formalPrintTree(tree))
			//fmt.Println("----------------------------")
			continue
		}
	}
}

func readInput() []Node {
	rawInput := helpers.GetFileArray("./input")
	nodes := []Node{}

	for _, line := range rawInput {
		mixed := []interface{}{}
		json.Unmarshal([]byte(line), &mixed)
		// fmt.Println("line", mixed)
		node := Node{}
		nestedArrToTree(mixed, &node)
		nodes = append(nodes, node)
	}
	return nodes
}

func addTrees(node *Node, node2 *Node) *Node {
	// fmt.Println(node, node2)
	newTrunk := Node{
		lnode: node,
		rnode: node2,
	}
	node.parent = &newTrunk
	node2.parent = &newTrunk
	return &newTrunk
}

func explodeTreeSimply(node *Node) {
	explodeTree(node, 0, GetBoolPointer(false), node)
}

func explodeTree(node *Node, level int, hasExploded *bool, fullTree *Node) {
	if *hasExploded || node == nil || node.data != nil {
		return
	}

	sortedNodeList := []*Node{}
	getSortedLeafs(fullTree, &sortedNodeList)

	if level >= 4 && node.lnode.data != nil && node.rnode.data != nil {

		//fmt.Println("exploding", formalPrintTree(node))
		//fmt.Println("exploding", prettyPrintQueue(&sortedNodeList))
		leftIndex := findIndexOfNode(sortedNodeList, node.lnode)
		rightIndex := findIndexOfNode(sortedNodeList, node.rnode)
		if leftIndex > 0 {
			data := sortedNodeList[leftIndex-1].data
			sortedNodeList[leftIndex-1].data = GetIntPointer(*data + *node.lnode.data)
		}
		if rightIndex < len(sortedNodeList)-1 {
			data := sortedNodeList[rightIndex+1].data
			sortedNodeList[rightIndex+1].data = GetIntPointer(*data + *node.rnode.data)
		}
		node.lnode = nil
		node.rnode = nil
		node.data = GetIntPointer(0)
		*hasExploded = true
		return
	} else {
		explodeTree(node.lnode, level+1, hasExploded, fullTree)
		explodeTree(node.rnode, level+1, hasExploded, fullTree)
	}
}

func splitTreeSimply(node *Node) {
	splitTree(node, GetBoolPointer(false))
}

func splitTree(node *Node, hasSplit *bool) {
	if *hasSplit {
		return
	}
	if node.lnode.data != nil && *node.lnode.data > 9 {
		// fmt.Println("Splitting:", *node.lnode.data)
		data := float64(*node.lnode.data)
		newLeft := Node{
			data:   GetIntPointer(int(math.Floor(data / 2))),
			parent: node.lnode,
		}
		newRight := Node{
			data:   GetIntPointer(int(math.Ceil(data / 2))),
			parent: node.lnode,
		}
		node.lnode.data = nil
		node.lnode.lnode = &newLeft
		node.lnode.rnode = &newRight
		*hasSplit = true
		return
	}
	if node.lnode.data == nil {
		splitTree(node.lnode, hasSplit)
	}
	if node.rnode.data != nil && *node.rnode.data > 9 && !*hasSplit {
		// fmt.Println("Splitting:", *node.rnode.data)
		data := float64(*node.rnode.data)
		//fmt.Println("spliting right")
		newLeft := Node{
			data:   GetIntPointer(int(math.Floor(data / 2))),
			parent: node.rnode,
		}
		newRight := Node{
			data:   GetIntPointer(int(math.Ceil(data / 2))),
			parent: node.rnode,
		}
		node.rnode.data = nil
		node.rnode.lnode = &newLeft
		node.rnode.rnode = &newRight
		*hasSplit = true
		return
	}
	if node.rnode.data == nil {
		splitTree(node.rnode, hasSplit)
	}

}
func treeNeedsExplode(node *Node, level int) bool {
	if level > 4 {
		return true
	}

	needsExplode := false
	if node.lnode != nil && treeNeedsExplode(node.lnode, level+1) {
		needsExplode = true
	}
	if node.rnode != nil && treeNeedsExplode(node.rnode, level+1) {
		needsExplode = true
	}
	return needsExplode
}

func treeNeedsSplit(node *Node) bool {
	if node.data != nil && *node.data > 9 {
		return true
	}

	needsSplit := false
	if node.lnode != nil && treeNeedsSplit(node.lnode) {
		needsSplit = true
	}
	if node.rnode != nil && treeNeedsSplit(node.rnode) {
		needsSplit = true
	}
	return needsSplit
}

func healTree(node *Node) {
	if node.lnode != nil {
		node.lnode.parent = node
		healTree(node.lnode)
	}
	if node.rnode != nil {
		node.rnode.parent = node
		healTree(node.rnode)
	}
}

func printTree(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}

	if node.data == nil {
		fmt.Fprintf(w, "%c(%p):%v\n", ch, node, " ")
	} else {
		fmt.Fprintf(w, "%c(%p):%v\n", ch, node, *node.data)
	}

	printTree(w, node.lnode, ns+2, 'L')
	printTree(w, node.rnode, ns+2, 'R')

}

func getMagnitude(node *Node) int {
	if node.data != nil {
		return *node.data
	}

	return getMagnitude(node.lnode)*3 + getMagnitude(node.rnode)*2

}

func formalPrintTree(node *Node) string {
	if node == nil {
		return ""
	}
	if node.data != nil {
		return fmt.Sprintf("%d", *node.data)
	} else {
		return fmt.Sprintf("[%s,%s]", formalPrintTree(node.lnode), formalPrintTree(node.rnode))
	}

}

func nestedArrToTree(nestedArr []interface{}, parentNode *Node) {
	for i, val := range nestedArr {
		//fmt.Println("building", i, val)
		nodeType := "left"
		if i == 1 {
			nodeType = "right"
		}

		newNode := Node{
			parent: parentNode,
		}
		switch v := val.(type) {
		case float64:
			//fmt.Println("int")
			newNode.data = GetIntPointer(int(v))
		case []interface{}:
			//fmt.Println("interface")
			nestedArrToTree(v, &newNode)
		}
		//fmt.Println("adding node", nodeType, newNode)
		parentNode.AddNode(nodeType, &newNode)
	}

}

func prettyPrintQueue(sortedNodeList *[]*Node) string {
	str := ""
	for _, node := range *sortedNodeList {
		str += fmt.Sprintf("%d,", *node.data)
	}
	return str
}

func GetIntPointer(value int) *int {
	return &value
}
func GetBoolPointer(b bool) *bool {
	boolVar := b
	return &boolVar
}

func getSortedLeafs(node *Node, sortedNodeList *[]*Node) {

	if node.data != nil {
		// fmt.Println("adding to list", *node.data)
		*sortedNodeList = append(*sortedNodeList, node)
	} else {
		// fmt.Println("processingNext", node.ToString())
		getSortedLeafs(node.lnode, sortedNodeList)
		getSortedLeafs(node.rnode, sortedNodeList)
	}
}

func findIndexOfNode(sortedList []*Node, node *Node) int {
	for i, nodeInList := range sortedList {
		if node == nodeInList {
			return i
		}
	}
	return -1
}
