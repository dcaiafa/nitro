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

%type <ast> module
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
%type <asts> object_elifs_opt
%type <asts> object_elifs
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

S: module                             { yylex.(*lex).Module = $1.(*ast.Module) }

module: stmts_opt                     
      { 
        $$ = &ast.Module{Stmts: $1}
        $$.SetPos($1.Pos())
      }

stmts_opt: stmts
         | stmt                        { $$ = ast.ASTs{$1}; }
         | /*empty*/                   { $$ = ast.ASTs{} }

stmts: stmts stmt ';'                  { l := $1; $$ = append(l, $2) }
     | stmts ';'
     | stmt ';'                        { $$ = ast.ASTs{$1} }
     | ';'                             { $$ = ast.ASTs{} }

stmt: assignment_stmt
    | var_decl_stmt
    | for_stmt
    | if_stmt
    | func_stmt
    | return_stmt
    | while_stmt
    | expr                             
      { 
        $$ = &ast.ExprStmt{Expr:$1}
        $$.SetPos($1.Pos())
      }

while_stmt: kWHILE expr kDO 
              stmts_opt
            kEND
            { 
              $$ = &ast.WhileStmt{Predicate: $2, Stmts: $4}
              $$.SetPos($1.Pos)
            }

return_stmt: kRETURN return_list       
             { 
               $$ = &ast.ReturnStmt{Values: $2.(ast.Exprs)}
               $$.SetPos($1.Pos)
             }

return_list: return_list ',' expr      
             {
               l := $1.(ast.Exprs)
               $$ = append(l, $3)
             }
           | expr
             { 
               $$ = ast.Exprs{$1}
               $$.SetPos($1.Pos())
             }

func_stmt: kFN ID '(' param_list_opt ')'
             stmts_opt
           kEND
           { 
             fn := &ast.LambdaExpr{FuncParams:$4, Stmts:$6}
             $$ = &ast.VarDeclStmt{VarName:$2, InitValue:fn}
             $$.SetPos($1.Pos)
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
           $$.SetPos($1.Pos)
         }

elifs_opt: elifs
         | /*empty*/ 
           {
             $$ = ast.ASTs{}
           }

elifs: elifs elif
       {
         l := $1
         $$ = append(l, $2)
       }
     | elif
       {
         $$ = ast.ASTs{$1}
       }

elif: kELIF expr kTHEN
        stmts_opt
      {
        $$ = &ast.IfBlock{Pred:$2, Stmts:$4}
        $$.SetPos($1.Pos)
      }

else_opt: else
        | /*empty*/
          {
            $$ = nil
          }

else: kELSE
        stmts_opt
      {
        $$ = &ast.IfBlock{Pred:nil, Stmts:$2}
        $$.SetPos($1.Pos)
      }

for_stmt: kFOR for_vars kIN expr kDO
            stmts_opt
          kEND
          {
            $$ = &ast.ForStmt{ForVars:$2, IterExpr:$4, Stmts: $6}
            $$.SetPos($1.Pos)
          }

for_vars: for_vars ',' ID
          {
            $$ = append($1, &ast.ForVar{VarName:$3})
          }
        | ID
          {
            $$ = ast.ASTs{&ast.ForVar{VarName:$1}}
            $$.SetPos($1.Pos)
          }

assignment_stmt: expr '=' expr
                 {
                   $$ = &ast.AssignStmt{Lvalue:$1, Rvalue:$3}
                   $$.SetPos($1.Pos())
                 }

var_decl_stmt: kVAR ID
               {
                 $$ = &ast.VarDeclStmt{VarName: $2, InitValue:nil}
                 $$.SetPos($1.Pos)
               }
             | kVAR ID '=' expr
               {
                 $$ = &ast.VarDeclStmt{VarName: $2, InitValue:$4}
                 $$.SetPos($1.Pos)
               }

