// Code generated from NitroParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NitroParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNitroParserListener is a complete listener for a parse tree produced by NitroParser.
type BaseNitroParserListener struct{}

var _ NitroParserListener = &BaseNitroParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNitroParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNitroParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNitroParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNitroParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseNitroParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseNitroParserListener) ExitStart(ctx *StartContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseNitroParserListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseNitroParserListener) ExitUnit(ctx *UnitContext) {}

// EnterPrologue is called when production prologue is entered.
func (s *BaseNitroParserListener) EnterPrologue(ctx *PrologueContext) {}

// ExitPrologue is called when production prologue is exited.
func (s *BaseNitroParserListener) ExitPrologue(ctx *PrologueContext) {}

// EnterMeta_directive is called when production meta_directive is entered.
func (s *BaseNitroParserListener) EnterMeta_directive(ctx *Meta_directiveContext) {}

// ExitMeta_directive is called when production meta_directive is exited.
func (s *BaseNitroParserListener) ExitMeta_directive(ctx *Meta_directiveContext) {}

// EnterMeta_info is called when production meta_info is entered.
func (s *BaseNitroParserListener) EnterMeta_info(ctx *Meta_infoContext) {}

// ExitMeta_info is called when production meta_info is exited.
func (s *BaseNitroParserListener) ExitMeta_info(ctx *Meta_infoContext) {}

// EnterMeta_param is called when production meta_param is entered.
func (s *BaseNitroParserListener) EnterMeta_param(ctx *Meta_paramContext) {}

// ExitMeta_param is called when production meta_param is exited.
func (s *BaseNitroParserListener) ExitMeta_param(ctx *Meta_paramContext) {}

// EnterMeta_flag is called when production meta_flag is entered.
func (s *BaseNitroParserListener) EnterMeta_flag(ctx *Meta_flagContext) {}

// ExitMeta_flag is called when production meta_flag is exited.
func (s *BaseNitroParserListener) ExitMeta_flag(ctx *Meta_flagContext) {}

// EnterMeta_attribs is called when production meta_attribs is entered.
func (s *BaseNitroParserListener) EnterMeta_attribs(ctx *Meta_attribsContext) {}

// ExitMeta_attribs is called when production meta_attribs is exited.
func (s *BaseNitroParserListener) ExitMeta_attribs(ctx *Meta_attribsContext) {}

// EnterMeta_attrib is called when production meta_attrib is entered.
func (s *BaseNitroParserListener) EnterMeta_attrib(ctx *Meta_attribContext) {}

// ExitMeta_attrib is called when production meta_attrib is exited.
func (s *BaseNitroParserListener) ExitMeta_attrib(ctx *Meta_attribContext) {}

// EnterMeta_literal is called when production meta_literal is entered.
func (s *BaseNitroParserListener) EnterMeta_literal(ctx *Meta_literalContext) {}

// ExitMeta_literal is called when production meta_literal is exited.
func (s *BaseNitroParserListener) ExitMeta_literal(ctx *Meta_literalContext) {}

// EnterImport_stmt is called when production import_stmt is entered.
func (s *BaseNitroParserListener) EnterImport_stmt(ctx *Import_stmtContext) {}

// ExitImport_stmt is called when production import_stmt is exited.
func (s *BaseNitroParserListener) ExitImport_stmt(ctx *Import_stmtContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BaseNitroParserListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BaseNitroParserListener) ExitStmts(ctx *StmtsContext) {}

// EnterStmt_list is called when production stmt_list is entered.
func (s *BaseNitroParserListener) EnterStmt_list(ctx *Stmt_listContext) {}

// ExitStmt_list is called when production stmt_list is exited.
func (s *BaseNitroParserListener) ExitStmt_list(ctx *Stmt_listContext) {}

// EnterStmt_assignment is called when production stmt_assignment is entered.
func (s *BaseNitroParserListener) EnterStmt_assignment(ctx *Stmt_assignmentContext) {}

// ExitStmt_assignment is called when production stmt_assignment is exited.
func (s *BaseNitroParserListener) ExitStmt_assignment(ctx *Stmt_assignmentContext) {}

// EnterStmt_op_assign is called when production stmt_op_assign is entered.
func (s *BaseNitroParserListener) EnterStmt_op_assign(ctx *Stmt_op_assignContext) {}

