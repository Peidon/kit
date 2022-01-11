package main

import (
	"colx/sql"
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/parser/test_driver"
)

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}

func main() {
	astNode, err := parse("Create Table test_tab (`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id'," +
		"`saas_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'saas id'," +
		"`region` varchar(10) NOT NULL DEFAULT '' COMMENT 'region')")
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	v := &sql.Generator{
		FieldTypeMap: make(map[string]string),
	}
	(*astNode).Accept(v)

	fmt.Printf("%v\n", *v)
}
