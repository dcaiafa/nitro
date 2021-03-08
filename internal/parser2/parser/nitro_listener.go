// Code generated from Nitro.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Nitro

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NitroListener is a complete listener for a parse tree produced by NitroParser.
type NitroListener interface {
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

	// EnterAssignment_rvalues is called when entering the assignment_rvalues production.
	EnterAssignment_rvalues(c *Assignment_rvaluesContext)

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

	// EnterFunc_call_stmt is called when entering the func_call_stmt production.
	EnterFunc_call_stmt(c *Func_call_stmtContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterUnary_expr is called when entering the unary_expr production.
	EnterUnary_expr(c *Unary_exprContext)

	// EnterPrimary_expr is called when entering the primary_expr production.
	EnterPrimary_expr(c *Primary_exprContext)

	// EnterArg_list is called when entering the arg_list production.
	EnterArg_list(c *Arg_listContext)

	// EnterLvalue_expr is called when entering the lvalue_expr production.
	EnterLvalue_expr(c *Lvalue_exprContext)

	// EnterObject_literal is called when entering the object_literal production.
	EnterObject_literal(c *Object_literalContext)

	// EnterObject_fields is called when entering the object_fields production.
	EnterObject_fields(c *Object_fieldsContext)

	// EnterObject_field is called when entering the object_field production.
	EnterObject_field(c *Object_fieldContext)

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

	// ExitAssignment_rvalues is called when exiting the assignment_rvalues production.
	ExitAssignment_rvalues(c *Assignment_rvaluesContext)

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

	// ExitFunc_call_stmt is called when exiting the func_call_stmt production.
	ExitFunc_call_stmt(c *Func_call_stmtContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitUnary_expr is called when exiting the unary_expr production.
	ExitUnary_expr(c *Unary_exprContext)

	// ExitPrimary_expr is called when exiting the primary_expr production.
	ExitPrimary_expr(c *Primary_exprContext)

	// ExitArg_list is called when exiting the arg_list production.
	ExitArg_list(c *Arg_listContext)

	// ExitLvalue_expr is called when exiting the lvalue_expr production.
	ExitLvalue_expr(c *Lvalue_exprContext)

	// ExitObject_literal is called when exiting the object_literal production.
	ExitObject_literal(c *Object_literalContext)

	// ExitObject_fields is called when exiting the object_fields production.
	ExitObject_fields(c *Object_fieldsContext)

	// ExitObject_field is called when exiting the object_field production.
	ExitObject_field(c *Object_fieldContext)

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
