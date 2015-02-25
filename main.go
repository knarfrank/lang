package main
import (
  "fmt"
  "lang/lexer"
  "io/ioutil"
)




func main() {
  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Compilation Failed.", r)
    }
  }()
  // Read in file (should move to function...)
  dat, err := ioutil.ReadFile("test.lg")
  checkError(err)
  program := string(dat)
  lexer.Lex(program)
}


func checkError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
