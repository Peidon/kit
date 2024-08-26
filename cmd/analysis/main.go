/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2024/8/26 10:15 AM
 * @Version: kit
 */

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	v := visitor{fset: token.NewFileSet()}

	for _, dir := range os.Args[1:] {
		if dir == "--" { // to be able to run this like "go run main.go -- /usr/src"
			continue
		}

		files, err := listGoFiles(dir)
		if err != nil {
			log.Fatalf("Scan directory %s: %s", dir, err)
		}

		for _, filePath := range files {
			f, pErr := parser.ParseFile(v.fset, filePath, nil, 0)
			if pErr != nil {
				log.Fatalf("Failed to parse file %s: %s", filePath, pErr)
			}

			ast.Walk(&v, f)
		}
	}
}

type visitor struct {
	fset *token.FileSet

	fName string
}

func (v *visitor) SetFuncName(name string) {
	v.fName = name
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	funcDecl, ok := node.(*ast.FuncDecl)
	if !ok {
		return v
	}

	//var buf bytes.Buffer
	//err := printer.Fprint(&buf, v.fset, node)
	//fmt.Printf("Begin analysis [%s] | %d | %v \n", funcDecl.Name.String(), len(funcDecl.Body.List), err)
	//
	//fmt.Println("function called list: ")
	name := funcDecl.Name.String()
	v.SetFuncName(name)
	expressions := v.scanStatements(funcDecl.Body.List)
	for _, expr := range expressions {
		callExpr, cal := expr.(*ast.CallExpr)
		if !cal {
			continue
		}
		expressions = v.handleExpression(callExpr)
		for _, e := range expressions {
			sub := v.handleExpression(e)

			// 这里最多扫描三层
			for _, b := range sub {
				_ = v.handleExpression(b)
			}
		}
	}
	//fmt.Printf("%s | %#v | %v\n", buf.String(), node, err)
	//println()
	return v
}

func (v *visitor) scanStatements(stats []ast.Stmt) []ast.Expr {
	expressions := make([]ast.Expr, 0)
	for _, stmt := range stats {

		switch block := stmt.(type) {
		case *ast.AssignStmt:
			expressions = append(expressions, block.Rhs...)

		case *ast.IfStmt:
			if block.Body != nil {
				expressions = append(expressions, v.scanStatements(block.Body.List)...)
			}

		case *ast.ReturnStmt:
			expressions = append(expressions, block.Results...)

		case *ast.ExprStmt:
			expressions = append(expressions, block.X)

		case *ast.RangeStmt:
			if block.Body != nil {
				expressions = append(expressions, v.scanStatements(block.Body.List)...)
			}

		case *ast.SwitchStmt:
			if block.Body != nil {
				expressions = append(expressions, v.scanStatements(block.Body.List)...)
			}

		case *ast.CaseClause:
			if block.Body != nil {
				expressions = append(expressions, v.scanStatements(block.Body)...)
			}

			for _, expr := range block.List {
				v.handleExpression(expr)
			}

		default:
			//fmt.Printf("Statement Type %T \n", stmt)
		}
	}

	return expressions
}

func (v *visitor) handleExpression(expr ast.Expr) []ast.Expr {

	var buf bytes.Buffer
	expressions := make([]ast.Expr, 0)

	switch callExpr := expr.(type) {
	case *ast.CallExpr:

		for _, e := range callExpr.Args {
			ca, ok := e.(*ast.CallExpr)
			if !ok {
				continue
			}

			expressions = append(expressions, ca)
		}

		err := printer.Fprint(&buf, v.fset, callExpr.Fun)
		if err != nil {
			fmt.Printf("print call function error %v", err)
			return nil
		}

		if strings.Contains(buf.String(), v.fName) {
			fmt.Printf("%v \n", buf.String())
		}

	default:
		//fmt.Printf("Expression Type %T", expr)
		return nil
	}

	return expressions
}

//func (v *visitor) Visit(node ast.Node) ast.Visitor {
//	funcDecl, ok := node.(*ast.FuncDecl)
//	if !ok {
//		return v
//	}
//
//	params := funcDecl.Type.Params.List
//	if len(params) != 2 { // [0] must be format (string), [1] must be args (...interface{})
//		return v
//	}
//
//	firstParamType, ok := params[0].Type.(*ast.Ident)
//	if !ok { // first param type isn't identificator so it can't be of type "string"
//		return v
//	}
//
//	if firstParamType.Name != "string" { // first param (format) type is not string
//		return v
//	}
//
//	secondParamType, ok := params[1].Type.(*ast.Ellipsis)
//	if !ok { // args are not ellipsis (...args)
//		return v
//	}
//
//	elementType, ok := secondParamType.Elt.(*ast.InterfaceType)
//	if !ok { // args are not of interface type, but we need interface{}
//		return v
//	}
//
//	if elementType.Methods != nil && len(elementType.Methods.List) != 0 {
//		return v // has >= 1 method in interface, but we need an empty interface "interface{}"
//	}
//
//	if strings.HasSuffix(funcDecl.Name.Name, "f") {
//		return v
//	}
//
//	fmt.Printf("%s: printf-like formatting function '%s' should be named '%sf'\n",
//		v.fset.Position(node.Pos()), funcDecl.Name.Name, funcDecl.Name.Name)
//	return v
//}

func listGoFiles(dir string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "vendor" {
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	return goFiles, err
}
