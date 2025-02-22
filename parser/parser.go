package parser

import (
	"math"
	parser "numscript/parser/antlr"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ParserError struct {
	Range Range
	Msg   string
}

type ParseResult[T any] struct {
	Value  T
	Errors []ParserError
}

type ErrorListener struct {
	antlr.DefaultErrorListener
	Errors []ParserError
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, startL, startC int, msg string, e antlr.RecognitionException) {
	length := 1
	if token, ok := offendingSymbol.(antlr.Token); ok {
		length = len(token.GetText())
	}
	endL := startL
	endC := startC + length - 1 // -1 so that end character is inside the offending token
	l.Errors = append(l.Errors, ParserError{
		Msg: msg,
		Range: Range{
			Start: Position{Character: startC, Line: startL - 1},
			End:   Position{Character: endC, Line: endL - 1},
		},
	})
}

func Parse(input string) ParseResult[Program] {
	// TODO handle lexer errors
	listener := &ErrorListener{}

	is := antlr.NewInputStream(input)
	lexer := parser.NewNumscriptLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewNumscriptParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)

	parsed := parseProgram(parser.Program())

	return ParseResult[Program]{
		Value:  parsed,
		Errors: listener.Errors,
	}
}

func parseVarsDeclaration(varsCtx parser.IVarsDeclarationContext) []VarDeclaration {
	if varsCtx == nil {
		return nil
	}

	var vars []VarDeclaration
	for _, varDecl := range varsCtx.AllVarDeclaration() {
		decl := parseVarDeclaration(varDecl)
		if decl != nil {
			vars = append(vars, *decl)
		}
	}
	return vars
}

func parseProgram(programCtx parser.IProgramContext) Program {
	vars := parseVarsDeclaration(programCtx.VarsDeclaration())

	var statements []Statement
	for _, statementCtx := range programCtx.AllStatement() {
		statements = append(statements, parseStatement(statementCtx))
	}

	return Program{
		Statements: statements,
		Vars:       vars,
	}
}

func parseVarDeclaration(varDecl parser.IVarDeclarationContext) *VarDeclaration {
	if varDecl == nil {
		return nil
	}

	var fnCallStatement *FnCall
	if varDecl.VarOrigin() != nil {
		fnCallStatement = parseFnCall(varDecl.VarOrigin().FunctionCall())
	}

	return &VarDeclaration{
		Range:  ctxToRange(varDecl),
		Type:   parseVarType(varDecl.GetType_()),
		Name:   parseVarLiteral(varDecl.GetName()),
		Origin: fnCallStatement,
	}
}

func parseVarLiteral(tk antlr.Token) *VariableLiteral {
	if tk == nil || tk.GetTokenIndex() == -1 {
		return nil
	}

	name := tk.GetText()[1:]

	return &VariableLiteral{
		Range: tokenToRange(tk),
		Name:  name,
	}
}

func parseVarType(tk antlr.Token) *TypeDecl {
	if tk == nil || tk.GetTokenIndex() == -1 {
		return nil
	}

	return &TypeDecl{
		Range: tokenToRange(tk),
		Name:  tk.GetText(),
	}
}

func parseCapLit(capCtx parser.ICapContext) Literal {
	switch capCtx := capCtx.(type) {
	case *parser.LitCapContext:
		return parseMonetaryLit(capCtx.MonetaryLit())

	case *parser.VarCapContext:
		return variableLiteralFromCtx(capCtx)

	case *parser.CapContext:
		return nil

	default:
		panic("Invalid ctx")
	}
}

func parseSource(sourceCtx parser.ISourceContext) Source {
	if sourceCtx == nil {
		return nil
	}

	range_ := ctxToRange(sourceCtx)

	switch sourceCtx := sourceCtx.(type) {
	case *parser.SrcAccountContext:
		return accountLiteralFromCtx(sourceCtx)

	case *parser.SrcVariableContext:
		return variableLiteralFromCtx(sourceCtx)

	case *parser.SrcInorderContext:
		var sources []Source
		for _, sourceCtx := range sourceCtx.AllSource() {
			sources = append(sources, parseSource(sourceCtx))
		}
		return &SourceInorder{
			Range:   range_,
			Sources: sources,
		}

	case *parser.SrcAllotmentContext:
		var items []SourceAllotmentItem
		for _, itemCtx := range sourceCtx.AllAllotmentClauseSrc() {
			item := SourceAllotmentItem{
				Range:     ctxToRange(itemCtx),
				Allotment: parseAllotment(itemCtx.Allotment()),
				From:      parseSource(itemCtx.Source()),
			}
			items = append(items, item)
		}
		return &SourceAllotment{
			Range: range_,
			Items: items,
		}

	case *parser.SrcCappedContext:
		return &SourceCapped{
			Range: range_,
			From:  parseSource(sourceCtx.Source()),
			Cap:   parseCapLit(sourceCtx.Cap_()),
		}

	case *parser.SrcAccountUnboundedOverdraftContext:
		return &SourceOverdraft{
			Range:   ctxToRange(sourceCtx),
			Address: parseVariableAccount(sourceCtx.VariableAccount()),
		}

	case *parser.SrcAccountBoundedOverdraftContext:
		varMon := parseVariableMonetary(sourceCtx.VariableMonetary())

		return &SourceOverdraft{
			Range:   ctxToRange(sourceCtx),
			Address: parseVariableAccount(sourceCtx.VariableAccount()),
			Bounded: &varMon,
		}

	case *parser.SourceContext:
		return nil

	default:
		panic("unhandled context: " + sourceCtx.GetText())
	}
}

