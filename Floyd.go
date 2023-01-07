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
	grafo["V1"] = new(node)
	grafo["V2"] = new(node)
	grafo["V3"] = new(node)
	grafo["V4"] = new(node)
	grafo["V5"] = new(node)
	archi := make(map[string][]arco)
	archi["V1"] = append(archi["V1"], arco{"V2", 5})
	archi["V2"] = append(archi["V2"], arco{"V4", 4})
	archi["V2"] = append(archi["V2"], arco{"V3", 7})
	archi["V3"] = append(archi["V3"], arco{"V2", 1})
	archi["V4"] = append(archi["V4"], arco{"V1", 2})
	archi["V4"] = append(archi["V4"], arco{"V3", 2})
	fmt.Println("Grafo", grafo)
	fmt.Println("Archi", archi)
	distanzeMin, policy := FloydWarshall(grafo, archi)
	fmt.Println(distanzeMin)
	fmt.Println(policy)
	printMinPath(grafo, "V3", "V1", policy)
	printMinPath(grafo, "V5", "V1", policy)
}

type from_to struct {
	from string
	to   string
}

func FloydWarshall(grafo map[string]*node, archi map[string][]arco) (d map[from_to]int, policy map[from_to]string) {
	d = make(map[from_to]int)
	policy = make(map[from_to]string)
	for key_from, _ := range grafo { //per ogni nodo guardo gli adiacenti
		for key_to, _ := range grafo {
			new_from_to := from_to{key_from, key_to}
			if key_from == key_to { //SE stesso nodo di partenza e di arrivo
				d[new_from_to] = 0
				policy[new_from_to] = "self"
			} else {
				archiUscenti := archi[key_from] //prendo tutti gli archi uscenti da from
				found := false
				for _, arco := range archiUscenti {
					if arco.to == key_to { //se l'arco attuale è l'arco d'arrivo fissato nel for 2
						d[new_from_to] = arco.cost //il costo per andare da from a to è l'arco diretto
						found = true
						break
					}
				}
				if !found {
					d[new_from_to] = math.MaxInt / 2 //SE non c è un arco da from a to setto infinito in d
					policy[new_from_to] = ""         //non so ancora cosa devo fare quando sono in from
				} else {
					policy[new_from_to] = key_to
				}
			}
		}
	}
	for key_from, _ := range grafo { //fisso partenza
		for key_middle, _ := range grafo { //fisso nodo intermedio
			for key_to, _ := range grafo { //fisso nodo destinazione
				costo_cammino := d[from_to{key_from, key_middle}] + d[from_to{key_middle, key_to}] //nuovo cammino costa from-middle + middle-to
				if costo_cammino < d[from_to{key_from, key_to}] {                                  //se il nuovo è minore dell'arco diretto
					d[from_to{key_from, key_to}] = costo_cammino   // aggiungo a d tale costo
					policy[from_to{key_from, key_to}] = key_middle // aggiungo a policy il middle quando voglio andare da from a to
				}
			}
		}
	}
	return
}
func printMinPath(grafo map[string]*node, start string, goal string, policy map[from_to]string) {
	if grafo[start] == nil || grafo[goal] == nil || policy[from_to{start, goal}] == "" {
		return
	}
	if start == goal { //sono arrivato a destinazione
		fmt.Println(goal)
		return
	}
	fmt.Println(start)
	printMinPath(grafo, policy[from_to{start, goal}], goal, policy) //ogni volta start diventa cio che mi dice di fare la policy
}
