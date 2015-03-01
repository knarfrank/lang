package main
import (
  "fmt"
  "lang/lexer"
  "lang/parser"
  "lang/compiler"
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
  fmt.Println("File Loaded. Compiling Code...")
  tokens := lexer.Lex(program)
  fmt.Println("Lexical Analysis Completed...")
  tree := parser.Parse(tokens)
  fmt.Println("Parse Tree Generation Completed...")
  compiler.Compile(tree)
  fmt.Println("Compilation Complete")
}


func checkError(err error) {
  if err != nil {
    panic(err)
  }
}
