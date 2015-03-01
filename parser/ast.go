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
  Children []*Tree
  child0 *Tree
  child1 *Tree
  child2 *Tree
  child3 *Tree
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
func displayTree(t *Tree, i int) {
  // Indent the line.
  for j:=0; j < i; j++ {
    fmt.Print("   ")
  }

  // Print label and value
  PrintTree(t)

  // Recursivly call displayTree on children.
  if t.Children != nil {
    if len(t.Children) != 0 {
      for j:=0; j < len(t.Children); j++ {
        displayTree(t.Children[j], i+1)
      }
    }
  } else { // Should get rid of this after i get rid of crude way of doing children.
    if t.child0 != nil {
      displayTree(t.child0, i+1)
    }
    if t.child1 != nil {
      displayTree(t.child1, i+1)
    }
    if t.child2 != nil {
      displayTree(t.child2, i+1)
    }
    if t.child3 != nil {
      displayTree(t.child3, i+1)
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


// Don't really need this funcitons. should delete them...
func addChildren1(parent *Tree, c0 *Tree) {
  parent.child0 = c0
}

func addChildren2(parent *Tree, c0 *Tree, c1 *Tree) {
  parent.child0 = c0
  parent.child1 = c1
}

func addChildren3(parent *Tree, c0 *Tree, c1 *Tree, c2 *Tree) {
  parent.child0 = c0
  parent.child1 = c1
  parent.child2 = c2
}
