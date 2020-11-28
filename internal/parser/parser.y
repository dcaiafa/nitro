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

program: stmts                            {}

stmts: stmts stmt                         {}
     | stmt                               {}

comma_opt: ','
         | /*empty*/

stmt: object_literal_stmt                 {}
    | assignment_stmt                     {}
    | var_decl_stmt                       {}
    | if_stmt                             {}
    | for_stmt                            {}

if_stmt: kIF expr kTHEN
           stmts
         elifs_opt
         else_opt
         kEND
         {}

elifs_opt: elifs {}
         | /*empty*/ {}

elifs: elifs elif {}
     | elif {}

elif: kELIF expr kTHEN
        stmts
      {}

else_opt: else {}
        | /*empty*/ {}

else: kELSE expr 
        stmts
      {}

for_stmt: kFOR for_args kIN expr kDO
            stmts
          kEND
          {}

for_args: for_args ',' ID {}
        | ID {}

object_literal_stmt: object_literal       {}

assignment_stmt: symbol_ref '=' expr          {}

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

symbol_ref: term_no_number '.' ID             {}
          | ID                                {}

unary_expr: kNOT term                     {}
          | term                          {}

term: number                              {}
    | term_no_number                      {}

number: '-' NUMBER                        {}
      | NUMBER                            {}

term_no_number: STRING               {}
              | array_literal        {}
              | object_literal       {}
              | symbol_ref           {}
              | '(' expr ')'         {}

//invocation: term_no_number '(' ')'

object_literal: '{' object_fields_opt '}' {}

object_fields_opt: object_fields            {}
                 |                          {}

object_fields: object_fields object_field   {}
             | object_field                 {}

object_field: ID ':' expr comma_opt {}
            | '[' expr ']' ':' expr comma_opt {}
            | object_if comma_opt {}

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

array_elems_opt: array_elems expr_opt {}
               | expr {}
               | {}

array_elems: array_elems array_elem {}
           | array_elem {}

array_elem: expr ',' {}
          | array_if {}

array_if: kIF expr kTHEN
            array_elems_opt
          array_elifs_opt
          array_else_opt
          kEND
          {}

expr_opt: expr {}
        |  {}

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
