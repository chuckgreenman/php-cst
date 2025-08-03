package main

import (
	"fmt"
	"reflect"
	"strings"

	"chuckgreenman.com/php-ast/src/node_types"
)

func printNode(node node_types.Node, depth int) {
	// Print indentation
	indent := strings.Repeat("  ", depth)
	
	// Print node type and location
	nodeType := reflect.TypeOf(node).Elem().Name()
	start := node.GetStart()
	end := node.GetEnd()
	fmt.Printf("%s%s (Line %d:%d -> %d:%d)\n", indent, nodeType, 
		start.Line, start.Column, end.Line, end.Column)
	
	// Recursively print children
	for _, child := range node.GetChildren() {
		printNode(child, depth+1)
	}
}

func main() {
	program := &node_types.Program{
		BaseNode: node_types.BaseNode{
			Start: node_types.Location{Line: 1, Column: 1},
			End:   node_types.Location{Line: 1, Column: 1},
		},
		Path: "example.php",
	}

	expr := &node_types.ExpressionStatement{
		BaseNode: node_types.BaseNode{
			Start: node_types.Location{Line: 1, Column: 1},
			End:   node_types.Location{Line: 1, Column: 10},
		},
		Expression: node_types.Expression{
			BaseNode: node_types.BaseNode{
				Start: node_types.Location{Line: 1, Column: 1},
				End:   node_types.Location{Line: 1, Column: 10},
			},
			Left: &node_types.Number{
				BaseNode: node_types.BaseNode{
					Start: node_types.Location{Line: 1, Column: 1},
					End:   node_types.Location{Line: 1, Column: 2},
				},
				Value: 1,
			},
			Right: &node_types.Number{
				BaseNode: node_types.BaseNode{
					Start: node_types.Location{Line: 1, Column: 9},
					End:   node_types.Location{Line: 1, Column: 10},
				},
				Value: 2,
			},
		},
	}

	program.Children = []node_types.Node{expr}
	printNode(program, 0)
}
