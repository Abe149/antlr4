package antlr4

import (
	"strings"
)

type TokenSourceInputStreamPair struct {
	tokenSource *TokenSource
	inputStream *InputStream
}

// A token has properties: text, type, line, character position in the line
// (so we can ignore tabs), token channel, index, and source from which
// we obtained this token.

type Token struct {
	source *TokenSourceInputStreamPair
	tokenType int // token type of the token
	channel int // The parser ignores everything not on DEFAULT_CHANNEL
	start int // optional return -1 if not implemented.
	stop int // optional return -1 if not implemented.
	tokenIndex int // from 0..n-1 of the token object in the input stream
	line int // line=1..n of the 1st character
	column int // beginning of the line at which it occurs, 0..n-1
	_text string // text of the token.

//	String getText();
//	int getType();
//	int getLine();
//	int getCharPositionInLine();
//	int getChannel();
//	int getTokenIndex();
//	int getStartIndex();
//	int getStopIndex();
//	TokenSource getTokenSource();
//	CharStream getInputStream();
}

const (
	TokenInvalidType = 0

	// During lookahead operations, this "token" signifies we hit rule end ATN state
	// and did not follow it despite needing to.
	TokenEpsilon = -2

	TokenMinUserTokenType = 1

	TokenEOF = -1

	// All tokens go to the parser (unless skip() is called in that rule)
	// on a particular "channel". The parser tunes to a particular channel
	// so that whitespace etc... can go to the parser on a "hidden" channel.

	TokenDefaultChannel = 0

	// Anything on different channel than DEFAULT_CHANNEL is not parsed
	// by parser.

	TokenHiddenChannel = 1
)

// Explicitly set the text for this token. If {code text} is not
// {@code nil}, then {@link //getText} will return this value rather than
// extracting the text from the input.
//
// @param text The explicit text of the token, or {@code nil} if the text
// should be obtained from the input along with the start and stop indexes
// of the token.

func (this *Token) text() string{
	return this._text
}

func (this *Token) setText(s string) {
	this._text = s
}

func (this *Token) getTokenSource() *TokenSource {
	return this.source.tokenSource
}

func (this *Token) getInputStream() *InputStream {
	return this.source.inputStream
}

type CommonToken struct {
	Token
}

func NewCommonToken(source *TokenSourceInputStreamPair, tokenType, channel, start, stop int) *CommonToken {

	t := CommonToken{Token{}}

	t.source = source
	t.tokenType = -1
	t.channel = channel
	t.start = start
	t.stop = stop
	t.tokenIndex = -1
	if (t.source[0] != nil) {
		t.line = source.tokenSource.line()
		t.column = source.tokenSource.column()
	} else {
		t.column = -1
	}
	return t
}

// An empty {@link Pair} which is used as the default value of
// {@link //source} for tokens that do not have a source.

//CommonToken.EMPTY_SOURCE = [ nil, nil ]

// Constructs a New{@link CommonToken} as a copy of another {@link Token}.
//
// <p>
// If {@code oldToken} is also a {@link CommonToken} instance, the newly
// constructed token will share a reference to the {@link //text} field and
// the {@link Pair} stored in {@link //source}. Otherwise, {@link //text} will
// be assigned the result of calling {@link //getText}, and {@link //source}
// will be constructed from the result of {@link Token//getTokenSource} and
// {@link Token//getInputStream}.</p>
//
// @param oldToken The token to copy.
//
func (ct *CommonToken) clone() {
	var t = NewCommonToken(ct.source, ct.tokenType, ct.channel, ct.start,
			ct.stop)
	t.tokenIndex = ct.tokenIndex
	t.line = ct.line
	t.column = ct.column
	t.text = ct.text
	return t
}

func (this *CommonToken) text() string {
	if (this._text != nil) {
		return this._text
	}
	var input = this.getInputStream()
	if (input == nil) {
		return nil
	}
	var n = input.size
	if (this.start < n && this.stop < n) {
		return input.getText(this.start, this.stop)
	} else {
		return "<EOF>"
	}
}

func (this *CommonToken) setText(text string) {
	this._text = text
}

func (this *CommonToken) toString() string {
	var txt = this.text
	if (txt != nil) {
		txt = strings.Replace(txt, "\n", "", -1)
		txt = strings.Replace(txt, "\r", "", -1)
		txt = strings.Replace(txt, "\t", "", -1)
	} else {
		txt = "<no text>"
	}

	var ch string;
	if (this.channel > 0){
		ch = ",channel=" + this.channel
	} else {
		ch = ""
	}

	return "[@" + this.tokenIndex + "," + this.start + ":" + this.stop + "='" +
			txt + "',<" + this.tokenType + ">" +
			ch + "," + this.line + ":" + this.column + "]"
}



