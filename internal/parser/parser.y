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
%type <ast> object_elifs_opt
%type <ast> object_else_opt
%type <ast> object_fields_opt
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

stmts: stmts stmt                         {}
     | stmt                               {}

comma_opt: ','
         |

stmt: object_literal_stmt                 {}
    | assignment_stmt                     {}
    | if_stmt                             {}

if_stmt: kIF expr kTHEN
           stmts
         kEND
         {}

object_literal_stmt: object_literal       {}

assignment_stmt: lvalue '=' expr          {}

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

lvalue: term_no_number '.' ID             {}
      | ID                                {}

unary_expr: kNOT term                     {}
          | term                          {}

term: number                              {}
    | term_no_number                      {}

term_no_number: object_literal       {}
              | array_literal        {}
              | STRING               {}
              | lvalue               {}
              | '(' expr ')'         {}

number: '-' NUMBER                        {}
      | NUMBER                            {}

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

array_literal: '[' array_elements_opt ']' {}

array_elements_opt: array_elements array_last_elem_opt {}
                  | {}

array_elements: array_elements array_element {}
              | array_element {}

array_element: expr ',' {}
             | array_if {}

array_if: kIF expr kTHEN
            array_elements_opt
          kEND
          {}

array_last_elem_opt: expr {}
                   |  {}



