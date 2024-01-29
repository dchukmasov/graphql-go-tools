package ast

import "github.com/dchukmasov/graphql-go-tools/pkg/lexer/position"

type SchemaExtension struct {
	ExtendLiteral position.Position
	SchemaDefinition
}
