package compiler

import (
  "fmt"
  "lang/parser"
)
const (
  public = 1
  private = 2
)

/*
  Stucture for a class
*/
type Class struct {
  name string
  //parent string
  attributes []Attribute
  methods []Method

}

/*
  Stucture for an attribute within a class
*/
type Attribute struct {
  identifier string
  attributeType string
  // Will be default private
  visibility int
}

type Method struct {
  identifier string
  parameters []Parameter
  returns []string
  // Will be default private
  visibilty int
}

type Parameter struct {
  // The parameter identifier that will be used within the method
  identifier string
  // String to primative or object
  paramType string
  // If it has a default value
  optional bool
  // The default value
  defaultValue string
}

func Compile(ast *parser.Tree) {
  children := parser.GetChildren(ast)

  for _,stmt := range children {
    switch(stmt.Label) {
      case parser.CLASS:
        fmt.Println("class")
      default:
        fmt.Println("Nothing")
    }
  }
}
