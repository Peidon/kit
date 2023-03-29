// package ast
// alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
grammar variate;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
DOT: '.';
LP: '(';
RP: ')';
COMMA:',';

WHITESPACE : [ \r\n\t]+ -> skip;

// Rules
start: expression EOF;

expression
   : expression (DOT expression)+                       # Selector
   | IDENTIFIER LP expression (COMMA expression)* RP    # Function
   | expression op=(MUL|DIV) expression                 # MulDiv
   | expression op=(ADD|SUB) expression                 # AddSub
   | IDENTIFIER                                         # Identifier
   | NUMBER                                             # Number
   | STRING                                             # String
   ;

//PROPERTY: FUNCTION | IDENTIFIER;
//
//FUNCTION : IDENTIFIER LP PARAM (COMMA PARAM)* RP;
//
//PARAM
//   : IDENTIFIER
//   | STRING
//   | NUMBER
//   ;

IDENTIFIER : ([a-zA-Z_] [a-zA-Z_0-9]*) ;

NUMBER    : [0-9]+ ( DOT [0-9]+ )? ;

STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;

fragment BINARY: MUL|DIV|ADD|SUB;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;

fragment HEX
   : [0-9a-fA-F]
   ;