expr: unary_expr
    | expr '+' expr 
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpPlus}
        $$.SetPos($1.Pos())
      }
    | expr '-' expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpMinus}
        $$.SetPos($1.Pos())
      }
    | expr '*' expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpMult}
        $$.SetPos($1.Pos())
      }
    | expr '/' expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpDiv}
        $$.SetPos($1.Pos())
      }
    | expr '<' expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpLT}
        $$.SetPos($1.Pos())
      }
    | expr LE expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpLE}
        $$.SetPos($1.Pos())
      }
    | expr '>' expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpGT}
        $$.SetPos($1.Pos())
      }
    | expr GE expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpGE}
        $$.SetPos($1.Pos())
      }
    | expr EQ expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpEq}
        $$.SetPos($1.Pos())
      }
    | expr NE expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpNE}
        $$.SetPos($1.Pos())
      }
    | expr kAND expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpAnd}
        $$.SetPos($1.Pos())
      }
    | expr kOR expr
      {
        $$ = &ast.BinaryExpr{Left:$1, Right:$3, Op:ast.OpOr}
        $$.SetPos($1.Pos())
      }

unary_expr: kNOT unary_expr
            {
              $$ = &ast.UnaryExpr{Term:$2, Op:ast.OpNot}
              $$.SetPos($1.Pos)
            }
          | '+' unary_expr
            {
              $$ = &ast.UnaryExpr{Term:$2, Op:ast.OpPlus}
              $$.SetPos($1.Pos)
            }
          | '-' unary_expr
            {
              $$ = &ast.UnaryExpr{Term:$2, Op:ast.OpMinus}
              $$.SetPos($1.Pos)
            }
          | primary_expr

primary_expr: STRING
              {
                $$ = &ast.LiteralExpr{Val:$1}
                $$.SetPos($1.Pos)
              }
            | NUMBER
              {
                $$ = &ast.LiteralExpr{Val:$1}
                $$.SetPos($1.Pos)
              }
            | ID
              {
                $$ = &ast.SimpleRef{ID:$1}
                $$.SetPos($1.Pos)
              }
            | kTRUE
              {
                $$ = &ast.LiteralExpr{Val:$1}
                $$.SetPos($1.Pos)
              }
            | kFALSE
              {
                $$ = &ast.LiteralExpr{Val:$1}
                $$.SetPos($1.Pos)
              }
            | primary_expr '[' expr ']'
              {
                $$ = &ast.IndexExpr{Target:$1, Index:$3}
                $$.SetPos($1.Pos())
              }
            | primary_expr '.' ID
              {
                $$ = &ast.MemberAccess{Target:$1, Member:$3 }
                $$.SetPos($1.Pos())
              }
            | primary_expr '(' arg_list_opt ')'
              {
                $$ = &ast.FuncCall{Target:$1, Args:$3}
                $$.SetPos($1.Pos())
              }
            | '(' expr ')'
              {
                $$ = $2
              }
            | lambda_expr
            | array_literal
            | object_literal

lambda_expr: kFN '(' param_list_opt ')'
               stmts_opt
             kEND
             { 
               $$ = &ast.LambdaExpr{FuncParams:$3, Stmts:$5}
               $$.SetPos($1.Pos)
             }

param_list_opt: param_list
              | /*empty*/
              {
                $$ = ast.ASTs{}
              }

param_list: param_list ',' ID
            {
              param := &ast.FuncParam{Name:$3}
              param.SetPos($3.Pos)
              $$ = append($1, param)
            }
          | ID
            {
              $$ = ast.ASTs{&ast.FuncParam{Name:$1}}
              $$.SetPos($1.Pos)
            }

arg_list_opt: arg_list trailing_comma
            | /*empty*/
              {
                $$ = ast.Exprs{}
              }

arg_list: arg_list ',' expr
          { 
            $$ = append($1, $3)
          }
        | expr
          {
            $$ = ast.Exprs{$1}
          }

