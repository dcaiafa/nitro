// Code generated from NitroParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // NitroParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NitroParserListener is a complete listener for a parse tree produced by NitroParser.
type NitroParserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterMeta_section is called when entering the meta_section production.
	EnterMeta_section(c *Meta_sectionContext)

	// EnterMeta_entry is called when entering the meta_entry production.
	EnterMeta_entry(c *Meta_entryContext)

	// EnterMeta_object is called when entering the meta_object production.
	EnterMeta_object(c *Meta_objectContext)

	// EnterMeta_fields is called when entering the meta_fields production.
	EnterMeta_fields(c *Meta_fieldsContext)

	// EnterMeta_field is called when entering the meta_field production.
	EnterMeta_field(c *Meta_fieldContext)

	// EnterMeta_field_value is called when entering the meta_field_value production.
	EnterMeta_field_value(c *Meta_field_valueContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

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

	// EnterExpr_binary is called when entering the expr_binary production.
	EnterExpr_binary(c *Expr_binaryContext)

	// EnterExpr_short_lambda is called when entering the expr_short_lambda production.
	EnterExpr_short_lambda(c *Expr_short_lambdaContext)

	// EnterExpr_pipeline is called when entering the expr_pipeline production.
	EnterExpr_pipeline(c *Expr_pipelineContext)

	// EnterPipeline_expr is called when entering the pipeline_expr production.
	EnterPipeline_expr(c *Pipeline_exprContext)

	// EnterBinary_expr is called when entering the binary_expr production.
	EnterBinary_expr(c *Binary_exprContext)

	// EnterUnary_expr is called when entering the unary_expr production.
	EnterUnary_expr(c *Unary_exprContext)

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

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitMeta_section is called when exiting the meta_section production.
	ExitMeta_section(c *Meta_sectionContext)

	// ExitMeta_entry is called when exiting the meta_entry production.
	ExitMeta_entry(c *Meta_entryContext)

	// ExitMeta_object is called when exiting the meta_object production.
	ExitMeta_object(c *Meta_objectContext)

	// ExitMeta_fields is called when exiting the meta_fields production.
	ExitMeta_fields(c *Meta_fieldsContext)

	// ExitMeta_field is called when exiting the meta_field production.
	ExitMeta_field(c *Meta_fieldContext)

	// ExitMeta_field_value is called when exiting the meta_field_value production.
	ExitMeta_field_value(c *Meta_field_valueContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

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

	// ExitExpr_binary is called when exiting the expr_binary production.
	ExitExpr_binary(c *Expr_binaryContext)

	// ExitExpr_short_lambda is called when exiting the expr_short_lambda production.
	ExitExpr_short_lambda(c *Expr_short_lambdaContext)

	// ExitExpr_pipeline is called when exiting the expr_pipeline production.
	ExitExpr_pipeline(c *Expr_pipelineContext)

	// ExitPipeline_expr is called when exiting the pipeline_expr production.
	ExitPipeline_expr(c *Pipeline_exprContext)

	// ExitBinary_expr is called when exiting the binary_expr production.
	ExitBinary_expr(c *Binary_exprContext)

	// ExitUnary_expr is called when exiting the unary_expr production.
	ExitUnary_expr(c *Unary_exprContext)

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
