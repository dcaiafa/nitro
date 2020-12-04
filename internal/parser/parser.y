%{
package parser

import (
  "github.com/dcaiafa/nitro/internal/ast"
  "github.com/dcaiafa/nitro/internal/token"
)

%}

%union {
  tok token.Token
  ast ast.AST
  asts ast.ASTs
  expr ast.Expr
  exprs ast.Exprs
  other interface{}
}

%token LEXERR

%token <tok> kAND 
%token <tok> kDO
%token <tok> kELIF 
%token <tok> kELSE 
%token <tok> kEND 
%token <tok> kFALSE 
%token <tok> kFN
%token <tok> kFOR
%token <tok> kIF 
%token <tok> kIN 
%token <tok> kNOT 
%token <tok> kOR 
%token <tok> kRETURN
%token <tok> kTHEN
%token <tok> kTRUE 
%token <tok> kVAR
%token <tok> kWHILE

%token <tok> NUMBER
%token <tok> STRING
%token <tok> ID

%token <tok> '=' EQ NE '<' LE '>' GE '"' '|' '+' '-' '*' '/' ';' '(' ',' '['
%token <tok> ':' '.' '{' ')' ']' '}'

%left kOR
%left kAND
%nonassoc '<' LE '>' GE EQ NE
%left '+' '-'
%left '*' '/'

%type <ast> program
%type <asts> stmts_opt
%type <asts> stmts
%type <ast> stmt
%type <ast> while_stmt
%type <ast> return_stmt
%type <ast> return_list
%type <ast> func_stmt
%type <ast> if_stmt
%type <asts> elifs_opt
%type <asts> elifs
%type <ast> elif
%type <ast> else_opt
%type <ast> else
%type <ast> for_stmt
%type <asts> for_vars
%type <ast> assignment_stmt
%type <ast> var_decl_stmt
%type <expr> expr
%type <expr> unary_expr
%type <expr> primary_expr
%type <expr> lambda_expr
%type <asts> param_list_opt
%type <asts> param_list
%type <exprs> arg_list_opt
%type <exprs> arg_list
%type <expr> object_literal
%type <asts> object_fields_opt
%type <asts> object_fields
%type <ast> object_field
%type <ast> object_if
%type <ast> object_for
%type <ast> object_elifs_opt
%type <ast> object_elifs
%type <ast> object_elif
%type <ast> object_else_opt
%type <ast> object_else
%type <expr> array_literal
%type <asts> array_elems_opt
%type <asts> array_elems
%type <ast> array_elem
%type <ast> array_if
%type <ast> array_for
%type <asts> array_elifs_opt
%type <asts> array_elifs
%type <ast> array_elif
%type <ast> array_else_opt
%type <ast> array_else

%start S

%%

S: program                             { yylex.(*lex).Program = $1.(*ast.Program) }

program: stmts_opt                     { $$ = &ast.Program{Stmts: $1} }

stmts_opt: stmts                       { $$ = $1 }
         | stmt                        { $$ = ast.ASTs{$1} }
         | /*empty*/                   { $$ = ast.ASTs{} }

stmts: stmts stmt ';'                  { l := $1; $$ = append(l, $2) }
     | stmts ';'                       { $$ = $1 }
     | stmt ';'                        { $$ = ast.ASTs{$1} }
     | ';'                             { $$ = ast.ASTs{} }

stmt: assignment_stmt
    | var_decl_stmt
    | for_stmt
    | if_stmt
    | func_stmt
    | return_stmt
    | while_stmt
    | expr                             { $$ = &ast.ExprStmt{Expr:$1} }

while_stmt: kWHILE expr kDO 
              stmts_opt
            kEND
            { $$ = &ast.WhileStmt{ Predicate: $2, Stmts: $4 } }

return_stmt: kRETURN return_list       { $$ = &ast.ReturnStmt{Values: $2.(ast.Exprs)} }

return_list: return_list ',' expr      { e := $1.(ast.Exprs); e = append(e, $3); $$ = e }
           | expr                      { $$ = ast.Exprs{ $1 } }

func_stmt: kFN ID '(' param_list_opt ')'
             stmts_opt
           kEND
           { 
             fn := &ast.LambdaExpr{FuncParams:$4, Stmts:$6}
             $$ = &ast.VarDeclStmt{VarName:$2, InitValue:fn}
           }

