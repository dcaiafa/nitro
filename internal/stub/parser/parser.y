%{
package parser

import (
  "github.com/dcaiafa/nitro/internal/stub/ast"
  "github.com/dcaiafa/nitro/internal/token"
)

func cast[T ast.AST](v ast.AST) T {
	cv, _ := v.(T)
	return cv
}

%}

%union {
  tok   token.Token
  ast   ast.AST
  asts  ast.ASTs
  other interface{}
}

%token LEXERR

%token <tok> kCONST
%token <tok> kFUNC
%token <tok> kPACKAGE
%token <tok> kTYPE
%token <tok> kSTRUCT
%token <tok> kNIL
%token <tok> kIMPORT

%token <tok> INT
%token <tok> FLOAT
%token <tok> STRING
%token <tok> ID
%token <tok> DOC
%token <tok> ELLIPSIS

%token <tok> '=' ':' ';' '(' ')' '[' ']' '*' '.' '?' '-' '{' '}'

/*   :sort /\S\+$/ r   */
%type <ast> const_decl
%type <ast> const_value
%type <ast> decl
%type <asts> decls
%type <asts> decls_opt
%type <ast> doc_opt
%type <ast> func_decl
%type <ast> func_param
%type <ast> func_param_d
%type <asts> func_params
%type <asts> func_params_opt
%type <asts> func_rets
%type <asts> func_rets_opt
%type <ast> go_type
%type <ast> import_decl
%type <ast> package
%type <ast> simple_go_type
%type <ast> struct_decl
%type <ast> struct_field
%type <asts> struct_fields
%type <asts> struct_fields_opt
%type <ast> type_decl
%type <ast> type_ref
%type <ast> type_ref_simple
%type <ast> unit

%start S

%%

S: unit                                      { yylex.(*lex).Unit = $1.(*ast.Unit) }
 ;

unit: package decls_opt          
    { 
      unit := $$.(*ast.Unit)
      unit.Decls = $2
    }
    ;

package: kPACKAGE STRING
         {
           $$ = &ast.Unit{ Package: $2 }
         }
       ;

decls_opt: decls
         | /* empty */              { $$ = nil }
         ;

decls: decls decl                   { $$ = append($1, $2) }
     | decl                         { $$ = ast.ASTs{$1} }
     ;

decl: const_decl
    | import_decl
    | type_decl
    | struct_decl
    | func_decl
    ;

doc_opt: DOC                        { $$ = &ast.Doc{} }
       | /* empty */                { $$ = nil }
       ;


const_decl: doc_opt kCONST ID type_ref '=' const_value
            { 
              $$ = &ast.ConstDecl{
                Doc: cast[*ast.Doc]($1),
                ID: $3,
                TypeRef: cast[*ast.TypeRef]($4),
                Value: cast[*ast.ConstValue]($6),
              }
            }
          ;

const_value: STRING
             {
               $$ = &ast.ConstValue{ Expr: $1.Str }
             }
           ;

import_decl: kIMPORT ID STRING
             {
               $$ = &ast.ImportDecl{
                 Alias: $2.Str,
                 Import: $3.Str,
               }
             }
           ;

struct_decl: kSTRUCT ID '{' struct_fields_opt '}'
             {
               $$ = &ast.StructDecl{
                Name: $2.Str,
                Fields: $4,
               }
             }
           ;

struct_fields_opt: struct_fields
                 | /* empty */
                   {
                     $$ = nil
                   }
                 ;

struct_fields: struct_fields struct_field ';'
               {
                 $$ = append($1, $2)
               }
             | struct_field ';'
               {
                 $$ = ast.ASTs{$1}
               }
             ;

struct_field: ID type_ref
              {
                $$ = &ast.StructField{
                  Name: $1.Str,
                  Type: cast[*ast.TypeRef]($2),
                }
              }
            ;

type_decl: doc_opt kTYPE ID go_type
           { 
             $$ = &ast.TypeDecl{
               Doc: cast[*ast.Doc]($1),
               ID: $3,
               GoType: cast[*ast.GoType]($4),
             }
           }
         ;

go_type: '*' simple_go_type
             { 
               $$ = $2
               $$.(*ast.GoType).Ref = true
             }
           | simple_go_type
           ;

simple_go_type: STRING '.' ID
                { 
                  $$ = &ast.GoType{Package:$1, ID:$3}
                }
              | ID
                { 
                  $$ = &ast.GoType{ID:$1}
                }
              ;

func_decl: doc_opt kFUNC ID '(' func_params_opt ')' func_rets_opt
           { 
             $$ = &ast.FuncDecl{
               Doc: cast[*ast.Doc]($1),
               ID: $3,
               Params: $5,
               Rets: $7,
             }
           }
         ;

func_params_opt: func_params
               | /* empty */
                { $$ = nil }
               ;

func_params: func_params ',' func_param_d  { $$ = append($1, $3) }      
           | func_param_d                  { $$ = ast.ASTs{$1} }        
           ;

func_param_d: func_param
            | func_param '=' const_value
              {
                $$.(*ast.FuncParam).DefaultValue = $3.(*ast.ConstValue)
              }
            ;

func_param: ID type_ref
            {
              $$ = &ast.FuncParam{
                ID: $1,
                Type: cast[*ast.TypeRef]($2),
              }
            }
          | ID ELLIPSIS type_ref
            {
              $$ = &ast.FuncParam{
                ID: $1,
                Type: cast[*ast.TypeRef]($3),
                VarArg: true,
              }
            }
          ;

func_rets_opt: '(' func_rets ')'
               { 
                 $$ = $2
               }
             | type_ref
               {
                 $$ = ast.ASTs{$1}
               }
             | /* empty */          { $$ = nil }
             ;

func_rets: func_rets ',' type_ref   { $$ = append($1, $3) }  
         | type_ref                 { $$ = ast.ASTs{$1} }    
         ;

type_ref: type_ref_simple
        | type_ref_simple '?'
          {
            $$.(*ast.TypeRef).Nilable = true
          }
        ;

type_ref_simple: ID
                 {
                   $$ = &ast.TypeRef{ID: $1}
                 }
               ;
