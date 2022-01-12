package main

import (
	"flag"
	_ "github.com/pingcap/parser/test_driver"
	"kit/sql"
	"log"
	"os"
	"strings"
)

func main() {
	inputSQL := flag.String("d", "", "DDL statement")
	//inputSQLFile := flag.String("sql-file", "", "sql file path")
	templateFile := flag.String("t", "", "repository template file")
	entityDir := flag.String("entity-output", "", "entity struct output dir")
	modelDir := flag.String("model-output", "", "model struct output dir")
	repoDir := flag.String("repo-output", "", "repository file output dir")
	flag.Parse()

	if inputSQL == nil {
		log.Fatalf("no sql input")
	}

	astNode, err := sql.Parse(*inputSQL)
	if err != nil {
		log.Fatalf("sql parse error[%v]", err)
	}

	// init generator
	gen := &sql.Generator{
		FieldTypeMap:  make(map[string]string),
		OutputEntity:  trimDirPath(entityDir),
		OutputModel:   trimDirPath(modelDir),
		OutputRepo:    trimDirPath(repoDir),
		InputTemplate: *templateFile,
	}

	(*astNode).Accept(gen)

	err = gen.InitFieldType()
	if err != nil {
		log.Fatalf("convert struct field type error[%v]", err)
	}

	// generate code
	gen.Do()
}

func trimDirPath(path *string) string {
	return strings.TrimRight(*path, string(os.PathSeparator))
}
