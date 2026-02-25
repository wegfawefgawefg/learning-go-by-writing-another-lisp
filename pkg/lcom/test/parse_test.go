package lcom_test

import (
	"reflect"
	"testing"
	"trying-mo/pkg/lcom"
)

func TestParseNumber(t *testing.T) {
	input := lcom.Node{"NumberLiteral", "12345", []lcom.Node{}}
	result, err := parseNumber(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := lcom.ParseResult{
		nextPosition: 1,
		node: lcom.Node{
			typ:    "NumberLiteral",
			value:  "12345",
			params: []lcom.Node{},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseString(t *testing.T) {
	input := lcom.Node{"StringLiteral", "abc", []lcom.Node{}}
	result, err := parseString(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := lcom.ParseResult{
		nextPosition: 1,
		node: lcom.Node{
			typ:    "StringLiteral",
			value:  "abc",
			params: []lcom.Node{},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseExpression(t *testing.T) {

}

func TestParseToken(t *testing.T) {
	input := "12345bhjkhuil"
	result, err := tokenizePattern("number", "[0-9]+", input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := MToken{5, Token{"number", "12345"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
