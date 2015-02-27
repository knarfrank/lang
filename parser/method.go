package parser
import (
  //"fmt"
  //"strconv"
  "lang/tokens"
)



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
      // parse parameters
      if b, t := methodParams(list[3:i]); b {
        addChild(ast, t)
      } else {
        // ?
      }
      break
    }
    i++;
  }
  // check arrow for return type (Maybe i should make a new token?)
  if list[i+1].Token != tokens.SUB || list[i+2].Token != tokens.GRT {
    return false, ast
  }

  // Check for return type brackets
  if list[i+3].Token != tokens.RPAREN || list[i+5].Token != tokens.LPAREN {
    return false, ast
  }

  // Check return type
  if c, t := methodReturnType(list[i+3]); c {
    addChild(ast, t)
  } else {
    return false, ast
  }
  if len(list[i+7:len(list)-1]) != 0 {
    stmts := new(Tree)
    statements(stmts, list[i+7:len(list)-1])
    addChild(ast, stmts)
  } else {
    // empty function
  }

  return true, ast
}

func methodReturnType(t tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = RETURNTYPE
  return true, ast
}


func methodParams(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.label = PARAMETERS
  //fmt.Println(list)

  return true, ast
}
