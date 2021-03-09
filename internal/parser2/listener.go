package parser2

import (
	"log"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/parser2/parser"
	"github.com/dcaiafa/nitro/internal/token"
)

type listener struct {
	parser.BaseNitroParserListener

	Module *ast.Module

	filename string
	values   map[antlr.RuleContext]interface{}
}

func newListener(filename string) *listener {
	return &listener{
		filename: filename,
		values:   make(map[antlr.RuleContext]interface{}),
	}
}

func (l *listener) tokenToNitro(at antlr.Token) token.Token {
	t := token.Token{}
	switch at.GetTokenType() {
	case parser.NitroLexerSTRING:
		s := at.GetText()
		// Remove quotes.
		s = s[1 : len(s)-1]
		t.Type = token.String
		t.Str = s
		// TODO: expand escaped sequences.

	case parser.NitroLexerNUMBER:
		if strings.IndexByte(at.GetText(), '.') == -1 {
			t.Type = token.Int
			t.Int, _ = strconv.ParseInt(at.GetText(), 10, 64)
		} else {
			t.Type = token.Float
			t.Float, _ = strconv.ParseFloat(at.GetText(), 64)
		}
	case parser.NitroLexerTRUE:
		t.Type = token.Bool
		t.Bool = true
	case parser.NitroLexerFALSE:
		t.Type = token.Bool
		t.Bool = false
	default:
		t.Type = token.String
		t.Str = at.GetText()
	}
	t.Pos = l.tokenToPos(at)
	return t
}

func (l *listener) tokenToPos(at antlr.Token) token.Pos {
	return token.Pos{
		Col:      at.GetColumn(),
		Line:     at.GetLine(),
		Filename: l.filename,
	}
}

func (l *listener) put(k antlr.RuleContext, v interface{}) {
	l.values[k] = v
}

func (l *listener) take(k antlr.RuleContext) (interface{}, bool) {
	if k == nil {
		return nil, false
	}
	v, ok := l.values[k]
	if !ok {
		return nil, false
	}
	delete(l.values, k)
	return v, true
}

func (l *listener) takeAST(k antlr.RuleContext) ast.AST {
	v, ok := l.take(k)
	if !ok {
		return nil
	}
	return v.(ast.AST)
}

func (l *listener) takeASTs(k antlr.RuleContext) ast.ASTs {
	v, ok := l.take(k)
	if !ok {
		return nil
	}
	return v.(ast.ASTs)
}

func (l *listener) takeExpr(k antlr.RuleContext) ast.Expr {
	v, ok := l.take(k)
	if !ok {
		return nil
	}
	return v.(ast.Expr)
}

func (l *listener) takeExprs(k antlr.RuleContext) ast.Exprs {
	v, ok := l.take(k)
	if !ok {
		return nil
	}
	return v.(ast.Exprs)
}

// ExitEveryRule is called when any rule is exited.
func (l *listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	v, ok := l.values[ctx]
	if !ok {
		return
	}
	ast, ok := v.(ast.AST)
	if !ok {
		return
	}
	ast.SetPos(l.tokenToPos(ctx.GetStart()))
}

// start: module;
func (l *listener) ExitStart(ctx *parser.StartContext) {
	l.Module = l.takeAST(ctx.Module()).(*ast.Module)
}

// module: meta_section? stmts;
func (l *listener) ExitModule(ctx *parser.ModuleContext) {
	m := &ast.Module{}
	m.Stmts = l.takeASTs(ctx.Stmts())
	l.put(ctx, m)
}

// meta_section: META meta_entry* END;
func (l *listener) ExitMeta_section(ctx *parser.Meta_sectionContext) {}

// meta_entry: ID meta_object;
func (l *listener) ExitMeta_entry(ctx *parser.Meta_entryContext) {}

// meta_object: '{' meta_fields? '}';
func (l *listener) ExitMeta_object(ctx *parser.Meta_objectContext) {}

// meta_fields: meta_field ((','|';') meta_field)*;
func (l *listener) ExitMeta_fields(ctx *parser.Meta_fieldsContext) {}

// meta_field: ID ':' meta_field_value;
func (l *listener) ExitMeta_field(ctx *parser.Meta_fieldContext) {}

// meta_field_value: STRING | NUMBER | TRUE | FALSE | meta_object;
func (l *listener) ExitMeta_field_value(ctx *parser.Meta_field_valueContext) {}

