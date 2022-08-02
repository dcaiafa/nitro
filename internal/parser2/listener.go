package parser2

import (
	"log"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/parser2/parser"
	"github.com/dcaiafa/nitro/internal/token"
	"github.com/dcaiafa/nitro/internal/vm"
)

type listener struct {
	parser.BaseNitroParserListener

	Unit *ast.Unit

	filename    string
	errListener *errorListener
	values      map[antlr.RuleContext]interface{}
}

func newListener(filename string, errListener *errorListener) *listener {
	return &listener{
		filename:    filename,
		errListener: errListener,
		values:      make(map[antlr.RuleContext]interface{}),
	}
}

func (l *listener) tokenToNitro(at antlr.Token) token.Token {
	var err error

	t := token.Token{}
	switch at.GetTokenType() {
	case parser.NitroLexerNIL:
		t.Type = token.Nil
	case parser.NitroLexerSTRING,
		parser.NitroLexerEXEC_DQUOTE_LITERAL,
		parser.NitroLexerEXEC_SQUOTE_LITERAL:
		s := at.GetText()
		s = s[1 : len(s)-1] // remove quotes
		t.Type = token.String
		t.Str = s
		t.Str, err = expandEscapeSequences(s)
		if err != nil {
			l.errListener.LogError(
				at.GetLine(),
				at.GetColumn(),
				"Invalid string literal: %v", err)
		}
	case parser.NitroLexerCHAR:
		s := at.GetText()
		s = s[1 : len(s)-1] // remove quotes
		r, err := expandCharEscapeSequences(s)
		if err != nil {
			l.errListener.LogError(
				at.GetLine(),
				at.GetColumn(),
				"Invalid character literal: %v", err)
		}
		t.Type = token.Int
		t.Int = int64(r)
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

// start: unit EOF;
func (l *listener) ExitStart(ctx *parser.StartContext) {
	l.Unit = l.takeAST(ctx.Unit()).(*ast.Unit)
}

// prologue: meta_directive* import_stmt*;
func (l *listener) ExitPrologue(ctx *parser.PrologueContext) {
	p := &Prologue{}

	allMeta := ctx.AllMeta_directive()
	p.Meta = make(ast.ASTs, 0, len(allMeta))
	for _, meta := range allMeta {
		p.Meta = append(p.Meta, l.takeAST(meta))
	}
	allImport := ctx.AllImport_stmt()
	p.Imports = make(ast.ASTs, 0, len(allImport))
	for _, imp := range allImport {
		p.Imports = append(p.Imports, l.takeAST(imp))
	}
	l.put(ctx, p)
}

// unit: prologue stmts;
func (l *listener) ExitUnit(ctx *parser.UnitContext) {
	m := &ast.Unit{}

	pRaw, _ := l.take(ctx.Prologue())
	p := pRaw.(*Prologue)

	m.Meta = p.Meta
	m.Imports = p.Imports
	block := l.takeAST(ctx.Stmts()).(*ast.StmtBlock)
	m.Block = block.Stmts

	l.put(ctx, m)
}

func (l *listener) ExitMeta_directive(ctx *parser.Meta_directiveContext) {
	var ast ast.AST
	if ctx.Meta_info() != nil {
		ast = l.takeAST(ctx.Meta_info())
	} else if ctx.Meta_param() != nil {
		ast = l.takeAST(ctx.Meta_param())
	} else if ctx.Meta_flag() != nil {
		ast = l.takeAST(ctx.Meta_flag())
	} else {
		panic("not reached")
	}
	l.put(ctx, ast)
}

func (l *listener) ExitMeta_info(ctx *parser.Meta_infoContext) {
	l.errListener.LogError(
		ctx.M_INFO().GetSymbol().GetLine(),
		ctx.M_INFO().GetSymbol().GetColumn(),
		"Info metadata not implemented")
}

func (l *listener) ExitMeta_param(ctx *parser.Meta_paramContext) {
	param := &ast.MetaParam{
		Name:   ctx.ID().GetText(),
		IsFlag: false,
	}
	if ctx.Expr() != nil {
		param.Default = l.takeExpr(ctx.Expr())
	}
	if ctx.Meta_attribs() != nil {
		param.Attribs = l.takeASTs(ctx.Meta_attribs())
	}
	l.put(ctx, param)
}

func (l *listener) ExitMeta_flag(ctx *parser.Meta_flagContext) {
	param := &ast.MetaParam{
		Name:   ctx.ID().GetText(),
		IsFlag: true,
	}
	if ctx.Expr() != nil {
		param.Default = l.takeExpr(ctx.Expr())
	}
	if ctx.Meta_attribs() != nil {
		param.Attribs = l.takeASTs(ctx.Meta_attribs())
	}
	l.put(ctx, param)
}

func (l *listener) ExitMeta_attribs(ctx *parser.Meta_attribsContext) {
	allAttribs := ctx.AllMeta_attrib()
	attribs := make(ast.ASTs, 0, len(allAttribs))
	for _, attrib := range allAttribs {
		attribs = append(attribs, l.takeAST(attrib))
	}
	l.put(ctx, attribs)
}

func (l *listener) ExitMeta_attrib(ctx *parser.Meta_attribContext) {
	idOrKeyw, _ := l.take(ctx.Id_or_keyword())
	attrib := &ast.MetaAttrib{
		Name: idOrKeyw.(string),
	}
	if ctx.Meta_literal() != nil {
		v, _ := l.take(ctx.Meta_literal())
		attrib.Value = v.(vm.Value)
	}
	l.put(ctx, attrib)
}

func (l *listener) ExitMeta_literal(ctx *parser.Meta_literalContext) {
	var v vm.Value
	t := l.tokenToNitro(ctx.GetVal())
	switch t.Type {
	case token.Nil:
		v = nil
	case token.Int:
		v = vm.NewInt(t.Int)
	case token.Float:
		v = vm.NewFloat(t.Float)
	case token.String:
		v = vm.NewString(t.Str)
	case token.Bool:
		v = vm.NewBool(t.Bool)
	default:
		panic("unreachable")
	}
	l.put(ctx, v)
}

// import_stmt: IMPORT ID? STRING ';';
func (l *listener) ExitImport_stmt(ctx *parser.Import_stmtContext) {
	alias := ""
	if ctx.ID() != nil {
		alias = l.tokenToNitro(ctx.ID().GetSymbol()).Str
	}
	importName := l.tokenToNitro(ctx.STRING().GetSymbol()).Str
	l.put(ctx, &ast.Import{
		Alias:   alias,
		Package: importName,
	})
}

// stmts: stmt_list?;
func (l *listener) ExitStmts(ctx *parser.StmtsContext) {
	if ctx.Stmt_list() == nil {
		l.put(ctx, &ast.StmtBlock{})
		return
	}
	l.put(ctx, l.takeAST(ctx.Stmt_list()))
}

// stmts: stmt*;
func (l *listener) ExitStmt_list(ctx *parser.Stmt_listContext) {
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
	block := &ast.StmtBlock{
		Stmts: stmts,
	}
	l.put(ctx, block)
}

// stmt: assignment_stmt      # stmt_assignment
//     | assignment_op_stmt   # stmt_op_assign
//     | var_decl_stmt        # stmt_var_dec
//     | for_stmt             # stmt_for
//     | while_stmt           # stmt_while
//     | if_stmt              # stmt_if
//     | func_stmt            # stmt_func
//     | return_stmt          # stmt_return
//     | expr                 # stmt_expr
//     | try_catch_stmt       # stmt_try_catch
//     | throw_stmt           # stmt_throw
//     | defer_stmt           # stmt_defer
//     | yield_stmt           # stmt_yield
//     | break_stmt           # stmt_break
//     | continue_stmt        # stmt_continue
//     | inc_dec_stmt         # stmt_inc_dec
// ;

func (l *listener) ExitStmt_assignment(ctx *parser.Stmt_assignmentContext) {
	l.put(ctx, l.takeAST(ctx.Assignment_stmt()))
}

func (l *listener) ExitStmt_op_assign(ctx *parser.Stmt_op_assignContext) {
	l.put(ctx, l.takeAST(ctx.Assignment_op_stmt()))
}

func (l *listener) ExitStmt_var_dec(ctx *parser.Stmt_var_decContext) {
	l.put(ctx, l.takeAST(ctx.Var_decl_stmt()))
}

func (l *listener) ExitStmt_for(ctx *parser.Stmt_forContext) {
	l.put(ctx, l.takeAST(ctx.For_stmt()))
}

func (l *listener) ExitStmt_while(ctx *parser.Stmt_whileContext) {
	l.put(ctx, l.takeAST(ctx.While_stmt()))
}

func (l *listener) ExitStmt_if(ctx *parser.Stmt_ifContext) {
	l.put(ctx, l.takeAST(ctx.If_stmt()))
}

func (l *listener) ExitStmt_func(ctx *parser.Stmt_funcContext) {
	l.put(ctx, l.takeAST(ctx.Func_stmt()))
}

func (l *listener) ExitStmt_return(ctx *parser.Stmt_returnContext) {
	l.put(ctx, l.takeAST(ctx.Return_stmt()))
}

func (l *listener) ExitStmt_expr(ctx *parser.Stmt_exprContext) {
	e := l.takeExpr(ctx.Expr())
	if objLit, ok := e.(*ast.ObjectLiteral); ok {
		emit := token.Token{
			Pos:  objLit.Pos(),
			Type: token.String,
			Str:  "emit",
		}
		emitRef := &ast.SimpleRef{ID: emit}
		emitRef.SetPos(objLit.Pos())
		e = &ast.FuncCallExpr{
			Target: emitRef,
			Args:   ast.Exprs{objLit},
			RetN:   0,
		}
		e.SetPos(objLit.Pos())
	}
	l.put(ctx, &ast.ExprStmt{Expr: e})
}

func (l *listener) ExitStmt_try_catch(ctx *parser.Stmt_try_catchContext) {
	l.put(ctx, l.takeAST(ctx.Try_catch_stmt()))
}

func (l *listener) ExitStmt_throw(ctx *parser.Stmt_throwContext) {
	l.put(ctx, l.takeAST(ctx.Throw_stmt()))
}

func (l *listener) ExitStmt_defer(ctx *parser.Stmt_deferContext) {
	l.put(ctx, l.takeAST(ctx.Defer_stmt()))
}

func (l *listener) ExitStmt_yield(ctx *parser.Stmt_yieldContext) {
	l.put(ctx, l.takeAST(ctx.Yield_stmt()))
}

func (l *listener) ExitStmt_break(ctx *parser.Stmt_breakContext) {
	l.put(ctx, l.takeAST(ctx.Break_stmt()))
}

func (l *listener) ExitStmt_continue(ctx *parser.Stmt_continueContext) {
	l.put(ctx, l.takeAST(ctx.Continue_stmt()))
}

func (l *listener) ExitStmt_inc_dec(ctx *parser.Stmt_inc_decContext) {
	l.put(ctx, l.takeAST(ctx.Inc_dec_stmt()))
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

func (l *listener) ExitAssignment_op_stmt(ctx *parser.Assignment_op_stmtContext) {
	assignOpAST := &ast.AssignOpStmt{
		LValue: &ast.LValue{Expr: l.takeExpr(ctx.Lvalue_expr())},
		RValue: l.takeExpr(ctx.Expr()),
	}

	switch ctx.GetOp().GetTokenType() {
	case parser.NitroLexerASSIGN_ADD:
		assignOpAST.Op = ast.BinOpPlus
	case parser.NitroLexerASSIGN_SUB:
		assignOpAST.Op = ast.BinOpMinus
	case parser.NitroLexerASSIGN_MUL:
		assignOpAST.Op = ast.BinOpMult
	case parser.NitroLexerASSIGN_DIV:
		assignOpAST.Op = ast.BinOpDiv
	case parser.NitroLexerASSIGN_MOD:
		assignOpAST.Op = ast.BinOpMod
	default:
		log.Panicf("Invalid operator %v", ctx.GetOp().GetTokenType())
	}

	l.put(ctx, assignOpAST)
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

// for_stmt: FOR for_vars ID expr DO stmts END;
func (l *listener) ExitFor_stmt(ctx *parser.For_stmtContext) {
	if ctx.ID().GetText() != "in" {
		l.errListener.LogError(
			ctx.ID().GetSymbol().GetLine(),
			ctx.ID().GetSymbol().GetColumn(),
			"Unexpected identifier %v. Expected 'in' instead.",
			ctx.ID().GetText())
		return
	}

	l.put(ctx, &ast.ForStmt{
		ForVars:  l.takeASTs(ctx.For_vars()),
		IterExpr: l.takeExpr(ctx.Expr()),
		Block:    l.takeAST(ctx.Stmts()),
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

// while_stmt: WHILE expr '{' stmts '}';
func (l *listener) ExitWhile_stmt(ctx *parser.While_stmtContext) {
	l.put(ctx, &ast.WhileStmt{
		Predicate: l.takeExpr(ctx.Expr()),
		Block:     l.takeAST(ctx.Stmts()).(*ast.StmtBlock),
	})
}

// if_stmt: IF expr THEN stmts if_elif* if_else?;
func (l *listener) ExitIf_stmt(ctx *parser.If_stmtContext) {
	ifBlock := &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Block: l.takeAST(ctx.Stmts()),
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
		Sections: blocks,
	})
}

// if_elif: ELIF expr THEN stmts;
func (l *listener) ExitIf_elif(ctx *parser.If_elifContext) {
	l.put(ctx, &ast.IfBlock{
		Pred:  l.takeExpr(ctx.Expr()),
		Block: l.takeAST(ctx.Stmts()),
	})
}

// if_else: ELSE stmts END;
func (l *listener) ExitIf_else(ctx *parser.If_elseContext) {
	l.put(ctx, &ast.IfBlock{
		Block: l.takeAST(ctx.Stmts()),
	})
}

// func_stmt: FN ID '(' param_list ')' stmts END;
func (l *listener) ExitFunc_stmt(ctx *parser.Func_stmtContext) {
	f := &ast.FuncStmt{}
	f.Name = ctx.ID().GetText()
	f.Params = l.takeASTs(ctx.Param_list())
	f.Block = l.takeAST(ctx.Stmts()).(*ast.StmtBlock)
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

// try_catch_stmt: TRY stmts CATCH ID? ';' stmts END;
func (l *listener) ExitTry_catch_stmt(ctx *parser.Try_catch_stmtContext) {
	var id *token.Token
	if ctx.ID() != nil {
		id = new(token.Token)
		*id = l.tokenToNitro(ctx.ID().GetSymbol())
	}
	l.put(ctx, ast.NewTryCatchStmt(
		l.takeAST(ctx.Stmts(0)), id, l.takeAST(ctx.Stmts(1))))
}

// throw_stmt: THROW expr;
func (l *listener) ExitThrow_stmt(ctx *parser.Throw_stmtContext) {
	l.put(ctx, &ast.ThrowStmt{
		Expr: l.takeExpr(ctx.Expr()),
	})
}

// defer_stmt: DEFER primary_expr;
func (l *listener) ExitDefer_stmt(ctx *parser.Defer_stmtContext) {
	primaryExpr := l.takeExpr(ctx.Primary_expr())
	callExpr, ok := primaryExpr.(*ast.FuncCallExpr)
	if !ok {
		l.errListener.LogError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Deferred expression must be a call")
	}
	l.put(ctx, ast.NewDeferStmt(callExpr))
}

// yield_stmt: YIELD rvalues;
func (l *listener) ExitYield_stmt(ctx *parser.Yield_stmtContext) {
	l.put(ctx, &ast.YieldStmt{
		Values: l.takeExprs(ctx.Rvalues()),
	})
}

// break_stmt: BREAK;
func (l *listener) ExitBreak_stmt(ctx *parser.Break_stmtContext) {
	l.put(ctx, &ast.BreakStmt{})
}

// continue_stmt: CONTINUE;
func (l *listener) ExitContinue_stmt(ctx *parser.Continue_stmtContext) {
	l.put(ctx, &ast.ContinueStmt{})
}

// inc_dec_stmt: lvalue_expr op=('++'|'--');
func (l *listener) ExitInc_dec_stmt(ctx *parser.Inc_dec_stmtContext) {
	if ctx.GetOp() == nil {
		return
	}
	var op ast.IncDecOp
	switch ctx.GetOp().GetTokenType() {
	case parser.NitroLexerINC:
		op = ast.IncDecOpInc
	case parser.NitroLexerDEC:
		op = ast.IncDecOpDec
	default:
		panic("unreachable")
	}

	l.put(ctx, &ast.IncDecStmt{
		LValue: &ast.LValue{Expr: l.takeExpr(ctx.Lvalue_expr())},
		Op:     op,
	})
}

// expr: short_lambda_expr
//     | expr3
//     ;
func (l *listener) ExitExpr(ctx *parser.ExprContext) {
	if ctx.Short_lambda_expr() != nil {
		l.put(ctx, l.takeExpr(ctx.Short_lambda_expr()))
	} else {
		l.put(ctx, l.takeExpr(ctx.Expr3()))
	}
}

// expr3: <assoc=right> expr3 '?' expr3 ':' expr3
//      | binary_expr
//      ;
func (l *listener) ExitExpr3(ctx *parser.Expr3Context) {
	allExpr := ctx.AllExpr3()
	if len(allExpr) != 0 {
		l.put(ctx, &ast.TernaryExpr{
			Condition: l.takeExpr(allExpr[0]),
			Then:      l.takeExpr(allExpr[1]),
			Else:      l.takeExpr(allExpr[2]),
		})
		return
	}

	l.put(ctx, l.takeExpr(ctx.Binary_expr()))
}

// binary_expr: unary_expr
// | binary_expr op=PIPE binary_expr
// | binary_expr op=(MUL|DIV|MOD) binary_expr
// | binary_expr op=(ADD|SUB) binary_expr
// | binary_expr op=(LT|LE|GT|GE|EQ|NE) binary_expr
// | binary_expr op=AND binary_expr
// | binary_expr op=OR binary_expr
// ;
func (l *listener) ExitBinary_expr(ctx *parser.Binary_exprContext) {
	if ctx.Unary_expr() != nil {
		l.put(ctx, l.takeExpr(ctx.Unary_expr()))
		return
	}

	left := l.takeExpr(ctx.Binary_expr(0))
	right := l.takeExpr(ctx.Binary_expr(1))

	simpleBinExp := func(op ast.Operator) {
		l.put(ctx, &ast.BinaryExpr{Left: left, Right: right, Op: op})
	}

	switch ctx.GetOp().GetTokenType() {
	case parser.NitroLexerPIPE:
		switch fn := right.(type) {
		case *ast.FuncCallExpr:
			fn.Args = append(ast.Exprs{left}, fn.Args...)
			fn.Pipeline = true
			l.put(ctx, fn)

		default:
			fnCall := &ast.FuncCallExpr{
				Target: fn,
				Args:   ast.Exprs{left},
				RetN:   1,
			}
			fnCall.Pipeline = true
			l.put(ctx, fnCall)
		}

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
	if ctx.GetOp() == nil {
		l.put(ctx, l.takeExpr(ctx.Primary_expr()))
		return
	}

	var op ast.UnaryOp
	if ctx.GetOp() != nil {
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
	}

	l.put(ctx, &ast.UnaryExpr{
		Term: l.takeExpr(ctx.Primary_expr()),
		Op:   op,
	})
}

// primary_expr: ID                                       # primary_expr_simple_ref
//             | primary_expr '.' ID                      # primary_expr_member_access
//             | primary_expr '[' expr ']'                # primary_expr_index
//             | primary_expr '[' expr? ':' expr? ']'     # primary_expr_slice
//             | primary_expr '(' arg_list? ')'           # primary_expr_call
//             | lambda_expr                              # primary_expr_lambda
//             | exec_expr                                # primary_exec_expr
//             | object_literal                           # primary_expr_object
//             | array_literal                            # primary_expr_array
//             | simple_literal                           # primary_expr_literal
//             | REGEX                                    # primary_expr_regex
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

func (l *listener) ExitPrimary_exec_expr(ctx *parser.Primary_exec_exprContext) {
	l.put(ctx, l.takeExpr(ctx.Exec_expr()))
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
		Expand: ctx.EXPAND() != nil,
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

func (l *listener) ExitPrimary_expr_regex(ctx *parser.Primary_expr_regexContext) {
	l.put(ctx, &ast.RegexLiteral{
		RegexDef: ctx.REGEX().GetText(),
	})
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
	lambda := &ast.AnonFuncExpr{}
	lambda.Params = l.takeASTs(ctx.Param_list())
	lambda.Block = l.takeAST(ctx.Stmts()).(*ast.StmtBlock)
	l.put(ctx, lambda)
}

// short_lambda_expr: '&' param_list? '->' expr2;
func (l *listener) ExitShort_lambda_expr(ctx *parser.Short_lambda_exprContext) {
	lambda := &ast.AnonFuncExpr{}
	lambda.Params = l.takeASTs(ctx.Param_list())
	lambda.Block = &ast.StmtBlock{
		Stmts: ast.ASTs{
			&ast.ReturnStmt{
				Values: ast.Exprs{l.takeExpr(ctx.Expr3())},
			},
		},
	}
	l.put(ctx, lambda)
}

// exec_expr: EXEC_PREFIX exec_expr_arg+ EXEC_SUFFIX;
func (l *listener) ExitExec_expr(ctx *parser.Exec_exprContext) {
	// TODO: set position.
	exprArgs := &ast.ArrayLiteral{
		Block: &ast.ArrayElementBlock{
			Elements: make(ast.ASTs, len(ctx.AllExec_expr_arg())),
		},
	}
	for i, argCtx := range ctx.AllExec_expr_arg() {
		exprArgs.Block.Elements[i] = &ast.ArrayElement{
			Val: l.takeExpr(argCtx),
		}
	}

	l.put(ctx, &ast.FuncCallExpr{
		Target: &ast.SimpleRef{
			ID: token.Token{
				Str:  "$exec",
				Type: token.String,
			},
		},
		Args: ast.Exprs{exprArgs},
		RetN: 1,
	})
}

// exec_expr_arg: EXEC_LITERAL
//              | EXEC_DQUOTE_LITERAL
//              | EXEC_SQUOTE_LITERAL
//              | EXEC_EXPR_PREFIX expr EXEC_EXPR_SUFFIX;
func (l *listener) ExitExec_expr_arg(ctx *parser.Exec_expr_argContext) {
	if expr := ctx.Expr(); expr != nil {
		l.put(ctx, l.takeExpr(expr))
		return
	}

	t := ctx.EXEC_LITERAL()
	if t == nil {
		t = ctx.EXEC_DQUOTE_LITERAL()
	}
	if t == nil {
		t = ctx.EXEC_SQUOTE_LITERAL()
	}
	if t == nil {
		t = ctx.EXEC_SQUOTE_LITERAL()
	}

	if t == nil {
		// The parser is probably in error recovery.
		l.put(ctx, &ast.LiteralExpr{
			Val: token.Token{
				Type: token.Nil,
			},
		})
		return
	}

	l.put(ctx, &ast.LiteralExpr{
		Val: l.tokenToNitro(t.GetSymbol()),
	})
}

// object_literal: '{' object_fields? '}';
func (l *listener) ExitObject_literal(ctx *parser.Object_literalContext) {
	l.put(ctx, &ast.ObjectLiteral{
		FieldBlock: l.takeAST(ctx.Object_fields()).(*ast.ObjectFieldBlock),
	})
}

// object_fields: (object_field ((','|';') object_field)* (','|';')*)?;
func (l *listener) ExitObject_fields(ctx *parser.Object_fieldsContext) {
	allFields := ctx.AllObject_field()
	fields := make(ast.ASTs, len(allFields))
	for i, entry := range allFields {
		fields[i] = l.takeAST(entry)
	}
	block := &ast.ObjectFieldBlock{
		Fields: fields,
	}
	l.put(ctx, block)
}

// object_field: id_or_keyword ':' expr     # object_field_id_key
//             | '[' expr ']' ':' expr      # object_field_expr_key
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

// array_literal: '[' array_elems? ']';
func (l *listener) ExitArray_literal(ctx *parser.Array_literalContext) {
	l.put(ctx, &ast.ArrayLiteral{
		Block: l.takeAST(ctx.Array_elems()).(*ast.ArrayElementBlock),
	})
}

// array_elems: array_elem ((','|';') array_elem)* (','|';')*;
func (l *listener) ExitArray_elems(ctx *parser.Array_elemsContext) {
	all := ctx.AllArray_elem()
	elems := make(ast.ASTs, len(all))
	for i, entry := range all {
		elems[i] = l.takeAST(entry)
	}
	block := &ast.ArrayElementBlock{
		Elements: elems,
	}
	l.put(ctx, block)
}

// array_elem: expr;
func (l *listener) ExitArray_elem(ctx *parser.Array_elemContext) {
	switch {
	case ctx.Expr() != nil:
		l.put(ctx, &ast.ArrayElement{Val: l.takeExpr(ctx.Expr())})
	default:
		panic("unreachable")
	}
}

func (l *listener) ExitId_or_keyword(ctx *parser.Id_or_keywordContext) {
	id := ctx.GetT().GetText()
	l.put(ctx, id)
}
