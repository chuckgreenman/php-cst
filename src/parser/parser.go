package parser

import (
	"chuckgreenman.com/php-ast/src/node_types"
)

type Parser struct {
	input string
}

func (p *Parser) ParseProgram() *node_types.Program {
	return &node_types.Program{
		BaseNode: node_types.BaseNode{
			Start: node_types.Location{Line: 1, Column: 1},
			End:   node_types.Location{Line: 1, Column: 1},
		},
		Path: "example.php",
	}
}

func NewParser(input string) *Parser {
	return &Parser{
		input: input,
	}
}
