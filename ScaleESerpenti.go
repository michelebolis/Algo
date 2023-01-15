package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	MODELLAZIONE:
	Utilizzo un grafo non orientato non pesato in cui
	- i nodi rappresentano le caselle raggiungibili da una mossa, quindi da 1 a n
	- gli archi rappresentano le possibili mosse, quindi i lanci di dadi
	Ogni lancio di un dado ci puo portare
	- in un nodo esistente che puo essere quello di goal/finale o meno
	- in un nodo non esistente che quindi causa la sconfitta
	obiettivo: trovare il cammino da start a goal con la lunghezza minimo
*/
func main() {
	board := createBoard()
	start := 1
	n_moves, moves := minMovesToWin(board, start, false)
	fmt.Println("Mosse minime per vincere da", start, ":", n_moves, "; Possibile set di dadi:", moves)

	start = 2
	n_moves, moves = minMovesToWin(board, start, false)
	fmt.Println("Mosse minime per vincere da", start, ":", n_moves, "; Possibile set di dadi:", moves)

	start = 1
	n_moves, moves = minMovesToWin(board, start, true)
	fmt.Println("Mosse minime per vincere senza scale e serpenti da", start, ":", n_moves, "; Possibile set di dadi:", moves)
}
func createBoard() []int {
	file, _ := os.Open("board.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	dim := strings.Split(scanner.Text(), " ")
	y, _ := strconv.Atoi(dim[0])
	x, _ := strconv.Atoi(dim[1])
	size := x * y
	board := make([]int, size) //il board lo vedo come un array, da matrice --> array

	addShortcuts(board, scanner) //aggiunge i serpenti
	addShortcuts(board, scanner) //aggiunge le scale
	return board
}
func addShortcuts(board []int, scanner *bufio.Scanner) {
	scanner.Scan()
	n_shortcut, _ := strconv.Atoi(scanner.Text())
	for i := n_shortcut; i > 0; i-- {
		scanner.Scan()
		parametri := strings.Split(scanner.Text(), " ")
		from, _ := strconv.Atoi(parametri[0])
		to, _ := strconv.Atoi(parametri[1])
		board[from-1] = to //in posizione from-1 prescrivo di saltare a to
	}
}
func minMovesToWin(board []int, start int, avoidShortcuts bool) (n_moves int, moves string) {
	type Configurazione struct {
		n_casella  int
		moves      int    //mi salvo le mosse fatte dall inizio per quella combinazioni
		dadiUsciti string //mi salvo i dadi usciti in quella combinazione come stringa
	}
	dado := []int{1, 2, 3, 4, 5, 6}
	coda := []Configurazione{Configurazione{start, 0, "("}}
	//uso coda che riempio dal fondo con tutte le combinazioni prelevando sempre quella di indice 0

	for len(coda) > 0 {
		configurazione := coda[0]
		n_casella := configurazione.n_casella - 1
		coda = coda[1:]
		for i := 0; i < len(dado); i++ {
			new_casella := n_casella + dado[i] //fissata la configurazione ci aggiungo tutti i lanci di dadi possibili
			if new_casella > len(board) {
				continue //finirei fuori dal tabellone
			}
			if !avoidShortcuts {
				prevJump := new_casella
				jumpTo := board[prevJump] - 1 //off by one problem
				for jumpTo > 0 {              //considero la possibilita di avere piu scale/serpenti consecutive
					prevJump = jumpTo
					jumpTo = board[prevJump] - 1 //off by one problem
				}
				new_casella = prevJump //se non ci sono salti, la casella resta new_casella
			}
			newConfigurazione := Configurazione{new_casella + 1, configurazione.moves + 1, configurazione.dadiUsciti + " " + strconv.Itoa(dado[i])}
			if newConfigurazione.n_casella == len(board) { //con la nuova configurazione si è arrivati alla fine
				return newConfigurazione.moves, newConfigurazione.dadiUsciti + " )"
			}
			coda = append(coda, newConfigurazione)
		}
	}
	return -1, ""
}
