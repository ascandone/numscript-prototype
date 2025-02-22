// Code generated from Numscript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Numscript

import "github.com/antlr4-go/antlr/v4"

// BaseNumscriptListener is a complete listener for a parse tree produced by NumscriptParser.
type BaseNumscriptListener struct{}

var _ NumscriptListener = &BaseNumscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNumscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNumscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNumscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNumscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAssetLiteral is called when production assetLiteral is entered.
func (s *BaseNumscriptListener) EnterAssetLiteral(ctx *AssetLiteralContext) {}

// ExitAssetLiteral is called when production assetLiteral is exited.
func (s *BaseNumscriptListener) ExitAssetLiteral(ctx *AssetLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseNumscriptListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseNumscriptListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterMonetaryLiteral is called when production monetaryLiteral is entered.
func (s *BaseNumscriptListener) EnterMonetaryLiteral(ctx *MonetaryLiteralContext) {}

// ExitMonetaryLiteral is called when production monetaryLiteral is exited.
func (s *BaseNumscriptListener) ExitMonetaryLiteral(ctx *MonetaryLiteralContext) {}

// EnterAccountLiteral is called when production accountLiteral is entered.
func (s *BaseNumscriptListener) EnterAccountLiteral(ctx *AccountLiteralContext) {}

// ExitAccountLiteral is called when production accountLiteral is exited.
func (s *BaseNumscriptListener) ExitAccountLiteral(ctx *AccountLiteralContext) {}

// EnterVariableLiteral is called when production variableLiteral is entered.
func (s *BaseNumscriptListener) EnterVariableLiteral(ctx *VariableLiteralContext) {}

// ExitVariableLiteral is called when production variableLiteral is exited.
func (s *BaseNumscriptListener) ExitVariableLiteral(ctx *VariableLiteralContext) {}

// EnterPortionLiteral is called when production portionLiteral is entered.
func (s *BaseNumscriptListener) EnterPortionLiteral(ctx *PortionLiteralContext) {}

// ExitPortionLiteral is called when production portionLiteral is exited.
func (s *BaseNumscriptListener) ExitPortionLiteral(ctx *PortionLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseNumscriptListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseNumscriptListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterRatio is called when production ratio is entered.
func (s *BaseNumscriptListener) EnterRatio(ctx *RatioContext) {}

// ExitRatio is called when production ratio is exited.
func (s *BaseNumscriptListener) ExitRatio(ctx *RatioContext) {}

// EnterPercentage is called when production percentage is entered.
func (s *BaseNumscriptListener) EnterPercentage(ctx *PercentageContext) {}

// ExitPercentage is called when production percentage is exited.
func (s *BaseNumscriptListener) ExitPercentage(ctx *PercentageContext) {}

// EnterFunctionCallArgs is called when production functionCallArgs is entered.
func (s *BaseNumscriptListener) EnterFunctionCallArgs(ctx *FunctionCallArgsContext) {}

// ExitFunctionCallArgs is called when production functionCallArgs is exited.
func (s *BaseNumscriptListener) ExitFunctionCallArgs(ctx *FunctionCallArgsContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseNumscriptListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseNumscriptListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterVarOrigin is called when production varOrigin is entered.
func (s *BaseNumscriptListener) EnterVarOrigin(ctx *VarOriginContext) {}

// ExitVarOrigin is called when production varOrigin is exited.
func (s *BaseNumscriptListener) ExitVarOrigin(ctx *VarOriginContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseNumscriptListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseNumscriptListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterVarsDeclaration is called when production varsDeclaration is entered.
func (s *BaseNumscriptListener) EnterVarsDeclaration(ctx *VarsDeclarationContext) {}

// ExitVarsDeclaration is called when production varsDeclaration is exited.
func (s *BaseNumscriptListener) ExitVarsDeclaration(ctx *VarsDeclarationContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseNumscriptListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseNumscriptListener) ExitProgram(ctx *ProgramContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseNumscriptListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseNumscriptListener) ExitNumber(ctx *NumberContext) {}

// EnterNumberVariable is called when production numberVariable is entered.
func (s *BaseNumscriptListener) EnterNumberVariable(ctx *NumberVariableContext) {}

// ExitNumberVariable is called when production numberVariable is exited.
func (s *BaseNumscriptListener) ExitNumberVariable(ctx *NumberVariableContext) {}

// EnterAsset is called when production asset is entered.
func (s *BaseNumscriptListener) EnterAsset(ctx *AssetContext) {}

// ExitAsset is called when production asset is exited.
func (s *BaseNumscriptListener) ExitAsset(ctx *AssetContext) {}

// EnterAssetVariable is called when production assetVariable is entered.
func (s *BaseNumscriptListener) EnterAssetVariable(ctx *AssetVariableContext) {}

// ExitAssetVariable is called when production assetVariable is exited.
func (s *BaseNumscriptListener) ExitAssetVariable(ctx *AssetVariableContext) {}

// EnterAccountName is called when production accountName is entered.
func (s *BaseNumscriptListener) EnterAccountName(ctx *AccountNameContext) {}

// ExitAccountName is called when production accountName is exited.
func (s *BaseNumscriptListener) ExitAccountName(ctx *AccountNameContext) {}

// EnterAccountVariable is called when production accountVariable is entered.
func (s *BaseNumscriptListener) EnterAccountVariable(ctx *AccountVariableContext) {}

// ExitAccountVariable is called when production accountVariable is exited.
func (s *BaseNumscriptListener) ExitAccountVariable(ctx *AccountVariableContext) {}

// EnterMonetary is called when production monetary is entered.
func (s *BaseNumscriptListener) EnterMonetary(ctx *MonetaryContext) {}

// ExitMonetary is called when production monetary is exited.
func (s *BaseNumscriptListener) ExitMonetary(ctx *MonetaryContext) {}

// EnterMonetaryVariable is called when production monetaryVariable is entered.
func (s *BaseNumscriptListener) EnterMonetaryVariable(ctx *MonetaryVariableContext) {}

// ExitMonetaryVariable is called when production monetaryVariable is exited.
func (s *BaseNumscriptListener) ExitMonetaryVariable(ctx *MonetaryVariableContext) {}

// EnterMonetaryLit is called when production monetaryLit is entered.
func (s *BaseNumscriptListener) EnterMonetaryLit(ctx *MonetaryLitContext) {}

// ExitMonetaryLit is called when production monetaryLit is exited.
func (s *BaseNumscriptListener) ExitMonetaryLit(ctx *MonetaryLitContext) {}

// EnterSentAllLit is called when production sentAllLit is entered.
func (s *BaseNumscriptListener) EnterSentAllLit(ctx *SentAllLitContext) {}

// ExitSentAllLit is called when production sentAllLit is exited.
func (s *BaseNumscriptListener) ExitSentAllLit(ctx *SentAllLitContext) {}

// EnterLitCap is called when production litCap is entered.
func (s *BaseNumscriptListener) EnterLitCap(ctx *LitCapContext) {}

// ExitLitCap is called when production litCap is exited.
func (s *BaseNumscriptListener) ExitLitCap(ctx *LitCapContext) {}

// EnterVarCap is called when production varCap is entered.
func (s *BaseNumscriptListener) EnterVarCap(ctx *VarCapContext) {}

// ExitVarCap is called when production varCap is exited.
func (s *BaseNumscriptListener) ExitVarCap(ctx *VarCapContext) {}

// EnterPortionedAllotment is called when production portionedAllotment is entered.
func (s *BaseNumscriptListener) EnterPortionedAllotment(ctx *PortionedAllotmentContext) {}

// ExitPortionedAllotment is called when production portionedAllotment is exited.
func (s *BaseNumscriptListener) ExitPortionedAllotment(ctx *PortionedAllotmentContext) {}

// EnterPortionVariable is called when production portionVariable is entered.
func (s *BaseNumscriptListener) EnterPortionVariable(ctx *PortionVariableContext) {}

// ExitPortionVariable is called when production portionVariable is exited.
func (s *BaseNumscriptListener) ExitPortionVariable(ctx *PortionVariableContext) {}

// EnterRemainingAllotment is called when production remainingAllotment is entered.
func (s *BaseNumscriptListener) EnterRemainingAllotment(ctx *RemainingAllotmentContext) {}

// ExitRemainingAllotment is called when production remainingAllotment is exited.
func (s *BaseNumscriptListener) ExitRemainingAllotment(ctx *RemainingAllotmentContext) {}

// EnterSrcAccountUnboundedOverdraft is called when production srcAccountUnboundedOverdraft is entered.
func (s *BaseNumscriptListener) EnterSrcAccountUnboundedOverdraft(ctx *SrcAccountUnboundedOverdraftContext) {
}

// ExitSrcAccountUnboundedOverdraft is called when production srcAccountUnboundedOverdraft is exited.
func (s *BaseNumscriptListener) ExitSrcAccountUnboundedOverdraft(ctx *SrcAccountUnboundedOverdraftContext) {
}

// EnterSrcAccountBoundedOverdraft is called when production srcAccountBoundedOverdraft is entered.
func (s *BaseNumscriptListener) EnterSrcAccountBoundedOverdraft(ctx *SrcAccountBoundedOverdraftContext) {
}

// ExitSrcAccountBoundedOverdraft is called when production srcAccountBoundedOverdraft is exited.
func (s *BaseNumscriptListener) ExitSrcAccountBoundedOverdraft(ctx *SrcAccountBoundedOverdraftContext) {
}

// EnterSrcAccount is called when production srcAccount is entered.
func (s *BaseNumscriptListener) EnterSrcAccount(ctx *SrcAccountContext) {}

// ExitSrcAccount is called when production srcAccount is exited.
func (s *BaseNumscriptListener) ExitSrcAccount(ctx *SrcAccountContext) {}

// EnterSrcVariable is called when production srcVariable is entered.
func (s *BaseNumscriptListener) EnterSrcVariable(ctx *SrcVariableContext) {}

// ExitSrcVariable is called when production srcVariable is exited.
func (s *BaseNumscriptListener) ExitSrcVariable(ctx *SrcVariableContext) {}

// EnterSrcAllotment is called when production srcAllotment is entered.
func (s *BaseNumscriptListener) EnterSrcAllotment(ctx *SrcAllotmentContext) {}

// ExitSrcAllotment is called when production srcAllotment is exited.
func (s *BaseNumscriptListener) ExitSrcAllotment(ctx *SrcAllotmentContext) {}

// EnterSrcInorder is called when production srcInorder is entered.
func (s *BaseNumscriptListener) EnterSrcInorder(ctx *SrcInorderContext) {}

// ExitSrcInorder is called when production srcInorder is exited.
func (s *BaseNumscriptListener) ExitSrcInorder(ctx *SrcInorderContext) {}

// EnterSrcCapped is called when production srcCapped is entered.
func (s *BaseNumscriptListener) EnterSrcCapped(ctx *SrcCappedContext) {}

// ExitSrcCapped is called when production srcCapped is exited.
func (s *BaseNumscriptListener) ExitSrcCapped(ctx *SrcCappedContext) {}

// EnterAllotmentClauseSrc is called when production allotmentClauseSrc is entered.
func (s *BaseNumscriptListener) EnterAllotmentClauseSrc(ctx *AllotmentClauseSrcContext) {}

// ExitAllotmentClauseSrc is called when production allotmentClauseSrc is exited.
func (s *BaseNumscriptListener) ExitAllotmentClauseSrc(ctx *AllotmentClauseSrcContext) {}

// EnterDestinationTo is called when production destinationTo is entered.
func (s *BaseNumscriptListener) EnterDestinationTo(ctx *DestinationToContext) {}

// ExitDestinationTo is called when production destinationTo is exited.
func (s *BaseNumscriptListener) ExitDestinationTo(ctx *DestinationToContext) {}

// EnterDestinationKept is called when production destinationKept is entered.
func (s *BaseNumscriptListener) EnterDestinationKept(ctx *DestinationKeptContext) {}

// ExitDestinationKept is called when production destinationKept is exited.
func (s *BaseNumscriptListener) ExitDestinationKept(ctx *DestinationKeptContext) {}

// EnterDestinationInOrderClause is called when production destinationInOrderClause is entered.
func (s *BaseNumscriptListener) EnterDestinationInOrderClause(ctx *DestinationInOrderClauseContext) {}

// ExitDestinationInOrderClause is called when production destinationInOrderClause is exited.
func (s *BaseNumscriptListener) ExitDestinationInOrderClause(ctx *DestinationInOrderClauseContext) {}

// EnterDestAccount is called when production destAccount is entered.
func (s *BaseNumscriptListener) EnterDestAccount(ctx *DestAccountContext) {}

// ExitDestAccount is called when production destAccount is exited.
func (s *BaseNumscriptListener) ExitDestAccount(ctx *DestAccountContext) {}

// EnterDestVariable is called when production destVariable is entered.
func (s *BaseNumscriptListener) EnterDestVariable(ctx *DestVariableContext) {}

// ExitDestVariable is called when production destVariable is exited.
func (s *BaseNumscriptListener) ExitDestVariable(ctx *DestVariableContext) {}

// EnterDestAllotment is called when production destAllotment is entered.
func (s *BaseNumscriptListener) EnterDestAllotment(ctx *DestAllotmentContext) {}

// ExitDestAllotment is called when production destAllotment is exited.
func (s *BaseNumscriptListener) ExitDestAllotment(ctx *DestAllotmentContext) {}

// EnterDestInorder is called when production destInorder is entered.
func (s *BaseNumscriptListener) EnterDestInorder(ctx *DestInorderContext) {}

// ExitDestInorder is called when production destInorder is exited.
func (s *BaseNumscriptListener) ExitDestInorder(ctx *DestInorderContext) {}

// EnterAllotmentClauseDest is called when production allotmentClauseDest is entered.
func (s *BaseNumscriptListener) EnterAllotmentClauseDest(ctx *AllotmentClauseDestContext) {}

// ExitAllotmentClauseDest is called when production allotmentClauseDest is exited.
func (s *BaseNumscriptListener) ExitAllotmentClauseDest(ctx *AllotmentClauseDestContext) {}

// EnterSentLiteral is called when production sentLiteral is entered.
func (s *BaseNumscriptListener) EnterSentLiteral(ctx *SentLiteralContext) {}

// ExitSentLiteral is called when production sentLiteral is exited.
func (s *BaseNumscriptListener) ExitSentLiteral(ctx *SentLiteralContext) {}

// EnterSentAll is called when production sentAll is entered.
func (s *BaseNumscriptListener) EnterSentAll(ctx *SentAllContext) {}

// ExitSentAll is called when production sentAll is exited.
func (s *BaseNumscriptListener) ExitSentAll(ctx *SentAllContext) {}

// EnterSendStatement is called when production sendStatement is entered.
func (s *BaseNumscriptListener) EnterSendStatement(ctx *SendStatementContext) {}

// ExitSendStatement is called when production sendStatement is exited.
func (s *BaseNumscriptListener) ExitSendStatement(ctx *SendStatementContext) {}

// EnterFnCallStatement is called when production fnCallStatement is entered.
func (s *BaseNumscriptListener) EnterFnCallStatement(ctx *FnCallStatementContext) {}

// ExitFnCallStatement is called when production fnCallStatement is exited.
func (s *BaseNumscriptListener) ExitFnCallStatement(ctx *FnCallStatementContext) {}
