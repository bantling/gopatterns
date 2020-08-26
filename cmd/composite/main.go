package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node represents individual nodes in the tree. According to this view, every node can have children.
type Node interface {
	fmt.Stringer
	Children() []Node
}

// LeafNode represents nodes that have no children
type LeafNode struct {}

// Children returns nil
func (n LeafNode) Children() []Node {
	return nil
}

// ParentNode represents nodes that do have children
type ParentNode struct {
	children []Node
}

// NewParentNode constructs a ParentNode
func NewParentNode(children ...Node) *ParentNode {
	return &ParentNode{children: children}
}

// Children returns the children, which may be nil
func (n ParentNode) Children() []Node {
	return n.children
}

// Number represents a numeric leaf node
type NumericNode struct {
	LeafNode
	Number int
}

func NewNumericNode(number int) *NumericNode {
	return &NumericNode{Number: number}
}

// String is the NumericNode Stringer
func (n NumericNode) String() string {
	return strconv.Itoa(n.Number)
}

// Operator is an enum of operators
type Operator uint

// Operators
const (
	Addition Operator = iota
	Subtraction
	Multiplication
	Division
)

var (
	operatorToString = map[Operator]string{
		Addition: "+",
		Subtraction: "-",
		Multiplication: "*",
		Division: "/",
	}
)

// String is Operator Stringer
func (op Operator) String() string {
	return operatorToString[op]
}

// Operator represents an operator leaf node
type OperatorNode struct {
	LeafNode
	Operator Operator
}

// NewOperatorNode costructs an OperatorNode
func NewOperatorNode(operator Operator) *OperatorNode {
	return &OperatorNode{Operator: operator}
}

// String is the OperatorNode Stringer
func (n OperatorNode) String() string {
	return n.Operator.String()
}

// BinaryOperation nodes contain two numbers and an operator
type BinaryOperation struct {
	*ParentNode
}

// NewBinaryOperation constructs a BinaryOperation
func NewBinaryOperation(left Node, operator Operator, right Node) *BinaryOperation {
	return &BinaryOperation{
		ParentNode: NewParentNode(left, NewOperatorNode(operator), right),
	}
}

// String is BinaryOperation Stringer
func (b BinaryOperation) String() string {
	return fmt.Sprintf("%s, %s, %s", b.children[0], b.children[1], b.children[2])
}

// InfixTraversal does an infix traversal of the tree
func InfixTraversal(root Node) string {
	var bldr strings.Builder
	
	// infix recursion
	var infix func(n Node)
	infix = func(n Node) {
		if bin, isa := n.(*BinaryOperation); isa {
			children := bin.Children()
			bldr.WriteRune('(')
			infix(children[0])
			bldr.WriteString(children[1].String())
			infix(children[2])
			bldr.WriteRune(')')
		} else {
			bldr.WriteString(n.String())
		}
	}
	infix(root)
	
	return bldr.String()
}

func main() {
	fmt.Println(
		InfixTraversal(
			NewNumericNode(1),
		),
	)
}
