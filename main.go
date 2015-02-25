package main
import (

  "lang/lexer"
)




func main() {
/*  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Compilation Failed.", r)
    }
  }()*/
  lexer.Lex()
}
