package main

func main() {

}
func IntegerSort(A []int) []int {
	max := max(A)
	Y := make([]int, max+1)
	for i := 0; i < len(A); i++ {
		Y[A[i]]++ //riempio Y con le occorrenza dei val di A
	}
	k := 0
	for i := 0; i < len(A); i++ {
		for Y[k] == 0 {
			k++ //salto le posizioni con valore (occorrenze) 0
		}
		A[i] = k //riempio A con k (valore)
		Y[k]--   //tolgo un'occorrenza di k in Y
	}
	return A
}
func BucketSort(A []int) []int { //invece che []int ho []*Node
	type Node struct {
		key int
		//*item
		next *Node
	}
	max := max(A)
	Y := make([]*Node, max+1)
	for i := 0; i < len(A); i++ {
		newNode := new(Node)
		newNode.key = A[i]
		newNode.next = Y[A[i]]
		Y[A[i]] = newNode
	}
	j := 0
	for i := 0; i < len(Y); i++ {
		node := Y[i]
		for node != nil {
			A[j] = node.key
			j++
			node = node.next
		}
	}
	return A
}
func max(A []int) (max int) {
	for i := 0; i < len(A); i++ {
		if max < A[i] {
			max = A[i]
		}
	}
	return
}
