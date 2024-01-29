package unsafeprinter

import (
	"github.com/dchukmasov/graphql-go-tools/v2/internal/pkg/unsafeparser"
	"github.com/dchukmasov/graphql-go-tools/v2/pkg/ast"
	"github.com/dchukmasov/graphql-go-tools/v2/pkg/astprinter"
)

func Print(document, definition *ast.Document) string {
	str, err := astprinter.PrintString(document, definition)
	if err != nil {
		panic(err)
	}
	return str
}

func PrettyPrint(document, definition *ast.Document) string {
	str, err := astprinter.PrintStringIndent(document, definition, "  ")
	if err != nil {
		panic(err)
	}
	return str
}

func Prettify(document string) string {
	doc := unsafeparser.ParseGraphqlDocumentString(document)
	return PrettyPrint(&doc, nil)
}
