package parser
import (
  "fmt"
  "lang/tokens"
)


/*
  Function to parse a method.
*/
func method(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.Label = FUNCTION
  fmt.Println("")
  if list[0].Token != tokens.FUNC {
    return false, ast
  }
  if list[1].Token != tokens.IDFR {
    return false, ast
  }
  ast.Value = list[1].Value

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
  if c, t := methodReturnType(list[i+4]); c {
    addChild(ast, t)
  } else {
    return false, ast
  }
  // Parse statements inside function
  if len(list[i+7:len(list)-1]) != 0 {
    stmts := new(Tree)
    statements(stmts, list[i+7:len(list)-1])
    addChild(ast, stmts)
  } else {
    // empty function
  }

  return true, ast
}

/*
  Returns tree of return types of a method
*/
func methodReturnType(t tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.Label = RETURNTYPE
  ast.Value = t.Value
  return true, ast
}

/*
  Parses fuction parameters
*/
func methodParams(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.Label = PARAMETERS
  i := 0
  for {
    // If parameter is the last one

    if len(list) <= 5 {
      if c, t := methodParam(list[0: len(list)]); c {
        addChild(ast, t)
      }
      break
    }
    // Each parameter is seperated by a comma
    if list[i].Token == tokens.COMMA {
      if c, t := methodParam(list[0: i]); c {
        addChild(ast, t)
      }
      list = list[i+1:len(list)]
      i = 0
    } else {
      i++
    }
  }
  return true, ast
}

func methodParam(list []tokens.Token) (bool, *Tree) {
  ast := new(Tree)
  ast.Label = VAR
  if list[0].Token != tokens.IDFR || list[1].Token != tokens.COLN || list[2].Token != tokens.IDFR {
    generateError("Invalid Function Parameter", list[0].Line, list[0].Col, "")
  }
  addChild(ast, node(IDFR, list[0].Value))
  addChild(ast, node(IDFR, list[2].Value))
  if len(list) > 3 {

    if len(list) != 5 {
      generateError("Invalid Function Parameter", list[0].Line, list[0].Col, "")
    }

    if list[3].Token == tokens.ASMT {
      addChild(ast, node(INT, list[4].Value))
    } else {
      generateError("Invalid Function Parameter", list[0].Line, list[0].Col, "")
    }
  }
  return true, ast
}
