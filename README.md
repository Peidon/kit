# kit
some useful tools  

## code generator
### usage

1. go install
2. kit --help

example:  
kit -d "Create Table test_tab (id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',saas_id bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'saas id',region varchar(10) NOT NULL DEFAULT '' COMMENT 'region')" -entity-output /Users/peidongxu/kit/tmp
