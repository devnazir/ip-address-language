package semantics

import (
	"fmt"
)

type ScopeResolver struct {
	symbolTable *SymbolTable
}

func NewScopeResolver(symbolTable *SymbolTable) *ScopeResolver {
	return &ScopeResolver{symbolTable: symbolTable}
}

func (sr *ScopeResolver) ResolveScope(name string) *SymbolInfo {
	info, exists := sr.symbolTable.Get(name)
	if !exists {
		panic(fmt.Sprintf("Variable %s is not defined", name))
	}
	return &info
}

func (sr *ScopeResolver) EnterScope() {
	sr.symbolTable.PushScope()
}

func (sr *ScopeResolver) ExitScope() {
	sr.symbolTable.PopScope()
}
