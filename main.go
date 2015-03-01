package main
import (
  "fmt"
  "os"
  "io/ioutil"
  "lang/lexer"
  "lang/parser"
  "lang/compiler"

)




func main() {
  fmt.Println(len(os.Args), os.Args)
  // Read in file (should move to function...)
  dat, err := ioutil.ReadFile("test.lg")
  checkError(err)
  program := string(dat)
  fmt.Println("File Loaded. Compiling Code...")
  tokens := lexer.Lex(program)
  fmt.Println("Lexical Analysis Completed...")

  displayAst := false
  for _, a := range os.Args {
    if a == "-ast" {
      displayAst = true
      break
    }
  }
  tree := parser.Parse(tokens, displayAst)


  fmt.Println("Parse Tree Generation Completed...")
  compiler.Compile(tree)
  fmt.Println("Compilation Complete")
}


func checkError(err error) {
  if err != nil {
    panic(err)
  }
}
