package main

import (
	"fmt"
	"math"
)

type node struct {
	key string
	//campi ...
}
type arco struct {
	to   string
	cost int
}

func main() {
	grafo := make(map[string]*node)
	archi := make(map[string][]arco)
	grafo["A"] = new(node)
	grafo["B"] = new(node)
	grafo["C"] = new(node)
	grafo["D"] = new(node)
	grafo["E"] = new(node)
	grafo["F"] = new(node)
	grafo["G"] = new(node)

	archi["A"] = []arco{arco{"B", 7}, arco{"C", 6}}
	archi["B"] = []arco{arco{"D", 7}}
	archi["C"] = []arco{arco{"A", 7}, arco{"E", 4}}
	archi["D"] = []arco{arco{"C", 2}, arco{"G", 2}}
	archi["E"] = []arco{arco{"D", 3}, arco{"F", 2}}
	archi["F"] = []arco{arco{"G", 4}}
	archi["G"] = []arco{arco{"E", 4}}
	distanze, policy := Dijkstra(grafo, archi, "A")
	fmt.Println(distanze)
	fmt.Println(policy)
	printMinPath(grafo, "A", "F", policy)
}

func Dijkstra(grafo map[string]*node, archi map[string][]arco, s string) (distanze map[string]int, policy map[string]string) {
	distanze = make(map[string]int)
	policy = make(map[string]string)
	var coda []arco
	for key, _ := range grafo { //tutti i nodi tranne quello di partenza hanno costo infinito
		if key != s {
			distanze[key] = math.MaxInt / 2
		} else {
			distanze[key] = 0
		}
	}
	/*
		Sia C una coda con priorità vuota
		FOR EACH v Є V DO C.insert(v, d[v])
	*/
	for key, cost := range distanze { //creo una coda con priorita semplice con il min in indice 0
		coda = append(coda, arco{key, cost})
		if cost == 0 {
			coda[0], coda[len(coda)-1] = coda[len(coda)-1], coda[0]
		}
	}
	for len(coda) != 0 {
		//u <- C.deleteMin()
		u := coda[0]
		coda = coda[1:]
		if len(coda) != 0 {
			coda = fix(coda) //metto il min della coda in 0
		}

		for _, arco := range archi[u.to] { //analizzo ogni arco uscente da u (da u a u.to)
			if distanze[u.to]+arco.cost < distanze[arco.to] {
				distanze[arco.to] = distanze[u.to] + arco.cost
				//C.changeKey(v, d[v])
				coda[find(coda, arco.to)].cost = distanze[arco.to]
				coda = fix(coda)
				policy[arco.to] = u.to
			}
		}
	}
	return
}
func fix(coda []arco) []arco {
	posmin := 0
	min := coda[0].cost
	for pos := 1; pos < len(coda); pos++ {
		if min > coda[pos].cost {
			posmin = pos
			min = coda[pos].cost
		}
	}
	if posmin != 0 {
		coda[0], coda[posmin] = coda[posmin], coda[0]
	}
	return coda
}
func find(coda []arco, x string) (index int) {
	for i := 0; i < len(coda); i++ {
		if coda[i].to == x {
			return i
		}
	}
	return -1
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
