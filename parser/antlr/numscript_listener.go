// Code generated from Numscript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Numscript

import "github.com/antlr4-go/antlr/v4"

// NumscriptListener is a complete listener for a parse tree produced by NumscriptParser.
type NumscriptListener interface {
	antlr.ParseTreeListener

	// EnterRatio is called when entering the ratio production.
	EnterRatio(c *RatioContext)

	// EnterPercentage is called when entering the percentage production.
	EnterPercentage(c *PercentageContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterVarsDeclaration is called when entering the varsDeclaration production.
	EnterVarsDeclaration(c *VarsDeclarationContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterMonetaryLit is called when entering the monetaryLit production.
	EnterMonetaryLit(c *MonetaryLitContext)

	// EnterLitCap is called when entering the litCap production.
	EnterLitCap(c *LitCapContext)

	// EnterVarCap is called when entering the varCap production.
	EnterVarCap(c *VarCapContext)

	// EnterSrcAccount is called when entering the srcAccount production.
	EnterSrcAccount(c *SrcAccountContext)

	// EnterSrcVariable is called when entering the srcVariable production.
	EnterSrcVariable(c *SrcVariableContext)

	// EnterSrcAllotment is called when entering the srcAllotment production.
	EnterSrcAllotment(c *SrcAllotmentContext)

	// EnterSrcSeq is called when entering the srcSeq production.
	EnterSrcSeq(c *SrcSeqContext)

	// EnterSrcCapped is called when entering the srcCapped production.
	EnterSrcCapped(c *SrcCappedContext)

	// EnterAllotmentClauseSrc is called when entering the allotmentClauseSrc production.
	EnterAllotmentClauseSrc(c *AllotmentClauseSrcContext)

	// EnterDestAccount is called when entering the destAccount production.
	EnterDestAccount(c *DestAccountContext)

	// EnterDestVariable is called when entering the destVariable production.
	EnterDestVariable(c *DestVariableContext)

	// EnterDestAllotment is called when entering the destAllotment production.
	EnterDestAllotment(c *DestAllotmentContext)

	// EnterDestSeq is called when entering the destSeq production.
	EnterDestSeq(c *DestSeqContext)

	// EnterAllotmentClauseDest is called when entering the allotmentClauseDest production.
	EnterAllotmentClauseDest(c *AllotmentClauseDestContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// ExitRatio is called when exiting the ratio production.
	ExitRatio(c *RatioContext)

	// ExitPercentage is called when exiting the percentage production.
	ExitPercentage(c *PercentageContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitVarsDeclaration is called when exiting the varsDeclaration production.
	ExitVarsDeclaration(c *VarsDeclarationContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitMonetaryLit is called when exiting the monetaryLit production.
	ExitMonetaryLit(c *MonetaryLitContext)

	// ExitLitCap is called when exiting the litCap production.
	ExitLitCap(c *LitCapContext)

	// ExitVarCap is called when exiting the varCap production.
	ExitVarCap(c *VarCapContext)

	// ExitSrcAccount is called when exiting the srcAccount production.
	ExitSrcAccount(c *SrcAccountContext)

	// ExitSrcVariable is called when exiting the srcVariable production.
	ExitSrcVariable(c *SrcVariableContext)

	// ExitSrcAllotment is called when exiting the srcAllotment production.
	ExitSrcAllotment(c *SrcAllotmentContext)

	// ExitSrcSeq is called when exiting the srcSeq production.
	ExitSrcSeq(c *SrcSeqContext)

	// ExitSrcCapped is called when exiting the srcCapped production.
	ExitSrcCapped(c *SrcCappedContext)

	// ExitAllotmentClauseSrc is called when exiting the allotmentClauseSrc production.
	ExitAllotmentClauseSrc(c *AllotmentClauseSrcContext)

	// ExitDestAccount is called when exiting the destAccount production.
	ExitDestAccount(c *DestAccountContext)

	// ExitDestVariable is called when exiting the destVariable production.
	ExitDestVariable(c *DestVariableContext)

	// ExitDestAllotment is called when exiting the destAllotment production.
	ExitDestAllotment(c *DestAllotmentContext)

	// ExitDestSeq is called when exiting the destSeq production.
	ExitDestSeq(c *DestSeqContext)

	// ExitAllotmentClauseDest is called when exiting the allotmentClauseDest production.
	ExitAllotmentClauseDest(c *AllotmentClauseDestContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)
}
