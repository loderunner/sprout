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
		"", "'('", "':'", "')'", "'||'", "'&&'", "'+'", "'-'", "'*'", "'/'",
		"'!'", "'let'", "'in'", "'if'", "'then'", "'else'", "'true'", "'false'",
		"'fun'", "'->'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "LET", "IN", "IF", "THEN",
		"ELSE", "TRUE", "FALSE", "FUN", "ARROW", "EQUALS", "INT", "IDENT", "COMPOP",
		"WS",
	}
	staticData.RuleNames = []string{
		"expr", "letExpr", "ifExpr", "funExpr", "compExpr", "orExpr", "andExpr",
		"termExpr", "factorExpr", "unaryExpr", "appExpr", "primaryExpr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 133, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 29, 8, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 59, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 67,
		8, 5, 10, 5, 12, 5, 70, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6,
		78, 8, 6, 10, 6, 12, 6, 81, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 89, 8, 7, 10, 7, 12, 7, 92, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		5, 8, 100, 8, 8, 10, 8, 12, 8, 103, 9, 8, 1, 9, 1, 9, 1, 9, 3, 9, 108,
		8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 118,
		8, 10, 10, 10, 12, 10, 121, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 131, 8, 11, 1, 11, 0, 5, 10, 12, 14, 16, 20, 12,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 3, 1, 0, 6, 7, 1, 0, 8, 9,
		2, 0, 7, 7, 10, 10, 134, 0, 28, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4, 37, 1,
		0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 60, 1, 0, 0, 0, 12,
		71, 1, 0, 0, 0, 14, 82, 1, 0, 0, 0, 16, 93, 1, 0, 0, 0, 18, 107, 1, 0,
		0, 0, 20, 109, 1, 0, 0, 0, 22, 130, 1, 0, 0, 0, 24, 29, 3, 2, 1, 0, 25,
		29, 3, 4, 2, 0, 26, 29, 3, 6, 3, 0, 27, 29, 3, 8, 4, 0, 28, 24, 1, 0, 0,
		0, 28, 25, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 27, 1, 0, 0, 0, 29, 1, 1,
		0, 0, 0, 30, 31, 5, 11, 0, 0, 31, 32, 5, 22, 0, 0, 32, 33, 5, 20, 0, 0,
		33, 34, 3, 0, 0, 0, 34, 35, 5, 12, 0, 0, 35, 36, 3, 0, 0, 0, 36, 3, 1,
		0, 0, 0, 37, 38, 5, 13, 0, 0, 38, 39, 3, 0, 0, 0, 39, 40, 5, 14, 0, 0,
		40, 41, 3, 0, 0, 0, 41, 42, 5, 15, 0, 0, 42, 43, 3, 0, 0, 0, 43, 5, 1,
		0, 0, 0, 44, 45, 5, 18, 0, 0, 45, 46, 5, 1, 0, 0, 46, 47, 5, 22, 0, 0,
		47, 48, 5, 2, 0, 0, 48, 49, 5, 22, 0, 0, 49, 50, 5, 3, 0, 0, 50, 51, 5,
		19, 0, 0, 51, 52, 3, 0, 0, 0, 52, 7, 1, 0, 0, 0, 53, 59, 3, 10, 5, 0, 54,
		55, 3, 10, 5, 0, 55, 56, 5, 23, 0, 0, 56, 57, 3, 10, 5, 0, 57, 59, 1, 0,
		0, 0, 58, 53, 1, 0, 0, 0, 58, 54, 1, 0, 0, 0, 59, 9, 1, 0, 0, 0, 60, 61,
		6, 5, -1, 0, 61, 62, 3, 12, 6, 0, 62, 68, 1, 0, 0, 0, 63, 64, 10, 1, 0,
		0, 64, 65, 5, 4, 0, 0, 65, 67, 3, 12, 6, 0, 66, 63, 1, 0, 0, 0, 67, 70,
		1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 11, 1, 0, 0, 0,
		70, 68, 1, 0, 0, 0, 71, 72, 6, 6, -1, 0, 72, 73, 3, 14, 7, 0, 73, 79, 1,
		0, 0, 0, 74, 75, 10, 1, 0, 0, 75, 76, 5, 5, 0, 0, 76, 78, 3, 14, 7, 0,
		77, 74, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1,
		0, 0, 0, 80, 13, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 83, 6, 7, -1, 0, 83,
		84, 3, 16, 8, 0, 84, 90, 1, 0, 0, 0, 85, 86, 10, 1, 0, 0, 86, 87, 7, 0,
		0, 0, 87, 89, 3, 16, 8, 0, 88, 85, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90,
		88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 15, 1, 0, 0, 0, 92, 90, 1, 0, 0,
		0, 93, 94, 6, 8, -1, 0, 94, 95, 3, 18, 9, 0, 95, 101, 1, 0, 0, 0, 96, 97,
		10, 1, 0, 0, 97, 98, 7, 1, 0, 0, 98, 100, 3, 18, 9, 0, 99, 96, 1, 0, 0,
		0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102,
		17, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 108, 3, 20, 10, 0, 105, 106,
		7, 2, 0, 0, 106, 108, 3, 18, 9, 0, 107, 104, 1, 0, 0, 0, 107, 105, 1, 0,
		0, 0, 108, 19, 1, 0, 0, 0, 109, 110, 6, 10, -1, 0, 110, 111, 3, 22, 11,
		0, 111, 119, 1, 0, 0, 0, 112, 113, 10, 1, 0, 0, 113, 114, 5, 1, 0, 0, 114,
		115, 3, 0, 0, 0, 115, 116, 5, 3, 0, 0, 116, 118, 1, 0, 0, 0, 117, 112,
		1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0,
		0, 0, 120, 21, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 131, 5, 21, 0, 0,
		123, 131, 5, 16, 0, 0, 124, 131, 5, 17, 0, 0, 125, 131, 5, 22, 0, 0, 126,
		127, 5, 1, 0, 0, 127, 128, 3, 0, 0, 0, 128, 129, 5, 3, 0, 0, 129, 131,
		1, 0, 0, 0, 130, 122, 1, 0, 0, 0, 130, 123, 1, 0, 0, 0, 130, 124, 1, 0,
		0, 0, 130, 125, 1, 0, 0, 0, 130, 126, 1, 0, 0, 0, 131, 23, 1, 0, 0, 0,
		9, 28, 58, 68, 79, 90, 101, 107, 119, 130,
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
	SproutParserT__3   = 4
	SproutParserT__4   = 5
	SproutParserT__5   = 6
	SproutParserT__6   = 7
	SproutParserT__7   = 8
	SproutParserT__8   = 9
	SproutParserT__9   = 10
	SproutParserLET    = 11
	SproutParserIN     = 12
	SproutParserIF     = 13
	SproutParserTHEN   = 14
	SproutParserELSE   = 15
	SproutParserTRUE   = 16
	SproutParserFALSE  = 17
	SproutParserFUN    = 18
	SproutParserARROW  = 19
	SproutParserEQUALS = 20
	SproutParserINT    = 21
	SproutParserIDENT  = 22
	SproutParserCOMPOP = 23
	SproutParserWS     = 24
)

