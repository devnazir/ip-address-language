package ast

const (
	SourceDeclarationTree = "SourceDeclaration"
	SourceTree            = "Source"
)

type SourceDeclaration struct {
	BaseNode
	Sources []Source
}

type Source struct {
	StringLiteral
	Alias string
}
