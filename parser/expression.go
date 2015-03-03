package parser

import (
///  "fmt"
//  "strconv"
  "lang/tokens"
)



/*


*/
func expression(list []tokens.Token) (bool, *Tree) {
  var fn func(*Tree, []tokens.Token) *Tree
  ast := new(Tree)
  ast.Label = EXPRESSION
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
            case tokens.ADD: tmp.Label = ADD
            case tokens.SUB: tmp.Label = SUB
            case tokens.MUL: tmp.Label = MUL
            case tokens.DIV: tmp.Label = DIV
          }
          for j:=2; j <= len(l); j++ {
            if s,t := factor(l[1:j]); s {
              l = l[j:len(l)]
              if len(l) == 0 {
                addChild(tmp, f)
                addChild(tmp, t)
              } else {
                addChild(tmp, f)
                addChild(tmp, fn(t, l))
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
        addChild(ast, f)
      } else {
        addChild(ast, fn(f, list))
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
    ast.Label = IDFR
    ast.Value = list[0].Value
    return true, ast
  } else if list[0].Token == tokens.INT {
    ast := new(Tree)
    ast.Label = INT
    ast.Value = list[0].Value
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
