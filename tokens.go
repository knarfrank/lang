package main
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
  IF     // "if"
  ELSE   // "else"
  FUNC    // "func"




)

type token struct {
  token int
  value string
  regex string
}


func getTokens() []token {
  var tokens []token
  tokens = make([]token, 0, 0)
  tokens = append(tokens, token{NL, "", "^\\n$"},
                          token{WS, "", "^(\\s)+$"},
                          token{SCMNT, "", "^/\\*$"},
                          token{ECMNT, "", "^\\*/$"},
                          token{IF, "", "^(if)$"},
                          token{ELSE, "", "^(else)$"},
                          token{FUNC, "", "^(func)$"},
                          token{RPAREN, "", "^\\($"},
                          token{LPAREN, "", "^\\)$"},
                          token{RCURL, "", "^\\{$"},
                          token{LCURL, "", "^\\}$"},
                          token{COLN, "", "^:$"},
                          token{COLN, "", "^;$"},
                          token{COMMA, "", "^,$"},
                          token{IDFR, "", "^[a-z|A-Z][a-z|A-Z]*$"},
                          token{INT, "", "^[0-9]+$"},
                          token{BOOL, "", "^(true|false)$"},
                          token{ADD, "", "^\\+$"},
                          token{SUB, "", "^-$"},
                          token{MUL, "", "^\\*$"},
                          token{DIV, "", "^/$"},
                          token{MOD, "", "^%$"},
                          token{AND, "", "^&$"},
                          token{OR, "", "^\\|$"},
                          token{XOR, "", "^\\^$"},
                          token{BNOT, "", "^~$"},
                          token{RSHFT, "", "^>>$"},
                          token{LSHFT, "", "^<<$"},
                          token{EQU, "", "^(==)$"},
                          token{LST, "", "^<$"},
                          token{GRT, "", "^>$"},
                          token{LNOT, "", "^!$"},
                          token{ASMT, "", "^=$"},
                          token{IGNR, "", "^.$"})

  return tokens
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



    default: return "unknow token: " + strconv.Itoa(id)
  }
}
