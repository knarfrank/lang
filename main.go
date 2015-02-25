package main
import (
  "fmt"
  "strconv"
)
type Tree struct {
  label int
  value string
  child0 *Tree
  child1 *Tree
  child2 *Tree
  child3 *Tree
}

const (
  STATEMENT = iota


)




func main() {
/*  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Compilation Failed.", r)
    }
  }()*/
  lex()
}


func generateError(s string, lineNumber int, pos int, line string) {
  fmt.Println(line)
  panic("Error: " + s + " On Line " + strconv.Itoa(lineNumber) + " column " + strconv.Itoa(pos))
}


func checkError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
