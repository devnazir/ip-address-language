package semantics

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type SymbolInfo struct {
	Kind           string
	Value          interface{}
	TypeAnnotation string
	Line           int
	Address        string
}

func (si SymbolInfo) GetLine() int {
	return si.Line
}

type SymbolTable struct {
	Scopes  []map[string]SymbolInfo
	Address map[string]SymbolInfo
}

func (st *SymbolTable) GetScopes() []map[string]SymbolInfo {
	scopes := make([]map[string]SymbolInfo, len(st.Scopes))
	return scopes
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Scopes:  []map[string]SymbolInfo{{}},
		Address: make(map[string]SymbolInfo),
	}
}

func (st *SymbolTable) Insert(name string, info SymbolInfo) {
	currentScope := st.Scopes[len(st.Scopes)-1]
	info.Address = st.MakeAddress(info)
	currentScope[name] = info
}

func (st *SymbolTable) Update(name string, info SymbolInfo) {
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

func (st *SymbolTable) MakeAddress(info SymbolInfo) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", info)))

	if _, exists := st.Address[fmt.Sprintf("%x", hash.Sum(nil))]; exists {
		return fmt.Sprintf("%x", hash.Sum(nil))
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func (st *SymbolTable) InsertAddress(address string, info SymbolInfo) {
	st.Address[address] = info
}
