package parser

import (
  "fmt"
  "strconv"
  "lang/tokens"
)

const (
  ROOT = iota
  STATEMENT
  ASSIGNMENT
  EXPRESSION
  IDFR
  INT
  ADD
  SUB
  MUL
  DIV
)
func getLabel(l int) string {
  switch(l) {
    case ROOT: return "ROOT"
    case STATEMENT: return "STATEMENT"
    case ASSIGNMENT: return "ASSIGNMENT"
    case EXPRESSION: return "EXPRESSION"
    case IDFR: return "IDFR"
    case INT: return "INT"
    case ADD: return "ADD"
    case SUB: return "SUB"
    case MUL: return "MUL"
    case DIV: return "DIV"
    default: return "BLEH"
  }
}
func GenerateTree(list []tokens.Token) bool {
  ast := new(Tree)
  ast.label = ROOT
  _, t0, t1:= statements(list)
  addChildren2(ast, t0, t1)
  displayTree(ast, 0)
  return true
}



func statements(list []tokens.Token) (bool, *Tree, *Tree) {
  if len(list) == 0 {
    return true, nil, nil
  }
  for i:=0; i<= len(list); i++ {
    if(list[i].Token == tokens.SCOLN) {
      t3 := new(Tree)
      t3.label = STATEMENT
      _, t0 := statement(list[0:i])
      _, t1, t2 := statements(list[i+1:len(list)])
      addChildren2(t3, t1, t2)
      return true, t0, t3
      break
    }
  }
  return true, nil, nil
}

func statement(list []tokens.Token) (bool, *Tree) {
  if s,t := assignment(list); s == true {
    return true, t
  }
  return false, nil
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
  ast.label = ASSIGNMENT
  if list[0].Token != tokens.IDFR {
    return false, ast
  }
  if list[1].Token != tokens.ASMT {
    return false, ast
  }
  b, t := expression(list[2:len(list)])
  if !b {
    return false, ast
  }
  addChildren2(ast, node(IDFR, list[0].Value), t)
  return true, ast
}

func expression(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = EXPRESSION
  var fn func(*Tree, []tokens.Token) *Tree
  for i:=1; i <= len(list); i++ {
    // If the first characters are a factor
    if s, f := factor(list[0:i]); s {
      list = list[i:len(list)]


      fn = func(f *Tree, l []tokens.Token) *Tree {
        tmp := new(Tree)
        if len(l) == 0 {
          return nil
        }
        if l[0].Token == tokens.ADD || l[0].Token == tokens.SUB || l[0].Token == tokens.MUL || l[0].Token == tokens.DIV {
          switch(list[0].Token) {
            case tokens.ADD: tmp.label = ADD
            case tokens.SUB: tmp.label = SUB
            case tokens.MUL: tmp.label = MUL
            case tokens.DIV: tmp.label = DIV
          }
          for j:=2; j <= len(l); j++ {
            if s,t := factor(l[1:j]); s {

              l = l[j:len(l)]
              if len(list) == 0 {
                addChildren2(tmp, f, t)
              } else {
                addChildren2(tmp, f, fn(t, l))
              }
              return tmp
            }
          }

        }
        return nil
      }
      if len(list) == 0 {
        addChildren1(ast, f)
      } else {
        addChildren1(ast, fn(f, list))
      }
      return true, ast
    } else {
      if i == len(list) {
        generateError("Invalid expression", -1, -1, "")
      }

    }
  }
  return false, nil
}

func factor(list []tokens.Token) (bool, *Tree) {

  if list[0].Token == tokens.IDFR {
    ast := new(Tree)
    ast.label = IDFR
    ast.value = list[0].Value
    return true, ast
  } else if list[0].Token == tokens.INT {
    ast := new(Tree)
    ast.label = INT
    ast.value = list[0].Value
    return true, ast
  } else {
    if list[0].Token == tokens.RPAREN {
      if list[len(list)-1].Token == tokens.LPAREN {
        if b, t := expression(list[1:len(list)-1]); b {
          return true, t
        }
      } else {
        // return error mismatched brackets...
      }
    }
    return false, nil
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
