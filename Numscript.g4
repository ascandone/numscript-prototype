grammar Numscript;

// Tokens
WS: [ \t\r\n]+ -> skip;
NEWLINE: [\r\n]+;
MULTILINE_COMMENT: '/*' (MULTILINE_COMMENT | .)*? '*/' -> skip;
LINE_COMMENT: '//' .*? NEWLINE -> skip;

VARS: 'vars';
MAX: 'max';
SOURCE: 'source';
DESTINATION: 'destination';
SEND: 'send';
FROM: 'from';
TO: 'to';
LPARENS: '(';
RPARENS: ')';
LBRACKET: '[';
RBRACKET: ']';
LBRACE: '{';
RBRACE: '}';
EQ: '=';

RATIO_PORTION_LITERAL: [0-9]+ [ ]? '/' [ ]? [0-9]+;
PERCENTAGE_PORTION_LITERAL: [0-9]+ ('.' [0-9]+)? '%';

TYPE_IDENT: [a-z]+;
NUMBER: [0-9]+;
VARIABLE_NAME: '$' [a-z_]+ [a-z0-9_]*;
ACCOUNT: '@' [a-zA-Z0-9_-]+ (':' [a-zA-Z0-9_-]+)*;
ASSET: [A-Z/0-9]+;

portion:
	RATIO_PORTION_LITERAL			# ratio
	| PERCENTAGE_PORTION_LITERAL	# percentage;

varDeclaration: type_ = TYPE_IDENT name = VARIABLE_NAME;
varsDeclaration: VARS LBRACE varDeclaration* RBRACE;

program: varsDeclaration? statement*;

monetaryLit: LBRACKET (asset = ASSET) (amt = NUMBER) RBRACKET;

cap: monetaryLit # litCap | VARIABLE_NAME # varCap;

source:
	ACCOUNT								# srcAccount
	| VARIABLE_NAME						# srcVariable
	| LBRACE allotmentClauseSrc+ RBRACE	# srcAllotment
	| LBRACE source* RBRACE				# srcSeq
	| MAX cap FROM source				# srcCapped;
allotmentClauseSrc: portion FROM source;

destination:
	ACCOUNT									# destAccount
	| VARIABLE_NAME							# destVariable
	| LBRACE allotmentClauseDest+ RBRACE	# destAllotment
	| LBRACE destination* RBRACE			# destSeq;
allotmentClauseDest: portion TO destination;

statement:
	SEND monetaryLit LPARENS SOURCE EQ source DESTINATION EQ destination RPARENS;