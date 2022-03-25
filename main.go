package main

import (
	"flag"
	_ "github.com/pingcap/parser/test_driver"
	"kit/sql"
	"os"
	"strings"
)

func main() {
	ddl := flag.String("ddl-file", "", "ddl file path")
	//inputSQLFile := flag.String("sql-file", "", "sql file path")
	templateFile := flag.String("repo-tpl", "", "repository template file")
	entityDir := flag.String("entity-output", "", "entity struct output dir")
	modelDir := flag.String("model-output", "", "model struct output dir")
	repoDir := flag.String("repo-output", "", "repository file output dir")
	flag.Parse()

	if ddl != nil && len(*ddl) != 0 {

		// init generator
		gen := &sql.Generator{
			FieldTypeMap: make(map[string]string),
			OutputEntity: trimDirPath(entityDir),
			OutputModel:  trimDirPath(modelDir),
			OutputRepo:   trimDirPath(repoDir),
			RepoTemplate: *templateFile,
			DDLFile:      *ddl,
		}

		// generate code
		gen.Do()
	}

}

func trimDirPath(path *string) string {
	return strings.TrimRight(*path, string(os.PathSeparator))
}
