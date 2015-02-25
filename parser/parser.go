package parser

import (
  "fmt"
  "strconv"
  "lang/tokens"
)

const (
  STATEMENT = iota


)
func GenerateTree(list []tokens.Token) bool {
  ast := new(Tree)
  _, t := statements(list)
  addChildren1(ast, t)
  return true
}



func statements(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)

  if len(list) == 0 {
    return true, ast
  }
  for i:=0; i<= len(list); i++ {
    if(list[i].Token == tokens.SCOLN) {
      _, t0 := statement(list[0:i])
      _, t1 := statements(list[i+1:len(list)])
      addChildren2(ast, t0, t1)
      break
    }
  }
  return true, ast
}

func statement(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = STATEMENT
  if s,t := assignment(list); s == true {
    addChildren1(ast, t)
  }
  return true, ast
}
/*
STATEMENTS := <STATEMENT> (STATEMENTS)*
STATEMENT := <ASSIGNMENT>

ASSIGNMENT := <IDFR> <ASMT> <EXPRSSION> <SCOLN>
EXPRSSION := <FACTOR> ((<ADD> | <SUB> | <MUL> | <DIV>) <FACTOR>)*
FACTOR := <IDFR> | <INT> | <RPAREN><EXPRSSION><LPAREN>
FUNCTION := <FUNC> <IDFR> <COLN> <RPAREN> <LPAREN>

*/
func assignment(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  if list[0].Token != tokens.IDFR {
    return false, ast
  }
  if list[1].Token != tokens.ASMT {
    return false, ast
  }
  if !expression(list[2:len(list)]) {
    return false, ast
  }
  if list[len(list)-1].Token != tokens.SCOLN {
    return false, ast
  }
  return true, ast
}

func expression(list []tokens.Token) bool {
  for i:=1; i <= len(list); i++ {
    // If the first characters are a factor
    if factor(list[0:i]) {
      list = list[i:len(list)]
      for {
        if len(list) == 0 {
          break
        }
        // If the next character is a arithmetic operator
        if list[0].Token == tokens.ADD || list[0].Token == tokens.SUB || list[0].Token == tokens.MUL || list[0].Token == tokens.DIV {

          for j:=2; j <= len(list); j++ {

            if factor(list[1:j]) {
              list = list[j:len(list)]
              break
            }
          }
        } else {
          break
        }
        if len(list) != 0 {
          list = list[1:len(list)]
        } else {
          return true
        }
      }
      return true
    } else {
      if i == len(list) {
        generateError("Invalid expression", -1, -1, "")
      }

    }
  }
  return false
}

func factor(list []tokens.Token) bool {
  if list[0].Token == tokens.IDFR || list[0].Token == tokens.INT {
    return true
  } else {
    if list[0].Token == tokens.RPAREN {
      if list[len(list)-1].Token == tokens.LPAREN {
        if expression(list[1:len(list)]){
          return true
        }
      } else {
        // return error mismatched brackets...
      }
    }
    return false
  }
}


func generateError(s string, lineNumber int, pos int, line string) {
  fmt.Println(line)
  panic("Parse Error: " + s + " On Line " + strconv.Itoa(lineNumber) + " column " + strconv.Itoa(pos))
}


func checkError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
