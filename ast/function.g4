// function.g4
grammar function;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
IDENTIFIER : ([a-zA-Z_] [a-zA-Z_0-9]*)('.'([a-zA-Z_] [a-zA-Z_0-9]*))* ;
DECIMAL    : [0-9]+ ( '.' [0-9]+ )? ;
NUMBER     : [0-9]+;
WHITESPACE : [ \r\n\t]+ -> skip;

// Rules
f : expression EOF;

expression
   : IDENTIFIER '(' expression (',' expression)* ')'    # Function
   | expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | DECIMAL                            # Decimal
   | IDENTIFIER                         # Identifier
   | STRING                             # String
   ;


STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;

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