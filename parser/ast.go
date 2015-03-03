package parser
import (
  "fmt"
  //"strconv"
  //"io/ioutil"

)

/*
  Struct for abstract syntax tree
*/
type Tree struct {
  Label int
  Value string
  Line int
  Column int
  Children []*Tree
}



/*
  Get child from array of children
*/
func GetChild(t *Tree, i int) *Tree {
  return t.Children[i]
}

func GetChildren(t *Tree) []*Tree {
  return t.Children
}
/*
  Prints out information about a tree node
*/
func PrintTree(t *Tree) {
  fmt.Println(getLabel(t.Label) + " (" + t.Value +")")
}

/*
  Function that displays a given tree
*/
func DisplayTree(t *Tree, i int) {
  // Indent the line.
  for j:=0; j < i; j++ {
    fmt.Print("   ")
  }

  // Print label and value
  PrintTree(t)

  // Recursivly call displayTree on children.
  if len(t.Children) != 0 {
    for j:=0; j < len(t.Children); j++ {
      DisplayTree(t.Children[j], i+1)
    }
  }
}

/*
  Adds a child tree to a given tree
*/
func addChild(parent *Tree, child *Tree) {
  parent.Children = append(parent.Children, child)
}

/*
  Returns new tree with given label and value
*/
func node(label int, value string) *Tree {
  node := new(Tree)
  node.Label = label
  node.Value = value
  return node
}
