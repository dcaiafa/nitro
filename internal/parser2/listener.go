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

func antlrTokenToNitro(at antlr.Token) token.Token {
	t := token.Token{}
	switch at.GetTokenType() {
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
	t.Pos = antlrTokenToPos(at)
	return t
}

func antlrTokenToPos(at antlr.Token) token.Pos {
	return token.Pos{
		Col:  at.GetColumn(),
		Line: at.GetLine(),
	}
	// TODO: filename
}

type listener struct {
	parser.BaseNitroListener

	values map[antlr.RuleContext]interface{}
	module *ast.Module
}

func newListener() *listener {
	return &listener{
		values: make(map[antlr.RuleContext]interface{}),
	}
}

func (l *listener) put(k antlr.RuleContext, v interface{}) {
	l.values[k] = v
}

func (l *listener) take(k antlr.RuleContext) (interface{}, bool) {
	if k == nil {
		return nil, false
	}
	v := l.values[k]
	if v != nil {
		delete(l.values, k)
	}
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
	ast.SetPos(antlrTokenToPos(ctx.GetStart()))
}

// start: module;
func (l *listener) ExitStart(ctx *parser.StartContext) {
	l.module = l.takeAST(ctx.Module()).(*ast.Module)
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
	stmts := make(ast.ASTs, len(allStmt))
	for i, stmtCtx := range allStmt {
		stmts[i] = l.takeAST(stmtCtx)
	}
	l.put(ctx, stmts)
}

// stmt: assignment_stmt ';'
//     | var_decl_stmt ';'
//     | for_stmt ';'
//     | while_stmt ';'
//     | if_stmt ';'
//     | func_stmt ';'
//     | func_call_stmt ';'
//     | return_stmt ';'
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
	} else if s = l.takeAST(ctx.Func_call_stmt()); s != nil {
	} else if s = l.takeAST(ctx.Return_stmt()); s != nil {
	}
	if s != nil {
		l.put(ctx, s)
	}
}

// assignment_stmt: assignment_lvalues '=' assignment_rvalues;
func (l *listener) ExitAssignment_stmt(ctx *parser.Assignment_stmtContext) {
	l.put(ctx, &ast.AssignStmt{
		Lvalues: l.takeASTs(ctx.Assignment_lvalues()),
		Rvalues: l.takeExprs(ctx.Assignment_rvalues()),
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

// assignment_rvalues: expr (',' expr)*;
func (l *listener) ExitAssignment_rvalues(ctx *parser.Assignment_rvaluesContext) {
	all := ctx.AllExpr()
	exprs := make(ast.Exprs, len(all))
	for i, child := range all {
		exprs[i] = l.takeExpr(child)
	}
	l.put(ctx, exprs)
}

// var_decl_stmt: VAR var_decl_vars ('=' assignment_rvalues)?;
func (l *listener) ExitVar_decl_stmt(ctx *parser.Var_decl_stmtContext) {
	vars, _ := l.take(ctx.Var_decl_vars())
	s := &ast.VarDeclStmt{
		Vars:       vars.([]token.Token),
		InitValues: l.takeExprs(ctx.Assignment_rvalues()),
	}
	l.put(ctx, s)
}

// var_decl_vars: ID (',' ID)*;
func (l *listener) ExitVar_decl_vars(ctx *parser.Var_decl_varsContext) {
	all := ctx.AllID()
	ids := make([]token.Token, len(all))
	for i, child := range all {
		ids[i] = antlrTokenToNitro(child.GetSymbol())
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
		vars[i] = &ast.ForVar{VarName: antlrTokenToNitro(child.GetSymbol())}
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
		params[i].SetPos(antlrTokenToPos(child.GetSymbol()))
	}
	l.put(ctx, params)
}

// func_call_stmt: primary_expr '(' arg_list? ')';
func (l *listener) ExitFunc_call_stmt(ctx *parser.Func_call_stmtContext) {
	l.put(ctx, &ast.FuncCallStmt{
		Target: l.takeExpr(ctx.Primary_expr()),
		Args:   l.takeExprs(ctx.Arg_list()),
	})
}

// return_stmt: RETURN expr*;
func (l *listener) ExitReturn_stmt(ctx *parser.Return_stmtContext) {
	all := ctx.AllExpr()
	exprs := make(ast.Exprs, len(all))
	for i, child := range all {
		exprs[i] = l.takeExpr(child)
	}
	l.put(ctx, &ast.ReturnStmt{
		Values: exprs,
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

func (l *listener) ExitPrimary_expr(ctx *parser.Primary_exprContext) {}

func (l *listener) ExitArg_list(ctx *parser.Arg_listContext) {}

func (l *listener) ExitLvalue_expr(ctx *parser.Lvalue_exprContext) {}

func (l *listener) ExitObject_literal(ctx *parser.Object_literalContext) {}

func (l *listener) ExitObject_fields(ctx *parser.Object_fieldsContext) {}

func (l *listener) ExitObject_field(ctx *parser.Object_fieldContext) {}

func (l *listener) ExitObject_if(ctx *parser.Object_ifContext) {}

func (l *listener) ExitObject_elif(ctx *parser.Object_elifContext) {}

func (l *listener) ExitObject_else(ctx *parser.Object_elseContext) {}

func (l *listener) ExitObject_for(ctx *parser.Object_forContext) {}

func (l *listener) ExitArray_literal(ctx *parser.Array_literalContext) {}

func (l *listener) ExitArray_elems(ctx *parser.Array_elemsContext) {}

func (l *listener) ExitArray_elem(ctx *parser.Array_elemContext) {}

func (l *listener) ExitArray_if(ctx *parser.Array_ifContext) {}

func (l *listener) ExitArray_elif(ctx *parser.Array_elifContext) {}

func (l *listener) ExitArray_else(ctx *parser.Array_elseContext) {}

func (l *listener) ExitArray_for(ctx *parser.Array_forContext) {}

func (l *listener) ExitId_or_keyword(ctx *parser.Id_or_keywordContext) {}
