package parser

import (
  "fmt"
  "strconv"
  "lang/tokens"
)

const (
  STATEMENTS = iota
  STATEMENT
  IFSTATEMENT
  CLASS
  FUNCTION
  BOOLEXPRESSION
  ASSIGNMENT
  EXPRESSION
  IDFR
  INT
  BOOL
  ADD
  SUB
  MUL
  DIV
  NOT

)
func getLabel(l int) string {
  switch(l) {
    case STATEMENTS: return "STATEMENTS"
    case STATEMENT: return "STATEMENT"
    case IFSTATEMENT: return "IFSTATEMENT"
    case CLASS: return "CLASS"
    case FUNCTION: return "FUNCTION"
    case BOOLEXPRESSION: return "BOOLEXPRESSION"
    case ASSIGNMENT: return "ASSIGNMENT"
    case EXPRESSION: return "EXPRESSION"
    case IDFR: return "IDFR"
    case INT: return "INT"
    case BOOL: return "BOOL"
    case ADD: return "ADD"
    case SUB: return "SUB"
    case MUL: return "MUL"
    case DIV: return "DIV"
    case NOT: return "NOT"
    default: return "BLEH"
  }
}


func Parse(list []tokens.Token) bool {
  ast := new(Tree)
  statements(ast, list)
  displayTree(ast, 0)
  return true
}



func statements(ast *Tree, list []tokens.Token) {
  ast.label = STATEMENTS
  if len(list) == 0 {
    return
  }
  for i:=0; i< len(list); i++ {
    if list[i].Token == tokens.SCOLN {
      t3 := new(Tree)
      t3.label = STATEMENT
      _, t0 := statement(list[0:i])
      addChild(ast, t0)
      statements(ast, list[i+1:len(list)])
      return

    } else if list[i].Token == tokens.RCURL {
      count := 1
      for {
        // If there is an open bracket but no code.
        if len(list[i:len(list)-1]) == 0 {
          generateError("Curly Bracket mismatch.", list[0].Line, list[0].Col, "")
        }
        i++
        if list[i].Token == tokens.LCURL {
          count--
          if count == 0 {
            t3 := new(Tree)
            t3.label = STATEMENT
            _, t0 := statement(list[0:i+1])
            addChild(ast, t0)
            statements(ast, list[i+1:len(list)])
            return
          }
        } else if list[i].Token == tokens.RCURL {
          count++
        }

      }

    } else {
     // nothing... error here? Nope..
    }
  }
  return
}



/*


*/
func statement(list []tokens.Token) (bool, *Tree) {
  if s,t := assignment(list); s {
    return true, t
  } else if s,t := ifStatement(list); s{
    return true, t
  } else if s,t := class(list); s {
    return true, t
  } else {
    generateError("Invalid Statement", list[0].Line, list[0].Col, "")
  }
  return false, nil
}



/*


*/
func class(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = CLASS
  if list[0].Token != tokens.CLASS {
    return false, ast
  }

  if list[1].Token != tokens.IDFR {
    generateError("Invalid Class Name", list[1].Line, list[1].Col, "")
  }
  if list[2].Token != tokens.RCURL || list[len(list)-1].Token != tokens.LCURL {
    generateError("Curly Bracket Mismatch", list[1].Line, list[1].Col, "")
  }
  if b, t := classBody(list[3:len(list)-1]); b {
    addChild(ast, t)
  } else {
    return false, ast
  }

  ast.value = list[1].Value
  return true, ast
}


func classBody(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = STATEMENTS
  // If class is empty
  if len(list) == 0 {
    return true, ast
  }

  for i:=0; i< len(list); i++ {

    if list[i].Token == tokens.SCOLN {
      fmt.Println("stmt")

    } else if list[i].Token == tokens.RCURL {

      count := 1
      for {
        // If there is an open bracket but no code.
        if len(list) == 0 {
          break
        }
        i++
        if list[i].Token == tokens.LCURL {
          count--
          if count == 0 {
            if b, t := method(list[0:i+1]); b {
              addChild(ast, t)
            }



            // Set list to the rest of the class body
            list = list[i+1:len(list)]

            i = 0
          }
        } else if list[i].Token == tokens.RCURL {
          count++
        }

      }
    } else {
     // nothing... error here? Nope..
    }
  }

  return true, ast
}


func method(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = FUNCTION
  if list[0].Token != tokens.FUNC {
    return false, ast
  }
  if list[1].Token != tokens.IDFR {
    return false, ast
  }
  ast.value = list[1].Value

  if list[2].Token != tokens.RPAREN {
    return false, ast
  }
  i := 3
  for {
    if list[i].Token == tokens.LPAREN {
      break
    }
    i++;
  }

  return true, ast
}

/*


*/
func ifStatement(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = IFSTATEMENT

  // If statement starts off with IF keyword
  if list[0].Token != tokens.IF {
    return false, ast
  }

  // If next character is an open bracket
  if list[1].Token != tokens.RPAREN {
    return false, ast
  }

  // Search for boolean expression
  i := 2
  for i=2; i < len(list); i++ {
    if list[i].Token == tokens.LPAREN {
      if c, t := booleanExpression(list[2:i]); c {
        addChild(ast, t)
        break
      } else {
        return false, ast
      }
    }
  }

  // Test for curly bracket starting statement code.
  if list[i+1].Token != tokens.RCURL {
    return false, ast
  }
  stmt := new(Tree)
  statements(stmt, list[i+2:len(list)-1])
  addChild(ast, stmt)
  return true, ast
}





/*


*/
func assignment(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = ASSIGNMENT
  if len(list) < 3 {
    return false, ast
  }
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


/*


*/
func booleanExpression(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = BOOLEXPRESSION
  if list[0].Token == tokens.BOOL {
    addChildren1(ast, node(BOOL, list[0].Value))
    return true, ast
  } else {
    // If first character is equal to a unary operation such as negation
    if list[0].Token == tokens.LNOT {
      _, t := booleanExpression(list[1:len(list)])
      b := node(NOT, "")
      addChildren1(b, t)
      addChildren1(ast, b)
      return true, ast
    } else {
      for i:=1; i <= len(list); i++ {
        if s, _ := factor(list[0:i]); s {
          list = list[i:len(list)]
        }
      }
    }
  }
  return true, ast
}





/*


*/
func expression(list []tokens.Token) (bool, *Tree) {
  var fn func(*Tree, []tokens.Token) *Tree
  ast := new(Tree)
  ast.label = EXPRESSION
  for i:=1; i <= len(list); i++ {
    // If the first characters are a factor
    if s, f := factor(list[0:i]); s {
      list = list[i:len(list)]

      // This nested function deals with repeated operations such as 5+2+d+5
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
              if len(l) == 0 {
                addChildren2(tmp, f, t)
              } else {
                addChildren2(tmp, f, fn(t, l))
              }
              return tmp
            }
          }
          generateError("Missing operand in expression", l[0].Line, l[0].Col, "")
        } else {
          generateError("Invalid Operation in Statement", l[0].Line, l[0].Col, "")
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
        generateError("Invalid Statement", list[0].Line, list[0].Col, "")
      }
    }
  }
  return false, nil
}



/*


*/
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
