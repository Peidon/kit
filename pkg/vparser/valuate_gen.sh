alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -o ../valuate/parser -visitor -no-listener -Dlanguage=Go ValuateLexer.g4
antlr4 -o ../valuate/parser -visitor -no-listener -Dlanguage=Go ValuateParser.g4
antlr4 -o . -visitor -no-listener -Dlanguage=Go ValuateLexer.g4
antlr4 -o . -visitor -no-listener -Dlanguage=Go ValuateParser.g4