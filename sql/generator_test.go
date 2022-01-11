package sql

import (
	"testing"
)

func TestGenCode(t *testing.T) {
	astNode, err := parse("Create Table test_tab (`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id'," +
		"`saas_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'saas id'," +
		"`region` varchar(10) NOT NULL DEFAULT '' COMMENT 'region')")
	if err != nil {
		t.Errorf("parse error: %v\n", err.Error())
		return
	}

	// init generator
	gen := &Generator{
		FieldTypeMap: make(map[string]string),
		OutputEntity: "../tmp",
		OutputModel:  "../tmp",
		OutputRepo:   "",
	}
	(*astNode).Accept(gen)
	err = gen.initFieldTypeMap()
	if err != nil {
		t.Errorf("parse error: %v\n", err.Error())
		return
	}

	// generate code
	gen.writeEntityFile()
	gen.writeModelStruct()

}
