package main

import "fmt"

func main() {
	var h []int
	//var A []int = []int{1, 2, 5, 10, 3, 7, 11, 15, 17, 20, 9, 15, 8, 16}
	var A []int = []int{7, 9, 14, 12, 11, 13}
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}

	for _, val := range A {
		h = append(h, val)
		heapify_up(h)
	}
	arrayToHeap(A)
	fmt.Println("A", A)
	stampaAlberoASommario(A)
	fmt.Println(h)
	stampaAlberoASommario(h)
	fmt.Println()
	_, h = deleteMin(h)
	fmt.Println(h)
	stampaAlberoASommario(h)
}
func arrayToHeap(A []int) {
	for j, _ := range A {
		for {
			i := (j - 1) / 2             //prendo il padre
			if i == j || (A[j] > A[i]) { //figlio > padre
				break //giusto
			}
			//vuol dire che padre>figlio, quindi li scambio
			A[i], A[j] = A[j], A[i]
			j = i //ora ripartiro esaminando il padre di i
		}
	}
}

func heapify_up(heap []int) { //lo richiamo se aggiungo in coda allo heap
	j := len(heap) - 1
	for {
		i := (j - 1) / 2                   //prendo il padre
		if i == j || (heap[j] > heap[i]) { //figlio > padre
			break //giusto
		}
		//vuol dire che padre>figlio, quindi li scambio
		heap[i], heap[j] = heap[j], heap[i]
		j = i //ora ripartiro esaminando il padre di i
	}
}
func down(heap []int) {
	i := 0
	for {
		j := 2*i + 1 // left child
		if j >= len(heap) {
			break
		}
		j2 := j + 1
		if j2 < len(heap) && (heap[j2] < heap[j]) {
			j = j2 // right child
		}
		if !(heap[j] < heap[i]) {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
}
func stampaAlberoASommario(heap []int) {
	for i := 0; i < len(heap); i++ {
		fmt.Print(heap[i], " figli ")
		if 2*i+1 < len(heap) {
			fmt.Print(heap[2*i+1])
			if 2*i+2 < len(heap) {
				fmt.Print(" ", heap[2*i+2])
			}
		}
		fmt.Println()
	}
}
func min(h []int) int {
	return h[0]
}
func deleteMin(h []int) (int, []int) {
	min := h[0]
	h[0] = h[len(h)-1]
	h = h[:len(h)-1]
	down(h)
	return min, h
}