// ExitStmt_op_assign is called when production stmt_op_assign is exited.
func (s *BaseNitroParserListener) ExitStmt_op_assign(ctx *Stmt_op_assignContext) {}

// EnterStmt_var_dec is called when production stmt_var_dec is entered.
func (s *BaseNitroParserListener) EnterStmt_var_dec(ctx *Stmt_var_decContext) {}

// ExitStmt_var_dec is called when production stmt_var_dec is exited.
func (s *BaseNitroParserListener) ExitStmt_var_dec(ctx *Stmt_var_decContext) {}

// EnterStmt_for is called when production stmt_for is entered.
func (s *BaseNitroParserListener) EnterStmt_for(ctx *Stmt_forContext) {}

// ExitStmt_for is called when production stmt_for is exited.
func (s *BaseNitroParserListener) ExitStmt_for(ctx *Stmt_forContext) {}

// EnterStmt_while is called when production stmt_while is entered.
func (s *BaseNitroParserListener) EnterStmt_while(ctx *Stmt_whileContext) {}

// ExitStmt_while is called when production stmt_while is exited.
func (s *BaseNitroParserListener) ExitStmt_while(ctx *Stmt_whileContext) {}

// EnterStmt_if is called when production stmt_if is entered.
func (s *BaseNitroParserListener) EnterStmt_if(ctx *Stmt_ifContext) {}

// ExitStmt_if is called when production stmt_if is exited.
func (s *BaseNitroParserListener) ExitStmt_if(ctx *Stmt_ifContext) {}

// EnterStmt_func is called when production stmt_func is entered.
func (s *BaseNitroParserListener) EnterStmt_func(ctx *Stmt_funcContext) {}

// ExitStmt_func is called when production stmt_func is exited.
func (s *BaseNitroParserListener) ExitStmt_func(ctx *Stmt_funcContext) {}

// EnterStmt_return is called when production stmt_return is entered.
func (s *BaseNitroParserListener) EnterStmt_return(ctx *Stmt_returnContext) {}

// ExitStmt_return is called when production stmt_return is exited.
func (s *BaseNitroParserListener) ExitStmt_return(ctx *Stmt_returnContext) {}

// EnterStmt_expr is called when production stmt_expr is entered.
func (s *BaseNitroParserListener) EnterStmt_expr(ctx *Stmt_exprContext) {}

// ExitStmt_expr is called when production stmt_expr is exited.
func (s *BaseNitroParserListener) ExitStmt_expr(ctx *Stmt_exprContext) {}

// EnterStmt_try_catch is called when production stmt_try_catch is entered.
func (s *BaseNitroParserListener) EnterStmt_try_catch(ctx *Stmt_try_catchContext) {}

// ExitStmt_try_catch is called when production stmt_try_catch is exited.
func (s *BaseNitroParserListener) ExitStmt_try_catch(ctx *Stmt_try_catchContext) {}

// EnterStmt_throw is called when production stmt_throw is entered.
func (s *BaseNitroParserListener) EnterStmt_throw(ctx *Stmt_throwContext) {}

// ExitStmt_throw is called when production stmt_throw is exited.
func (s *BaseNitroParserListener) ExitStmt_throw(ctx *Stmt_throwContext) {}

// EnterStmt_defer is called when production stmt_defer is entered.
func (s *BaseNitroParserListener) EnterStmt_defer(ctx *Stmt_deferContext) {}

// ExitStmt_defer is called when production stmt_defer is exited.
func (s *BaseNitroParserListener) ExitStmt_defer(ctx *Stmt_deferContext) {}

// EnterStmt_yield is called when production stmt_yield is entered.
func (s *BaseNitroParserListener) EnterStmt_yield(ctx *Stmt_yieldContext) {}

// ExitStmt_yield is called when production stmt_yield is exited.
func (s *BaseNitroParserListener) ExitStmt_yield(ctx *Stmt_yieldContext) {}

// EnterStmt_break is called when production stmt_break is entered.
func (s *BaseNitroParserListener) EnterStmt_break(ctx *Stmt_breakContext) {}

// ExitStmt_break is called when production stmt_break is exited.
func (s *BaseNitroParserListener) ExitStmt_break(ctx *Stmt_breakContext) {}

// EnterStmt_continue is called when production stmt_continue is entered.
func (s *BaseNitroParserListener) EnterStmt_continue(ctx *Stmt_continueContext) {}

