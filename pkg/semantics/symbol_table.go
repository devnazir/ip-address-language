package semantics

import "github.com/devnazir/gosh-script/pkg/ast"

type SymbolInfo struct {
	Type       string
	IsFunction bool
	Parameters []ast.Identifier
	Value      interface{}
}

type SymbolTable struct {
	scopes []map[string]SymbolInfo
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{scopes: []map[string]SymbolInfo{{}}}
}

func (st *SymbolTable) Insert(name string, info SymbolInfo) {
	currentScope := st.scopes[len(st.scopes)-1]
	currentScope[name] = info
}

func (st *SymbolTable) Exists(name string) bool {
	currentScope := st.scopes[len(st.scopes)-1]
	_, exists := currentScope[name]
	return exists
}

func (st *SymbolTable) Get(name string) (SymbolInfo, bool) {
	for i := len(st.scopes) - 1; i >= 0; i-- {
		if info, exists := st.scopes[i][name]; exists {
			return info, true
		}
	}
	return SymbolInfo{}, false
}

func (st *SymbolTable) ExistsInAnyScope(name string) bool {
	for i := len(st.scopes) - 1; i >= 0; i-- {
		if _, exists := st.scopes[i][name]; exists {
			return true
		}
	}
	return false
}

func (st *SymbolTable) PushScope() {
	st.scopes = append(st.scopes, make(map[string]SymbolInfo))
}

func (st *SymbolTable) PopScope() {
	if len(st.scopes) > 1 {
		st.scopes = st.scopes[:len(st.scopes)-1]
	}
}