if_stmt: kIF expr kTHEN
           stmts_opt
         elifs_opt
         else_opt
         kEND
         { 
           ifBlock := &ast.IfBlock{Pred: $2, Stmts: $4}
           ifStmt := &ast.IfStmt{Blocks: ast.ASTs{ifBlock}}
           if $5 != nil {
             ifStmt.Blocks = append(ifStmt.Blocks, $5...)
           }
           if $6 != nil {
             ifStmt.Blocks = append(ifStmt.Blocks, $6)
           }
           $$ = ifBlock
         }

elifs_opt: elifs     { $$ = $1 }
         | /*empty*/ { $$ = ast.ASTs{} }

elifs: elifs elif                      { l := $1; $$ = append(l, $2) }
     | elif                            { $$ = ast.ASTs{$1} }

elif: kELIF expr kTHEN
        stmts_opt
      { $$ = &ast.IfBlock{Pred:$2, Stmts:$4} }

else_opt: else                         { $$ = $1 }
        | /*empty*/                    { $$ = nil }

else: kELSE
        stmts_opt
      { $$ = &ast.IfBlock{Pred:nil, Stmts:$2} }

for_stmt: kFOR for_vars kIN expr kDO
            stmts_opt
          kEND
          { $$ = &ast.ForStmt{ForVars:$2, IterExpr:$4, Stmts: $6} }

for_vars: for_vars ',' ID              { l := $1; $$ = append(l, &ast.ForVar{VarName:$3}) }
        | ID                           { $$ = ast.ASTs{&ast.ForVar{VarName:$1}} }

assignment_stmt: expr '=' expr         { $$ = &ast.AssignStmt{Lvalue:$1, Rvalue:$3} }

var_decl_stmt: kVAR ID                 { $$ = &ast.VarDeclStmt{VarName: $2, InitValue:nil} }
             | kVAR ID '=' expr        { $$ = &ast.VarDeclStmt{VarName: $2, InitValue:$4} }

expr: unary_expr                       { $$ = $1 }
    | expr '+' expr                    { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpPlus} }
    | expr '-' expr                    { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpMinus} }
    | expr '*' expr                    { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpMult} }
    | expr '/' expr                    { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpDiv} }
    | expr '<' expr                    { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpLT} }
    | expr LE expr                     { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpLE} }
    | expr '>' expr                    { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpGT} }
    | expr GE expr                     { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpGE} }
    | expr EQ expr                     { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpEq} }
    | expr NE expr                     { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpNE} }
    | expr kAND expr                   { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpAnd} }
    | expr kOR expr                    { $$ = &ast.BinaryExpr{Left:$1, Right: $3, Op:ast.OpOr} }

unary_expr: kNOT unary_expr            { $$ = &ast.UnaryExpr{ Term: $2, Op: ast.OpNot} }
        | '+' unary_expr               { $$ = &ast.UnaryExpr{ Term: $2, Op: ast.OpPlus} }
        | '-' unary_expr               { $$ = &ast.UnaryExpr{ Term: $2, Op: ast.OpMinus} }
        | primary_expr                 { $$ = $1 }

primary_expr: STRING                             { $$ = &ast.LiteralExpr{Val:$1} }
            | NUMBER                             { $$ = &ast.LiteralExpr{Val:$1} }
            | ID                                 { $$ = &ast.SimpleRef{ID:$1} }
            | kTRUE                              { $$ = &ast.LiteralExpr{Val:$1} }
            | kFALSE                             { $$ = &ast.LiteralExpr{Val:$1} }
            | array_literal                      { $$ = $1 }
            | object_literal                     { $$ = $1 }
            | primary_expr '[' expr ']'          { $$ = &ast.IndexExpr{Target:$1, Index:$3} }
            | primary_expr '.' ID                { $$ = &ast.MemberAccess{Target:$1, Member:$3 } }
            | primary_expr '(' arg_list_opt ')'  { $$ = &ast.FuncCall{Target:$1, Args:$3} }
            | lambda_expr                        { $$ = $1 }
            | '(' expr ')'                       { $$ = $2 }

lambda_expr: kFN '(' param_list_opt ')'
               stmts_opt
             kEND
             { $$ = &ast.LambdaExpr{FuncParams:$3, Stmts:$5} }

param_list_opt: param_list             { $$ = $1 }
              | /*empty*/              { $$ = ast.ASTs{} }

param_list: param_list ',' ID          { p := $1; p = append(p, $1); $$ = p }
          | ID                         { $$ = ast.ASTs{&ast.FuncParam{Name:$1}} }

arg_list_opt: arg_list trailing_comma  { $$ = $1 }
            | /*empty*/                { $$ = ast.Exprs{} }