func parseVariableAccount(variableAccountCtx parser.IVariableAccountContext) Literal {
	switch variableAccountCtx := variableAccountCtx.(type) {
	case *parser.AccountNameContext:
		return accountLiteralFromCtx(variableAccountCtx)

	case *parser.AccountVariableContext:
		return variableLiteralFromCtx(variableAccountCtx)

	case *parser.VariableAccountContext:
		return nil

	default:
		panic("Unhandled clause")
	}

}

func parseRatio(source string, range_ Range) *RatioLiteral {
	split := strings.Split(source, "/")

	num, err := strconv.ParseUint(split[0], 0, 64)
	if err != nil {
		panic(err)
	}

	den, err := strconv.ParseUint(split[1], 0, 64)
	if err != nil {
		panic(err)
	}

	return &RatioLiteral{
		Range:       range_,
		Numerator:   num,
		Denominator: den,
	}
}

func parsePercentageRatio(source string, range_ Range) *RatioLiteral {
	str := strings.TrimSuffix(source, "%")
	num, err := strconv.ParseUint(strings.Replace(str, ".", "", -1), 0, 64)
	if err != nil {
		panic(err)
	}

	var denominator uint64
	split := strings.Split(str, ".")
	if len(split) > 1 {
		// TODO verify this is always correct
		floatingDigits := len(split[1])
		denominator = (uint64)(math.Pow10(2 + floatingDigits))
	} else {
		denominator = 100
	}

	return &RatioLiteral{
		Range:       range_,
		Numerator:   num,
		Denominator: denominator,
	}
}

func parseAllotment(allotmentCtx parser.IAllotmentContext) SourceAllotmentValue {
	switch allotmentCtx := allotmentCtx.(type) {
	case *parser.PortionedAllotmentContext:
		return parsePortionSource(allotmentCtx.Portion())

	case *parser.RemainingAllotmentContext:
		return &RemainingAllotment{
			Range: ctxToRange(allotmentCtx),
		}

	case *parser.PortionVariableContext:
		return variableLiteralFromCtx(allotmentCtx)

	case *parser.AllotmentContext:
		return nil

	default:
		panic("Invalid allotment")
	}
}

func parseStringLiteralCtx(stringCtx *parser.StringLiteralContext) *StringLiteral {
	rawStr := stringCtx.GetText()
	// Remove leading and trailing '"'
	innerStr := rawStr[1 : len(rawStr)-1]
	return &StringLiteral{
		Range:  ctxToRange(stringCtx),
		String: innerStr,
	}
}

func parseLiteral(literalCtx parser.ILiteralContext) Literal {
	switch literalCtx := literalCtx.(type) {
	case *parser.AccountLiteralContext:
		return accountLiteralFromCtx(literalCtx)

	case *parser.MonetaryLiteralContext:
		return parseMonetaryLit(literalCtx.MonetaryLit())

	case *parser.AssetLiteralContext:
		return &AssetLiteral{
			Range: ctxToRange(literalCtx),
			Asset: literalCtx.GetText(),
		}

	case *parser.NumberLiteralContext:
		return parseNumberLiteral(literalCtx.NUMBER())

	case *parser.PortionLiteralContext:
		return parsePortionSource(literalCtx.Portion())

	case *parser.VariableLiteralContext:
		return variableLiteralFromCtx(literalCtx)

	case *parser.StringLiteralContext:
		return parseStringLiteralCtx(literalCtx)

	case *parser.LiteralContext:
		return nil

	default:
		panic("unhandled clause")
	}
}

func variableLiteralFromCtx(ctx antlr.ParserRuleContext) *VariableLiteral {
	// Discard the '$'
	name := ctx.GetText()[1:]

	return &VariableLiteral{
		Range: ctxToRange(ctx),
		Name:  name,
	}
}

func accountLiteralFromCtx(ctx antlr.ParserRuleContext) *AccountLiteral {
	// Discard the '@'
	name := ctx.GetText()[1:]

	return &AccountLiteral{
		Range: ctxToRange(ctx),
		Name:  name,
	}
}

