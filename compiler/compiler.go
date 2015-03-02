package compiler

import (
  "fmt"
  "lang/parser"
  "strconv"
)
const (
  public = 1
  private = 2
)


type Program struct {
  classes []Class
}

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
  memoryLocation int
}

type Method struct {
  identifier string
  parameters []Parameter
  returns []string
  // Will be default private
  visibilty int
  memoryLocation int
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
  program := new(Program)
  children := parser.GetChildren(ast)

  for _,stmt := range children {
    switch(stmt.Label) {
      case parser.CLASS:
        program.classes = append(program.classes, class(stmt))
      default:
        fmt.Println("Nothing")
    }
  }
}





func class(c *parser.Tree) Class {
  class := new(Class)
  statements := parser.GetChildren(parser.GetChild(c, 0))

  // Loop through each statement in a class (either a declaration or method)
  for _, s := range statements {
    switch(s.Label) {
      case parser.VAR:
        class.attributes = append(class.attributes, attribute(class, s))
      case parser.FUNCTION:
        class.methods = append(class.methods, method(class, s))

    }
  }
  return *class
}

func checkAttributeExists(class *Class, id string) bool {
  exists := false
  for _, e := range class.attributes {
    if e.identifier == id {
      exists = true
      break
    }
  }
  return exists
}


func paramCompare(p1 []Parameter, p2 []Parameter) bool {
  if len(p1) != len(p2) {
    return false
  }
  for i, e := range p1 {
    /*
    if e.identifier != p2[i].identifier {
      return false
    }*/
    if e.paramType != p2[i].paramType {
      return false
    }
  }
  return true
}

func checkMethodExists(class *Class, id string, p []Parameter) bool {
  exists := false
  for _, e := range class.methods {
    if e.identifier == id {
      if paramCompare(e.parameters, p) {
        exists = true
        break
      }

    }
  }
  return exists
}



func method(class *Class, c *parser.Tree) Method {
  method := new(Method)
  parameters := parser.GetChildren(parser.GetChild(c, 0))
  for _, p := range parameters {
    method.parameters = append(method.parameters, methodParameter(p))
  }
  fmt.Println(method.parameters)
  method.identifier = c.Value
  if !checkMethodExists(class, method.identifier, method.parameters) {
    // Deal with params and return types here...
    fmt.Println("boom")
  } else {
    generateError("Method already exists", -1, -1, "")
  }

  return *method
}

func methodParameter(p *parser.Tree) Parameter {
  parameter := new(Parameter)
  parser.DisplayTree(p, 0)
  parameter.identifier = parser.GetChild(p, 0).Value
  if t := parser.GetChild(p, 1).Value; !checkType(t) {
    generateError("Unknown type " + t, -1, -1, "")
  }
  parameter.paramType = parser.GetChild(p, 1).Value
  if len(parser.GetChildren(p)) > 2 {
    // Set default value...
    parameter.optional = true
    parameter.defaultValue = parser.GetChild(p, 2).Value
  }
  return *parameter
}


func checkType(t string) bool {
  return true
}

func attribute(class *Class, c *parser.Tree) Attribute {
  attribute := new(Attribute)
  attribute.identifier = parser.GetChild(c, 0).Value

  // Check if the atribute has already been declared.
  if !checkAttributeExists(class, attribute.identifier) {
    attribute.attributeType = parser.GetChild(c, 1).Value

    fmt.Println(attribute.identifier)
    if len(parser.GetChildren(c)) > 2 {
      if parser.GetChild(c, 2).Label == parser.EXPRESSION {
        // do expression...
      } else {
        // do typing...
        // do expression on 3?
      }
    }
  } else {
    generateError("Class attribute already declared", -1, -1, "")
  }


  return *attribute
}


func generateError(s string, lineNumber int, pos int, line string) {
  fmt.Println(line)
  panic("Compiler Error: " + s + " On Line " + strconv.Itoa(lineNumber) + " column " + strconv.Itoa(pos))
}
