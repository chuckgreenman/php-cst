package tests

import (
	"reflect"

	"chuckgreenman.com/php-ast/src/node_types"
	"chuckgreenman.com/php-ast/src/parser"

	"testing"
)

func TestSimpleAdditionToTree(t *testing.T) {
	php_string := "1 + 2"

	parser := parser.NewParser(php_string)

	goldenProgram := &node_types.Program{
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

	goldenProgram.Children = []node_types.Node{expr}

	if !reflect.DeepEqual(parser.ParseProgram(), goldenProgram) {
		t.Errorf("ParseProgram() did not return the golden program")
	}
}