// stmts: stmt*;
func (l *listener) ExitStmts(ctx *parser.StmtsContext) {
	allStmt := ctx.AllStmt()
	stmts := make(ast.ASTs, 0, len(allStmt))
	for _, stmtCtx := range allStmt {
		stmt := l.takeAST(stmtCtx)
		// `stmt` can be the empty statement. In this case it will have a parser
		// tree node, but the `stmt` rule will not have a value produced for it.
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
	}
	l.put(ctx, stmts)
}

// stmt: assignment_stmt ';'
//     | var_decl_stmt ';'
//     | for_stmt ';'
//     | while_stmt ';'
//     | if_stmt ';'
//     | func_stmt ';'
//     | return_stmt ';'
//     | primary_expr ';'
//     | ';'
//     ;
func (l *listener) ExitStmt(ctx *parser.StmtContext) {
	var s ast.AST
	if s = l.takeAST(ctx.Assignment_stmt()); s != nil {
	} else if s = l.takeAST(ctx.Var_decl_stmt()); s != nil {
	} else if s = l.takeAST(ctx.For_stmt()); s != nil {
	} else if s = l.takeAST(ctx.While_stmt()); s != nil {
	} else if s = l.takeAST(ctx.If_stmt()); s != nil {
	} else if s = l.takeAST(ctx.Func_stmt()); s != nil {
	} else if s = l.takeAST(ctx.Return_stmt()); s != nil {
	} else if e := l.takeExpr(ctx.Primary_expr()); e != nil {
		s = &ast.ExprStmt{Expr: e}
	}
	if s != nil {
		l.put(ctx, s)
	}
}

// assignment_stmt: assignment_lvalues '=' rvalues;
func (l *listener) ExitAssignment_stmt(ctx *parser.Assignment_stmtContext) {
	l.put(ctx, &ast.AssignStmt{
		Lvalues: l.takeASTs(ctx.Assignment_lvalues()),
		Rvalues: l.takeExprs(ctx.Rvalues()),
	})
}

// assignment_lvalues: lvalue_expr (',' lvalue_expr)*;
func (l *listener) ExitAssignment_lvalues(ctx *parser.Assignment_lvaluesContext) {
	all := ctx.AllLvalue_expr()
	asts := make(ast.ASTs, len(all))
	for i, child := range all {
		asts[i] = &ast.LValue{Expr: l.takeExpr(child)}
	}
	l.put(ctx, asts)
}

// rvalues: expr (',' expr)*;
func (l *listener) ExitRvalues(ctx *parser.RvaluesContext) {
	all := ctx.AllExpr()
	exprs := make(ast.Exprs, len(all))
	for i, child := range all {
		exprs[i] = l.takeExpr(child)
	}
	l.put(ctx, exprs)
}

// var_decl_stmt: VAR var_decl_vars ('=' rvalues)?;
func (l *listener) ExitVar_decl_stmt(ctx *parser.Var_decl_stmtContext) {
	vars, _ := l.take(ctx.Var_decl_vars())
	s := &ast.VarDeclStmt{
		Vars:       vars.([]token.Token),
		InitValues: l.takeExprs(ctx.Rvalues()),
	}
	l.put(ctx, s)
}

// var_decl_vars: ID (',' ID)*;
func (l *listener) ExitVar_decl_vars(ctx *parser.Var_decl_varsContext) {
	all := ctx.AllID()
	ids := make([]token.Token, len(all))
	for i, child := range all {
		ids[i] = l.tokenToNitro(child.GetSymbol())
	}
	l.put(ctx, ids)
}

// for_stmt: FOR for_vars IN expr DO stmts END;
func (l *listener) ExitFor_stmt(ctx *parser.For_stmtContext) {
	l.put(ctx, &ast.ForStmt{
		ForVars:  l.takeASTs(ctx.For_vars()),
		IterExpr: l.takeExpr(ctx.Expr()),
		Stmts:    l.takeASTs(ctx.Stmts()),
	})
}

// for_vars: ID (',' ID)*;
func (l *listener) ExitFor_vars(ctx *parser.For_varsContext) {
	all := ctx.AllID()
	vars := make(ast.ASTs, len(all))
	for i, child := range all {
		vars[i] = &ast.ForVar{VarName: l.tokenToNitro(child.GetSymbol())}
	}
	l.put(ctx, vars)
}