// ExitStmt_continue is called when production stmt_continue is exited.
func (s *BaseNitroParserListener) ExitStmt_continue(ctx *Stmt_continueContext) {}

// EnterStmt_inc_dec is called when production stmt_inc_dec is entered.
func (s *BaseNitroParserListener) EnterStmt_inc_dec(ctx *Stmt_inc_decContext) {}

// ExitStmt_inc_dec is called when production stmt_inc_dec is exited.
func (s *BaseNitroParserListener) ExitStmt_inc_dec(ctx *Stmt_inc_decContext) {}

// EnterAssignment_stmt is called when production assignment_stmt is entered.
func (s *BaseNitroParserListener) EnterAssignment_stmt(ctx *Assignment_stmtContext) {}

// ExitAssignment_stmt is called when production assignment_stmt is exited.
func (s *BaseNitroParserListener) ExitAssignment_stmt(ctx *Assignment_stmtContext) {}

// EnterAssignment_lvalues is called when production assignment_lvalues is entered.
func (s *BaseNitroParserListener) EnterAssignment_lvalues(ctx *Assignment_lvaluesContext) {}

// ExitAssignment_lvalues is called when production assignment_lvalues is exited.
func (s *BaseNitroParserListener) ExitAssignment_lvalues(ctx *Assignment_lvaluesContext) {}

// EnterRvalues is called when production rvalues is entered.
func (s *BaseNitroParserListener) EnterRvalues(ctx *RvaluesContext) {}

// ExitRvalues is called when production rvalues is exited.
func (s *BaseNitroParserListener) ExitRvalues(ctx *RvaluesContext) {}

// EnterAssignment_op_stmt is called when production assignment_op_stmt is entered.
func (s *BaseNitroParserListener) EnterAssignment_op_stmt(ctx *Assignment_op_stmtContext) {}

// ExitAssignment_op_stmt is called when production assignment_op_stmt is exited.
func (s *BaseNitroParserListener) ExitAssignment_op_stmt(ctx *Assignment_op_stmtContext) {}

// EnterVar_decl_stmt is called when production var_decl_stmt is entered.
func (s *BaseNitroParserListener) EnterVar_decl_stmt(ctx *Var_decl_stmtContext) {}

// ExitVar_decl_stmt is called when production var_decl_stmt is exited.
func (s *BaseNitroParserListener) ExitVar_decl_stmt(ctx *Var_decl_stmtContext) {}

// EnterVar_decl_vars is called when production var_decl_vars is entered.
func (s *BaseNitroParserListener) EnterVar_decl_vars(ctx *Var_decl_varsContext) {}

