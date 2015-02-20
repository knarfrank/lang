package main


const (
  IGNR = iota
  /*
        Identifier
  */
  IDFR

  /*
        Constant Types
  */
  INT     // 343
  BOOL    // true/false
  STRG    // "sdf"

  /*
        Arthimetic Operators
  */
  ADD     // +
  SUB     // -
  MUL     // *
  DIV     // /
  MOD     // %

  /*
        Bitwise Operators
  */
  AND     // &
  OR      // |
  XOR     // ^
  BNOT    // ~
  RSHFT   // >>
  LSHFT   // <<


  /*
      Boolean Logical Operators
  */
  EQU    // ==
  LST    // <
  GRT    // >
  LNOT   // !

  /*
      Keywords
  */
  IF     // "if"
  ELSE   // "else"
  VAR    // "var"


)

type token struct {
  token int
  value string
  regex string
}


func getTokens() []token {
  var tokens []token
  tokens = make([]token, 0, 0)
  tokens = append(tokens, token{IDFR, "", "^[a-z][a-z]*$"},
                          token{INT, "", "^[0-9]+$"},
                          token{BOOL, "", "^(true|false)$"},
                          token{BOOL, "", "^(true|false)$"},
                          token{ADD, "", "^\\+$"},
                          token{SUB, "", "^-$"},
                          token{MUL, "", "^\\*$"},
                          token{DIV, "", "^/$"},
                          token{MOD, "", "^%$"},
                          token{AND, "", "^&$"},
                          token{OR, "", "^\\|$"},
                          token{XOR, "", "^\\^$"},
                          token{BNOT, "", "^~$"})

  return tokens;
}
