package main
import (
  //"fmt"
  //"regexp"
  //"strconv"
  //"io/ioutil"
)

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
func generateTree(list []token) bool {
  ast := new(Tree)
  _, t := statements(list)
  addChildren1(ast, t)
  return true
}