// ExitVar_decl_vars is called when production var_decl_vars is exited.
func (s *BaseNitroParserListener) ExitVar_decl_vars(ctx *Var_decl_varsContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BaseNitroParserListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BaseNitroParserListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterFor_vars is called when production for_vars is entered.
func (s *BaseNitroParserListener) EnterFor_vars(ctx *For_varsContext) {}

// ExitFor_vars is called when production for_vars is exited.
func (s *BaseNitroParserListener) ExitFor_vars(ctx *For_varsContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BaseNitroParserListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BaseNitroParserListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseNitroParserListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseNitroParserListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterIf_elif is called when production if_elif is entered.
func (s *BaseNitroParserListener) EnterIf_elif(ctx *If_elifContext) {}

// ExitIf_elif is called when production if_elif is exited.
func (s *BaseNitroParserListener) ExitIf_elif(ctx *If_elifContext) {}

// EnterIf_else is called when production if_else is entered.
func (s *BaseNitroParserListener) EnterIf_else(ctx *If_elseContext) {}

// ExitIf_else is called when production if_else is exited.
func (s *BaseNitroParserListener) ExitIf_else(ctx *If_elseContext) {}

// EnterFunc_stmt is called when production func_stmt is entered.
func (s *BaseNitroParserListener) EnterFunc_stmt(ctx *Func_stmtContext) {}

// ExitFunc_stmt is called when production func_stmt is exited.
func (s *BaseNitroParserListener) ExitFunc_stmt(ctx *Func_stmtContext) {}

// EnterParam_list is called when production param_list is entered.
func (s *BaseNitroParserListener) EnterParam_list(ctx *Param_listContext) {}

// ExitParam_list is called when production param_list is exited.
func (s *BaseNitroParserListener) ExitParam_list(ctx *Param_listContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseNitroParserListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseNitroParserListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterTry_catch_stmt is called when production try_catch_stmt is entered.
func (s *BaseNitroParserListener) EnterTry_catch_stmt(ctx *Try_catch_stmtContext) {}

// ExitTry_catch_stmt is called when production try_catch_stmt is exited.
func (s *BaseNitroParserListener) ExitTry_catch_stmt(ctx *Try_catch_stmtContext) {}

// EnterThrow_stmt is called when production throw_stmt is entered.
func (s *BaseNitroParserListener) EnterThrow_stmt(ctx *Throw_stmtContext) {}

// ExitThrow_stmt is called when production throw_stmt is exited.
func (s *BaseNitroParserListener) ExitThrow_stmt(ctx *Throw_stmtContext) {}

// EnterDefer_stmt is called when production defer_stmt is entered.
func (s *BaseNitroParserListener) EnterDefer_stmt(ctx *Defer_stmtContext) {}

// ExitDefer_stmt is called when production defer_stmt is exited.
func (s *BaseNitroParserListener) ExitDefer_stmt(ctx *Defer_stmtContext) {}

// EnterYield_stmt is called when production yield_stmt is entered.
func (s *BaseNitroParserListener) EnterYield_stmt(ctx *Yield_stmtContext) {}

// ExitYield_stmt is called when production yield_stmt is exited.
func (s *BaseNitroParserListener) ExitYield_stmt(ctx *Yield_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BaseNitroParserListener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BaseNitroParserListener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BaseNitroParserListener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BaseNitroParserListener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterInc_dec_stmt is called when production inc_dec_stmt is entered.
func (s *BaseNitroParserListener) EnterInc_dec_stmt(ctx *Inc_dec_stmtContext) {}

// ExitInc_dec_stmt is called when production inc_dec_stmt is exited.
func (s *BaseNitroParserListener) ExitInc_dec_stmt(ctx *Inc_dec_stmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseNitroParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseNitroParserListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr2 is called when production expr2 is entered.
func (s *BaseNitroParserListener) EnterExpr2(ctx *Expr2Context) {}

// ExitExpr2 is called when production expr2 is exited.
func (s *BaseNitroParserListener) ExitExpr2(ctx *Expr2Context) {}

// EnterExpr3 is called when production expr3 is entered.
func (s *BaseNitroParserListener) EnterExpr3(ctx *Expr3Context) {}

// ExitExpr3 is called when production expr3 is exited.
func (s *BaseNitroParserListener) ExitExpr3(ctx *Expr3Context) {}

// EnterBinary_expr is called when production binary_expr is entered.
func (s *BaseNitroParserListener) EnterBinary_expr(ctx *Binary_exprContext) {}

// ExitBinary_expr is called when production binary_expr is exited.
func (s *BaseNitroParserListener) ExitBinary_expr(ctx *Binary_exprContext) {}

// EnterUnary_expr is called when production unary_expr is entered.
func (s *BaseNitroParserListener) EnterUnary_expr(ctx *Unary_exprContext) {}

// ExitUnary_expr is called when production unary_expr is exited.
func (s *BaseNitroParserListener) ExitUnary_expr(ctx *Unary_exprContext) {}

// EnterPrimary_expr_regex is called when production primary_expr_regex is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_regex(ctx *Primary_expr_regexContext) {}

// ExitPrimary_expr_regex is called when production primary_expr_regex is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_regex(ctx *Primary_expr_regexContext) {}

// EnterPrimary_expr_object is called when production primary_expr_object is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_object(ctx *Primary_expr_objectContext) {}

// ExitPrimary_expr_object is called when production primary_expr_object is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_object(ctx *Primary_expr_objectContext) {}

// EnterPrimary_expr_parenthesis is called when production primary_expr_parenthesis is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_parenthesis(ctx *Primary_expr_parenthesisContext) {
}

// ExitPrimary_expr_parenthesis is called when production primary_expr_parenthesis is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_parenthesis(ctx *Primary_expr_parenthesisContext) {
}

// EnterPrimary_expr_simple_ref is called when production primary_expr_simple_ref is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_simple_ref(ctx *Primary_expr_simple_refContext) {}

// ExitPrimary_expr_simple_ref is called when production primary_expr_simple_ref is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_simple_ref(ctx *Primary_expr_simple_refContext) {}

// EnterPrimary_expr_lambda is called when production primary_expr_lambda is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_lambda(ctx *Primary_expr_lambdaContext) {}

// ExitPrimary_expr_lambda is called when production primary_expr_lambda is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_lambda(ctx *Primary_expr_lambdaContext) {}

// EnterPrimary_expr_index is called when production primary_expr_index is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_index(ctx *Primary_expr_indexContext) {}

// ExitPrimary_expr_index is called when production primary_expr_index is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_index(ctx *Primary_expr_indexContext) {}

// EnterPrimary_expr_literal is called when production primary_expr_literal is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_literal(ctx *Primary_expr_literalContext) {}

// ExitPrimary_expr_literal is called when production primary_expr_literal is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_literal(ctx *Primary_expr_literalContext) {}

// EnterPrimary_expr_member_access is called when production primary_expr_member_access is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_member_access(ctx *Primary_expr_member_accessContext) {
}

// ExitPrimary_expr_member_access is called when production primary_expr_member_access is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_member_access(ctx *Primary_expr_member_accessContext) {
}

// EnterPrimary_exec_expr is called when production primary_exec_expr is entered.
func (s *BaseNitroParserListener) EnterPrimary_exec_expr(ctx *Primary_exec_exprContext) {}

// ExitPrimary_exec_expr is called when production primary_exec_expr is exited.
func (s *BaseNitroParserListener) ExitPrimary_exec_expr(ctx *Primary_exec_exprContext) {}

// EnterPrimary_expr_array is called when production primary_expr_array is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_array(ctx *Primary_expr_arrayContext) {}

// ExitPrimary_expr_array is called when production primary_expr_array is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_array(ctx *Primary_expr_arrayContext) {}

// EnterPrimary_expr_call is called when production primary_expr_call is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_call(ctx *Primary_expr_callContext) {}

// ExitPrimary_expr_call is called when production primary_expr_call is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_call(ctx *Primary_expr_callContext) {}

// EnterPrimary_expr_slice is called when production primary_expr_slice is entered.
func (s *BaseNitroParserListener) EnterPrimary_expr_slice(ctx *Primary_expr_sliceContext) {}

// ExitPrimary_expr_slice is called when production primary_expr_slice is exited.
func (s *BaseNitroParserListener) ExitPrimary_expr_slice(ctx *Primary_expr_sliceContext) {}

// EnterSimple_literal is called when production simple_literal is entered.
func (s *BaseNitroParserListener) EnterSimple_literal(ctx *Simple_literalContext) {}

// ExitSimple_literal is called when production simple_literal is exited.
func (s *BaseNitroParserListener) ExitSimple_literal(ctx *Simple_literalContext) {}

// EnterArg_list is called when production arg_list is entered.
func (s *BaseNitroParserListener) EnterArg_list(ctx *Arg_listContext) {}

// ExitArg_list is called when production arg_list is exited.
func (s *BaseNitroParserListener) ExitArg_list(ctx *Arg_listContext) {}

// EnterLvalue_expr_simple_ref is called when production lvalue_expr_simple_ref is entered.
func (s *BaseNitroParserListener) EnterLvalue_expr_simple_ref(ctx *Lvalue_expr_simple_refContext) {}

// ExitLvalue_expr_simple_ref is called when production lvalue_expr_simple_ref is exited.
func (s *BaseNitroParserListener) ExitLvalue_expr_simple_ref(ctx *Lvalue_expr_simple_refContext) {}

// EnterLvalue_expr_member_access is called when production lvalue_expr_member_access is entered.
func (s *BaseNitroParserListener) EnterLvalue_expr_member_access(ctx *Lvalue_expr_member_accessContext) {
}

// ExitLvalue_expr_member_access is called when production lvalue_expr_member_access is exited.
func (s *BaseNitroParserListener) ExitLvalue_expr_member_access(ctx *Lvalue_expr_member_accessContext) {
}

// EnterLvalue_expr_index is called when production lvalue_expr_index is entered.
func (s *BaseNitroParserListener) EnterLvalue_expr_index(ctx *Lvalue_expr_indexContext) {}

// ExitLvalue_expr_index is called when production lvalue_expr_index is exited.
func (s *BaseNitroParserListener) ExitLvalue_expr_index(ctx *Lvalue_expr_indexContext) {}

// EnterLambda_expr is called when production lambda_expr is entered.
func (s *BaseNitroParserListener) EnterLambda_expr(ctx *Lambda_exprContext) {}

// ExitLambda_expr is called when production lambda_expr is exited.
func (s *BaseNitroParserListener) ExitLambda_expr(ctx *Lambda_exprContext) {}

// EnterShort_lambda_expr is called when production short_lambda_expr is entered.
func (s *BaseNitroParserListener) EnterShort_lambda_expr(ctx *Short_lambda_exprContext) {}

// ExitShort_lambda_expr is called when production short_lambda_expr is exited.
func (s *BaseNitroParserListener) ExitShort_lambda_expr(ctx *Short_lambda_exprContext) {}

// EnterExec_expr is called when production exec_expr is entered.
func (s *BaseNitroParserListener) EnterExec_expr(ctx *Exec_exprContext) {}

// ExitExec_expr is called when production exec_expr is exited.
func (s *BaseNitroParserListener) ExitExec_expr(ctx *Exec_exprContext) {}

// EnterExec_expr_arg is called when production exec_expr_arg is entered.
func (s *BaseNitroParserListener) EnterExec_expr_arg(ctx *Exec_expr_argContext) {}

// ExitExec_expr_arg is called when production exec_expr_arg is exited.
func (s *BaseNitroParserListener) ExitExec_expr_arg(ctx *Exec_expr_argContext) {}

// EnterObject_literal is called when production object_literal is entered.
func (s *BaseNitroParserListener) EnterObject_literal(ctx *Object_literalContext) {}

// ExitObject_literal is called when production object_literal is exited.
func (s *BaseNitroParserListener) ExitObject_literal(ctx *Object_literalContext) {}

// EnterObject_fields is called when production object_fields is entered.
func (s *BaseNitroParserListener) EnterObject_fields(ctx *Object_fieldsContext) {}

// ExitObject_fields is called when production object_fields is exited.
func (s *BaseNitroParserListener) ExitObject_fields(ctx *Object_fieldsContext) {}

// EnterObject_field_id_key is called when production object_field_id_key is entered.
func (s *BaseNitroParserListener) EnterObject_field_id_key(ctx *Object_field_id_keyContext) {}

// ExitObject_field_id_key is called when production object_field_id_key is exited.
func (s *BaseNitroParserListener) ExitObject_field_id_key(ctx *Object_field_id_keyContext) {}

// EnterObject_field_expr_key is called when production object_field_expr_key is entered.
func (s *BaseNitroParserListener) EnterObject_field_expr_key(ctx *Object_field_expr_keyContext) {}

// ExitObject_field_expr_key is called when production object_field_expr_key is exited.
func (s *BaseNitroParserListener) ExitObject_field_expr_key(ctx *Object_field_expr_keyContext) {}

// EnterObject_field_expansion is called when production object_field_expansion is entered.
func (s *BaseNitroParserListener) EnterObject_field_expansion(ctx *Object_field_expansionContext) {}

// ExitObject_field_expansion is called when production object_field_expansion is exited.
func (s *BaseNitroParserListener) ExitObject_field_expansion(ctx *Object_field_expansionContext) {}

// EnterArray_literal is called when production array_literal is entered.
func (s *BaseNitroParserListener) EnterArray_literal(ctx *Array_literalContext) {}

// ExitArray_literal is called when production array_literal is exited.
func (s *BaseNitroParserListener) ExitArray_literal(ctx *Array_literalContext) {}

// EnterArray_elems is called when production array_elems is entered.
func (s *BaseNitroParserListener) EnterArray_elems(ctx *Array_elemsContext) {}

// ExitArray_elems is called when production array_elems is exited.
func (s *BaseNitroParserListener) ExitArray_elems(ctx *Array_elemsContext) {}

// EnterArray_elem is called when production array_elem is entered.
func (s *BaseNitroParserListener) EnterArray_elem(ctx *Array_elemContext) {}

// ExitArray_elem is called when production array_elem is exited.
func (s *BaseNitroParserListener) ExitArray_elem(ctx *Array_elemContext) {}

// EnterId_or_keyword is called when production id_or_keyword is entered.
func (s *BaseNitroParserListener) EnterId_or_keyword(ctx *Id_or_keywordContext) {}

// ExitId_or_keyword is called when production id_or_keyword is exited.
func (s *BaseNitroParserListener) ExitId_or_keyword(ctx *Id_or_keywordContext) {}
