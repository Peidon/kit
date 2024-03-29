// package ast
// alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
grammar valuate;

// Keywords
TRUE :  'true';
FALSE : 'false';

NIL_LIT:    'nil';

// Punctuation
LP                     : '(';
RP                     : ')';
L_CURLY                : '{';
R_CURLY                : '}';
L_BRACKET              : '[';
R_BRACKET              : ']';
COMMA                  : ',';
COLON                  : ':';
DOT                    : '.';

// Logical
LOGICAL_OR             : '||';
LOGICAL_AND            : '&&';

// Relation operators
EQUALS                 : '==';
NOT_EQUALS             : '!=';
LESS                   : '<';
LESS_OR_EQUALS         : '<=';
GREATER                : '>';
GREATER_OR_EQUALS      : '>=';

// Arithmetic operators
DIV                    : '/';
PLUS                   : '+';
MINUS                  : '-';
STAR                   : '*';
MODULUS                : '%';

SKIP_
 : ( SPACES | LINE_JOINING ) -> skip
 ;
WHITESPACE
 : [ \r\n\t]+ -> skip
 ;

IDENTIFIER : LETTER (LETTER | DIGIT)* ;

/*
 *  parser Rules
 */
plan: expression EOF;

expression
    : primaryExpr
    | unaryExpr
    | expression ('*' | '/' | '%' ) expression
    | expression ('+' | '-' ) expression
    | expression ('==' | '!=' | '<' | '<=' | '>' | '>=') expression
    | expression '&&' expression
    | expression '||' expression
    ;

primaryExpr
    : operand
    | variate
    | primaryExpr ( DOT IDENTIFIER
                  | index
                  | arguments)
    ;

unaryExpr
    : ('-' | '!') expression
    ;

arguments
    : LP (expressionList)? RP
    ;

expressionList
    : expression (COMMA expression)*
    ;

operand
    : basicLit
    | LP expression RP
    ;

basicLit
    : STRING
    | INT
    | FLOAT_NUMBER
    | obj
    | arr
    | TRUE
    | FALSE
    | NIL_LIT
    ;


obj
    : L_CURLY pair (',' pair)* R_CURLY
    | '{' '}'
    ;

pair
    : STRING_LITERAL ':' basicLit
    ;


arr
    : '[' basicLit (',' basicLit)* ']'
    | '[' ']'
    ;

// variate literal : _abc ab_c {ab.c}
variate
    : IDENTIFIER
    | L_CURLY IDENTIFIER (DOT IDENTIFIER)* R_CURLY
    ;

index
    : '[' expression ']'
    ;

/*
 * lexer rules
 */
STRING_LITERAL
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;

/// bytesliteral   ::=  bytesprefix(shortbytes | longbytes)
/// bytesprefix    ::=  "b" | "B" | "br" | "Br" | "bR" | "BR" | "rb" | "rB" | "Rb" | "RB"
BYTES_LITERAL
 : ( [bB] | ( [bB] [rR] ) | ( [rR] [bB] ) ) ( SHORT_BYTES | LONG_BYTES )
 ;

INT
 : DECIMAL_LIT
 | '0'
 ;

DECIMAL_LIT            : NON_ZERO_DIGIT DIGIT*;

/// floatnumber   ::=  pointfloat | exponentfloat
FLOAT_NUMBER
 : POINT_FLOAT
 | EXPONENT_FLOAT
 ;

STRING
 : STRING_LITERAL
 | SHORT_BYTES
 | LONG_BYTES
 ;

/*
 * fragments
 */
fragment LETTER
 : [a-zA-Z]
 | '_'
 ;

/// nonzerodigit   ::=  "1"..."9"
fragment NON_ZERO_DIGIT
 : [1-9]
 ;

/// digit          ::=  "0"..."9"
fragment DIGIT
 : [0-9]
 ;

/// octdigit       ::=  "0"..."7"
fragment OCT_DIGIT
 : [0-7]
 ;

/// hexdigit       ::=  digit | "a"..."f" | "A"..."F"
fragment HEX_DIGIT
 : [0-9a-fA-F]
 ;

/// bindigit       ::=  "0" | "1"
fragment BIN_DIGIT
 : [01]
 ;

/// pointfloat    ::=  [intpart] fraction | intpart "."
fragment POINT_FLOAT
 : INT_PART? FRACTION
 | INT_PART '.'
 ;

/// exponentfloat ::=  (intpart | pointfloat) exponent
fragment EXPONENT_FLOAT
 : ( INT_PART | POINT_FLOAT ) EXPONENT
 ;

/// intpart       ::=  digit+
fragment INT_PART
 : DIGIT+
 ;

/// fraction      ::=  "." digit+
fragment FRACTION
 : '.' DIGIT+
 ;

/// exponent      ::=  ("e" | "E") ["+" | "-"] digit+
fragment EXPONENT
 : [eE] [+-]? DIGIT+
 ;

/// shortbytes     ::=  "'" shortbytesitem* "'" | '"' shortbytesitem* '"'
/// shortbytesitem ::=  shortbyteschar | bytesescapeseq
fragment SHORT_BYTES
 : '\'' ( SHORT_BYTES_CHAR_NO_SINGLE_QUOTE | BYTES_ESCAPE_SEQ )* '\''
 | '"' ( SHORT_BYTES_CHAR_NO_DOUBLE_QUOTE | BYTES_ESCAPE_SEQ )* '"'
 ;

/// longbytes      ::=  "'''" longbytesitem* "'''" | '"""' longbytesitem* '"""'
fragment LONG_BYTES
 : '\'\'\'' LONG_BYTES_ITEM*? '\'\'\''
 | '"""' LONG_BYTES_ITEM*? '"""'
 ;

/// longbytesitem  ::=  longbyteschar | bytesescapeseq
fragment LONG_BYTES_ITEM
 : LONG_BYTES_CHAR
 | BYTES_ESCAPE_SEQ
 ;

/// shortbyteschar ::=  <any ASCII character except "\" or newline or the quote>
fragment SHORT_BYTES_CHAR_NO_SINGLE_QUOTE
 : [\u0000-\u0009]
 | [\u000B-\u000C]
 | [\u000E-\u0026]
 | [\u0028-\u005B]
 | [\u005D-\u007F]
 ;

fragment SHORT_BYTES_CHAR_NO_DOUBLE_QUOTE
 : [\u0000-\u0009]
 | [\u000B-\u000C]
 | [\u000E-\u0021]
 | [\u0023-\u005B]
 | [\u005D-\u007F]
 ;

/// longbyteschar  ::=  <any ASCII character except "\">
fragment LONG_BYTES_CHAR
 : [\u0000-\u005B]
 | [\u005D-\u007F]
 ;

/// bytesescapeseq ::=  "\" <any ASCII character>
fragment BYTES_ESCAPE_SEQ
 : '\\' [\u0000-\u007F]
 ;

fragment SPACES
 : [ \t]+
 ;


fragment LINE_JOINING
 : '\\' SPACES? ( '\r'? '\n' | '\r' | '\f')
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