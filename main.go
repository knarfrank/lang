
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
  // Make parser not diaply the ast by default
  displayAst := false


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


  // If command line option "-ast" is set
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
