grammar Sprout;

// ============ Lexer Rules ============

LET: 'let';
IN: 'in';
IF: 'if';
THEN: 'then';
ELSE: 'else';
TRUE: 'true';
FALSE: 'false';
FUN: 'fun';
ARROW: '->';
EQUALS: '=';

INT: [0-9]+;
IDENT: [A-Za-z_][0-9A-Za-z_]*;

WS: [ \t\r\n]+ -> skip;

// ============ Parser Rules ============

expr: letExpr | ifExpr | funExpr | appExpr;
letExpr: LET IDENT EQUALS expr IN expr;
ifExpr: IF expr THEN expr ELSE expr;
funExpr: FUN '(' IDENT ':' IDENT ')' ARROW expr;
appExpr: primaryExpr | appExpr '(' expr ')';
primaryExpr: INT | TRUE | FALSE | IDENT | '(' expr ')';