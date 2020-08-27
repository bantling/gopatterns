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
			// left
			bldr.WriteRune('(')
			infix(children[0])
			// operator
			bldr.WriteString(children[1].String())
			// right
			infix(children[2])
			bldr.WriteRune(')')
		} else {
			bldr.WriteString(n.String())
		}
	}
	infix(root)
	
	return bldr.String()
}

// PrefixTraversal does a prefix traversal of the tree
func PrefixTraversal(root Node) string {
	var bldr strings.Builder
	
	// prefix recursion
	var prefix func(n Node)
	prefix = func(n Node) {
		if bin, isa := n.(*BinaryOperation); isa {
			children := bin.Children()
			// operator
			bldr.WriteString(children[1].String())
			// left
			prefix(children[0])
			// right
			prefix(children[2])
		} else {
			bldr.WriteString(n.String())
		}
	}
	prefix(root)
	
	return bldr.String()
}

// PostfixTraversal does a prefix traversal of the tree
func PostfixTraversal(root Node) string {
	var bldr strings.Builder
	
	// postfix recursion
	var postfix func(n Node)
	postfix = func(n Node) {
		if bin, isa := n.(*BinaryOperation); isa {
			children := bin.Children()
			// left
			postfix(children[0])
			// right
			postfix(children[2])
			// operator
			bldr.WriteString(children[1].String())
		} else {
			bldr.WriteString(n.String())
		}
	}
	postfix(root)
	
	return bldr.String()
}

func main() {
	exprs := []Node {
		// 1
		NewNumericNode(1),
		
		// 2 + 3
		NewBinaryOperation(
			NewNumericNode(2),
			Addition,
			NewNumericNode(3),
		),
		
		// ((1 * 2) / 3) + (4 - ((5 / 6) * 7))
		NewBinaryOperation(
			NewBinaryOperation(
				NewBinaryOperation(
					NewNumericNode(1),
					Multiplication,
					NewNumericNode(2),
				),
				Division,
				NewNumericNode(3),
			),
			Addition,
			NewBinaryOperation(
				NewNumericNode(4),
				Subtraction,
				NewBinaryOperation(
					NewBinaryOperation(
						NewNumericNode(5),
						Division,
						NewNumericNode(6),
					),
					Multiplication,
					NewNumericNode(7),
				),
			),
		),
	}
	
	for _, expr := range exprs {
		fmt.Println("Expr =",expr)
		fmt.Println("  Infix =", InfixTraversal(expr))
		fmt.Println("  Prefix =", PrefixTraversal(expr))
		fmt.Println("  Postfix =", PostfixTraversal(expr))
	}
}
