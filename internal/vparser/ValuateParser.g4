parser grammar ValuateParser;

options
{
tokenVocab=ValuateLexer;
}

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
    : variate
    | operand
    | primaryExpr ( DOT IDENTIFIER
                  | index)
    | IDENTIFIER arguments
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

// variate literal : _abc ab_c ${ab.c}
variate
    : VARKEY_OPEN VARID VARKEY_CLOSE
    | IDENTIFIER
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
    | L_CURLY R_CURLY
    ;

pair
    : STRING COLON basicLit
    ;


arr
    : '[' basicLit (',' basicLit)* ']'
    | '[' ']'
    ;


index
    : '[' expression ']'
    ;