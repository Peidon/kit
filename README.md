# kit
some useful tools  

## code generator
### usage

1. go install
2. kit --help

Usage of generating code depending on ddl:
```cmd
  -ddl-file string
        ddl file path, is Required
  -entity-output string
        entity struct output dir
  -model-output string
        model struct output dir
  -repo-output string
        repository file output dir
  -repo-tpl string
        repository template file
```

example command:  
```cmd
kit -ddl-file /Users/peidongxu/kit/tmp/sql \  
    -entity-output /Users/peidongxu/home/kit/entity \  
    -model-output /Users/peidongxu/home/kit/model/tm \  
    -repo-tpl /Users/peidongxu/kit/tmp/tpl \  
    -repo-output /Users/peidongxu/home/
```