parser grammar NitroParser;

options { tokenVocab=NitroLexer; }

start: module;

module: meta_section? stmts;

// Metadata

meta_section: META meta_entry* END;

meta_entry: ID meta_object;

meta_object: '{' meta_fields? '}';

meta_fields: meta_field ((','|';') meta_field)*;

meta_field: ID ':' meta_field_value;

meta_field_value: STRING | NUMBER | TRUE | FALSE | meta_object;

// Statements

stmts: stmt*;

stmt: assignment_stmt ';'
    | var_decl_stmt ';'
    | for_stmt ';'
    | while_stmt ';'
    | if_stmt ';'
    | func_stmt ';'
    | func_call_stmt ';'
    | return_stmt ';'
    | ';'
    ;

assignment_stmt: assignment_lvalues '=' assignment_rvalues;
assignment_lvalues: lvalue_expr (',' lvalue_expr)*;
assignment_rvalues: expr (',' expr)*;

var_decl_stmt: VAR var_decl_vars ('=' assignment_rvalues)?;
var_decl_vars: ID (',' ID)*;

for_stmt: FOR for_vars IN expr DO stmts END;
for_vars: ID (',' ID)*;

while_stmt: WHILE expr DO stmts END;

if_stmt: IF expr THEN stmts if_elif* if_else?;
if_elif: ELIF expr THEN stmts;
if_else: ELSE stmts END;

func_stmt: FN ID '(' param_list? ')' stmts END;
param_list: ID (',' ID)*;

func_call_stmt: primary_expr '(' arg_list? ')';

return_stmt: RETURN expr*;

// Expressions

expr: unary_expr
    | expr op=('*'|'/'|'%') expr
    | expr op=('+'|'-') expr
    | expr op=('<'|'<='|'>'|'>='|'=='|'!=') expr
    | expr op=AND expr
    | expr op=OR expr
    ;

unary_expr: op=NOT unary_expr
          | op='+' unary_expr
          | op='-' unary_expr
          | primary_expr
          ;

primary_expr: ID                                       # primary_expr_simple_ref
            | primary_expr '.' ID                      # primary_expr_member_access
            | primary_expr '[' expr ']'                # primary_expr_index
            | primary_expr '[' b=expr? ':' e=expr? ']' # primary_expr_slice
            | primary_expr '(' arg_list? ')'           # primary_expr_call
            | lambda_expr                              # primary_expr_lambda
            | object_literal                           # primary_expr_object
            | array_literal                            # primary_expr_array
            | simple_literal                           # primary_expr_literal
            | '(' expr ')'                             # primary_expr_parenthesis
            ;

simple_literal: val=(STRING | NUMBER | TRUE | FALSE);

arg_list: expr (',' expr)*;

lvalue_expr: ID                          # lvalue_expr_simple_ref
           | primary_expr '.' ID         # lvalue_expr_member_access
           | primary_expr '[' expr ']'   # lvalue_expr_index
           ;

lambda_expr: FN '(' param_list? ')' stmts END;

object_literal: '{' object_fields? '}';
object_fields: object_field ((','|';') object_field)* (','|';')*;

object_field: id_or_keyword ':' expr     # object_field_id_key
            | '[' expr ']' ':' expr      # object_field_expr_key
            | object_if                  # object_field_if
            | object_for                 # object_field_for
            ;

object_if: IF expr THEN object_fields? object_elif* object_else? END;
object_elif: ELIF expr THEN object_fields?;
object_else: ELSE object_fields?;

object_for: FOR for_vars IN expr DO object_fields? END;

array_literal: '[' array_elems? ']';
array_elems: array_elem ((','|';') array_elem)* (','|';')*;
array_elem: expr | array_if | array_for;

array_if: IF expr THEN array_elems? array_elif* array_else? END;
array_elif: ELIF expr THEN array_elems?;
array_else: ELSE array_elems?;

array_for: FOR for_vars IN expr DO array_elems? END;

id_or_keyword: 
    t=(ID | AND | DO | ELIF | ELSE | END | FALSE |
       FN | FOR | IF | IN | META | NOT | OR | RETURN |
       THEN | TRUE | VAR | WHILE)
    ;
