package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertice struct {
	nome     string
	eta      int
	hobby    []string
	followed []*vertice
}
type grafo struct {
	vertici map[string]*vertice //la chiave è il nome dell'utente
}

func main() {
	g := readGrafo()
	printGrafo(g)
	printHobbyFollowed(*g, "Andrea")
	printHobbyFollower(*g, "Andrea")
}
func readGrafo() *grafo {
	file, _ := os.Open("utenti.txt")
	scanner := bufio.NewScanner(file)
	g := new(grafo)
	g.vertici = make(map[string]*vertice)
	for scanner.Scan() { //es input Mario;20;Cucina;Calcio
		input := strings.Split(scanner.Text(), ";")
		newVertice := new(vertice)
		newVertice.nome = input[0]
		newVertice.eta, _ = strconv.Atoi(input[1])
		newVertice.hobby = input[2:]
		g.vertici[input[0]] = newVertice
	}

	file, _ = os.Open("relazioni.txt")
	scanner = bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() { //arco A-->B SE utente A segue utente B
		input := strings.Split(scanner.Text(), ";")
		g.vertici[input[0]].followed = append(g.vertici[input[0]].followed, g.vertici[input[1]])
	}
	return g
}
func printGrafo(g *grafo) {
	for nome, info := range g.vertici {
		fmt.Println(nome + ", info: ")
		fmt.Print("\t", "Età: ", info.eta, "\n")
		fmt.Print("\t", "Hobby: ", info.hobby[0])
		for i := 1; i < len(info.hobby); i++ {
			fmt.Print(", ", info.hobby[i])
		}
		fmt.Println()
		fmt.Println("Followed: ")
		for i := 0; i < len(info.followed); i++ {
			fmt.Print("\t", info.followed[i].nome, "\n")
		}
	}
}
func printHobbyFollower(g grafo, s string) {
	fmt.Print("Gli hobby di ", s, " e di chi lo segue sono:", "\n")
	hobby := make(map[string]bool) //permette di non avere duplicati di hobby
	for i := 0; i < len(g.vertici[s].hobby); i++ {
		hobby[g.vertici[s].hobby[i]] = true //salvo tutti gli hobby della persona fissata
	}
	for nome, info := range g.vertici {
		for i := 0; i < len(info.followed); i++ {
			if info.followed[i].nome == s {
				for k := 0; k < len(g.vertici[nome].hobby); k++ {
					hobby[g.vertici[nome].hobby[k]] = true //salvo gli hobby di chi segue la persona fissata
				}
				break
			}
		}
	}

	for chiave, _ := range hobby {
		fmt.Println("\t", "-", chiave) //stampo gli hobby
	}
}
func printHobbyFollowed(g grafo, s string) {
	fmt.Print("Gli hobby di ", s, " e di chi segue sono:", "\n")
	hobby := make(map[string]bool)
	for i := 0; i < len(g.vertici[s].hobby); i++ {
		hobby[g.vertici[s].hobby[i]] = true //salvo tutti gli hobby della persona fissata
	}
	for i := 0; i < len(g.vertici[s].followed); i++ {
		followed := g.vertici[s].followed[i].nome
		for j := 0; j < len(g.vertici[followed].hobby); j++ {
			hobby[g.vertici[followed].hobby[j]] = true
		}
	}
	for chiave, _ := range hobby {
		fmt.Println("\t", "-", chiave) //stampo gli hobby
	}
}
