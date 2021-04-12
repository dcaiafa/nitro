// Code generated from NitroParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

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

// EnterShort_prog is called when production short_prog is entered.
func (s *BaseNitroParserListener) EnterShort_prog(ctx *Short_progContext) {}

// ExitShort_prog is called when production short_prog is exited.
func (s *BaseNitroParserListener) ExitShort_prog(ctx *Short_progContext) {}

// EnterStart is called when production start is entered.
func (s *BaseNitroParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseNitroParserListener) ExitStart(ctx *StartContext) {}

// EnterModule is called when production module is entered.
func (s *BaseNitroParserListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseNitroParserListener) ExitModule(ctx *ModuleContext) {}

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

// EnterExpr is called when production expr is entered.
func (s *BaseNitroParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseNitroParserListener) ExitExpr(ctx *ExprContext) {}

// EnterPipeline_term_expr_binary is called when production pipeline_term_expr_binary is entered.
func (s *BaseNitroParserListener) EnterPipeline_term_expr_binary(ctx *Pipeline_term_expr_binaryContext) {
}

// ExitPipeline_term_expr_binary is called when production pipeline_term_expr_binary is exited.
func (s *BaseNitroParserListener) ExitPipeline_term_expr_binary(ctx *Pipeline_term_expr_binaryContext) {
}

// EnterPipeline_term_expr_short_lambda is called when production pipeline_term_expr_short_lambda is entered.
func (s *BaseNitroParserListener) EnterPipeline_term_expr_short_lambda(ctx *Pipeline_term_expr_short_lambdaContext) {
}

// ExitPipeline_term_expr_short_lambda is called when production pipeline_term_expr_short_lambda is exited.
func (s *BaseNitroParserListener) ExitPipeline_term_expr_short_lambda(ctx *Pipeline_term_expr_short_lambdaContext) {
}

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

// EnterObject_field_if is called when production object_field_if is entered.
func (s *BaseNitroParserListener) EnterObject_field_if(ctx *Object_field_ifContext) {}

// ExitObject_field_if is called when production object_field_if is exited.
func (s *BaseNitroParserListener) ExitObject_field_if(ctx *Object_field_ifContext) {}

// EnterObject_field_for is called when production object_field_for is entered.
func (s *BaseNitroParserListener) EnterObject_field_for(ctx *Object_field_forContext) {}

// ExitObject_field_for is called when production object_field_for is exited.
func (s *BaseNitroParserListener) ExitObject_field_for(ctx *Object_field_forContext) {}

// EnterObject_if is called when production object_if is entered.
func (s *BaseNitroParserListener) EnterObject_if(ctx *Object_ifContext) {}

// ExitObject_if is called when production object_if is exited.
func (s *BaseNitroParserListener) ExitObject_if(ctx *Object_ifContext) {}

// EnterObject_elif is called when production object_elif is entered.
func (s *BaseNitroParserListener) EnterObject_elif(ctx *Object_elifContext) {}

// ExitObject_elif is called when production object_elif is exited.
func (s *BaseNitroParserListener) ExitObject_elif(ctx *Object_elifContext) {}

// EnterObject_else is called when production object_else is entered.
func (s *BaseNitroParserListener) EnterObject_else(ctx *Object_elseContext) {}

// ExitObject_else is called when production object_else is exited.
func (s *BaseNitroParserListener) ExitObject_else(ctx *Object_elseContext) {}

// EnterObject_for is called when production object_for is entered.
func (s *BaseNitroParserListener) EnterObject_for(ctx *Object_forContext) {}

// ExitObject_for is called when production object_for is exited.
func (s *BaseNitroParserListener) ExitObject_for(ctx *Object_forContext) {}

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

// EnterArray_if is called when production array_if is entered.
func (s *BaseNitroParserListener) EnterArray_if(ctx *Array_ifContext) {}

// ExitArray_if is called when production array_if is exited.
func (s *BaseNitroParserListener) ExitArray_if(ctx *Array_ifContext) {}

// EnterArray_elif is called when production array_elif is entered.
func (s *BaseNitroParserListener) EnterArray_elif(ctx *Array_elifContext) {}

// ExitArray_elif is called when production array_elif is exited.
func (s *BaseNitroParserListener) ExitArray_elif(ctx *Array_elifContext) {}

// EnterArray_else is called when production array_else is entered.
func (s *BaseNitroParserListener) EnterArray_else(ctx *Array_elseContext) {}

// ExitArray_else is called when production array_else is exited.
func (s *BaseNitroParserListener) ExitArray_else(ctx *Array_elseContext) {}

// EnterArray_for is called when production array_for is entered.
func (s *BaseNitroParserListener) EnterArray_for(ctx *Array_forContext) {}

// ExitArray_for is called when production array_for is exited.
func (s *BaseNitroParserListener) ExitArray_for(ctx *Array_forContext) {}

// EnterId_or_keyword is called when production id_or_keyword is entered.
func (s *BaseNitroParserListener) EnterId_or_keyword(ctx *Id_or_keywordContext) {}

// ExitId_or_keyword is called when production id_or_keyword is exited.
func (s *BaseNitroParserListener) ExitId_or_keyword(ctx *Id_or_keywordContext) {}
