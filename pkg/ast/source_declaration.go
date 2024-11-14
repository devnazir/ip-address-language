package ast

type SourceDeclaration struct {
	BaseNode
	Sources []Source
}

type Source struct {
	StringLiteral
	Alias string
}
