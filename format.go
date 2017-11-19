package pqlinq

import (
	"strings"
	"text/template"

	"github.com/lebenasa/pqprobe"
)

// Provides template struct.

type (
	// Data provided to pqlinq template.
	Data struct {
		PackageName string
		Table       Table
	}

	// Table extends pqprobe.Table with additional formatting functions for pqlinq template.
	Table struct {
		pqprobe.Table
	}
)

// NewData generates data for pqlinq template.
func NewData(packageName string, table pqprobe.Table) Data {
	return Data{
		PackageName: packageName,
		Table:       Table{Table: table},
	}
}

// GenFuncMap generate function map for pqlinq template.
func GenFuncMap() template.FuncMap {
	return template.FuncMap{
		"add":      Add,
		"sub":      Sub,
		"notAtEnd": NotAtEnd,
		"lower":    strings.ToLower,
	}
}

// Add = a + b
func Add(a, b int) int {
	return a + b
}

// Sub = a - b
func Sub(a, b int) int {
	return a - b
}

// NotAtEnd returns true when i < len-1.
func NotAtEnd(i, len int) bool {
	return i < len-1
}
