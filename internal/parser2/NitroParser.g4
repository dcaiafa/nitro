parser grammar NitroParser;

options { tokenVocab=NitroLexer; }

start: unit EOF;

unit: meta_directive* import_stmt* stmts;

// Meta

meta_directive: meta_info 
              | meta_param
              | meta_flag;

meta_info: M_INFO ('{' meta_attribs '}')? ';';
meta_param: M_PARAM ID ('=' expr)? ('{' meta_attribs '}')? ';';
meta_flag: M_FLAG ID ('=' expr)? ('{' meta_attribs '}')? ';';
meta_attribs: meta_attrib ((','|';') meta_attrib)* (','|';')?;
meta_attrib: id_or_keyword (':' meta_literal)?;
meta_literal: val=(STRING | NUMBER | CHAR | TRUE | FALSE | NIL);

// Statements

import_stmt: IMPORT ID? STRING ';';

stmts: stmt_list? ';'*;

stmt_list: stmt (';'+ stmt)*;

stmt: assignment_stmt      # stmt_assignment
    | var_decl_stmt        # stmt_var_dec
    | for_stmt             # stmt_for
    | while_stmt           # stmt_while
    | if_stmt              # stmt_if
    | func_stmt            # stmt_func
    | return_stmt          # stmt_return
    | expr                 # stmt_expr
    | try_catch_stmt       # stmt_try_catch
    | throw_stmt           # stmt_throw
    | defer_stmt           # stmt_defer
    | yield_stmt           # stmt_yield
    | break_stmt           # stmt_break
    | continue_stmt        # stmt_continue
    ;

assignment_stmt: assignment_lvalues '=' rvalues;
assignment_lvalues: lvalue_expr (',' lvalue_expr)*;
rvalues: expr (',' expr)*;

var_decl_stmt: VAR var_decl_vars ('=' rvalues)?;
var_decl_vars: ID (',' ID)*;

for_stmt: FOR for_vars ID expr '{' stmts '}';
for_vars: ID (',' ID)*;

while_stmt: WHILE expr '{' stmts '}';

if_stmt: IF expr '{' stmts '}' if_elif* if_else?;
if_elif: ELSE IF expr '{' stmts '}';
if_else: ELSE '{' stmts '}';

func_stmt: FUNC ID '(' param_list? ')' '{' stmts '}';
param_list: ID (',' ID)*;

return_stmt: RETURN rvalues?;

try_catch_stmt: TRY '{' stmts '}' CATCH ID? '{' stmts '}';

throw_stmt: THROW expr;

defer_stmt: DEFER primary_expr;

yield_stmt: YIELD rvalues;

break_stmt: BREAK;

continue_stmt: CONTINUE;

// Expressions

expr: expr '|' expr
    | expr2
    ;

expr2: short_lambda_expr
     | expr3
     ;

expr3: <assoc=right> expr3 '?' expr3 ':' expr3
     | binary_expr
     ;

binary_expr: unary_expr
           | binary_expr op=('*'|'/'|'%') binary_expr
           | binary_expr op=('+'|'-') binary_expr
           | binary_expr op=('<'|'<='|'>'|'>='|'=='|'!=') binary_expr
           | binary_expr op=AND binary_expr
           | binary_expr op=OR binary_expr
           ;

unary_expr: op=NOT primary_expr
          | op='+' primary_expr
          | op='-' primary_expr
          | primary_expr
          ;

primary_expr: ID                                         # primary_expr_simple_ref
            | primary_expr '.' ID                        # primary_expr_member_access
            | primary_expr '[' expr ']'                  # primary_expr_index
            | primary_expr '[' b=expr? ':' e=expr? ']'   # primary_expr_slice
            | primary_expr '(' (arg_list '...'? )? ')'   # primary_expr_call
            | lambda_expr                                # primary_expr_lambda
            | object_literal                             # primary_expr_object
            | array_literal                              # primary_expr_array
            | simple_literal                             # primary_expr_literal
            | REGEX                                      # primary_expr_regex
            | '(' expr ')'                               # primary_expr_parenthesis
            ;

simple_literal: val=(STRING | NUMBER | CHAR | TRUE | FALSE | NIL);

arg_list: expr (',' expr)* (','|';')?;

lvalue_expr: ID                          # lvalue_expr_simple_ref
           | primary_expr '.' ID         # lvalue_expr_member_access
           | primary_expr '[' expr ']'   # lvalue_expr_index
           ;

lambda_expr: FUNC '(' param_list? ')' '{' stmts '}';

short_lambda_expr: '&' param_list? '->' expr3;

object_literal: '{' object_fields '}';
object_fields: (object_field ((','|';') object_field)* (','|';')?)?;

object_field: id_or_keyword ':' expr     # object_field_id_key
            | '[' expr ']' ':' expr      # object_field_expr_key
            | primary_expr '...'         # object_field_expansion
            | object_if                  # object_field_if
            | object_for                 # object_field_for
            ;

object_if: IF expr '{' object_fields '}' object_elif* object_else?;
object_elif: ELSE IF expr '{' object_fields '}';
object_else: ELSE '{' object_fields '}';

object_for: FOR for_vars ID expr '{' object_fields '}';

array_literal: '[' array_elems ']';
array_elems: ( array_elem ((','|';') array_elem)* (','|';')* )?;
array_elem: expr | array_if | array_for;

array_if: IF expr '{' array_elems '}' array_elif* array_else?;
array_elif: ELSE IF expr '{' array_elems '}';
array_else: ELSE '{' array_elems '}';

array_for: FOR for_vars ID expr '{' array_elems '}';

id_or_keyword: 
    t=(
      ID |

      // Keywords
      AND |
      BREAK | 
      CATCH |
      CONTINUE |
      DEFER |
      ELSE |
      FALSE |
      FOR |
      FUNC |
      IF |
      IMPORT |
      NIL |
      NOT |
      OR |
      RETURN |
      THROW |
      TRUE |
      TRY |
      VAR |
      WHILE |
      YIELD
    )
    ;