// while_stmt: WHILE expr DO stmts END;
func (l *listener) ExitWhile_stmt(ctx *parser.While_stmtContext) {
	l.put(ctx, &ast.WhileStmt{
		Predicate: l.takeExpr(ctx.Expr()),
		Stmts:     l.takeASTs(ctx.Stmts()),
	})
}

// if_stmt: IF expr THEN stmts if_elif* if_else?;
func (l *listener) ExitIf_stmt(ctx *parser.If_stmtContext) {
	ifBlock := &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Stmts: l.takeASTs(ctx.Stmts()),
	}

	allElifs := ctx.AllIf_elif()
	blocks := make(ast.ASTs, 0, len(allElifs)+2)
	blocks = append(blocks, ifBlock)
	for _, elif := range allElifs {
		blocks = append(blocks, l.takeAST(elif))
	}
	if ctx.If_else() != nil {
		blocks = append(blocks, l.takeAST(ctx.If_else()))
	}

	l.put(ctx, &ast.IfStmt{
		Blocks: blocks,
	})
}

// if_elif: ELIF expr THEN stmts;
func (l *listener) ExitIf_elif(ctx *parser.If_elifContext) {
	l.put(ctx, &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Stmts: l.takeASTs(ctx.Stmts()),
	})
}

// if_else: ELSE stmts END;
func (l *listener) ExitIf_else(ctx *parser.If_elseContext) {
	l.put(ctx, &ast.IfBlock{
		Stmts: l.takeASTs(ctx.Stmts()),
	})
}

// func_stmt: FN ID '(' param_list ')' stmts END;
func (l *listener) ExitFunc_stmt(ctx *parser.Func_stmtContext) {
	f := &ast.FuncStmt{}
	f.Name = ctx.ID().GetText()
	f.Params = l.takeASTs(ctx.Param_list())
	f.Stmts = l.takeASTs(ctx.Stmts())
	l.put(ctx, f)
}

// param_list: ID (',' ID)*;
func (l *listener) ExitParam_list(ctx *parser.Param_listContext) {
	all := ctx.AllID()
	params := make(ast.ASTs, len(all))
	for i, child := range all {
		params[i] = &ast.FuncParam{
			Name: child.GetText(),
		}
		params[i].SetPos(l.tokenToPos(child.GetSymbol()))
	}
	l.put(ctx, params)
}

// return_stmt: RETURN rvalues?;
func (l *listener) ExitReturn_stmt(ctx *parser.Return_stmtContext) {
	l.put(ctx, &ast.ReturnStmt{
		Values: l.takeExprs(ctx.Rvalues()),
	})
}

// expr: unary_expr
//     | expr op=('*'|'/'|'%') expr
//     | expr op=('+'|'-') expr
//     | expr op=('<'|'<='|'>'|'>='|'=='|'!=') expr
//     | expr op=AND expr
//     | expr op=OR expr
//     ;
func (l *listener) ExitExpr(ctx *parser.ExprContext) {
	if ctx.Unary_expr() != nil {
		l.put(ctx, l.takeExpr(ctx.Unary_expr()))
		return
	}

	left := l.takeExpr(ctx.Expr(0))
	right := l.takeExpr(ctx.Expr(1))

	simpleBinExp := func(op ast.Operator) {
		l.put(ctx, &ast.BinaryExpr{Left: left, Right: right, Op: op})
	}

	switch ctx.GetOp().GetTokenType() {
	case parser.NitroLexerMUL:
		simpleBinExp(ast.BinOpMult)
	case parser.NitroLexerDIV:
		simpleBinExp(ast.BinOpDiv)
	case parser.NitroLexerMOD:
		simpleBinExp(ast.BinOpMod)
	case parser.NitroLexerADD:
		simpleBinExp(ast.BinOpPlus)
	case parser.NitroLexerSUB:
		simpleBinExp(ast.BinOpMinus)
	case parser.NitroLexerLT:
		simpleBinExp(ast.BinOpLT)
	case parser.NitroLexerLE:
		simpleBinExp(ast.BinOpLE)
	case parser.NitroLexerGT:
		simpleBinExp(ast.BinOpGT)
	case parser.NitroLexerGE:
		simpleBinExp(ast.BinOpGE)
	case parser.NitroLexerEQ:
		simpleBinExp(ast.BinOpEq)
	case parser.NitroLexerNE:
		simpleBinExp(ast.BinOpNE)
	case parser.NitroLexerAND:
		l.put(ctx, &ast.AndExpr{Left: left, Right: right})
	case parser.NitroLexerOR:
		l.put(ctx, &ast.OrExpr{Left: left, Right: right})
	default:
		log.Panicf("Invalid operator %v", ctx.GetOp().GetTokenType())
	}
}

