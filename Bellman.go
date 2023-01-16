package main

import (
	"fmt"
	"math"
)

type node struct {
	campi string
}
type arco struct {
	to   string
	cost int
}

func main() {
	grafo := make(map[string]*node)
	A := new(node)
	B := new(node)
	C := new(node)
	D := new(node)
	E := new(node)
	grafo["A"] = A
	grafo["B"] = B
	grafo["C"] = C
	grafo["D"] = D
	grafo["E"] = E
	archi := make(map[string][]arco)
	archi["A"] = []arco{{"B", 4}, {"C", 2}}
	archi["B"] = []arco{{"D", 2}, {"C", 3}, {"E", 3}}
	archi["C"] = []arco{{"B", 1}, {"E", 5}, {"D", 4}}
	archi["D"] = []arco{}
	archi["E"] = []arco{{"D", -5}}
	d, policy := Bellman(grafo, archi, "A")
	for chiave, val := range d {
		fmt.Println(chiave, val)
	}
	printMinPath(grafo, "A", "D", policy)
}
func Bellman(grafo map[string]*node, archi map[string][]arco, s string) (d map[string]int, policy map[string]string) {
	d = make(map[string]int)
	policy = make(map[string]string)
	//inizializzo d con infinito tranne per s
	for key, _ := range grafo {
		if key != s {
			d[key] = math.MaxInt / 2
		} else {
			d[key] = 0
		}
	}

	for k := 1; k < len(grafo); k++ {
		for nodoCorrente, _ := range grafo { //fissato un nodo
			for _, arco := range archi[nodoCorrente] { //guardo tutti i suoi archi uscenti
				newD := d[nodoCorrente] + arco.cost
				if newD < d[arco.to] {
					d[arco.to] = newD
					policy[arco.to] = nodoCorrente
				}
			}
		}
	}
	return d, policy
}
func printMinPath(grafo map[string]*node, start string, goal string, policy map[string]string) {
	path := goal
	for {
		if goal == start {
			break
		}
		goal = policy[goal]
		path = goal + " -> " + path
	}
	fmt.Println(path)
}
