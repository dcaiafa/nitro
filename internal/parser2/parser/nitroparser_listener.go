// Code generated from NitroParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // NitroParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NitroParserListener is a complete listener for a parse tree produced by NitroParser.
type NitroParserListener interface {
	antlr.ParseTreeListener

	// EnterShort_prog is called when entering the short_prog production.
	EnterShort_prog(c *Short_progContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

	// EnterStmt_list is called when entering the stmt_list production.
	EnterStmt_list(c *Stmt_listContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignment_stmt is called when entering the assignment_stmt production.
	EnterAssignment_stmt(c *Assignment_stmtContext)

	// EnterAssignment_lvalues is called when entering the assignment_lvalues production.
	EnterAssignment_lvalues(c *Assignment_lvaluesContext)

	// EnterRvalues is called when entering the rvalues production.
	EnterRvalues(c *RvaluesContext)

	// EnterVar_decl_stmt is called when entering the var_decl_stmt production.
	EnterVar_decl_stmt(c *Var_decl_stmtContext)

	// EnterVar_decl_vars is called when entering the var_decl_vars production.
	EnterVar_decl_vars(c *Var_decl_varsContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterFor_vars is called when entering the for_vars production.
	EnterFor_vars(c *For_varsContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterIf_elif is called when entering the if_elif production.
	EnterIf_elif(c *If_elifContext)

	// EnterIf_else is called when entering the if_else production.
	EnterIf_else(c *If_elseContext)

	// EnterFunc_stmt is called when entering the func_stmt production.
	EnterFunc_stmt(c *Func_stmtContext)

	// EnterParam_list is called when entering the param_list production.
	EnterParam_list(c *Param_listContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterTry_catch_stmt is called when entering the try_catch_stmt production.
	EnterTry_catch_stmt(c *Try_catch_stmtContext)

	// EnterThrow_stmt is called when entering the throw_stmt production.
	EnterThrow_stmt(c *Throw_stmtContext)

	// EnterDefer_stmt is called when entering the defer_stmt production.
	EnterDefer_stmt(c *Defer_stmtContext)

	// EnterYield_stmt is called when entering the yield_stmt production.
	EnterYield_stmt(c *Yield_stmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterPipeline_term_expr_binary is called when entering the pipeline_term_expr_binary production.
	EnterPipeline_term_expr_binary(c *Pipeline_term_expr_binaryContext)

	// EnterPipeline_term_expr_short_lambda is called when entering the pipeline_term_expr_short_lambda production.
	EnterPipeline_term_expr_short_lambda(c *Pipeline_term_expr_short_lambdaContext)

	// EnterBinary_expr is called when entering the binary_expr production.
	EnterBinary_expr(c *Binary_exprContext)

	// EnterUnary_expr is called when entering the unary_expr production.
	EnterUnary_expr(c *Unary_exprContext)

	// EnterPrimary_expr_regex is called when entering the primary_expr_regex production.
	EnterPrimary_expr_regex(c *Primary_expr_regexContext)

	// EnterPrimary_expr_object is called when entering the primary_expr_object production.
	EnterPrimary_expr_object(c *Primary_expr_objectContext)

	// EnterPrimary_expr_parenthesis is called when entering the primary_expr_parenthesis production.
	EnterPrimary_expr_parenthesis(c *Primary_expr_parenthesisContext)

	// EnterPrimary_expr_simple_ref is called when entering the primary_expr_simple_ref production.
	EnterPrimary_expr_simple_ref(c *Primary_expr_simple_refContext)

	// EnterPrimary_expr_lambda is called when entering the primary_expr_lambda production.
	EnterPrimary_expr_lambda(c *Primary_expr_lambdaContext)

	// EnterPrimary_expr_index is called when entering the primary_expr_index production.
	EnterPrimary_expr_index(c *Primary_expr_indexContext)

	// EnterPrimary_expr_literal is called when entering the primary_expr_literal production.
	EnterPrimary_expr_literal(c *Primary_expr_literalContext)

	// EnterPrimary_expr_member_access is called when entering the primary_expr_member_access production.
	EnterPrimary_expr_member_access(c *Primary_expr_member_accessContext)

	// EnterPrimary_expr_array is called when entering the primary_expr_array production.
	EnterPrimary_expr_array(c *Primary_expr_arrayContext)

	// EnterPrimary_expr_call is called when entering the primary_expr_call production.
	EnterPrimary_expr_call(c *Primary_expr_callContext)

	// EnterPrimary_expr_slice is called when entering the primary_expr_slice production.
	EnterPrimary_expr_slice(c *Primary_expr_sliceContext)

	// EnterSimple_literal is called when entering the simple_literal production.
	EnterSimple_literal(c *Simple_literalContext)

	// EnterArg_list is called when entering the arg_list production.
	EnterArg_list(c *Arg_listContext)

	// EnterLvalue_expr_simple_ref is called when entering the lvalue_expr_simple_ref production.
	EnterLvalue_expr_simple_ref(c *Lvalue_expr_simple_refContext)

	// EnterLvalue_expr_member_access is called when entering the lvalue_expr_member_access production.
	EnterLvalue_expr_member_access(c *Lvalue_expr_member_accessContext)

	// EnterLvalue_expr_index is called when entering the lvalue_expr_index production.
	EnterLvalue_expr_index(c *Lvalue_expr_indexContext)

	// EnterLambda_expr is called when entering the lambda_expr production.
	EnterLambda_expr(c *Lambda_exprContext)

	// EnterShort_lambda_expr is called when entering the short_lambda_expr production.
	EnterShort_lambda_expr(c *Short_lambda_exprContext)

	// EnterObject_literal is called when entering the object_literal production.
	EnterObject_literal(c *Object_literalContext)

	// EnterObject_fields is called when entering the object_fields production.
	EnterObject_fields(c *Object_fieldsContext)

	// EnterObject_field_id_key is called when entering the object_field_id_key production.
	EnterObject_field_id_key(c *Object_field_id_keyContext)

	// EnterObject_field_expr_key is called when entering the object_field_expr_key production.
	EnterObject_field_expr_key(c *Object_field_expr_keyContext)

	// EnterObject_field_expansion is called when entering the object_field_expansion production.
	EnterObject_field_expansion(c *Object_field_expansionContext)

	// EnterObject_field_if is called when entering the object_field_if production.
	EnterObject_field_if(c *Object_field_ifContext)

	// EnterObject_field_for is called when entering the object_field_for production.
	EnterObject_field_for(c *Object_field_forContext)

	// EnterObject_if is called when entering the object_if production.
	EnterObject_if(c *Object_ifContext)

	// EnterObject_elif is called when entering the object_elif production.
	EnterObject_elif(c *Object_elifContext)

	// EnterObject_else is called when entering the object_else production.
	EnterObject_else(c *Object_elseContext)

	// EnterObject_for is called when entering the object_for production.
	EnterObject_for(c *Object_forContext)

	// EnterArray_literal is called when entering the array_literal production.
	EnterArray_literal(c *Array_literalContext)

	// EnterArray_elems is called when entering the array_elems production.
	EnterArray_elems(c *Array_elemsContext)

	// EnterArray_elem is called when entering the array_elem production.
	EnterArray_elem(c *Array_elemContext)

	// EnterArray_if is called when entering the array_if production.
	EnterArray_if(c *Array_ifContext)

	// EnterArray_elif is called when entering the array_elif production.
	EnterArray_elif(c *Array_elifContext)

	// EnterArray_else is called when entering the array_else production.
	EnterArray_else(c *Array_elseContext)

	// EnterArray_for is called when entering the array_for production.
	EnterArray_for(c *Array_forContext)

	// EnterId_or_keyword is called when entering the id_or_keyword production.
	EnterId_or_keyword(c *Id_or_keywordContext)

	// ExitShort_prog is called when exiting the short_prog production.
	ExitShort_prog(c *Short_progContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

	// ExitStmt_list is called when exiting the stmt_list production.
	ExitStmt_list(c *Stmt_listContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignment_stmt is called when exiting the assignment_stmt production.
	ExitAssignment_stmt(c *Assignment_stmtContext)

	// ExitAssignment_lvalues is called when exiting the assignment_lvalues production.
	ExitAssignment_lvalues(c *Assignment_lvaluesContext)

	// ExitRvalues is called when exiting the rvalues production.
	ExitRvalues(c *RvaluesContext)

	// ExitVar_decl_stmt is called when exiting the var_decl_stmt production.
	ExitVar_decl_stmt(c *Var_decl_stmtContext)

	// ExitVar_decl_vars is called when exiting the var_decl_vars production.
	ExitVar_decl_vars(c *Var_decl_varsContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitFor_vars is called when exiting the for_vars production.
	ExitFor_vars(c *For_varsContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitIf_elif is called when exiting the if_elif production.
	ExitIf_elif(c *If_elifContext)

	// ExitIf_else is called when exiting the if_else production.
	ExitIf_else(c *If_elseContext)

	// ExitFunc_stmt is called when exiting the func_stmt production.
	ExitFunc_stmt(c *Func_stmtContext)

	// ExitParam_list is called when exiting the param_list production.
	ExitParam_list(c *Param_listContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitTry_catch_stmt is called when exiting the try_catch_stmt production.
	ExitTry_catch_stmt(c *Try_catch_stmtContext)

	// ExitThrow_stmt is called when exiting the throw_stmt production.
	ExitThrow_stmt(c *Throw_stmtContext)

	// ExitDefer_stmt is called when exiting the defer_stmt production.
	ExitDefer_stmt(c *Defer_stmtContext)

	// ExitYield_stmt is called when exiting the yield_stmt production.
	ExitYield_stmt(c *Yield_stmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitPipeline_term_expr_binary is called when exiting the pipeline_term_expr_binary production.
	ExitPipeline_term_expr_binary(c *Pipeline_term_expr_binaryContext)

	// ExitPipeline_term_expr_short_lambda is called when exiting the pipeline_term_expr_short_lambda production.
	ExitPipeline_term_expr_short_lambda(c *Pipeline_term_expr_short_lambdaContext)

	// ExitBinary_expr is called when exiting the binary_expr production.
	ExitBinary_expr(c *Binary_exprContext)

	// ExitUnary_expr is called when exiting the unary_expr production.
	ExitUnary_expr(c *Unary_exprContext)

	// ExitPrimary_expr_regex is called when exiting the primary_expr_regex production.
	ExitPrimary_expr_regex(c *Primary_expr_regexContext)

	// ExitPrimary_expr_object is called when exiting the primary_expr_object production.
	ExitPrimary_expr_object(c *Primary_expr_objectContext)

	// ExitPrimary_expr_parenthesis is called when exiting the primary_expr_parenthesis production.
	ExitPrimary_expr_parenthesis(c *Primary_expr_parenthesisContext)

	// ExitPrimary_expr_simple_ref is called when exiting the primary_expr_simple_ref production.
	ExitPrimary_expr_simple_ref(c *Primary_expr_simple_refContext)

	// ExitPrimary_expr_lambda is called when exiting the primary_expr_lambda production.
	ExitPrimary_expr_lambda(c *Primary_expr_lambdaContext)

	// ExitPrimary_expr_index is called when exiting the primary_expr_index production.
	ExitPrimary_expr_index(c *Primary_expr_indexContext)

	// ExitPrimary_expr_literal is called when exiting the primary_expr_literal production.
	ExitPrimary_expr_literal(c *Primary_expr_literalContext)

	// ExitPrimary_expr_member_access is called when exiting the primary_expr_member_access production.
	ExitPrimary_expr_member_access(c *Primary_expr_member_accessContext)

	// ExitPrimary_expr_array is called when exiting the primary_expr_array production.
	ExitPrimary_expr_array(c *Primary_expr_arrayContext)

	// ExitPrimary_expr_call is called when exiting the primary_expr_call production.
	ExitPrimary_expr_call(c *Primary_expr_callContext)

	// ExitPrimary_expr_slice is called when exiting the primary_expr_slice production.
	ExitPrimary_expr_slice(c *Primary_expr_sliceContext)

	// ExitSimple_literal is called when exiting the simple_literal production.
	ExitSimple_literal(c *Simple_literalContext)

	// ExitArg_list is called when exiting the arg_list production.
	ExitArg_list(c *Arg_listContext)

	// ExitLvalue_expr_simple_ref is called when exiting the lvalue_expr_simple_ref production.
	ExitLvalue_expr_simple_ref(c *Lvalue_expr_simple_refContext)

	// ExitLvalue_expr_member_access is called when exiting the lvalue_expr_member_access production.
	ExitLvalue_expr_member_access(c *Lvalue_expr_member_accessContext)

	// ExitLvalue_expr_index is called when exiting the lvalue_expr_index production.
	ExitLvalue_expr_index(c *Lvalue_expr_indexContext)

	// ExitLambda_expr is called when exiting the lambda_expr production.
	ExitLambda_expr(c *Lambda_exprContext)

	// ExitShort_lambda_expr is called when exiting the short_lambda_expr production.
	ExitShort_lambda_expr(c *Short_lambda_exprContext)

	// ExitObject_literal is called when exiting the object_literal production.
	ExitObject_literal(c *Object_literalContext)

	// ExitObject_fields is called when exiting the object_fields production.
	ExitObject_fields(c *Object_fieldsContext)

	// ExitObject_field_id_key is called when exiting the object_field_id_key production.
	ExitObject_field_id_key(c *Object_field_id_keyContext)

	// ExitObject_field_expr_key is called when exiting the object_field_expr_key production.
	ExitObject_field_expr_key(c *Object_field_expr_keyContext)

	// ExitObject_field_expansion is called when exiting the object_field_expansion production.
	ExitObject_field_expansion(c *Object_field_expansionContext)

	// ExitObject_field_if is called when exiting the object_field_if production.
	ExitObject_field_if(c *Object_field_ifContext)

	// ExitObject_field_for is called when exiting the object_field_for production.
	ExitObject_field_for(c *Object_field_forContext)

	// ExitObject_if is called when exiting the object_if production.
	ExitObject_if(c *Object_ifContext)

	// ExitObject_elif is called when exiting the object_elif production.
	ExitObject_elif(c *Object_elifContext)

	// ExitObject_else is called when exiting the object_else production.
	ExitObject_else(c *Object_elseContext)

	// ExitObject_for is called when exiting the object_for production.
	ExitObject_for(c *Object_forContext)

	// ExitArray_literal is called when exiting the array_literal production.
	ExitArray_literal(c *Array_literalContext)

	// ExitArray_elems is called when exiting the array_elems production.
	ExitArray_elems(c *Array_elemsContext)

	// ExitArray_elem is called when exiting the array_elem production.
	ExitArray_elem(c *Array_elemContext)

	// ExitArray_if is called when exiting the array_if production.
	ExitArray_if(c *Array_ifContext)

	// ExitArray_elif is called when exiting the array_elif production.
	ExitArray_elif(c *Array_elifContext)

	// ExitArray_else is called when exiting the array_else production.
	ExitArray_else(c *Array_elseContext)

	// ExitArray_for is called when exiting the array_for production.
	ExitArray_for(c *Array_forContext)

	// ExitId_or_keyword is called when exiting the id_or_keyword production.
	ExitId_or_keyword(c *Id_or_keywordContext)
}
