package semantics

import "github.com/devnazir/gosh-script/pkg/ast"

type SymbolInfo struct {
	Type       string
	IsFunction bool
	Parameters []ast.Identifier
	Value      interface{}
}

type SymbolTable struct {
	Scopes []map[string]SymbolInfo
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{Scopes: []map[string]SymbolInfo{{}}}
}

func (st *SymbolTable) Insert(name string, info SymbolInfo) {
	currentScope := st.Scopes[len(st.Scopes)-1]
	currentScope[name] = info
}

func (st *SymbolTable) Exists(name string) bool {
	currentScope := st.Scopes[len(st.Scopes)-1]
	_, exists := currentScope[name]
	return exists
}

func (st *SymbolTable) Get(name string) (SymbolInfo, bool) {
	for i := len(st.Scopes) - 1; i >= 0; i-- {
		if info, exists := st.Scopes[i][name]; exists {
			return info, true
		}
	}
	return SymbolInfo{}, false
}

func (st *SymbolTable) ExistsInAnyScope(name string) bool {
	for i := len(st.Scopes) - 1; i >= 0; i-- {
		if _, exists := st.Scopes[i][name]; exists {
			return true
		}
	}
	return false
}

func (st *SymbolTable) PushScope() {
	st.Scopes = append(st.Scopes, make(map[string]SymbolInfo))
}

func (st *SymbolTable) PopScope() {
	if len(st.Scopes) > 1 {
		st.Scopes = st.Scopes[:len(st.Scopes)-1]
	}
}
