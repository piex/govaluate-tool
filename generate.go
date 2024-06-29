package main

import (
	"fmt"
	"strings"
)

func generate(ast *ASTNode, indent int) string {
	indentStr := strings.Repeat("  ", indent)
	switch ast.Token.Kind {
	case LOGICALOP:
		code := fmt.Sprintf("%s\n%s%s\n%s%s %s",
			generate(ast.Children[0], indent+1),
			indentStr,
			ast.Token.Content,
			indentStr,
			generate(ast.Children[1], indent+1),
			indentStr,
		)
		if indent > 0 {
			return fmt.Sprintf("(\n%s%s\n%s)", indentStr, code, indentStr)
		}
		return code
	case COMPARATOR:
		return fmt.Sprintf("%s %s %s",
			generate(ast.Children[0], indent+1),
			ast.Token.Content,
			generate(ast.Children[1], indent+1),
		)
	case PREFIX:
		return fmt.Sprintf("%s(%s%s)",
			ast.Token.Content,
			generate(ast.Children[0], indent+1),
			indentStr,
		)
	case FUNCTION:
		params := []string{}
		for _, child := range ast.Children {
			params = append(params, generate(child, 0))
		}
		return fmt.Sprintf("%s( %s )",
			ast.Token.Content,
			strings.Join(params, ", "),
		)
	case VARIABLE, STRING, NUMERIC:
		return fmt.Sprintf("%v", ast.Token.Value)
	case ARRAY:
		elements := []string{}
		for _, child := range ast.Children {
			elements = append(elements, generate(child, indent+1))
		}
		return fmt.Sprintf("( %s )",
			strings.Join(elements, ", "),
		)
	default:
		return ""
	}
}