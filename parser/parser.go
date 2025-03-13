package parser

import (
	"fmt"
	"strconv"

	"meowlang/lexer"
)

// Parser holds the tokens, current position, and variable environment.
// The environment now maps variable names to interface{} so that values can be either int or string.
type Parser struct {
	tokens []lexer.Token
	pos    int
	vars   map[string]interface{}
}

// NewParser creates a new Parser instance.
func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
		vars:   make(map[string]interface{}),
	}
}

// Parse processes tokens until MEOWEND or EOF is encountered.
func (p *Parser) Parse() {
	 
	if p.tokens[p.pos].Type != lexer.MEOWSTART {
		fmt.Println("😾 Error: Program must start with MEOWSTART")
		return
	}
	p.pos++

 
	for p.pos < len(p.tokens) && p.tokens[p.pos].Type != lexer.MEOWEND && p.tokens[p.pos].Type != lexer.EOF {
		tok := p.tokens[p.pos]
		switch tok.Type {
		case lexer.MEOW:
			p.parseVariableDeclaration()
		case lexer.PURR:
			p.parsePrintStatement()
		default:
			p.pos++
		}
	}

	if p.pos < len(p.tokens) && p.tokens[p.pos].Type == lexer.MEOWEND {
		p.pos++
	}
}

// parseVariableDeclaration processes a variable declaration.
// Expected form: MEOW IDENT ASSIGN expression SEMICOLON
func (p *Parser) parseVariableDeclaration() {
	p.pos++ 

	if p.pos >= len(p.tokens) || p.tokens[p.pos].Type != lexer.IDENT {
		fmt.Println("😾 Error: Expected identifier after 'meow'")
		return
	}
	varName := p.tokens[p.pos].Literal
	p.pos++ 

	if p.pos >= len(p.tokens) || p.tokens[p.pos].Type != lexer.ASSIGN {
		fmt.Println("😾 Error: Expected '=' after variable name")
		return
	}
	p.pos++ 
 
	value := p.parseExpression()
	p.vars[varName] = value

	if p.pos < len(p.tokens) && p.tokens[p.pos].Type == lexer.SEMICOLON {
		p.pos++  
	}
}

// parsePrintStatement processes a print statement.
// Expected form: PURR LPAREN expression RPAREN SEMICOLON
func (p *Parser) parsePrintStatement() {
	p.pos++ // skip PURR token

	if p.pos >= len(p.tokens) || p.tokens[p.pos].Type != lexer.LPAREN {
		fmt.Println("😾 Error: Expected '(' after 'purr'")
		return
	}
	p.pos++ // skip LPAREN

	value := p.parseExpression()

	if p.pos >= len(p.tokens) || p.tokens[p.pos].Type != lexer.RPAREN {
		fmt.Println("😾 Error: Expected ')' after expression in 'purr'")
		return
	}
	p.pos++ // skip RPAREN

	if p.pos < len(p.tokens) && p.tokens[p.pos].Type == lexer.SEMICOLON {
		p.pos++ // skip SEMICOLON
	}
	fmt.Printf("😺: %v\n", value)
}

// parseExpression evaluates a sequence of terms connected by operators.
// It returns an interface{} that may be an int or string.
func (p *Parser) parseExpression() interface{} {
	left := p.parseTerm()

	// Process operators while the current token is an operator.
	for p.pos < len(p.tokens) && p.tokens[p.pos].Type == lexer.OPERATOR {
		opToken := p.tokens[p.pos]
		p.pos++ // skip operator

		right := p.parseTerm()
		switch opToken.Literal {
		case "+":
			// If both operands are integers, perform addition.
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li + ri
				} else {
					fmt.Println("😾 Type mismatch: cannot add int and non-int")
					return 0
				}
			} else if ls, ok := left.(string); ok {
				// If both operands are strings, perform concatenation.
				if rs, ok := right.(string); ok {
					left = ls + rs
				} else {
					fmt.Println("😾 Type mismatch: cannot concatenate string with non-string")
					return ""
				}
			} else {
				fmt.Println("😾 Unsupported types for '+' operator")
				return 0
			}
		case "-":
			// Subtraction only for integers.
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li - ri
				} else {
					fmt.Println("😾 Type mismatch: cannot subtract non-int from int")
					return 0
				}
			} else {
				fmt.Println("😾 Subtraction only supported on integers")
				return 0
			}
		case "*":
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li * ri
				} else {
					fmt.Println("😾 Type mismatch: cannot multiply int with non-int")
					return 0
				}
			} else {
				fmt.Println("😾 Multiplication only supported on integers")
				return 0
			}
		case "/":
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					if ri != 0 {
						left = li / ri
					} else {
						fmt.Println("😾 Division by zero")
						left = 0
					}
				} else {
					fmt.Println("😾 Type mismatch: cannot divide int by non-int")
					return 0
				}
			} else {
				fmt.Println("😾 Division only supported on integers")
				return 0
			}
		// Bitwise operations – these work only on integers.
		case "&":
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li & ri
				} else {
					fmt.Println("😾 Type mismatch: bitwise '&' requires ints")
					return 0
				}
			} else {
				fmt.Println("😾 Bitwise operations only supported on integers")
				return 0
			}
		case "|":
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li | ri
				} else {
					fmt.Println("😾 Type mismatch: bitwise '|' requires ints")
					return 0
				}
			} else {
				fmt.Println("😾 Bitwise operations only supported on integers")
				return 0
			}
		case "^":
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li ^ ri
				} else {
					fmt.Println("😾 Type mismatch: bitwise '^' requires ints")
					return 0
				}
			} else {
				fmt.Println("😾 Bitwise operations only supported on integers")
				return 0
			}
		case "<<":
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li << ri
				} else {
					fmt.Println("😾 Type mismatch: '<<' requires an integer shift amount")
					return 0
				}
			} else {
				fmt.Println("😾 Bitwise shift only supported on integers")
				return 0
			}
		case ">>":
			if li, ok := left.(int); ok {
				if ri, ok := right.(int); ok {
					left = li >> ri
				} else {
					fmt.Println("😾 Type mismatch: '>>' requires an integer shift amount")
					return 0
				}
			} else {
				fmt.Println("😾 Bitwise shift only supported on integers")
				return 0
			}
		default:
			fmt.Printf("😾 Unknown operator: %s\n", opToken.Literal)
		}
	}
	return left
}

// parseTerm processes a single term: number, string, identifier, or parenthesized expression.
func (p *Parser) parseTerm() interface{} {
	token := p.tokens[p.pos]
	var value interface{}

	switch token.Type {
	case lexer.NUMBER:
		v, err := strconv.Atoi(token.Literal)
		if err != nil {
			fmt.Printf("😾 Error converting '%s' to integer\n", token.Literal)
		}
		value = v
		p.pos++
	case lexer.STRING:
		value = token.Literal
		p.pos++
	case lexer.IDENT:
		// Look up variable value.
		if val, ok := p.vars[token.Literal]; ok {
			value = val
		} else {
			fmt.Printf("😾 Undefined variable '%s'\n", token.Literal)
			value = 0
		}
		p.pos++
	case lexer.LPAREN:
		p.pos++ // skip LPAREN
		value = p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos].Type != lexer.RPAREN {
			fmt.Println("😾 Error: Expected ')' after expression")
		} else {
			p.pos++ // skip RPAREN
		}
	default:
		fmt.Printf("😾 Unexpected token '%s' in expression\n", token.Literal)
		p.pos++
		value = 0
	}
	return value
}
