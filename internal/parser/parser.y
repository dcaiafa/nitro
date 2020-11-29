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
%token kVAR
%token kAND 
%token kDO
%token kELIF 
%token kELSE 
%token kEND 
%token kFALSE 
%token kFOR
%token kIF 
%token kIN 
%token kNOT 
%token kOR 
%token kTHEN
%token kTRUE 

%token <num> NUMBER
%token <str> STRING
%token <str> ID

%left OR kOR
%left AND kAND
%nonassoc kIN
%nonassoc '<' LE '>' GE EQ NE
%left '+' '-'
%left '*' '/'

%start program

%%

program: opt_stmts                        {}

opt_stmts: stmts                          {}
         | /*empty*/                      {}

stmts: stmts stmt ';'                     {}
     | stmt ';'                           {}

stmt: assignment_stmt                     {}
    | var_decl_stmt                       {}
    | expr                                {}
    | for_stmt                            {}
    | if_expr                             {}

if_expr: kIF expr kTHEN
           opt_stmts
         elifs_opt
         else_opt
         kEND
         {}

elifs_opt: elifs {}
         | /*empty*/ {}

elifs: elifs elif {}
     | elif {}

elif: kELIF expr kTHEN
        opt_stmts
      {}

else_opt: else {}
        | /*empty*/ {}

else: kELSE
        opt_stmts
      {}

for_stmt: kFOR for_args kIN expr kDO
            opt_stmts
          kEND
          {}

for_args: for_args ',' ID {}
        | ID {}

assignment_stmt: expr '=' expr          {}

var_decl_stmt: kVAR ID                    {}
             | kVAR ID '=' expr           {}

expr: unary_expr                          {}
    | expr '+' expr                       {}
    | expr '-' expr                       {}
    | expr '*' expr                       {}
    | expr '/' expr                       {}
    | expr '<' expr                       {}
    | expr LE expr                        {}
    | expr '>' expr                       {}
    | expr GE expr                        {}
    | expr EQ expr                        {}
    | expr NE expr                        {}
    | expr kAND expr                      {}
    | expr kOR expr                       {}

unary_expr: kNOT unary_expr               {}
        | '+' unary_expr                  {}
        | '-' unary_expr                  {}
        | expr3                           {}

expr3: STRING           {}
     | NUMBER           {}
     | ID               {}
     | kTRUE            {}
     | kFALSE           {}
     | array_literal    {}
     | object_literal   {}
     | expr3 '[' expr ']' {}
     | expr3 '.' ID     {}
     | expr3 '(' ')'    {}
     | '(' expr ')'     {}

object_literal: '{' object_fields_opt '}' {}

object_fields_opt: object_fields comma_or_semicollon_opt {}
                 |               {}

object_fields: object_fields ',' object_field   {}
             | object_fields ';' object_field   {} 
             | object_field                     {}

object_field: ID ':' expr {}
            | '[' expr ']' ':' expr {}
            | object_if {}

object_if: kIF expr kTHEN
             object_fields_opt
           object_elifs_opt
           object_else_opt
           kEND
           {}

object_elifs_opt: object_elifs {}
                |             {}

object_elifs: object_elifs object_elif  {}
            | object_elif               {}

object_elif: kELIF expr kTHEN
               object_fields_opt
             {}

object_else_opt: object_else {}
               |             {}

object_else: kELSE 
               object_fields_opt
             {}

array_literal: '[' array_elems_opt ']' {}

array_elems_opt: array_elems comma_or_semicollon_opt {}
               | {}

array_elems: array_elems ',' array_elem         {}
           | array_elems ';' array_elem         {} 
           | array_elem                         {}

array_elem: expr     {}
          | array_if {}

array_if: kIF expr kTHEN
            array_elems_opt
          array_elifs_opt
          array_else_opt
          kEND
          {}

array_elifs_opt: array_elifs {}
               | {}

array_elifs: array_elifs array_elif
           | array_elif

array_elif: kELIF expr kTHEN
              array_elems_opt
            {}

array_else_opt: array_else {}
              | {}

array_else: kELSE
              array_elems_opt
            {}

comma_or_semicollon_opt: ';' | ',' 
