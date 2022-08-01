parser grammar NitroParser;

options { tokenVocab=NitroLexer; }

start: unit EOF;

unit: prologue stmts;

prologue: meta_directive* import_stmt*;

// Meta

meta_directive: meta_info 
              | meta_param
              | meta_flag;

meta_info: M_INFO (OCURLY meta_attribs CCURLY)? ';';
meta_param: M_PARAM ID ('=' expr)? (OCURLY meta_attribs CCURLY)? ';';
meta_flag: M_FLAG ID ('=' expr)? (OCURLY meta_attribs CCURLY)? ';';
meta_attribs: meta_attrib ((COMMA|';') meta_attrib)* (COMMA|';')?;
meta_attrib: id_or_keyword (COLON meta_literal)?;
meta_literal: val=(STRING | NUMBER | CHAR | TRUE | FALSE | NIL);

// Statements

import_stmt: IMPORT ID? STRING ';';

stmts: stmt_list? ';'*;

stmt_list: stmt (';'+ stmt)*;

stmt: assignment_stmt      # stmt_assignment
    | assignment_op_stmt   # stmt_op_assign
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
    | inc_dec_stmt         # stmt_inc_dec
    ;

assignment_stmt: assignment_lvalues '=' rvalues;
assignment_lvalues: lvalue_expr (COMMA lvalue_expr)*;
rvalues: expr (COMMA expr)*;

assignment_op_stmt: lvalue_expr op=('+='|'-='|'*='|'/='|'%=') expr;

var_decl_stmt: VAR var_decl_vars ('=' rvalues)?;
var_decl_vars: ID (COMMA ID)*;

for_stmt: FOR for_vars ID expr OCURLY stmts CCURLY;
for_vars: ID (COMMA ID)*;

while_stmt: WHILE expr OCURLY stmts CCURLY;

if_stmt: IF expr OCURLY stmts CCURLY if_elif* if_else?;
if_elif: ELSE IF expr OCURLY stmts CCURLY;
if_else: ELSE OCURLY stmts CCURLY;

func_stmt: FUNC ID OPAREN param_list? CPAREN OCURLY stmts CCURLY;
param_list: ID (COMMA ID)*;

return_stmt: RETURN rvalues?;

try_catch_stmt: TRY OCURLY stmts CCURLY CATCH ID? OCURLY stmts CCURLY;

throw_stmt: THROW expr;

defer_stmt: DEFER primary_expr;

yield_stmt: YIELD rvalues;

break_stmt: BREAK;

continue_stmt: CONTINUE;

inc_dec_stmt: lvalue_expr op=('++'|'--');

// Expressions

expr: expr PIPE expr
    | expr2
    ;

expr2: short_lambda_expr
     | expr3
     ;

expr3: <assoc=right> expr3 QUESTION_MARK expr3 COLON expr3
     | binary_expr
     ;

binary_expr: unary_expr
           | binary_expr op=(MUL|DIV|MOD) binary_expr
           | binary_expr op=(ADD|SUB) binary_expr
           | binary_expr op=(LT|LE|GT|GE|EQ|NE) binary_expr
           | binary_expr op=AND binary_expr
           | binary_expr op=OR binary_expr
           ;

unary_expr: op=NOT primary_expr
          | op=ADD primary_expr
          | op=SUB primary_expr
          | primary_expr
          ;

primary_expr: ID                                                    # primary_expr_simple_ref
            | primary_expr PERIOD ID                                # primary_expr_member_access
            | primary_expr OBRACKET expr CBRACKET                   # primary_expr_index
            | primary_expr OBRACKET b=expr? COLON e=expr? CBRACKET  # primary_expr_slice
            | primary_expr OPAREN (arg_list EXPAND? )? CPAREN       # primary_expr_call
            | lambda_expr                                           # primary_expr_lambda
            | exec_expr                                             # primary_exec_expr
            | object_literal                                        # primary_expr_object
            | array_literal                                         # primary_expr_array
            | simple_literal                                        # primary_expr_literal
            | REGEX                                                 # primary_expr_regex
            | OPAREN expr CPAREN                                    # primary_expr_parenthesis
            ;

simple_literal: val=(STRING | NUMBER | CHAR | TRUE | FALSE | NIL);

arg_list: expr (COMMA expr)* (COMMA|';')?;

lvalue_expr: ID                                   # lvalue_expr_simple_ref
           | primary_expr PERIOD ID               # lvalue_expr_member_access
           | primary_expr OBRACKET expr CBRACKET  # lvalue_expr_index
           ;

lambda_expr: FUNC OPAREN param_list? CPAREN OCURLY stmts CCURLY;

short_lambda_expr: LAMBDA param_list? ARROW expr3;

exec_expr: EXEC_PREFIX exec_expr_arg+ EXEC_SUFFIX;
exec_expr_arg: EXEC_LITERAL
             | EXEC_DQUOTE_LITERAL
             | EXEC_SQUOTE_LITERAL
             | OCURLY expr CCURLY; 

object_literal: OCURLY object_fields CCURLY;
object_fields: (object_field ((COMMA|';') object_field)* (COMMA|';')?)?;

object_field: id_or_keyword COLON expr           # object_field_id_key
            | OBRACKET expr CBRACKET COLON expr  # object_field_expr_key
            | primary_expr EXPAND                # object_field_expansion
            ;

array_literal: OBRACKET array_elems CBRACKET;
array_elems: ( array_elem ((COMMA|';') array_elem)* (COMMA|';')* )?;
array_elem: expr;

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
