package main
import (
  "fmt"
  "lang/lexer"
  "lang/parser"
  "io/ioutil"
)




func main() {
  // Read in file (should move to function...)
  dat, err := ioutil.ReadFile("test.lg")
  checkError(err)
  program := string(dat)
  fmt.Println("File Loaded. Compiling Code...")
  tokens := lexer.Lex(program)
  fmt.Println("Lexical Analysis Completed...")
  parser.Parse(tokens)
  fmt.Println("Parse Tree Generation Completed...")
}


func checkError(err error) {
  if err != nil {
    panic(err)
  }
}
