%{
package parser

import (
)

%}

%union {
  num float64
  str string
  ast interface{}
  expr interface{}
}

%token LEXERR
%token ID kTRUE kFALSE kIN kAND kOR kNOT kIF kELIF kELSE kEND kTHEN
%token BLOCK_OPEN BLOCK_CLOSE 
%token <num> NUMBER
%token <str> STRING
%token <str> ID

%type <ast> if_stmt
%type <ast> assignment_stmt
%type <ast> object_literal_stmt
%type <ast> object_field
%type <ast> object_fields
%type <ast> object_if
%type <ast> opt_object_else
%type <ast> object_else
%type <ast> opt_object_fields
%type <ast> opt_object_elifs
%type <ast> object_elifs
%type <ast> object_elif
%type <ast> stmt
%type <ast> stmts
%type <expr> expr
%type <expr> number
%type <expr> object_literal 
%type <expr> ref_expr
%type <expr> term_expr
%type <expr> unary_expr

%left OR kOR
%left AND kAND
%nonassoc kIN
%nonassoc '<' LE '>' GE EQ NE
%left '+' '-'
%left '*' '/'

%start program

%%

program: stmts                            { }

stmts: stmts stmt                     { $$ = nil }
     | stmt { $$ = nil }

opt_comma: ','
         |

stmt: object_literal_stmt                 { $$ = nil }
    | assignment_stmt ';'                 { $$ = nil }
    | if_stmt                             { $$ = nil }

if_stmt: kIF expr kTHEN
           stmts
         kEND
         { $$ = nil }

object_literal_stmt: object_literal       { $$ = nil }

assignment_stmt: ref_expr '=' expr        { $$ = nil }

expr: unary_expr                          { $$ = nil }
    | expr '+' expr                       { $$ = nil }
    | expr '-' expr                       { $$ = nil }

ref_expr: term_expr '.' ID                { $$ = nil }
        | ID                              { $$ = nil }

unary_expr: kNOT term_expr                { $$ = nil }
          | term_expr                     { $$ = nil }

term_expr: number                         { $$ = nil }
         | STRING                         { $$ = nil }
         | ref_expr                       { $$ = nil }
         | object_literal                 { $$ = nil }
         | '(' expr ')'                   { $$ = nil }

number: '-' NUMBER                        { $$ = nil }
      | NUMBER                            { $$ = nil }

object_literal: '{' opt_object_fields '}' { $$ = nil }

opt_object_fields: object_fields            { $$ = nil }
                 |                          { $$ = nil }

object_fields: object_fields object_field   { $$ = nil }
             | object_field                 { $$ = nil }

object_field: ID ':' expr opt_comma { $$ = nil }
            | object_if             { $$ = nil }

object_if: kIF expr kTHEN
           opt_object_fields
           opt_object_elifs
           opt_object_else
           kEND
           { $$ = nil }

opt_object_elifs: object_elifs { $$ = nil }
                |             { $$ = nil }

object_elifs: object_elifs object_elif  { $$ = nil }
            | object_elif               { $$ = nil }

object_elif: kELIF expr kTHEN
             opt_object_fields
             { $$ = nil }

opt_object_else: object_else { $$ = nil }
               |             { $$ = nil }

object_else: kELSE 
             opt_object_fields
             { $$ = nil }