func parsePortionSource(portionCtx parser.IPortionContext) *RatioLiteral {
	switch portionCtx.(type) {
	case *parser.RatioContext:
		return parseRatio(portionCtx.GetText(), ctxToRange(portionCtx))

	case *parser.PercentageContext:
		return parsePercentageRatio(portionCtx.GetText(), ctxToRange(portionCtx))

	case *parser.PortionContext:
		return nil

	default:
		panic("unhandled portion ctx")
	}
}

func parseDestination(destCtx parser.IDestinationContext) Destination {
	if destCtx == nil {
		return nil
	}

	range_ := ctxToRange(destCtx)

	switch destCtx := destCtx.(type) {
	case *parser.DestAccountContext:
		return accountLiteralFromCtx(destCtx)

	case *parser.DestVariableContext:
		return variableLiteralFromCtx(destCtx)

	case *parser.DestInorderContext:
		var inorderClauses []DestinationInorderClause
		for _, destInorderClause := range destCtx.AllDestinationInOrderClause() {
			inorderClauses = append(inorderClauses, parseDestinationInorderClause(destInorderClause))
		}

		return &DestinationInorder{
			Range:     range_,
			Clauses:   inorderClauses,
			Remaining: parseKeptOrDestination(destCtx.KeptOrDestination()),
		}

	case *parser.DestAllotmentContext:
		var items []DestinationAllotmentItem
		for _, itemCtx := range destCtx.AllAllotmentClauseDest() {
			item := DestinationAllotmentItem{
				Range:     ctxToRange(itemCtx),
				Allotment: parseDestinationAllotment(itemCtx.Allotment()),
				To:        parseKeptOrDestination(itemCtx.KeptOrDestination()),
			}
			items = append(items, item)
		}
		return &DestinationAllotment{
			Range: range_,
			Items: items,
		}

	case *parser.DestinationContext:
		return nil

	default:
		panic("Unhandled dest" + destCtx.GetText())

	}

}

func parseDestinationInorderClause(clauseCtx parser.IDestinationInOrderClauseContext) DestinationInorderClause {
	return DestinationInorderClause{
		Range: ctxToRange(clauseCtx),
		Cap:   parseLiteral(clauseCtx.Literal()),
		To:    parseKeptOrDestination(clauseCtx.KeptOrDestination()),
	}
}

func parseKeptOrDestination(clauseCtx parser.IKeptOrDestinationContext) KeptOrDestination {
	if clauseCtx == nil {
		return nil
	}

	switch clauseCtx := clauseCtx.(type) {
	case *parser.DestinationToContext:
		return &DestinationTo{
			Destination: parseDestination(clauseCtx.Destination()),
		}
	case *parser.DestinationKeptContext:
		return &DestinationKept{
			Range: ctxToRange(clauseCtx),
		}
	case *parser.KeptOrDestinationContext:
		return nil

	default:
		panic("Unhandled clause")
	}

}

func parseDestinationAllotment(allotmentCtx parser.IAllotmentContext) DestinationAllotmentValue {
	switch allotmentCtx := allotmentCtx.(type) {
	case *parser.RemainingAllotmentContext:
		return &RemainingAllotment{
			Range: ctxToRange(allotmentCtx),
		}

	case *parser.PortionedAllotmentContext:
		return parseDestinationPortion(allotmentCtx.Portion())

	case *parser.PortionVariableContext:
		return variableLiteralFromCtx(allotmentCtx)

	case *parser.AllotmentContext:
		return nil

	default:
		panic("unhandled portion ctx")
	}
}

func parseDestinationPortion(portionCtx parser.IPortionContext) DestinationAllotmentValue {
	switch portionCtx.(type) {
	case *parser.RatioContext:
		return parseRatio(portionCtx.GetText(), ctxToRange(portionCtx))

	case *parser.PercentageContext:
		return parsePercentageRatio(portionCtx.GetText(), ctxToRange(portionCtx))

	case *parser.PortionContext:
		return nil

	default:
		panic("unhandled portion ctx")
	}
}

func parseFnArgs(fnCallArgCtx parser.IFunctionCallArgsContext) []Literal {
	if fnCallArgCtx == nil {
		return nil
	}

	var args []Literal
	for _, literal := range fnCallArgCtx.AllLiteral() {
		args = append(args, parseLiteral(literal))
	}
	return args
}

