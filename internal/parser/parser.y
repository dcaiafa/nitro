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
%token kAND 
%token kELIF 
%token kELSE 
%token kEND 
%token kFALSE 
%token kIF 
%token kIN 
%token kNOT 
%token kOR 
%token kTHEN
%token kTRUE 

%token <num> NUMBER
%token <str> STRING
%token <str> ID

%type <ast> assignment_stmt
%type <ast> expr
%type <ast> if_stmt
%type <ast> lvalue
%type <ast> number
%type <ast> object_elif
%type <ast> object_elifs
%type <ast> object_else
%type <ast> object_field
%type <ast> object_fields
%type <ast> object_if
%type <ast> object_literal 
%type <ast> object_literal_stmt
%type <ast> opt_object_elifs
%type <ast> opt_object_else
%type <ast> opt_object_fields
%type <ast> stmt
%type <ast> stmts
%type <ast> term
%type <ast> term_no_number
%type <ast> unary_expr

%left OR kOR
%left AND kAND
%nonassoc kIN
%nonassoc '<' LE '>' GE EQ NE
%left '+' '-'
%left '*' '/'

%start program

%%

program: stmts                            { }

stmts: stmts stmt                         { $$ = nil }
     | stmt                               { $$ = nil }

opt_comma: ','
         |

stmt: object_literal_stmt                 { $$ = nil }
    | assignment_stmt                     { $$ = nil }
    | if_stmt                             { $$ = nil }

if_stmt: kIF expr kTHEN
           stmts
         kEND
         { $$ = nil }

object_literal_stmt: object_literal       { $$ = nil }

assignment_stmt: lvalue '=' expr          { $$ = nil }

expr: unary_expr                          { $$ = nil }
    | expr '+' expr                       { $$ = nil }
    | expr '-' expr                       { $$ = nil }
    | expr '*' expr                       { $$ = nil }
    | expr '/' expr                       { $$ = nil }
    | expr '<' expr                       { $$ = nil }
    | expr LE expr                        { $$ = nil }
    | expr '>' expr                       { $$ = nil }
    | expr GE expr                        { $$ = nil }
    | expr EQ expr                        { $$ = nil }
    | expr NE expr                        { $$ = nil }
    | expr kAND expr                      { $$ = nil }
    | expr kOR expr                       { $$ = nil }

lvalue: term_no_number '.' ID             { $$ = nil }
      | ID                                { $$ = nil }

unary_expr: kNOT term                     { $$ = nil }
          | term                          { $$ = nil }

term: number                              { $$ = nil }
    | term_no_number                      { $$ = nil }

term_no_number: object_literal       { $$ = nil }
              | array_literal        {}
              | STRING               { $$ = nil }
              | lvalue               { $$ = nil }
              | '(' expr ')'         { $$ = nil }

number: '-' NUMBER                        { $$ = nil }
      | NUMBER                            { $$ = nil }

object_literal: '{' opt_object_fields '}' { $$ = nil }

opt_object_fields: object_fields            { $$ = nil }
                 |                          { $$ = nil }

object_fields: object_fields object_field   { $$ = nil }
             | object_field                 { $$ = nil }

object_field: ID ':' expr opt_comma { $$ = nil }
            | '[' expr ']' ':' expr opt_comma { $$ = nil }
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

array_literal: '[' opt_array_elements ']' {}

opt_array_elements: array_elements {}
                  | {}

array_elements: array_elements array_element {}
              | array_element {}

array_element: expr ',' {}
             | array_if {}

array_if: kIF expr kTHEN
            opt_array_elements
          kEND
          {}
