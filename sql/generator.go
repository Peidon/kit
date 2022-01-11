package sql

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/pingcap/parser/ast"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Generator struct {
	TableName     string
	FieldNames    []string
	FieldTypeMap  map[string]string
	InputSQL      string // sql file path
	InputTemplate string // template file path
	OutputRepo    string // repo layer dir path
	OutputModel   string // model file path
	OutputEntity  string // entity file path
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

func (g *Generator) initFieldTypeMap() error {
	if len(g.FieldTypeMap) == 0 {
		return errors.New("no field type map")
	}
	for k, v := range g.FieldTypeMap {
		g.FieldTypeMap[k] = getStructFieldType(v)
	}
	return nil
}

func getStructFieldType(define string) string {
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
		log.Fatalf("%s Entity File exists.", fileInfo.Name())
	}
	// create file and write code
	file := createFile(targetFile)
	defer closeFile(file)
	entityName := camelCaseName(tableName)
	fmt.Print(entityName)

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("package %s\n", packageName(g.OutputEntity)))
	buffer.WriteString(fmt.Sprintf("type %s struct {\n", entityName))
	var tail string
	for _, field := range g.FieldNames {
		if field == "id" {
			tail = fmt.Sprintf("gorm:\"column:%s;primary_key;AUTO_INCREMENT\"", field)
		} else {
			tail = fmt.Sprintf("gorm:\"column:%s\"", field)
		}
		buffer.WriteString(structFieldName(field))
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
	log.Printf("AUTO Generate SUCCESS ! Wrote %d bytes.\n", bytesWritten)
}

func (g *Generator) writeModelStruct() {
	tableName := g.TableName
	if len(g.OutputModel) == 0 || len(tableName) == 0 {
		return
	}
	fileName := modelFileName(tableName)
	// if model file exists, not create
	targetFile := fmt.Sprintf("%s/%s.go", g.OutputModel, fileName)
	_, err := os.Stat(targetFile)
	if !os.IsNotExist(err) {
		log.Fatal("Model File exists.")
	}
	// create file and write code
	file := createFile(targetFile)
	defer closeFile(file)

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("package %s\n", packageName(g.OutputModel)))
	buffer.WriteString(fmt.Sprintf("type %sModel struct {\n", modelStructName(tableName)))

	var tail string
	for _, field := range g.FieldNames {
		tail = fmt.Sprintf("json:\"%s\"", field)
		buffer.WriteString(structFieldName(field))
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

func modelStructName(tableName string) string {
	j := strings.LastIndex(tableName, underscoreStr)
	if j != -1 {
		tableName = tableName[:j]
	}
	return camelCaseName(tableName)
}

func packageName(filePath string) string {
	end := strings.LastIndex(filePath, string(os.PathSeparator))
	return filePath[end+1:]

}

func camelCaseName(tableName string) string {
	words := strings.Split(tableName, underscoreStr)
	var entityName string
	for i := range words {
		entityName += strings.Title(words[i])
	}
	return entityName
}

func modelFileName(tableName string) string {
	j := strings.LastIndex(tableName, underscoreStr)
	if j != -1 {
		return tableName[:j]
	}
	return tableName
}

func structFieldName(fieldName string) string {
	words := strings.Split(fieldName, underscoreStr)
	var field string
	for i := range words {
		field += strings.Title(words[i])
	}
	return field
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal("File close Failed.", err)
	}
}

func createFile(filePath string) *os.File {
	file, err := os.OpenFile(
		filePath,
		os.O_RDWR|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

const (
	underscoreStr = "_"
	placeHolder   = "XXX"
)

func (g *Generator) writeRepoFile(tableName string) {
	if len(g.InputTemplate) == 0 {
		return
	}
	// 1.open file
	dao, err := os.Open(g.InputTemplate)
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
	daoCode := strings.ReplaceAll(template, placeHolder, modelStructName(tableName))

	// 3.if dao file exists, not create
	targetFile := fmt.Sprintf("%s/%s.go", g.OutputRepo, tableName)
	_, err = os.Stat(targetFile)
	if !os.IsNotExist(err) {
		log.Fatal("Repository File exists.")
	}
	// 4.create and write code
	daoFile := createFile(targetFile)
	wrote, err := daoFile.WriteString(daoCode)
	if err != nil {
		log.Fatal("Write Repository File error.", zap.Error(err))
	}
	log.Printf("Wrote %d characters Successfully.", wrote)
}
