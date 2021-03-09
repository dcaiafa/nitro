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

// EnterStart is called when production start is entered.
func (s *BaseNitroParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseNitroParserListener) ExitStart(ctx *StartContext) {}

// EnterModule is called when production module is entered.
func (s *BaseNitroParserListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseNitroParserListener) ExitModule(ctx *ModuleContext) {}

// EnterMeta_section is called when production meta_section is entered.
func (s *BaseNitroParserListener) EnterMeta_section(ctx *Meta_sectionContext) {}

// ExitMeta_section is called when production meta_section is exited.
func (s *BaseNitroParserListener) ExitMeta_section(ctx *Meta_sectionContext) {}

// EnterMeta_entry is called when production meta_entry is entered.
func (s *BaseNitroParserListener) EnterMeta_entry(ctx *Meta_entryContext) {}

// ExitMeta_entry is called when production meta_entry is exited.
func (s *BaseNitroParserListener) ExitMeta_entry(ctx *Meta_entryContext) {}

// EnterMeta_object is called when production meta_object is entered.
func (s *BaseNitroParserListener) EnterMeta_object(ctx *Meta_objectContext) {}

// ExitMeta_object is called when production meta_object is exited.
func (s *BaseNitroParserListener) ExitMeta_object(ctx *Meta_objectContext) {}

// EnterMeta_fields is called when production meta_fields is entered.
func (s *BaseNitroParserListener) EnterMeta_fields(ctx *Meta_fieldsContext) {}

// ExitMeta_fields is called when production meta_fields is exited.
func (s *BaseNitroParserListener) ExitMeta_fields(ctx *Meta_fieldsContext) {}

// EnterMeta_field is called when production meta_field is entered.
func (s *BaseNitroParserListener) EnterMeta_field(ctx *Meta_fieldContext) {}

// ExitMeta_field is called when production meta_field is exited.
func (s *BaseNitroParserListener) ExitMeta_field(ctx *Meta_fieldContext) {}

// EnterMeta_field_value is called when production meta_field_value is entered.
func (s *BaseNitroParserListener) EnterMeta_field_value(ctx *Meta_field_valueContext) {}

// ExitMeta_field_value is called when production meta_field_value is exited.
func (s *BaseNitroParserListener) ExitMeta_field_value(ctx *Meta_field_valueContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BaseNitroParserListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BaseNitroParserListener) ExitStmts(ctx *StmtsContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseNitroParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseNitroParserListener) ExitStmt(ctx *StmtContext) {}

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

// EnterExpr is called when production expr is entered.
func (s *BaseNitroParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseNitroParserListener) ExitExpr(ctx *ExprContext) {}

// EnterUnary_expr is called when production unary_expr is entered.
func (s *BaseNitroParserListener) EnterUnary_expr(ctx *Unary_exprContext) {}

// ExitUnary_expr is called when production unary_expr is exited.
func (s *BaseNitroParserListener) ExitUnary_expr(ctx *Unary_exprContext) {}

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