func parseFnCall(fnCallCtx parser.IFunctionCallContext) *FnCall {
	if fnCallCtx == nil {
		return nil
	}

	ident := fnCallCtx.IDENTIFIER()
	if ident == nil {
		return nil
	}

	allArgs := fnCallCtx.FunctionCallArgs()

	return &FnCall{
		Range: ctxToRange(fnCallCtx),
		Caller: &FnCallIdentifier{
			Range: tokenToRange(ident.GetSymbol()),
			Name:  ident.GetSymbol().GetText(),
		},
		Args: parseFnArgs(allArgs),
	}
}

func parseStatement(statementCtx parser.IStatementContext) Statement {
	switch statementCtx := statementCtx.(type) {
	case *parser.SendStatementContext:
		return parseSendStatement(statementCtx)

	case *parser.FnCallStatementContext:

		return parseFnCall(statementCtx.FunctionCall())

	case *parser.StatementContext:
		return nil

	default:
		panic("unhandled clause")
	}
}

func parseSentValue(statementCtx parser.ISentValueContext) SentValue {
	switch statementCtx := statementCtx.(type) {
	case *parser.SentLiteralContext:
		return &SentValueLiteral{
			Monetary: parseLiteral(statementCtx.Literal()),
		}
	case *parser.SentAllContext:
		return &SentValueAll{
			Range: ctxToRange(statementCtx),
			Asset: parseVariableAsset(statementCtx.SentAllLit().VariableAsset()),
		}

	case *parser.SentValueContext:
		return nil

	default:
		panic("Unhandled clause")
	}

}

func parseSendStatement(statementCtx *parser.SendStatementContext) *SendStatement {
	return &SendStatement{
		Source:      parseSource(statementCtx.Source()),
		Destination: parseDestination(statementCtx.Destination()),
		Range:       ctxToRange(statementCtx),
		SentValue:   parseSentValue(statementCtx.SentValue()),
	}
}

func parseVariableMonetary(sendExpr parser.IVariableMonetaryContext) Literal {
	switch sendExpr := sendExpr.(type) {
	case *parser.MonetaryVariableContext:
		return parseVarLiteral(sendExpr.VARIABLE_NAME().GetSymbol())

	case *parser.MonetaryContext:
		return parseMonetaryLit(sendExpr.MonetaryLit())

	case *parser.VariableMonetaryContext:
		return nil

	default:
		panic("Unhandled clause")
	}
}

func parseNumberLiteral(numNode antlr.TerminalNode) *NumberLiteral {
	amtStr := numNode.GetText()

	amt, err := strconv.Atoi(amtStr)
	if err != nil {
		panic("Invalid amt: " + amtStr)
	}

	return &NumberLiteral{
		Range:  tokenToRange(numNode.GetSymbol()),
		Number: amt,
	}
}

func parseVariableNumber(numCtx parser.IVariableNumberContext) Literal {
	switch numCtx := numCtx.(type) {
	case *parser.NumberContext:
		return parseNumberLiteral(numCtx.NUMBER())

	case *parser.NumberVariableContext:
		return parseVarLiteral(numCtx.VARIABLE_NAME().GetSymbol())

	case *parser.VariableNumberContext:
		return nil

	default:
		panic("Unhandled clause")
	}
}

func parseVariableAsset(assetCtx parser.IVariableAssetContext) Literal {
	switch assetCtx := assetCtx.(type) {
	case *parser.AssetContext:
		return &AssetLiteral{
			Range: ctxToRange(assetCtx),
			Asset: assetCtx.GetText(),
		}

	case *parser.AssetVariableContext:
		return parseVarLiteral(assetCtx.VARIABLE_NAME().GetSymbol())

	case *parser.VariableAssetContext:
		return nil

	default:
		panic("Unhandled clause")
	}
}

func parseMonetaryLit(monetaryLitCtx parser.IMonetaryLitContext) *MonetaryLiteral {
	if monetaryLitCtx.GetAmt() == nil {
		return nil
	}

	return &MonetaryLiteral{
		Range:  ctxToRange(monetaryLitCtx),
		Asset:  parseVariableAsset(monetaryLitCtx.GetAsset()),
		Amount: parseVariableNumber(monetaryLitCtx.GetAmt()),
	}
}

func ctxToRange(ctx antlr.ParserRuleContext) Range {
	startTk := ctx.GetStart()
	endTk := ctx.GetStop()

	return Range{
		Start: Position{
			Line:      startTk.GetLine() - 1,
			Character: startTk.GetColumn(),
		},
		End: Position{
			Line: endTk.GetLine() - 1,

			// this is based on the assumption that a token cannot span multiple lines
			Character: endTk.GetColumn() + len(endTk.GetText()),
		},
	}
}

func tokenToRange(tk antlr.Token) Range {
	return Range{
		Start: Position{
			Line:      tk.GetLine() - 1,
			Character: tk.GetColumn(),
		},
		End: Position{
			Line:      tk.GetLine() - 1,
			Character: tk.GetColumn() + len(tk.GetText()),
		},
	}
}