object_literal: '{' object_fields_opt '}'
                {
                  $$ = &ast.ObjectLiteral{Fields:$2}
                  $$.SetPos($1.Pos)
                }

object_fields_opt: object_fields trailing_seps
                 | /*empty*/
                   {
                     $$ = ast.ASTs{}
                   }

object_fields: object_fields ',' object_field
               {
                 $$ = append($1, $3)
               }
             | object_fields ';' object_field
               {
                 $$ = append($1, $3)
               }
             | object_field
               {
                 $$ = ast.ASTs{$1}
               }

object_field: ID ':' expr
              {
                $$ = &ast.ObjectField{NameID:$1, Val:$3}
              }
            | '[' expr ']' ':' expr
              {
                $$ = &ast.ObjectField{NameExpr:$2, Val: $5}
              }
            | object_if
            | object_for

object_if: kIF expr kTHEN
             object_fields_opt
           object_elifs_opt
           object_else_opt
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
             $$.SetPos($1.Pos)
           }

object_for: kFOR for_vars kIN expr kDO
              object_fields_opt
            kEND
            {
              $$ = &ast.ForStmt{ForVars:$2, IterExpr:$4, Stmts: $6}
              $$.SetPos($1.Pos)
            }

object_elifs_opt: object_elifs
                | /*empty*/
                  {
                    $$ = ast.ASTs{}
                  }

object_elifs: object_elifs object_elif
              { 
                $$ = append($1, $2)
              }
            | object_elif
              {
                $$ = ast.ASTs{$1}
              }

object_elif: kELIF expr kTHEN        
               object_fields_opt
             {
               $$ = &ast.IfBlock{Pred:$2, Stmts:$4}
               $$.SetPos($1.Pos)
             }

object_else_opt: object_else
               | /*empty*/
               {
                 $$ = nil
               }

object_else: kELSE 
               object_fields_opt
             {
               $$ = &ast.IfBlock{Pred:nil, Stmts:$2}
               $$.SetPos($1.Pos)
             }

array_literal: '[' array_elems_opt ']'
               {
                 $$ = &ast.ArrayLiteral{Elements:$2}
               }

array_elems_opt: array_elems trailing_seps
               | /*empty*/ 
                 {
                   $$ = ast.ASTs{}
                 }

array_elems: array_elems ',' array_elem
             {
               $$ = append($1, $3)
             }
           | array_elems ';' array_elem
             { 
               $$ = append($1, $3)
             } 
           | array_elem
             {
               $$ = ast.ASTs{$1}
             }

array_elem: expr 
            {
              $$ = &ast.ArrayElement{Val:$1}
              $$.SetPos($1.Pos())
            }
          | array_if
          | array_for

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
            $$.SetPos($1.Pos)
          }

array_for: kFOR for_vars kIN expr kDO
             array_elems_opt
           kEND
           {
             $$ = &ast.ForStmt{ForVars:$2, IterExpr:$4, Stmts:$6}
             $$.SetPos($1.Pos)
           }

array_elifs_opt: array_elifs
               | /*empty*/
                 {
                   $$ = ast.ASTs{}
                 }

array_elifs: array_elifs array_elif
             { 
               $$ = append($1, $2)
             }
           | array_elif
             {
               $$ = ast.ASTs{$1}
             }

array_elif: kELIF expr kTHEN
              array_elems_opt
            {
              $$ = &ast.IfBlock{Pred:$2, Stmts:$4}
              $$.SetPos($1.Pos)
            }

array_else_opt: array_else
              | /*empty*/
                {
                  $$ = nil
                }

array_else: kELSE
              array_elems_opt
            {
              $$ = &ast.IfBlock{Pred:nil, Stmts:$2}
              $$.SetPos($1.Pos)
            }

trailing_seps: trailing_seps1 | /*empty*/

trailing_seps1: trailing_seps1 ';'
              | trailing_seps1 ','
              | ';'
              | ','

trailing_comma: ',' | /*empty*/
