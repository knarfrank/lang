package tokens
import (  "strconv")

const (
  IGNR = iota

  WS
  NL
  SCMNT
  ECMNT
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
      Assignment
  */
  ASMT   // =


  /*
      Brackets, symbols and the other stuff
  */
  RPAREN  // (
  LPAREN  // )
  RCURL   // {
  LCURL   // }
  COLN    // :
  COMMA   // ,
  SCOLN   // ;

  /*
      Keywords
  */
  IF     // if
  ELSE   // else
  FUNC    // func
  LET    // let



)

type Token struct {
  Token int
  Value string
  Regex string
}


func GetTokens() []Token {
  var Tokens []Token
  Tokens = make([]Token, 0, 0)
  Tokens = append(Tokens, Token{NL, "", "^\\n$"},
                          Token{WS, "", "^(\\s)+$"},
                          Token{SCMNT, "", "^/\\*$"},
                          Token{ECMNT, "", "^\\*/$"},
                          Token{IF, "", "^(if)$"},
                          Token{ELSE, "", "^(else)$"},
                          Token{FUNC, "", "^(func)$"},
                          Token{LET, "", "^(let)$"},
                          Token{RPAREN, "", "^\\($"},
                          Token{LPAREN, "", "^\\)$"},
                          Token{RCURL, "", "^\\{$"},
                          Token{LCURL, "", "^\\}$"},
                          Token{COLN, "", "^:$"},
                          Token{SCOLN, "", "^;$"},
                          Token{COMMA, "", "^,$"},
                          Token{IDFR, "", "^[a-z|A-Z][a-z|A-Z]*$"},
                          Token{INT, "", "^[0-9]+$"},
                          Token{BOOL, "", "^(true|false)$"},
                          Token{ADD, "", "^\\+$"},
                          Token{SUB, "", "^-$"},
                          Token{MUL, "", "^\\*$"},
                          Token{DIV, "", "^/$"},
                          Token{MOD, "", "^%$"},
                          Token{AND, "", "^&$"},
                          Token{OR, "", "^\\|$"},
                          Token{XOR, "", "^\\^$"},
                          Token{BNOT, "", "^~$"},
                          Token{RSHFT, "", "^>>$"},
                          Token{LSHFT, "", "^<<$"},
                          Token{EQU, "", "^(==)$"},
                          Token{LST, "", "^<$"},
                          Token{GRT, "", "^>$"},
                          Token{LNOT, "", "^!$"},
                          Token{ASMT, "", "^=$"},
                          Token{IGNR, "", "^.$"})

  return Tokens
}


func getTokenName(id int) string {
  switch(id) {
    case IGNR: return "IGNR"
    case WS: return "WS"
    case NL: return "NL"
    case IDFR: return "IDFR"

    case INT: return "INT"
    case BOOL: return "BOOL"
    case STRG: return "STRG"

    case ADD: return "ADD"
    case SUB: return "SUB"
    case MUL: return "MUL"
    case DIV: return "DIV"
    case MOD: return "MOD"


    case EQU: return "EQU"
    case LST: return "LST"
    case GRT: return "GRT"
    case LNOT: return "LNOT"

    case ASMT: return "ASMT"
    case RPAREN: return "RPAREN"
    case LPAREN: return "LPAREN"
    case RCURL: return "RCURL"
    case LCURL: return "LCURL"
    case COLN: return "COLN"
    case COMMA: return "COMMA"
    case SCOLN: return "SCOLN"
    case IF: return "IF"
    case ELSE: return "ELSE"
    case FUNC: return "FUNC"

    default: return "Unknown Token: " + strconv.Itoa(id)
  }
}
