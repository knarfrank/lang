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
  Line  int
  Col   int
}


func GetTokens() []Token {
  var Tokens []Token
  Tokens = make([]Token, 0, 0)
  Tokens = append(Tokens, Token{NL, "", "^\\n$", 0, 0},
                          Token{WS, "", "^(\\s)+$", 0, 0},
                          Token{SCMNT, "", "^/\\*$", 0, 0},
                          Token{ECMNT, "", "^\\*/$", 0, 0},
                          Token{IF, "", "^(if)$", 0, 0},
                          Token{ELSE, "", "^(else)$", 0, 0},
                          Token{FUNC, "", "^(func)$", 0, 0},
                          Token{LET, "", "^(let)$", 0, 0},
                          Token{RPAREN, "", "^\\($", 0, 0},
                          Token{LPAREN, "", "^\\)$", 0, 0},
                          Token{RCURL, "", "^\\{$", 0, 0},
                          Token{LCURL, "", "^\\}$", 0, 0},
                          Token{COLN, "", "^:$", 0, 0},
                          Token{SCOLN, "", "^;$", 0, 0},
                          Token{COMMA, "", "^,$", 0, 0},
                          Token{BOOL, "", "^(true|false)$", 0, 0},
                          Token{IDFR, "", "^[a-z|A-Z][a-z|A-Z]*$", 0, 0},
                          Token{INT, "", "^[0-9]+$", 0, 0},
                          Token{ADD, "", "^\\+$", 0, 0},
                          Token{SUB, "", "^-$", 0, 0},
                          Token{MUL, "", "^\\*$", 0, 0},
                          Token{DIV, "", "^/$", 0, 0},
                          Token{MOD, "", "^%$", 0, 0},
                          Token{AND, "", "^&$", 0, 0},
                          Token{OR, "", "^\\|$", 0, 0},
                          Token{XOR, "", "^\\^$", 0, 0},
                          Token{BNOT, "", "^~$", 0, 0},
                          Token{RSHFT, "", "^>>$", 0, 0},
                          Token{LSHFT, "", "^<<$", 0, 0},
                          Token{EQU, "", "^(==)$", 0, 0},
                          Token{LST, "", "^<$", 0, 0},
                          Token{GRT, "", "^>$", 0, 0},
                          Token{LNOT, "", "^!$", 0, 0},
                          Token{ASMT, "", "^=$", 0, 0},
                          Token{IGNR, "", "^.$", 0, 0})

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
