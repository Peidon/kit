alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -o ../internal/valuate/ast -Dlanguage=Go valuate.g4