arg_list: arg_list ',' expr            { l := $1; $$ = append(l, $3) }
        | expr                         { $$ = ast.Exprs{$1} }

object_literal: '{' object_fields_opt '}'   { $$ = &ast.ObjectLiteral{Fields:$2} }

object_fields_opt: object_fields trailing_seps  { $$ = $1 }
                 |                              { $$ = ast.ASTs{} }

object_fields: object_fields ',' object_field   { l := $1; $$ = append(l, $3) }
             | object_fields ';' object_field   { l := $1; $$ = append(l, $3) } 
             | object_field                     { $$ = ast.ASTs{$1} }

object_field: ID ':' expr                { $$ = &ast.ObjectField{NameID:$1, Val:$3} }
            | '[' expr ']' ':' expr      { $$ = &ast.ObjectField{NameExpr:$2, Val: $5} }
            | object_if                  { $$ = $1 }
            | object_for                 { $$ = $1 }

object_if: kIF expr kTHEN
             object_fields_opt
           object_elifs_opt
           object_else_opt
           kEND
           {
             ifBlock := &ast.IfBlock{Pred:$2, Stmts:$4}
             ifStmt := &ast.IfStmt{Blocks: ast.ASTs{ifBlock}}
             if $5 != nil {
               ifStmt.Blocks = append(ifStmt.Blocks, $5.(ast.ASTs)...)
             }
             if $6 != nil {
               ifStmt.Blocks = append(ifStmt.Blocks, $6)
             }
             $$ = ifStmt
           }

object_for: kFOR for_vars kIN expr kDO
              object_fields_opt
            kEND
            { $$ = &ast.ForStmt{ForVars:$2, IterExpr:$4, Stmts: $6} }

object_elifs_opt: object_elifs         { $$ = $1 }
                |                      { $$ = ast.ASTs{} }

object_elifs: object_elifs object_elif { l := $1.(ast.ASTs); $$ = append(l, $2) }
            | object_elif              { $$ = ast.ASTs{$1} }

object_elif: kELIF expr kTHEN        
               object_fields_opt
             { $$ = &ast.IfBlock{Pred:$2, Stmts:$4} }

object_else_opt: object_else { $$ = $1 }
               |             { $$ = nil }

object_else: kELSE 
               object_fields_opt
             { $$ = &ast.IfBlock{Pred:nil, Stmts:$2} }

array_literal: '[' array_elems_opt ']'          { $$ = &ast.ArrayLiteral{Elements:$2} }

array_elems_opt: array_elems trailing_seps      { $$ = $1 }
               |                                { $$ = ast.ASTs{} }

array_elems: array_elems ',' array_elem         { l := $1; $$ = append(l, $3) }
           | array_elems ';' array_elem         { l := $1; $$ = append(l, $3) } 
           | array_elem                         { $$ = ast.ASTs{$1} }

array_elem: expr                       { $$ = &ast.ArrayElement{Val:$1} }
          | array_if                   { $$ = $1 }
          | array_for                  { $$ = $1 }

array_if: kIF expr kTHEN
            array_elems_opt
          array_elifs_opt
          array_else_opt
          kEND
          {
            ifBlock := &ast.IfBlock{Pred:$2, Stmts:$4}
            ifStmt := &ast.IfStmt{Blocks: ast.ASTs{ifBlock}}
            if $5 != nil {
              ifStmt.Blocks = append(ifStmt.Blocks, $5...)
            }
            if $6 != nil {
              ifStmt.Blocks = append(ifStmt.Blocks, $6)
            }
            $$ = ifStmt
          }

array_for: kFOR for_vars kIN expr kDO
             array_elems_opt
           kEND
           { $$ = &ast.ForStmt{ForVars:$2, IterExpr:$4, Stmts: $6} }

array_elifs_opt: array_elifs           { $$ = $1 }
               |                       { $$ = ast.ASTs{} }

array_elifs: array_elifs array_elif    { l := $1; $$ = append(l, $2) }
           | array_elif                { $$ = ast.ASTs{$1} }

array_elif: kELIF expr kTHEN
              array_elems_opt
            { $$ = &ast.IfBlock{Pred:$2, Stmts:$4} }

array_else_opt: array_else             { $$ = $1 }
              |                        { $$ = nil }

array_else: kELSE
              array_elems_opt
            { $$ = &ast.IfBlock{Pred:nil, Stmts:$2} }

trailing_seps: /*empty*/ | trailing_seps1

trailing_seps1: trailing_seps1 ';'
              | trailing_seps1 ','
              | ';'
              | ','

trailing_comma: ',' | /*empty*/
