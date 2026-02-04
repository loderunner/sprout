// Code generated from Sprout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Sprout
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SproutParser struct {
	*antlr.BaseParser
}

var SproutParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sproutParserInit() {
	staticData := &SproutParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "':'", "')'", "'let'", "'in'", "'if'", "'then'", "'else'",
		"'true'", "'false'", "'fun'", "'->'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "LET", "IN", "IF", "THEN", "ELSE", "TRUE", "FALSE",
		"FUN", "ARROW", "EQUALS", "INT", "IDENT", "WS",
	}
	staticData.RuleNames = []string{
		"expr", "letExpr", "ifExpr", "funExpr", "appExpr", "primaryExpr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 65, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 17, 8, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 50, 8, 4, 10, 4, 12, 4, 53, 9, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 63, 8, 5, 1, 5, 0, 1, 8, 6,
		0, 2, 4, 6, 8, 10, 0, 0, 66, 0, 16, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4, 25,
		1, 0, 0, 0, 6, 32, 1, 0, 0, 0, 8, 41, 1, 0, 0, 0, 10, 62, 1, 0, 0, 0, 12,
		17, 3, 2, 1, 0, 13, 17, 3, 4, 2, 0, 14, 17, 3, 6, 3, 0, 15, 17, 3, 8, 4,
		0, 16, 12, 1, 0, 0, 0, 16, 13, 1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 16, 15,
		1, 0, 0, 0, 17, 1, 1, 0, 0, 0, 18, 19, 5, 4, 0, 0, 19, 20, 5, 15, 0, 0,
		20, 21, 5, 13, 0, 0, 21, 22, 3, 0, 0, 0, 22, 23, 5, 5, 0, 0, 23, 24, 3,
		0, 0, 0, 24, 3, 1, 0, 0, 0, 25, 26, 5, 6, 0, 0, 26, 27, 3, 0, 0, 0, 27,
		28, 5, 7, 0, 0, 28, 29, 3, 0, 0, 0, 29, 30, 5, 8, 0, 0, 30, 31, 3, 0, 0,
		0, 31, 5, 1, 0, 0, 0, 32, 33, 5, 11, 0, 0, 33, 34, 5, 1, 0, 0, 34, 35,
		5, 15, 0, 0, 35, 36, 5, 2, 0, 0, 36, 37, 5, 15, 0, 0, 37, 38, 5, 3, 0,
		0, 38, 39, 5, 12, 0, 0, 39, 40, 3, 0, 0, 0, 40, 7, 1, 0, 0, 0, 41, 42,
		6, 4, -1, 0, 42, 43, 3, 10, 5, 0, 43, 51, 1, 0, 0, 0, 44, 45, 10, 1, 0,
		0, 45, 46, 5, 1, 0, 0, 46, 47, 3, 0, 0, 0, 47, 48, 5, 3, 0, 0, 48, 50,
		1, 0, 0, 0, 49, 44, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		51, 52, 1, 0, 0, 0, 52, 9, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 63, 5, 14,
		0, 0, 55, 63, 5, 9, 0, 0, 56, 63, 5, 10, 0, 0, 57, 63, 5, 15, 0, 0, 58,
		59, 5, 1, 0, 0, 59, 60, 3, 0, 0, 0, 60, 61, 5, 3, 0, 0, 61, 63, 1, 0, 0,
		0, 62, 54, 1, 0, 0, 0, 62, 55, 1, 0, 0, 0, 62, 56, 1, 0, 0, 0, 62, 57,
		1, 0, 0, 0, 62, 58, 1, 0, 0, 0, 63, 11, 1, 0, 0, 0, 3, 16, 51, 62,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SproutParserInit initializes any static state used to implement SproutParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSproutParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SproutParserInit() {
	staticData := &SproutParserStaticData
	staticData.once.Do(sproutParserInit)
}

// NewSproutParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSproutParser(input antlr.TokenStream) *SproutParser {
	SproutParserInit()
	this := new(SproutParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SproutParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Sprout.g4"

	return this
}

// SproutParser tokens.
const (
	SproutParserEOF    = antlr.TokenEOF
	SproutParserT__0   = 1
	SproutParserT__1   = 2
	SproutParserT__2   = 3
	SproutParserLET    = 4
	SproutParserIN     = 5
	SproutParserIF     = 6
	SproutParserTHEN   = 7
	SproutParserELSE   = 8
	SproutParserTRUE   = 9
	SproutParserFALSE  = 10
	SproutParserFUN    = 11
	SproutParserARROW  = 12
	SproutParserEQUALS = 13
	SproutParserINT    = 14
	SproutParserIDENT  = 15
	SproutParserWS     = 16
)

// SproutParser rules.
const (
	SproutParserRULE_expr        = 0
	SproutParserRULE_letExpr     = 1
	SproutParserRULE_ifExpr      = 2
	SproutParserRULE_funExpr     = 3
	SproutParserRULE_appExpr     = 4
	SproutParserRULE_primaryExpr = 5
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LetExpr() ILetExprContext
	IfExpr() IIfExprContext
	FunExpr() IFunExprContext
	AppExpr() IAppExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) LetExpr() ILetExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetExprContext)
}

func (s *ExprContext) IfExpr() IIfExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExprContext)
}

func (s *ExprContext) FunExpr() IFunExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunExprContext)
}

func (s *ExprContext) AppExpr() IAppExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SproutParserRULE_expr)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SproutParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.LetExpr()
		}

	case SproutParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.IfExpr()
		}

	case SproutParserFUN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(14)
			p.FunExpr()
		}

	case SproutParserT__0, SproutParserTRUE, SproutParserFALSE, SproutParserINT, SproutParserIDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(15)
			p.appExpr(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILetExprContext is an interface to support dynamic dispatch.
type ILetExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	IN() antlr.TerminalNode

	// IsLetExprContext differentiates from other interfaces.
	IsLetExprContext()
}

type LetExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetExprContext() *LetExprContext {
	var p = new(LetExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_letExpr
	return p
}

func InitEmptyLetExprContext(p *LetExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_letExpr
}

func (*LetExprContext) IsLetExprContext() {}

func NewLetExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetExprContext {
	var p = new(LetExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_letExpr

	return p
}

func (s *LetExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LetExprContext) LET() antlr.TerminalNode {
	return s.GetToken(SproutParserLET, 0)
}

func (s *LetExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SproutParserIDENT, 0)
}

func (s *LetExprContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SproutParserEQUALS, 0)
}

func (s *LetExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LetExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LetExprContext) IN() antlr.TerminalNode {
	return s.GetToken(SproutParserIN, 0)
}

func (s *LetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitLetExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) LetExpr() (localctx ILetExprContext) {
	localctx = NewLetExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SproutParserRULE_letExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(SproutParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Match(SproutParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Match(SproutParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)
		p.Expr()
	}
	{
		p.SetState(22)
		p.Match(SproutParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(23)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfExprContext is an interface to support dynamic dispatch.
type IIfExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	THEN() antlr.TerminalNode
	ELSE() antlr.TerminalNode

	// IsIfExprContext differentiates from other interfaces.
	IsIfExprContext()
}

type IfExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExprContext() *IfExprContext {
	var p = new(IfExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_ifExpr
	return p
}

func InitEmptyIfExprContext(p *IfExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_ifExpr
}

func (*IfExprContext) IsIfExprContext() {}

func NewIfExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExprContext {
	var p = new(IfExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_ifExpr

	return p
}

func (s *IfExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExprContext) IF() antlr.TerminalNode {
	return s.GetToken(SproutParserIF, 0)
}

func (s *IfExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IfExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfExprContext) THEN() antlr.TerminalNode {
	return s.GetToken(SproutParserTHEN, 0)
}

func (s *IfExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SproutParserELSE, 0)
}

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) IfExpr() (localctx IIfExprContext) {
	localctx = NewIfExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SproutParserRULE_ifExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(SproutParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Expr()
	}
	{
		p.SetState(27)
		p.Match(SproutParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Expr()
	}
	{
		p.SetState(29)
		p.Match(SproutParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunExprContext is an interface to support dynamic dispatch.
type IFunExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUN() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	ARROW() antlr.TerminalNode
	Expr() IExprContext

	// IsFunExprContext differentiates from other interfaces.
	IsFunExprContext()
}

type FunExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunExprContext() *FunExprContext {
	var p = new(FunExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_funExpr
	return p
}

func InitEmptyFunExprContext(p *FunExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_funExpr
}

func (*FunExprContext) IsFunExprContext() {}

func NewFunExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunExprContext {
	var p = new(FunExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_funExpr

	return p
}

func (s *FunExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FunExprContext) FUN() antlr.TerminalNode {
	return s.GetToken(SproutParserFUN, 0)
}

func (s *FunExprContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SproutParserIDENT)
}

func (s *FunExprContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SproutParserIDENT, i)
}

func (s *FunExprContext) ARROW() antlr.TerminalNode {
	return s.GetToken(SproutParserARROW, 0)
}

func (s *FunExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitFunExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) FunExpr() (localctx IFunExprContext) {
	localctx = NewFunExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SproutParserRULE_funExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(SproutParserFUN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(SproutParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(SproutParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Match(SproutParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Match(SproutParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(SproutParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(SproutParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAppExprContext is an interface to support dynamic dispatch.
type IAppExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext
	AppExpr() IAppExprContext
	Expr() IExprContext

	// IsAppExprContext differentiates from other interfaces.
	IsAppExprContext()
}

type AppExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppExprContext() *AppExprContext {
	var p = new(AppExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_appExpr
	return p
}

func InitEmptyAppExprContext(p *AppExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_appExpr
}

func (*AppExprContext) IsAppExprContext() {}

func NewAppExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppExprContext {
	var p = new(AppExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_appExpr

	return p
}

func (s *AppExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AppExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *AppExprContext) AppExpr() IAppExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppExprContext)
}

func (s *AppExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AppExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitAppExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) AppExpr() (localctx IAppExprContext) {
	return p.appExpr(0)
}

func (p *SproutParser) appExpr(_p int) (localctx IAppExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAppExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAppExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, SproutParserRULE_appExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.PrimaryExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAppExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SproutParserRULE_appExpr)
			p.SetState(44)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(45)
				p.Match(SproutParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(46)
				p.Expr()
			}
			{
				p.SetState(47)
				p.Match(SproutParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	Expr() IExprContext

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_primaryExpr
	return p
}

func InitEmptyPrimaryExprContext(p *PrimaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_primaryExpr
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) INT() antlr.TerminalNode {
	return s.GetToken(SproutParserINT, 0)
}

func (s *PrimaryExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SproutParserTRUE, 0)
}

func (s *PrimaryExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SproutParserFALSE, 0)
}

func (s *PrimaryExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SproutParserIDENT, 0)
}

func (s *PrimaryExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SproutParserRULE_primaryExpr)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SproutParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(SproutParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserTRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Match(SproutParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Match(SproutParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserIDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.Match(SproutParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(58)
			p.Match(SproutParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Expr()
		}
		{
			p.SetState(60)
			p.Match(SproutParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SproutParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *AppExprContext = nil
		if localctx != nil {
			t = localctx.(*AppExprContext)
		}
		return p.AppExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SproutParser) AppExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
