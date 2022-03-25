package yapi

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

const (
	demo = `{
  "type": "object",
  "title": "empty object",
  "properties": {
    "code": {
      "type": "integer"
    },
    "msg": {
      "type": "string"
    },
    "data": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "id_list": {
            "type": "array",
            "items": {
              "type": "integer"
            }
          },
          "name": {
            "type": "string"
          },
          "data_classify": {
            "type": "number",
            "mock": {
              "mock": "1"
            },
            "description": "1，用来标注是级联数据，需要请求下一级，2为intent，不需要请求下一级"
          }
        },
        "required": [
          "id_list",
          "name",
          "data_classify"
        ]
      }
    }
  },
  "required": [
    "code",
    "msg",
    "data"
  ]
}`
)

type yApiObject struct {
	Title       string                `json:"title,omitempty"`
	YType       string                `json:"type"`
	Properties  map[string]yApiObject `json:"properties,omitempty"`
	Mock        string                `json:"mock,omitempty"`
	Description string                `json:"description,omitempty"`
}

func fieldType(ty string) string {
	return ""
}

func parseAST(srcCode string) *string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", srcCode, 0)
	if err != nil {
		log.Fatalf("src code[%s]", srcCode)
		return nil
	}

	ast.Inspect(f, func(n ast.Node) bool {
		// Called recursively.
		_ = ast.Print(fset, n)
		return true
	})

	return nil
}

type StructParser struct{}

func (p *StructParser) Visit(ast.Node) ast.Visitor {
	return nil
}
