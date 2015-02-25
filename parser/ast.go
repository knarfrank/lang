package parser
import (
  "fmt"
  //"strconv"
  //"io/ioutil"

)
type Tree struct {
  label int
  value string
  children []*Tree
  child0 *Tree
  child1 *Tree
  child2 *Tree
  child3 *Tree
}


func displayTree(t *Tree, i int) {
  if t.children != nil {
    if len(t.children) != 0 {
      for j:=0; j < len(t.children); j++ {
        displayTree(t.children[j], i+1)
      }
    }
  } else {
    for j:=0; j < i; j++ {
      fmt.Print("   ")
    }
    fmt.Println(getLabel(t.label) + " (" + t.value +")")
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


func addChild(parent *Tree, child *Tree) {
  parent.children = append(parent.children, child)
}

func node(label int, value string) *Tree {
  node := new(Tree)
  node.label = label
  node.value = value
  return node
}

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
