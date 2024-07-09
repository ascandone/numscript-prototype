// Code generated from Numscript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Numscript

import "github.com/antlr4-go/antlr/v4"

type BaseNumscriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseNumscriptVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNumscriptVisitor) VisitMonetaryLit(ctx *MonetaryLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNumscriptVisitor) VisitSrcAccount(ctx *SrcAccountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNumscriptVisitor) VisitSrcVariable(ctx *SrcVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNumscriptVisitor) VisitDestAccount(ctx *DestAccountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNumscriptVisitor) VisitDestVariable(ctx *DestVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNumscriptVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}
