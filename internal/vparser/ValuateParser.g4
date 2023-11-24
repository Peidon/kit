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
    | L_CURLY R_CURLY
    ;

pair
    : STRING COLON basicLit
    ;


arr
    : '[' basicLit (',' basicLit)* ']'
    | '[' ']'
    ;

// variate literal : _abc ab_c {ab.c}
variate
    : IDENTIFIER
    | VARKEY_OPEN VARKEY_CLOSE
    ;

index
    : '[' expression ']'
    ;