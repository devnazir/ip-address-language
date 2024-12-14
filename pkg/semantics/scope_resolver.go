package semantics

import (
	"fmt"

	"github.com/devnazir/ip-address-language/pkg/oops"
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
		panic(oops.SyntaxError(info, fmt.Sprintf("Undefined variable %s", name)))
	}
	return &info
}

func (sr *ScopeResolver) EnterScope() {
	sr.symbolTable.PushScope()
}

func (sr *ScopeResolver) ExitScope() {
	sr.symbolTable.PopScope()
}
