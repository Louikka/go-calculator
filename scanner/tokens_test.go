package scanner

import "testing"

func testInterfaceMethods(t *testing.T, token Token, typeExpected, stringExpected string) {
	tokenType := token.Type()
	if tokenType != typeExpected {
		t.Errorf("%T -> Type() -> got %q, expected %q", token, tokenType, typeExpected)
	}

	tokenString := token.String()
	if tokenString != stringExpected {
		t.Errorf("%T -> String() -> got %q, expected %q", token, tokenString, stringExpected)
	}
}

func TestTokenInvalid(t *testing.T) {
	token := NewTokenInvalid()
	testInterfaceMethods(t, token, "INVALID", "")
}

func TestTokenNumber(t *testing.T) {
	token := NewTokenNumber(12.3)
	testInterfaceMethods(t, token, "NUMBER", "12.3")
}

func TestTokenWord(t *testing.T) {
	tests := []struct {
		name           string
		token          TokenWord
		kind           string
		expectedType   string
		expectedString string
	}{
		{
			name:           "default",
			token:          NewTokenWord("A"),
			expectedType:   "WORD",
			expectedString: "A",
		},
		{
			name:           "kind function",
			token:          NewTokenWord("AB"),
			kind:           WORD_KIND_FUNCTION,
			expectedType:   WORD_KIND_FUNCTION,
			expectedString: "AB",
		},
		{
			name:           "kind unspecified",
			token:          NewTokenWord("ABC"),
			kind:           WORD_KIND_UNSPECIFIED,
			expectedType:   "WORD",
			expectedString: "ABC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.token.Kind = tt.kind
			testInterfaceMethods(t, tt.token, tt.expectedType, tt.expectedString)
		})
	}
}

func TestTokenOperatro(t *testing.T) {
	tests := []struct {
		name                  string
		token                 TokenOperator
		expectedPrecedence    int
		expectedAssociativity string
		expectedType          string
		expectedString        string
		expectedCanBeUnary    bool
	}{
		{
			name:                  "plus",
			token:                 NewTokenOperator("+"),
			expectedPrecedence:    1,
			expectedAssociativity: ASSOC_LEFT,
			expectedType:          "OPERATOR",
			expectedString:        "+",
			expectedCanBeUnary:    true,
		},
		{
			name:                  "slash",
			token:                 NewTokenOperator("/"),
			expectedPrecedence:    2,
			expectedAssociativity: ASSOC_LEFT,
			expectedType:          "OPERATOR",
			expectedString:        "/",
			expectedCanBeUnary:    false,
		},
		{
			name:                  "caret",
			token:                 NewTokenOperator("^"),
			expectedPrecedence:    3,
			expectedAssociativity: ASSOC_RIGHT,
			expectedType:          "OPERATOR",
			expectedString:        "^",
			expectedCanBeUnary:    false,
		},
		{
			name:                  "not an operator",
			token:                 NewTokenOperator("NOT_OPERATOR"),
			expectedPrecedence:    0,
			expectedAssociativity: ASSOC_UNDEFINED,
			expectedType:          "OPERATOR",
			expectedString:        "NOT_OPERATOR",
			expectedCanBeUnary:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testInterfaceMethods(t, tt.token, tt.expectedType, tt.expectedString)

			prec := tt.token.Precedence
			if prec != tt.expectedPrecedence {
				t.Errorf("%T -> Precedence -> got %d, expected %d", tt.token, prec, tt.expectedPrecedence)
			}

			assoc := tt.token.Associativity
			if assoc != tt.expectedAssociativity {
				t.Errorf("%T -> Associativity -> got %q, expected %q", tt.token, assoc, tt.expectedAssociativity)
			}

			canBeUnary := tt.token.CanBeUnary()
			if canBeUnary != tt.expectedCanBeUnary {
				t.Errorf("%T -> CanBeUnary() -> got %t, expected %t", tt.token, canBeUnary, tt.expectedCanBeUnary)
			}
		})
	}
}

func TestTokenPunctuation(t *testing.T) {
	tests := []struct {
		name             string
		token            TokenPunctuation
		expectedType     string
		expectedString   string
		expectedIsLParen bool
	}{
		{
			name:             "lparen",
			token:            NewTokenPunctuation("("),
			expectedType:     "PUNCTUATION",
			expectedString:   "(",
			expectedIsLParen: true,
		},
		{
			name:             "rparen",
			token:            NewTokenPunctuation(")"),
			expectedType:     "PUNCTUATION",
			expectedString:   ")",
			expectedIsLParen: false,
		},
		{
			name:             "comma",
			token:            NewTokenPunctuation(","),
			expectedType:     "PUNCTUATION",
			expectedString:   ",",
			expectedIsLParen: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testInterfaceMethods(t, tt.token, tt.expectedType, tt.expectedString)

			isLParen := tt.token.IsLeftParenthesis()
			if isLParen != tt.expectedIsLParen {
				t.Errorf("%T -> IsLeftParenthesis() -> got %t, expected %t", tt.token, isLParen, tt.expectedIsLParen)
			}
		})
	}
}
