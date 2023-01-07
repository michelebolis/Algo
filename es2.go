package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := riempiArray()
	fmt.Println(slice)
	selectionSort(slice)
	fmt.Println(slice)

	fmt.Println()

	//slice=riempiArray()
	//fmt.Println(slice)
	//selectionSortRicorsivo(slice)
	//fmt.Println(slice)

	slice = riempiArray()
	fmt.Println(slice)
	mergeSort(slice)
	fmt.Println(slice)
}
func riempiArray() []int {
	var slice []int
	dim := 4
	rand.Seed(time.Now().Unix())
	for i := 0; i < dim; i++ {
		slice = append(slice, rand.Intn(500))
	}
	return slice
}

/*
func selectionSortRicorsivo(slice []int){
  if len(slice)==1 || len(slice)==0{
    return
  }
//ho sbagliato
  for i:=0; i<len(slice)-1;i++{
    if slice[i]>slice[len(slice)-1]{
      slice[i], slice[len(slice)-1]=slice[len(slice)-1], slice[i]
    }
  }
  selectionSortRicorsivo(slice[1:])
}
*/

func selectionSort(slice []int) {
	for i := 0; i < len(slice)-2; i++ {
		m := i
		for k := i + 1; k < len(slice)-1; k++ {
			if slice[k] < slice[m] {
				m = k
			}
		}
		slice[m], slice[i] = slice[i], slice[m]
	}
}

func mergeSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	m := len(slice) / 2
	prima_meta := slice[0:m]
	seconda_meta := slice[m:len(slice)]
	mergeSort(prima_meta)
	mergeSort(seconda_meta)
	copy(slice, merge(prima_meta, seconda_meta)) //IMPORTANTE
}
func merge(slice1 []int, slice2 []int) (slice_merged []int) {
	i1 := 0
	i2 := 0
	for i1 < len(slice1) && i2 < len(slice2) {
		if slice1[i1] <= slice2[i2] {
			slice_merged = append(slice_merged, slice1[i1])
			i1++
		} else {
			slice_merged = append(slice_merged, slice2[i2])
			i2++
		}
	}
	if i1 < len(slice1) {
		for i := i1; i < len(slice1); i++ {
			slice_merged = append(slice_merged, slice1[i])
		}
	} else {
		for i := i2; i < len(slice2); i++ {
			slice_merged = append(slice_merged, slice2[i])
		}
	}
	return
}
