package main

import "github.com/bzick/tokenizer"

const (
	PlusToken = iota + 1
	MinusToken
	ChevronRightToken
	ChevronLeftToken
	OpenBracketsToken
	CloseBracketsToken
	DotToken
	CommaToken
	StarToken
)

var program_tokens = []tokenizer.TokenKey{PlusToken, MinusToken, ChevronRightToken, ChevronLeftToken, OpenBracketsToken, CloseBracketsToken, DotToken, CommaToken, StarToken}

func new_lexer() *tokenizer.Tokenizer {
	lexer := tokenizer.New()
	lexer.DefineTokens(PlusToken, []string{"+"})
	lexer.DefineTokens(MinusToken, []string{"-"})
	lexer.DefineTokens(ChevronRightToken, []string{">"})
	lexer.DefineTokens(ChevronLeftToken, []string{"<"})
	lexer.DefineTokens(OpenBracketsToken, []string{"["})
	lexer.DefineTokens(CloseBracketsToken, []string{"]"})
	lexer.DefineTokens(DotToken, []string{"."})
	lexer.DefineTokens(CommaToken, []string{","})
	lexer.DefineTokens(StarToken, []string{"*"})

	return lexer
}
