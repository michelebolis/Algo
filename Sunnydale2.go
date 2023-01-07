package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arco struct {
	destination, luminosita int
}

func main() {
	/*
		MODELLAZIONE:
		Utilizziamo un grafo non orientato pesato
		- I nodi del grafo rappresentano gli svincoli
		- Gli archi rappresentano le gallerie ognuna con la propria luminosita L, il peso dell'arco
		devo trovare SE esiste un cammino che rispetti la regola di Harmony ed eventualmente la sua lunghezza
	*/
	n_svincoli, luminosita, start, goal := createGallerie()
	//fmt.Println(n_svincoli, luminosita, start, goal)
	fmt.Println("Gallerie da attraversare:", countGallerie(start, goal, n_svincoli, luminosita))
}

/*
N è il numero degli svincoli (numerati da 1 a N);
M è il numero delle gallerie;
H è l’indice dello svincolo dove abita Harmony;
S è l’indice dello svincolo dove abita Sarah

Ognuna delle successive M righe descrive una galleria e contiene tre numeri interi A, B
e L separati da uno spazio: i primi due rappresentano gli svincoli collegati dalla galleria
mentre il terzo rappresenta il suo grado di luminosità.
*/
func createGallerie() (n_svincoli int, gallerie map[int]*arco, start int, goal int) {
	//La prima riga dell’input è composta dai quattro numeri interi N, M, H e S
	file, _ := os.Open("gallerie.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	parametri := strings.Split(scanner.Text(), " ")
	n_svincoli, _ = strconv.Atoi(parametri[0])
	n_gallerie, _ := strconv.Atoi(parametri[1])
	start, _ = strconv.Atoi(parametri[2])
	goal, _ = strconv.Atoi(parametri[3])

	gallerie = make(map[int]*arco)
	for i := 0; i < n_gallerie; i++ {
		scanner.Scan()
		parametri = strings.Split(scanner.Text(), " ")
		svincolo1, _ := strconv.Atoi(parametri[0])
		svincolo2, _ := strconv.Atoi(parametri[1])
		luce, _ := strconv.Atoi(parametri[2])
		if gallerie[svincolo1] == nil || luce < gallerie[svincolo1].luminosita {
			gallerie[svincolo1] = &arco{svincolo2, luce}
		}
		if gallerie[svincolo2] == nil || luce < gallerie[svincolo2].luminosita {
			gallerie[svincolo2] = &arco{svincolo1, luce}
		}
	}
	return
}
func countGallerie(start int, goal int, n_svincoli int, gallerie map[int]*arco) int {
	if goal > n_svincoli || start > n_svincoli || gallerie[start] == nil {
		return -1
	}
	passed := make(map[int]bool)
	return countGallerieRic(start, goal, 0, n_svincoli, gallerie, passed)
}
func countGallerieRic(start int, goal int, gallerie_passate int, n_svincoli int, gallerie map[int]*arco, passed map[int]bool) int {
	if gallerie[start] == nil {
		return -1
	}
	if gallerie[start].destination == goal {
		return gallerie_passate + 1
	}
	passed[start] = true
	if passed[gallerie[start].destination] {
		return -1 //SE sono tornato alla galleria di partenza
	}
	//faccio divantare lo start, start.destination
	return countGallerieRic(gallerie[start].destination, goal, gallerie_passate+1, n_svincoli, gallerie, passed)
}
