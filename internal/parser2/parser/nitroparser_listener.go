// Code generated from NitroParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NitroParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NitroParserListener is a complete listener for a parse tree produced by NitroParser.
type NitroParserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterMeta_directive is called when entering the meta_directive production.
	EnterMeta_directive(c *Meta_directiveContext)

	// EnterMeta_info is called when entering the meta_info production.
	EnterMeta_info(c *Meta_infoContext)

	// EnterMeta_param is called when entering the meta_param production.
	EnterMeta_param(c *Meta_paramContext)

	// EnterMeta_flag is called when entering the meta_flag production.
	EnterMeta_flag(c *Meta_flagContext)

	// EnterMeta_attribs is called when entering the meta_attribs production.
	EnterMeta_attribs(c *Meta_attribsContext)

	// EnterMeta_attrib is called when entering the meta_attrib production.
	EnterMeta_attrib(c *Meta_attribContext)

	// EnterMeta_literal is called when entering the meta_literal production.
	EnterMeta_literal(c *Meta_literalContext)

	// EnterImport_stmt is called when entering the import_stmt production.
	EnterImport_stmt(c *Import_stmtContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

	// EnterStmt_list is called when entering the stmt_list production.
	EnterStmt_list(c *Stmt_listContext)

	// EnterStmt_assignment is called when entering the stmt_assignment production.
	EnterStmt_assignment(c *Stmt_assignmentContext)

	// EnterStmt_op_assign is called when entering the stmt_op_assign production.
	EnterStmt_op_assign(c *Stmt_op_assignContext)

	// EnterStmt_var_dec is called when entering the stmt_var_dec production.
	EnterStmt_var_dec(c *Stmt_var_decContext)

	// EnterStmt_for is called when entering the stmt_for production.
	EnterStmt_for(c *Stmt_forContext)

	// EnterStmt_while is called when entering the stmt_while production.
	EnterStmt_while(c *Stmt_whileContext)

	// EnterStmt_if is called when entering the stmt_if production.
	EnterStmt_if(c *Stmt_ifContext)

	// EnterStmt_func is called when entering the stmt_func production.
	EnterStmt_func(c *Stmt_funcContext)

	// EnterStmt_return is called when entering the stmt_return production.
	EnterStmt_return(c *Stmt_returnContext)

	// EnterStmt_expr is called when entering the stmt_expr production.
	EnterStmt_expr(c *Stmt_exprContext)

	// EnterStmt_try_catch is called when entering the stmt_try_catch production.
	EnterStmt_try_catch(c *Stmt_try_catchContext)

	// EnterStmt_throw is called when entering the stmt_throw production.
	EnterStmt_throw(c *Stmt_throwContext)

	// EnterStmt_defer is called when entering the stmt_defer production.
	EnterStmt_defer(c *Stmt_deferContext)

	// EnterStmt_yield is called when entering the stmt_yield production.
	EnterStmt_yield(c *Stmt_yieldContext)

	// EnterStmt_break is called when entering the stmt_break production.
	EnterStmt_break(c *Stmt_breakContext)

	// EnterStmt_continue is called when entering the stmt_continue production.
	EnterStmt_continue(c *Stmt_continueContext)

	// EnterStmt_inc_dec is called when entering the stmt_inc_dec production.
	EnterStmt_inc_dec(c *Stmt_inc_decContext)

	// EnterAssignment_stmt is called when entering the assignment_stmt production.
	EnterAssignment_stmt(c *Assignment_stmtContext)

	// EnterAssignment_lvalues is called when entering the assignment_lvalues production.
	EnterAssignment_lvalues(c *Assignment_lvaluesContext)

	// EnterRvalues is called when entering the rvalues production.
	EnterRvalues(c *RvaluesContext)

	// EnterAssignment_op_stmt is called when entering the assignment_op_stmt production.
	EnterAssignment_op_stmt(c *Assignment_op_stmtContext)

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

	// EnterBreak_stmt is called when entering the break_stmt production.
	EnterBreak_stmt(c *Break_stmtContext)

	// EnterContinue_stmt is called when entering the continue_stmt production.
	EnterContinue_stmt(c *Continue_stmtContext)

	// EnterInc_dec_stmt is called when entering the inc_dec_stmt production.
	EnterInc_dec_stmt(c *Inc_dec_stmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr2 is called when entering the expr2 production.
	EnterExpr2(c *Expr2Context)

	// EnterExpr3 is called when entering the expr3 production.
	EnterExpr3(c *Expr3Context)

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

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitMeta_directive is called when exiting the meta_directive production.
	ExitMeta_directive(c *Meta_directiveContext)

	// ExitMeta_info is called when exiting the meta_info production.
	ExitMeta_info(c *Meta_infoContext)

	// ExitMeta_param is called when exiting the meta_param production.
	ExitMeta_param(c *Meta_paramContext)

	// ExitMeta_flag is called when exiting the meta_flag production.
	ExitMeta_flag(c *Meta_flagContext)

	// ExitMeta_attribs is called when exiting the meta_attribs production.
	ExitMeta_attribs(c *Meta_attribsContext)

	// ExitMeta_attrib is called when exiting the meta_attrib production.
	ExitMeta_attrib(c *Meta_attribContext)

	// ExitMeta_literal is called when exiting the meta_literal production.
	ExitMeta_literal(c *Meta_literalContext)

	// ExitImport_stmt is called when exiting the import_stmt production.
	ExitImport_stmt(c *Import_stmtContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

	// ExitStmt_list is called when exiting the stmt_list production.
	ExitStmt_list(c *Stmt_listContext)

	// ExitStmt_assignment is called when exiting the stmt_assignment production.
	ExitStmt_assignment(c *Stmt_assignmentContext)

	// ExitStmt_op_assign is called when exiting the stmt_op_assign production.
	ExitStmt_op_assign(c *Stmt_op_assignContext)

	// ExitStmt_var_dec is called when exiting the stmt_var_dec production.
	ExitStmt_var_dec(c *Stmt_var_decContext)

	// ExitStmt_for is called when exiting the stmt_for production.
	ExitStmt_for(c *Stmt_forContext)

	// ExitStmt_while is called when exiting the stmt_while production.
	ExitStmt_while(c *Stmt_whileContext)

	// ExitStmt_if is called when exiting the stmt_if production.
	ExitStmt_if(c *Stmt_ifContext)

	// ExitStmt_func is called when exiting the stmt_func production.
	ExitStmt_func(c *Stmt_funcContext)

	// ExitStmt_return is called when exiting the stmt_return production.
	ExitStmt_return(c *Stmt_returnContext)

	// ExitStmt_expr is called when exiting the stmt_expr production.
	ExitStmt_expr(c *Stmt_exprContext)

	// ExitStmt_try_catch is called when exiting the stmt_try_catch production.
	ExitStmt_try_catch(c *Stmt_try_catchContext)

	// ExitStmt_throw is called when exiting the stmt_throw production.
	ExitStmt_throw(c *Stmt_throwContext)

	// ExitStmt_defer is called when exiting the stmt_defer production.
	ExitStmt_defer(c *Stmt_deferContext)

	// ExitStmt_yield is called when exiting the stmt_yield production.
	ExitStmt_yield(c *Stmt_yieldContext)

	// ExitStmt_break is called when exiting the stmt_break production.
	ExitStmt_break(c *Stmt_breakContext)

	// ExitStmt_continue is called when exiting the stmt_continue production.
	ExitStmt_continue(c *Stmt_continueContext)

	// ExitStmt_inc_dec is called when exiting the stmt_inc_dec production.
	ExitStmt_inc_dec(c *Stmt_inc_decContext)

	// ExitAssignment_stmt is called when exiting the assignment_stmt production.
	ExitAssignment_stmt(c *Assignment_stmtContext)

	// ExitAssignment_lvalues is called when exiting the assignment_lvalues production.
	ExitAssignment_lvalues(c *Assignment_lvaluesContext)

	// ExitRvalues is called when exiting the rvalues production.
	ExitRvalues(c *RvaluesContext)

	// ExitAssignment_op_stmt is called when exiting the assignment_op_stmt production.
	ExitAssignment_op_stmt(c *Assignment_op_stmtContext)

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

	// ExitBreak_stmt is called when exiting the break_stmt production.
	ExitBreak_stmt(c *Break_stmtContext)

	// ExitContinue_stmt is called when exiting the continue_stmt production.
	ExitContinue_stmt(c *Continue_stmtContext)

	// ExitInc_dec_stmt is called when exiting the inc_dec_stmt production.
	ExitInc_dec_stmt(c *Inc_dec_stmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr2 is called when exiting the expr2 production.
	ExitExpr2(c *Expr2Context)

	// ExitExpr3 is called when exiting the expr3 production.
	ExitExpr3(c *Expr3Context)

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
