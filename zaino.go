package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
SE ho una soluzione ottima per uno zaino che regge P e tolgo dalla soluzione un elemento, trovo S' che avra peso P-pt e V-vt
SE da soluzione ottima tolgo un qualsiasi elemento, ho ancora la soluzione ottima
*/
type oggetto struct {
	key   int
	peso  int
	value int
}

func main() {
	file, _ := os.Open("input2.txt")
	scanner := bufio.NewScanner(file)
	var oggetti []*oggetto
	for scanner.Scan() {
		campi := strings.Split(scanner.Text(), " ")
		newoggetto := new(oggetto)
		newoggetto.key, _ = strconv.Atoi(campi[0])
		newoggetto.peso, _ = strconv.Atoi(campi[1])
		newoggetto.value, _ = strconv.Atoi(campi[2])
		oggetti = append(oggetti, newoggetto)
	}
	riempiZaino(20, oggetti)
	riempiZainoGenerale(10, oggetti)
}

/*
costruite un vettore s di lunghezza P, tale che s[i] contenga
il valore delle soluzioni ottime per uno zaino di peso i. Il vettore si può costruire partendo da i = 0 e incrementando i,
sfruttando la relazione vista sopra. Una volta completato il vettore s, il valore ottimo di una soluzione per P è chiaramente
dato da s[P]
//ho dei tipi di oggetti, magari riempio lo zaino di 20 oggetti del tipo 1
*/
/*Vogliamo riempire lo zaino non superando P con il peso complessivo degli oggetti,
ma allo stesso tempo massimizzando la somma dei valori degli oggetti nello zaino.
*/
func riempiZaino(P int, oggetti []*oggetto) {
	zaino := make([]int, P+1)
	for i := 0; i < len(zaino); i++ { //i rappresenta il peso di uno zaino da 0 a P
		for j := 0; j < len(oggetti); j++ {
			// controllo se lo zaino con peso i + il nuovo oggetto ci sta nello zaino massimo di peso P
			// aggiungo allo zaino di peso i l'oggetto[j] SE E SOLO SE lo zaino con peso i + oggetto.peso avra value maggiore di quanto ha ora
			if i+oggetti[j].peso <= P && zaino[i]+oggetti[j].value > zaino[i+oggetti[j].peso] {
				zaino[i+oggetti[j].peso] = zaino[i] + oggetti[j].value
			}
		}
	}
	fmt.Println(zaino)
}

//stavolta ho degli oggetti NON dei tipi
//riempire al massimo peso P lo zaino massimizzandone il value
func riempiZainoGenerale(P int, oggetti []*oggetto) {
	//creo lo zaino come matrice
	zaino := make([][]int, P+1)
	for i := 0; i < P+1; i++ {
		zaino[i] = make([]int, len(oggetti)+1)
	}

	//procedo colonna per colonna
	for j := 0; j < len(oggetti); j++ {
		for i := 0; i < len(zaino); i++ {
			//propago il value alla colonna sotto SE la colonna sotto ha value inferiore
			if zaino[i][j+1] <= zaino[i][j] {
				zaino[i][j+1] = zaino[i][j]
			}

			//controllo se lo zaino con peso i + il nuovo oggetto ci sta nello zaino massimo di peso P
			value := zaino[i][j] + oggetti[j].value
			if i+oggetti[j].peso <= P {
				//propago il risultato alla colonna adiacente, quindi all'oggetto adiacente
				zaino[i+oggetti[j].peso][j+1] = value
			}
		}
	}
	fmt.Println(zaino[len(zaino)-1][len(oggetti)]) //il value massimo sara nell'ultima riga, ultima colonna
}
