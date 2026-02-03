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
		"", "'let'", "'in'", "'if'", "'then'", "'else'", "'true'", "'false'",
		"'='",
	}
	staticData.SymbolicNames = []string{
		"", "LET", "IN", "IF", "THEN", "ELSE", "TRUE", "FALSE", "EQUALS", "INT",
		"IDENT", "WS",
	}
	staticData.RuleNames = []string{
		"expr", "letExpr", "ifExpr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 29, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 13, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 0,
		30, 0, 12, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 21, 1, 0, 0, 0, 6, 13, 3,
		2, 1, 0, 7, 13, 3, 4, 2, 0, 8, 13, 5, 9, 0, 0, 9, 13, 5, 6, 0, 0, 10, 13,
		5, 7, 0, 0, 11, 13, 5, 10, 0, 0, 12, 6, 1, 0, 0, 0, 12, 7, 1, 0, 0, 0,
		12, 8, 1, 0, 0, 0, 12, 9, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 12, 11, 1, 0,
		0, 0, 13, 1, 1, 0, 0, 0, 14, 15, 5, 1, 0, 0, 15, 16, 5, 10, 0, 0, 16, 17,
		5, 8, 0, 0, 17, 18, 3, 0, 0, 0, 18, 19, 5, 2, 0, 0, 19, 20, 3, 0, 0, 0,
		20, 3, 1, 0, 0, 0, 21, 22, 5, 3, 0, 0, 22, 23, 3, 0, 0, 0, 23, 24, 5, 4,
		0, 0, 24, 25, 3, 0, 0, 0, 25, 26, 5, 5, 0, 0, 26, 27, 3, 0, 0, 0, 27, 5,
		1, 0, 0, 0, 1, 12,
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
	SproutParserLET    = 1
	SproutParserIN     = 2
	SproutParserIF     = 3
	SproutParserTHEN   = 4
	SproutParserELSE   = 5
	SproutParserTRUE   = 6
	SproutParserFALSE  = 7
	SproutParserEQUALS = 8
	SproutParserINT    = 9
	SproutParserIDENT  = 10
	SproutParserWS     = 11
)

// SproutParser rules.
const (
	SproutParserRULE_expr    = 0
	SproutParserRULE_letExpr = 1
	SproutParserRULE_ifExpr  = 2
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LetExpr() ILetExprContext
	IfExpr() IIfExprContext
	INT() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	IDENT() antlr.TerminalNode

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

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(SproutParserINT, 0)
}

func (s *ExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SproutParserTRUE, 0)
}

func (s *ExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SproutParserFALSE, 0)
}

func (s *ExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SproutParserIDENT, 0)
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
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SproutParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)
			p.LetExpr()
		}

	case SproutParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(7)
			p.IfExpr()
		}

	case SproutParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(8)
			p.Match(SproutParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(9)
			p.Match(SproutParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserFALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(10)
			p.Match(SproutParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SproutParserIDENT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(11)
			p.Match(SproutParserIDENT)
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
		p.SetState(14)
		p.Match(SproutParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(15)
		p.Match(SproutParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(16)
		p.Match(SproutParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(17)
		p.Expr()
	}
	{
		p.SetState(18)
		p.Match(SproutParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(19)
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
		p.SetState(21)
		p.Match(SproutParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.Expr()
	}
	{
		p.SetState(23)
		p.Match(SproutParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Expr()
	}
	{
		p.SetState(25)
		p.Match(SproutParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
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