// unary_expr: op=NOT unary_expr
//           | op='+' unary_expr
//           | op='-' unary_expr
//           | primary_expr
//           ;
func (l *listener) ExitUnary_expr(ctx *parser.Unary_exprContext) {
	if ctx.Primary_expr() != nil {
		l.put(ctx, l.takeExpr(ctx.Primary_expr()))
		return
	}

	var op ast.UnaryOp
	switch ctx.GetOp().GetTokenType() {
	case parser.NitroLexerNOT:
		op = ast.UnaryOpNot
	case parser.NitroLexerADD:
		op = ast.UnaryOpPlus
	case parser.NitroLexerSUB:
		op = ast.UnaryOpMinus
	default:
		log.Panicf("Invalid operator %v", ctx.GetOp().GetTokenType())
	}

	l.put(ctx, &ast.UnaryExpr{
		Term: l.takeExpr(ctx.Unary_expr()),
		Op:   op,
	})
}

// primary_expr: ID                                       # primary_expr_simple_ref
//             | primary_expr '.' ID                      # primary_expr_member_access
//             | primary_expr '[' expr ']'                # primary_expr_index
//             | primary_expr '[' expr? ':' expr? ']'     # primary_expr_slice
//             | primary_expr '(' arg_list? ')'           # primary_expr_call
//             | lambda_expr                              # primary_expr_lambda
//             | object_literal                           # primary_expr_object
//             | array_literal                            # primary_expr_array
//             | simple_literal                           # primary_expr_literal
//             | '(' expr ')'                             # primary_expr_parenthesis
//             ;

func (l *listener) ExitPrimary_expr_simple_ref(ctx *parser.Primary_expr_simple_refContext) {
	l.put(ctx, &ast.SimpleRef{
		ID: l.tokenToNitro(ctx.ID().GetSymbol()),
	})
}
func (l *listener) ExitPrimary_expr_member_access(ctx *parser.Primary_expr_member_accessContext) {
	l.put(ctx, &ast.MemberAccess{
		Target: l.takeExpr(ctx.Primary_expr()),
		Member: l.tokenToNitro(ctx.ID().GetSymbol()),
	})
}

func (l *listener) ExitPrimary_expr_index(ctx *parser.Primary_expr_indexContext) {
	l.put(ctx, &ast.IndexExpr{
		Target: l.takeExpr(ctx.Primary_expr()),
		Index:  l.takeExpr(ctx.Expr()),
	})
}

func (l *listener) ExitPrimary_expr_slice(ctx *parser.Primary_expr_sliceContext) {
	l.put(ctx, &ast.SliceExpr{
		Target: l.takeExpr(ctx.Primary_expr()),
		Begin:  l.takeExpr(ctx.GetB()),
		End:    l.takeExpr(ctx.GetE()),
	})
}

func (l *listener) ExitPrimary_expr_call(ctx *parser.Primary_expr_callContext) {
	l.put(ctx, &ast.FuncCallExpr{
		Target: l.takeExpr(ctx.Primary_expr()),
		Args:   l.takeExprs(ctx.Arg_list()),
		RetN:   1,
	})
}

func (l *listener) ExitPrimary_expr_lambda(ctx *parser.Primary_expr_lambdaContext) {
	l.put(ctx, l.takeExpr(ctx.Lambda_expr()))
}

func (l *listener) ExitPrimary_expr_object(ctx *parser.Primary_expr_objectContext) {
	l.put(ctx, l.takeExpr(ctx.Object_literal()))
}

func (l *listener) ExitPrimary_expr_array(ctx *parser.Primary_expr_arrayContext) {
	l.put(ctx, l.takeExpr(ctx.Array_literal()))
}

func (l *listener) ExitPrimary_expr_literal(ctx *parser.Primary_expr_literalContext) {
	l.put(ctx, l.takeExpr(ctx.Simple_literal()))
}

