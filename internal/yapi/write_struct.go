package yapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kit/internal/author"
	"log"
	"os"
)

func structFieldType(ty string) string {
	return ""
}

func writeStruct(dirPath, jsonStr string) {
	obj := new(yApiObject)
	err := json.Unmarshal([]byte(jsonStr), obj)

	if err != nil {
		log.Fatalf("yapi json err[%v]", err)
	}

	log.Println(obj.Title)
	log.Println(obj.YType)
	log.Println(obj.Description)

	objTitle := obj.Title
	if len(dirPath) == 0 || len(objTitle) == 0 {
		return
	}
	fileName := author.ModelFileName(objTitle)
	// if model file exists, not create
	targetFile := fmt.Sprintf("%s/%s.go", dirPath, fileName)
	_, err = os.Stat(targetFile)
	var (
		file *os.File
	)
	if os.IsNotExist(err) {
		log.Println("struct File exists.")
		file = author.AppendFile(targetFile)
	} else {
		// create file and write code
		file = author.CreateFile(targetFile)
	}
	defer author.CloseFile(file)

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("package %s\n", author.PackageName(dirPath)))
	buffer.WriteString(fmt.Sprintf("type %sModel struct {\n", author.ModelStructName(objTitle)))

	var (
		tail string
	)
	for k, v := range obj.Properties {
		// field Name k
		// file type v.YType
		tail = fmt.Sprintf("json:\"%s\"", k)
		buffer.WriteString(author.StructFieldName(k))
		buffer.WriteString(" ")
		buffer.WriteString(structFieldType(v.YType))
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
