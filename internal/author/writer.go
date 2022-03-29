package author

import (
	"log"
	"os"
	"strings"
)

func ModelStructName(tableName string) string {
	j := strings.LastIndex(tableName, underscoreStr)
	if j != -1 {
		tableName = tableName[:j]
	}
	return CamelCaseName(tableName)
}

func PackageName(filePath string) string {
	end := strings.LastIndex(filePath, string(os.PathSeparator))
	return filePath[end+1:]

}

func CamelCaseName(tableName string) string {
	words := strings.Split(tableName, underscoreStr)
	var entityName string
	for i := range words {
		entityName += strings.Title(words[i])
	}
	return entityName
}

func ModelFileName(tableName string) string {
	j := strings.LastIndex(tableName, underscoreStr)
	if j != -1 {
		return tableName[:j]
	}
	return tableName
}

func StructFieldName(fieldName string) string {
	words := strings.Split(fieldName, underscoreStr)
	var field string
	for i := range words {
		field += strings.Title(words[i])
	}
	return field
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal("File close Failed.", err)
	}
}

func CreateFile(filePath string) *os.File {
	file, err := os.OpenFile(
		filePath,
		os.O_RDWR|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatalf("create %s error %v", filePath, err)
	}
	return file
}

func AppendFile(filePath string) *os.File {
	file, err := os.OpenFile(
		filePath,
		os.O_RDWR|os.O_TRUNC|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf("append %s %v", filePath, err)
	}
	return file
}

const (
	underscoreStr = "_"
)