func (l *listener) ExitPrimary_expr_parenthesis(ctx *parser.Primary_expr_parenthesisContext) {
	l.put(ctx, l.takeExpr(ctx.Expr()))
}

// simple_literal: val=(STRING | NUMBER | TRUE | FALSE);
func (l *listener) ExitSimple_literal(ctx *parser.Simple_literalContext) {
	l.put(ctx, &ast.LiteralExpr{
		Val: l.tokenToNitro(ctx.GetVal()),
	})
}

// arg_list: expr (',' expr)*;
func (l *listener) ExitArg_list(ctx *parser.Arg_listContext) {
	all := ctx.AllExpr()
	exprs := make(ast.Exprs, len(all))
	for i, child := range all {
		exprs[i] = l.takeExpr(child)
	}
	l.put(ctx, exprs)
}

// lvalue_expr: ID                          # lvalue_expr_simple_ref
//            | primary_expr '.' ID         # lvalue_expr_member_access
//            | primary_expr '[' expr ']'   # lvalue_expr_index
//            ;

func (l *listener) ExitLvalue_expr_simple_ref(ctx *parser.Lvalue_expr_simple_refContext) {
	l.put(ctx, &ast.SimpleRef{
		ID: l.tokenToNitro(ctx.ID().GetSymbol()),
	})
}

func (l *listener) ExitLvalue_expr_member_access(ctx *parser.Lvalue_expr_member_accessContext) {
	l.put(ctx, &ast.MemberAccess{
		Target: l.takeExpr(ctx.Primary_expr()),
		Member: l.tokenToNitro(ctx.ID().GetSymbol()),
	})
}

func (l *listener) ExitLvalue_expr_index(ctx *parser.Lvalue_expr_indexContext) {
	l.put(ctx, &ast.IndexExpr{
		Target: l.takeExpr(ctx.Primary_expr()),
		Index:  l.takeExpr(ctx.Expr()),
	})
}

// lambda_expr: FN '(' param_list? ')' stmts END;
func (l *listener) ExitLambda_expr(ctx *parser.Lambda_exprContext) {
	lambda := &ast.LambdaExpr{}
	lambda.Params = l.takeASTs(ctx.Param_list())
	lambda.Stmts = l.takeASTs(ctx.Stmts())
	l.put(ctx, lambda)
}

// object_literal: '{' object_fields? '}';
func (l *listener) ExitObject_literal(ctx *parser.Object_literalContext) {
	l.put(ctx, &ast.ObjectLiteral{
		Fields: l.takeASTs(ctx.Object_fields()),
	})
}

// object_fields: object_field ((','|';') object_field)* (','|';')*;
func (l *listener) ExitObject_fields(ctx *parser.Object_fieldsContext) {
	allFields := ctx.AllObject_field()
	fields := make(ast.ASTs, len(allFields))
	for i, entry := range allFields {
		fields[i] = l.takeAST(entry)
	}
	l.put(ctx, fields)
}

// object_field: id_or_keyword ':' expr     # object_field_id_key
//             | '[' expr ']' ':' expr      # object_field_expr_key
//             | object_if                  # object_field_if
//             | object_for                 # object_field_for
//             ;

func (l *listener) ExitObject_field_id_key(ctx *parser.Object_field_id_keyContext) {
	idOrKeyw, _ := l.take(ctx.Id_or_keyword())
	l.put(ctx, &ast.ObjectField{
		NameID: idOrKeyw.(string),
		Val:    l.takeExpr(ctx.Expr()),
	})
}

func (l *listener) ExitObject_field_expr_key(ctx *parser.Object_field_expr_keyContext) {
	l.put(ctx, &ast.ObjectField{
		NameExpr: l.takeExpr(ctx.Expr(0)),
		Val:      l.takeExpr(ctx.Expr(1)),
	})
}

func (l *listener) ExitObject_field_if(ctx *parser.Object_field_ifContext) {
	l.put(ctx, l.takeAST(ctx.Object_if()))
}

func (l *listener) ExitObject_field_for(ctx *parser.Object_field_forContext) {
	l.put(ctx, l.takeAST(ctx.Object_for()))
}

