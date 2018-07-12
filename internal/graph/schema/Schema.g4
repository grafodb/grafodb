grammar Schema;

document
    : schema EOF
    ;

schema
    : 'schema' schemaBody
    ;

schemaBody
    :   '{' schemaTypes '}'
    ;

schemaTypes
   : typeDecl (typeDecl)*
   ;

typeDecl
    : 'type' typeName typeBody
    ;

typeName
    :   TypeIdent
    ;

typeBody
    :   '{' ( field )+ '}'
    ;

field
    : fieldName fieldType
    ;

fieldName
    :   FieldIdent
    ;

fieldType
    :   (   'ID'
        |   'int'
        |   'float'
        |   'string'
        |   'bool'
        |   'datetime'
        |   'geopoint'
        |   TypeIdent
        |   TypeIdentArray
        )
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