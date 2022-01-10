package sql

import "github.com/pingcap/parser/ast"

type Generator struct {
	TableName    string
	FieldNames   []string
	FieldTypeMap map[string]string
	InputFile    string
	OutputFile   string
}

func (g *Generator) Enter(in ast.Node) (ast.Node, bool) {

	if name, ok := in.(*ast.TableName); ok {
		g.TableName = name.Name.O
	}

	if name, ok := in.(*ast.ColumnName); ok {
		g.FieldNames = append(g.FieldNames, name.Name.O)
	}

	if columnDef, ok := in.(*ast.ColumnDef); ok {
		name := columnDef.Name.Name.O
		fieldType := columnDef.Tp.String()
		g.FieldTypeMap[name] = fieldType
	}

	return in, false
}

func (g *Generator) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