// SproutParser rules.
const (
	SproutParserRULE_expr        = 0
	SproutParserRULE_letExpr     = 1
	SproutParserRULE_ifExpr      = 2
	SproutParserRULE_funExpr     = 3
	SproutParserRULE_compExpr    = 4
	SproutParserRULE_orExpr      = 5
	SproutParserRULE_andExpr     = 6
	SproutParserRULE_termExpr    = 7
	SproutParserRULE_factorExpr  = 8
	SproutParserRULE_unaryExpr   = 9
	SproutParserRULE_appExpr     = 10
	SproutParserRULE_primaryExpr = 11
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
	CompExpr() ICompExprContext

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

func (s *ExprContext) CompExpr() ICompExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompExprContext)
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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SproutParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.LetExpr()
		}

	case SproutParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.IfExpr()
		}

	case SproutParserFUN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.FunExpr()
		}

	case SproutParserT__0, SproutParserT__6, SproutParserT__9, SproutParserTRUE, SproutParserFALSE, SproutParserINT, SproutParserIDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(27)
			p.CompExpr()
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
		p.SetState(30)
		p.Match(SproutParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(SproutParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(SproutParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Expr()
	}
	{
		p.SetState(34)
		p.Match(SproutParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
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
		p.SetState(37)
		p.Match(SproutParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Expr()
	}
	{
		p.SetState(39)
		p.Match(SproutParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Expr()
	}
	{
		p.SetState(41)
		p.Match(SproutParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
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
		p.SetState(44)
		p.Match(SproutParserFUN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.Match(SproutParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(SproutParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(SproutParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(SproutParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(SproutParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
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

// ICompExprContext is an interface to support dynamic dispatch.
type ICompExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOrExpr() []IOrExprContext
	OrExpr(i int) IOrExprContext
	COMPOP() antlr.TerminalNode

	// IsCompExprContext differentiates from other interfaces.
	IsCompExprContext()
}

type CompExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompExprContext() *CompExprContext {
	var p = new(CompExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_compExpr
	return p
}

func InitEmptyCompExprContext(p *CompExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_compExpr
}

func (*CompExprContext) IsCompExprContext() {}

func NewCompExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompExprContext {
	var p = new(CompExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_compExpr

	return p
}

func (s *CompExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CompExprContext) AllOrExpr() []IOrExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrExprContext); ok {
			len++
		}
	}

	tst := make([]IOrExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrExprContext); ok {
			tst[i] = t.(IOrExprContext)
			i++
		}
	}

	return tst
}

func (s *CompExprContext) OrExpr(i int) IOrExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExprContext); ok {
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

	return t.(IOrExprContext)
}

func (s *CompExprContext) COMPOP() antlr.TerminalNode {
	return s.GetToken(SproutParserCOMPOP, 0)
}

func (s *CompExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitCompExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) CompExpr() (localctx ICompExprContext) {
	localctx = NewCompExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SproutParserRULE_compExpr)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.orExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.orExpr(0)
		}
		{
			p.SetState(55)
			p.Match(SproutParserCOMPOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.orExpr(0)
		}

	case antlr.ATNInvalidAltNumber:
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

// IOrExprContext is an interface to support dynamic dispatch.
type IOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AndExpr() IAndExprContext
	OrExpr() IOrExprContext

	// IsOrExprContext differentiates from other interfaces.
	IsOrExprContext()
}

type OrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExprContext() *OrExprContext {
	var p = new(OrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_orExpr
	return p
}

func InitEmptyOrExprContext(p *OrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_orExpr
}

func (*OrExprContext) IsOrExprContext() {}

func NewOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExprContext {
	var p = new(OrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_orExpr

	return p
}

func (s *OrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExprContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *OrExprContext) OrExpr() IOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) OrExpr() (localctx IOrExprContext) {
	return p.orExpr(0)
}

func (p *SproutParser) orExpr(_p int) (localctx IOrExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewOrExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SproutParserRULE_orExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.andExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SproutParserRULE_orExpr)
			p.SetState(63)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(64)
				p.Match(SproutParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(65)
				p.andExpr(0)
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
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

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TermExpr() ITermExprContext
	AndExpr() IAndExprContext

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_andExpr
	return p
}

func InitEmptyAndExprContext(p *AndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_andExpr
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) TermExpr() ITermExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermExprContext)
}

func (s *AndExprContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) AndExpr() (localctx IAndExprContext) {
	return p.andExpr(0)
}

func (p *SproutParser) andExpr(_p int) (localctx IAndExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAndExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, SproutParserRULE_andExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.termExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SproutParserRULE_andExpr)
			p.SetState(74)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(75)
				p.Match(SproutParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(76)
				p.termExpr(0)
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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

// ITermExprContext is an interface to support dynamic dispatch.
type ITermExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	FactorExpr() IFactorExprContext
	TermExpr() ITermExprContext

	// IsTermExprContext differentiates from other interfaces.
	IsTermExprContext()
}

type TermExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyTermExprContext() *TermExprContext {
	var p = new(TermExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_termExpr
	return p
}

func InitEmptyTermExprContext(p *TermExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_termExpr
}

func (*TermExprContext) IsTermExprContext() {}

func NewTermExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermExprContext {
	var p = new(TermExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_termExpr

	return p
}

func (s *TermExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TermExprContext) GetOp() antlr.Token { return s.op }

func (s *TermExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *TermExprContext) FactorExpr() IFactorExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorExprContext)
}

func (s *TermExprContext) TermExpr() ITermExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermExprContext)
}

func (s *TermExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitTermExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) TermExpr() (localctx ITermExprContext) {
	return p.termExpr(0)
}

func (p *SproutParser) termExpr(_p int) (localctx ITermExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTermExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITermExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, SproutParserRULE_termExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.factorExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTermExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SproutParserRULE_termExpr)
			p.SetState(85)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(86)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TermExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SproutParserT__5 || _la == SproutParserT__6) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TermExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(87)
				p.factorExpr(0)
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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

// IFactorExprContext is an interface to support dynamic dispatch.
type IFactorExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	UnaryExpr() IUnaryExprContext
	FactorExpr() IFactorExprContext

	// IsFactorExprContext differentiates from other interfaces.
	IsFactorExprContext()
}

type FactorExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyFactorExprContext() *FactorExprContext {
	var p = new(FactorExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_factorExpr
	return p
}

func InitEmptyFactorExprContext(p *FactorExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_factorExpr
}

func (*FactorExprContext) IsFactorExprContext() {}

func NewFactorExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorExprContext {
	var p = new(FactorExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_factorExpr

	return p
}

func (s *FactorExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorExprContext) GetOp() antlr.Token { return s.op }

func (s *FactorExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *FactorExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *FactorExprContext) FactorExpr() IFactorExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorExprContext)
}

func (s *FactorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitFactorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) FactorExpr() (localctx IFactorExprContext) {
	return p.factorExpr(0)
}

func (p *SproutParser) factorExpr(_p int) (localctx IFactorExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewFactorExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFactorExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, SproutParserRULE_factorExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.UnaryExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFactorExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SproutParserRULE_factorExpr)
			p.SetState(96)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(97)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*FactorExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SproutParserT__7 || _la == SproutParserT__8) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*FactorExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(98)
				p.UnaryExpr()
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AppExpr() IAppExprContext
	UnaryExpr() IUnaryExprContext

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_unaryExpr
	return p
}

func InitEmptyUnaryExprContext(p *UnaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SproutParserRULE_unaryExpr
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SproutParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) AppExpr() IAppExprContext {
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

func (s *UnaryExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SproutVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SproutParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SproutParserRULE_unaryExpr)
	var _la int

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SproutParserT__0, SproutParserTRUE, SproutParserFALSE, SproutParserINT, SproutParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.appExpr(0)
		}

	case SproutParserT__6, SproutParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SproutParserT__6 || _la == SproutParserT__9) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(106)
			p.UnaryExpr()
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SproutParserRULE_appExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.PrimaryExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
			p.SetState(112)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(113)
				p.Match(SproutParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(114)
				p.Expr()
			}
			{
				p.SetState(115)
				p.Match(SproutParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, SproutParserRULE_primaryExpr)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SproutParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(SproutParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserTRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Match(SproutParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(SproutParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserIDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
			p.Match(SproutParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.Match(SproutParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Expr()
		}
		{
			p.SetState(128)
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
	case 5:
		var t *OrExprContext = nil
		if localctx != nil {
			t = localctx.(*OrExprContext)
		}
		return p.OrExpr_Sempred(t, predIndex)

	case 6:
		var t *AndExprContext = nil
		if localctx != nil {
			t = localctx.(*AndExprContext)
		}
		return p.AndExpr_Sempred(t, predIndex)

	case 7:
		var t *TermExprContext = nil
		if localctx != nil {
			t = localctx.(*TermExprContext)
		}
		return p.TermExpr_Sempred(t, predIndex)

	case 8:
		var t *FactorExprContext = nil
		if localctx != nil {
			t = localctx.(*FactorExprContext)
		}
		return p.FactorExpr_Sempred(t, predIndex)

	case 10:
		var t *AppExprContext = nil
		if localctx != nil {
			t = localctx.(*AppExprContext)
		}
		return p.AppExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SproutParser) OrExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SproutParser) AndExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SproutParser) TermExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SproutParser) FactorExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SproutParser) AppExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