// object_if: IF expr THEN object_fields? object_elif* object_else? END;
func (l *listener) ExitObject_if(ctx *parser.Object_ifContext) {
	ifBlock := &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Stmts: l.takeASTs(ctx.Object_fields()),
	}

	allElifs := ctx.AllObject_elif()
	blocks := make(ast.ASTs, 0, len(allElifs)+2)
	blocks = append(blocks, ifBlock)
	for _, elif := range allElifs {
		blocks = append(blocks, l.takeAST(elif))
	}
	if ctx.Object_else() != nil {
		blocks = append(blocks, l.takeAST(ctx.Object_else()))
	}

	l.put(ctx, &ast.IfStmt{
		Blocks: blocks,
	})
}

// object_elif: ELIF expr THEN object_fields?;
func (l *listener) ExitObject_elif(ctx *parser.Object_elifContext) {
	l.put(ctx, &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Stmts: l.takeASTs(ctx.Object_fields()),
	})
}

// object_else: ELSE object_fields?;
func (l *listener) ExitObject_else(ctx *parser.Object_elseContext) {
	l.put(ctx, &ast.IfBlock{
		Stmts: l.takeASTs(ctx.Object_fields()),
	})
}

// object_for: FOR for_vars IN expr DO object_fields? END;
func (l *listener) ExitObject_for(ctx *parser.Object_forContext) {
	l.put(ctx, &ast.ForStmt{
		ForVars:  l.takeASTs(ctx.For_vars()),
		IterExpr: l.takeExpr(ctx.Expr()),
		Stmts:    l.takeASTs(ctx.Object_fields()),
	})
}

// array_literal: '[' array_elems? ']';
func (l *listener) ExitArray_literal(ctx *parser.Array_literalContext) {
	l.put(ctx, &ast.ArrayLiteral{
		Elements: l.takeASTs(ctx.Array_elems()),
	})
}

// array_elems: array_elem ((','|';') array_elem)* (','|';')*;
func (l *listener) ExitArray_elems(ctx *parser.Array_elemsContext) {
	all := ctx.AllArray_elem()
	elems := make(ast.ASTs, len(all))
	for i, entry := range all {
		elems[i] = l.takeAST(entry)
	}
	l.put(ctx, elems)
}

// array_elem: expr | array_if | array_for;
func (l *listener) ExitArray_elem(ctx *parser.Array_elemContext) {
	switch {
	case ctx.Expr() != nil:
		l.put(ctx, &ast.ArrayElement{Val: l.takeExpr(ctx.Expr())})
	case ctx.Array_if() != nil:
		l.put(ctx, l.takeAST(ctx.Array_if()))
	case ctx.Array_for() != nil:
		l.put(ctx, l.takeAST(ctx.Array_for()))
	default:
		panic("unreachable")
	}
}

// array_if: IF expr THEN array_elems? array_elif* array_else? END;
func (l *listener) ExitArray_if(ctx *parser.Array_ifContext) {
	ifBlock := &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Stmts: l.takeASTs(ctx.Array_elems()),
	}

	allElifs := ctx.AllArray_elif()
	blocks := make(ast.ASTs, 0, len(allElifs)+2)
	blocks = append(blocks, ifBlock)
	for _, elif := range allElifs {
		blocks = append(blocks, l.takeAST(elif))
	}
	if ctx.Array_else() != nil {
		blocks = append(blocks, l.takeAST(ctx.Array_else()))
	}
	l.put(ctx, &ast.IfStmt{
		Blocks: blocks,
	})
}

// array_elif: ELIF expr THEN array_elems?;
func (l *listener) ExitArray_elif(ctx *parser.Array_elifContext) {
	l.put(ctx, &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Stmts: l.takeASTs(ctx.Array_elems()),
	})
}

// array_else: ELSE array_elems?;
func (l *listener) ExitArray_else(ctx *parser.Array_elseContext) {
	l.put(ctx, &ast.IfBlock{
		Stmts: l.takeASTs(ctx.Array_elems()),
	})
}

// array_for: FOR for_vars IN expr DO array_elems? END;
func (l *listener) ExitArray_for(ctx *parser.Array_forContext) {
	l.put(ctx, &ast.ForStmt{
		ForVars:  l.takeASTs(ctx.For_vars()),
		IterExpr: l.takeExpr(ctx.Expr()),
		Stmts:    l.takeASTs(ctx.Array_elems()),
	})
}

func (l *listener) ExitId_or_keyword(ctx *parser.Id_or_keywordContext) {
	id := ctx.GetT().GetText()
	l.put(ctx, id)
}