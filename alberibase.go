package main

import (
	"fmt"
	"math/rand"
	"time"
)

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}
type bitree struct {
	root *bitreeNode
}

func main() {
	slice := creaRandomSlice(10)
	fmt.Println(slice)
	bitree := arr2tree(slice, 0)
	inorder(bitree)
	fmt.Println()
	preorder(bitree)
	fmt.Println()
	postorder(bitree)
	fmt.Println()

	stampaAlberoASommario(bitree, 0)
}
func creaRandomSlice(max int) (newSlice []int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < max; i++ {
		newSlice = append(newSlice, generaNum())
	}
	return
}
func generaNum() int {
	return rand.Intn(500)
}
func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)
	return root
}
func newNode(val int) *bitreeNode {
	return &bitreeNode{nil, nil, val}
}
func inorder(node *bitreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Println(node.val)
	inorder(node.right)
}
func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	preorder(node.left)
	preorder(node.right)
}
func postorder(node *bitreeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Println(node.val)
}

func stampaAlberoASommario(node *bitreeNode, spaces int) { //stampa ricorsiva
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	if node != nil {
		fmt.Println(node.val)
		if !(node.left == nil && node.right == nil) {
			stampaAlberoASommario(node.right, spaces+1)
			stampaAlberoASommario(node.left, spaces+1)
		}
	} else {
		fmt.Println()
	}
}
