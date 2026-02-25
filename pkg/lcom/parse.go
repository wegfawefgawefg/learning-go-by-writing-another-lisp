package lcom

import "fmt"

type ParseResult struct {
	nextPosition int
	node         Node
}

type Node struct {
	typ    string
	value  string
	params []Node
}

func parseNumber(tokens []Token, current int) (ParseResult, error) {
	return ParseResult{
		nextPosition: current + 1,
		node: Node{
			typ:    "NumberLiteral",
			value:  tokens[current].value,
			params: []Node{},
		},
	}, nil
}

func parseString(tokens []Token, current int) (ParseResult, error) {
	return ParseResult{
		nextPosition: current + 1,
		node: Node{
			typ:    "StringLiteral",
			value:  tokens[current].value,
			params: []Node{},
		},
	}, nil
}

func parseExpression(tokens []Token, current int) (ParseResult, error) {
	// steps:
	// skip opening parens
	// create base node with type CallExpression, and name from current token
	// recursively call parseToken until encountering a closing parens
	// skip the last token - the closing parens

	// skip opening parens
	current++
	token := tokens[current]
	node := Node{
		typ:    "CallExpression",
		value:  token.value,
		params: []Node{},
	}
	token = tokens[current+1]

	for !(token.typ == "paren" && token.value == ")") {
		// recursively call parseToken
		result, err := parseToken(tokens, current)
		if err != nil {
			return ParseResult{}, err
		}
		node.params = append(node.params, result.node)
		current = result.nextPosition
		token = tokens[current]
	}

	current++
	return ParseResult{
		nextPosition: current,
		node:         node,
	}, nil
}

func parseToken(tokens []Token, current int) (ParseResult, error) {
	token := tokens[current]

	switch token.typ {
	case "number":
		return parseNumber(tokens, current)
	case "string":
		return parseString(tokens, current)
	case "paren":
		if token.value == "(" {
			return parseExpression(tokens, current)
		}
		fallthrough
	default:
		return ParseResult{}, fmt.Errorf("unknown token type: %s", token.typ)
	}
}
