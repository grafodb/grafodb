grammar GrafoQL;

document
    : (query | mutation) EOF
    ;

query
    : 'query' schemaBody
    ;

schemaBody
    :   '{' .*? '}'
    ;

mutation
    : 'mutation' schemaBody
    ;

mutationBody
    :   '{' .*? '}'
    ;

// Letters and digits

fragment ULetter
    : [A-Z]
    ;

fragment LLetter
    : [a-z_]
    ;

fragment Letter
    : [A-Za-z_]
    ;

fragment Digit
    : [0-9]
    ;

TypeIdentArray
    : LBRACK TypeIdent RBRACK
    ;

TypeIdent
    : ULetter (Letter | Digit)*
    ;

FieldIdent
    : LLetter (Letter | Digit)*
    ;


// Separators

//LBRACE  : '{';
//RBRACE  : '}';
LBRACK  : '[';
RBRACK  : ']';
//COMMA   : ',';
//DOT     : '.';
//SEMI    : ';';


// Whitespace and comments

WS
    : [ \t\u000C]+ -> skip
    ;

NEWLINE
    : ('\r'? '\n' | '\r')+ -> skip
    ;

COMMENT
    : '/*' .*? '*/' -> skip
    ;

LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;