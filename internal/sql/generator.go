package sql

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/pingcap/parser/ast"
	"go.uber.org/zap"
	"io/ioutil"
	"kit/internal/author"
	"kit/pkg/slice"
	"log"
	"os"
	"strings"
)

type Generator struct {
	TableName    string
	FieldNames   []string
	FieldTypeMap map[string]string
	DDLFile      string // ddl file path
	RepoTemplate string // template file path
	OutputRepo   string // repo layer dir path
	OutputModel  string // model file path
	OutputEntity string // entity file path
}

func (g *Generator) Enter(in ast.Node) (ast.Node, bool) {

	if name, ok := in.(*ast.TableName); ok {
		g.TableName = name.Name.O
	}

	if name, ok := in.(*ast.ColumnName); ok {
		if !slice.Contains(g.FieldNames, name.Name.O) {
			g.FieldNames = append(g.FieldNames, name.Name.O)
		}
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

func (g *Generator) InitFieldType() error {
	if len(g.FieldTypeMap) == 0 {
		return errors.New("no field type map")
	}
	for k, v := range g.FieldTypeMap {
		g.FieldTypeMap[k] = dbToStructType(v)
	}
	return nil
}

func dbToStructType(define string) string {
	i := strings.Index(define, "(")
	if i == -1 {
		i = len(define)
	}
	switch strings.ToLower(define[:i]) {
	case "tinyint":
		return "int8"
	case "int":
		return "int32"
	case "bigint":
		return "int64"
	case "varchar":
		return "string"
	case "text":
		return "string"
	case "json":
		return "[]byte"
	default:
		return ""
	}
}

func (g *Generator) writeEntityFile() {
	tableName := g.TableName
	if len(g.OutputEntity) == 0 || len(tableName) == 0 {
		return
	}
	// if file exists, not create
	targetFile := fmt.Sprintf("%s/%s.go", g.OutputEntity, tableName)
	fileInfo, err := os.Stat(targetFile)
	if !os.IsNotExist(err) {
		log.Printf("%s Entity File exists.\n", fileInfo.Name())
		return
	}
	// create file and write code
	file := author.CreateFile(targetFile)
	defer author.CloseFile(file)
	entityName := author.CamelCaseName(tableName)
	fmt.Print(entityName)

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("package %s\n", author.PackageName(g.OutputEntity)))
	buffer.WriteString(fmt.Sprintf("type %s struct {\n", entityName))
	var tail string
	for _, field := range g.FieldNames {
		if field == "id" {
			tail = fmt.Sprintf("gorm:\"column:%s;primary_key;AUTO_INCREMENT\"", field)
		} else {
			tail = fmt.Sprintf("gorm:\"column:%s\"", field)
		}
		buffer.WriteString(author.StructFieldName(field))
		buffer.WriteString(" ")
		buffer.WriteString(g.FieldTypeMap[field])
		buffer.WriteString(fmt.Sprintf(" `%s`\n", tail))
	}
	buffer.WriteString("}\n")

	buffer.WriteString(fmt.Sprintf("func (*%s) TableName() string {\n", entityName))
	buffer.WriteString(fmt.Sprintf("return %q\n}", tableName))
	byteSlice := buffer.Bytes()
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AUTO Generate entity SUCCESS ! Wrote %d bytes.\n", bytesWritten)
}

func (g *Generator) writeModelStruct() {
	tableName := g.TableName
	if len(g.OutputModel) == 0 || len(tableName) == 0 {
		return
	}
	fileName := author.ModelFileName(tableName)
	// if model file exists, not create
	targetFile := fmt.Sprintf("%s/%s.go", g.OutputModel, fileName)
	_, err := os.Stat(targetFile)
	if !os.IsNotExist(err) {
		log.Println("Model File exists.")
		return
	}
	// create file and write code
	file := author.CreateFile(targetFile)
	defer author.CloseFile(file)

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("package %s\n", author.PackageName(g.OutputModel)))
	buffer.WriteString(fmt.Sprintf("type %sModel struct {\n", author.ModelStructName(tableName)))

	var tail string
	for _, field := range g.FieldNames {
		tail = fmt.Sprintf("json:\"%s\"", field)
		buffer.WriteString(author.StructFieldName(field))
		buffer.WriteString(" ")
		buffer.WriteString(g.FieldTypeMap[field])
		buffer.WriteString(fmt.Sprintf(" `%s`\n", tail))
	}
	buffer.WriteString("}\n")

	byteSlice := buffer.Bytes()
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes into model file.\n", bytesWritten)
}

func (g *Generator) writeRepoFile() {
	if len(g.RepoTemplate) == 0 || len(g.OutputRepo) == 0 {
		return
	}
	// 1.open file
	dao, err := os.Open(g.RepoTemplate)
	defer func() {
		err = dao.Close()
		if err != nil {
			log.Fatal("File close Failed.", err)
		}
	}()
	// 2.read code
	data, err := ioutil.ReadAll(dao)
	if err != nil {
		log.Fatal("Write Repository file error.", err)
	}
	template := string(data)
	daoCode := strings.ReplaceAll(template, placeHolder, author.ModelStructName(g.TableName))

	// 3.if dao file exists, not create
	targetFile := fmt.Sprintf("%s/%s.go", g.OutputRepo, g.TableName)
	_, err = os.Stat(targetFile)
	if !os.IsNotExist(err) {
		log.Println("Repository File exists.")
		return
	}
	// 4.create and write code
	log.Printf("create repo %s file\n", targetFile)
	daoFile := author.CreateFile(targetFile)
	wrote, err := daoFile.WriteString(daoCode)
	if err != nil {
		log.Fatal("Write Repository File error.", zap.Error(err))
	}
	log.Printf("Wrote %d characters into repo file Successfully.", wrote)
}

func (g *Generator) Do() {
	sql := readSQL(g.DDLFile)
	if len(sql) == 0 {
		log.Fatalln("sql is null")
		return
	}

	astNode, err := Parse(string(sql))
	if err != nil {
		log.Fatalf("sql parse error[%v]", err)
	}
	(*astNode).Accept(g)

	err = g.InitFieldType()
	if err != nil {
		log.Fatalf("convert struct field type error[%v]", err)
	}

	g.writeEntityFile()
	g.writeModelStruct()
	g.writeRepoFile()
}

func readSQL(filePath string) []byte {
	// 1.open file
	ddl, err := os.Open(filePath)
	defer func() {
		err = ddl.Close()
		if err != nil {
			log.Fatal("File close Failed.", err)
		}
	}()
	// 2.read code
	data, er := ioutil.ReadAll(ddl)
	if er != nil {
		log.Fatal("Read SQL file error.", er)
	}
	return data
}
