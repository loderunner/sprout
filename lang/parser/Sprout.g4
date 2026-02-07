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

INT: '0' | [1-9][0-9]*;
IDENT: [A-Za-z_][0-9A-Za-z_]*;
COMPOP: '==' | '!=' | '<=' | '>=' | '<' | '>';

WS: [ \t\r\n]+ -> skip;

// ============ Parser Rules ============

expr: letExpr | ifExpr | funExpr | compExpr;
letExpr: LET IDENT EQUALS expr IN expr;
ifExpr: IF expr THEN expr ELSE expr;
funExpr: FUN '(' IDENT (':' IDENT)? ')' ARROW expr;
compExpr: orExpr | orExpr COMPOP orExpr;
orExpr: andExpr | orExpr '||' andExpr;
andExpr: termExpr | andExpr '&&' termExpr;
termExpr: factorExpr | termExpr op = ('+' | '-') factorExpr;
factorExpr: unaryExpr | factorExpr op = ('*' | '/') unaryExpr;
unaryExpr: appExpr | op = ('-' | '!') unaryExpr;
appExpr: primaryExpr | appExpr '(' expr ')';
primaryExpr: INT | TRUE | FALSE | IDENT | '(' expr ')';