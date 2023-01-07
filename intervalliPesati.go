package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Programmazione dinamica perche i sottoproblemi si ripetono --> memorization
Soluzione ottima da trovare nei sottoproblemi --> struttura ricorsiva
trovo soluzione degli intervalli da 1 a j (sottoinsieme)
SE contiene j sara il valore dell intervallo j + gli altri
SE non c è j allora la soluzione ottima sara la soluzione ottima fino a j-1

Dobbiamo annotare per ogni j il best
MA dato che alla fine calcoliamo da 0 e poi vado avanti, posso farlo iterativo

*/
type intervallo struct {
	inizio, fine, peso int
}

func main() {
	intervalli := []*intervallo{}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		arg := strings.Split(scanner.Text(), " ")
		inizio_fine := strings.Split(arg[0], "-")
		newIntervallo := new(intervallo)
		newIntervallo.inizio, _ = strconv.Atoi(inizio_fine[0])
		newIntervallo.fine, _ = strconv.Atoi(inizio_fine[1])
		newIntervallo.peso, _ = strconv.Atoi(arg[1])
		intervalli = append(intervalli, newIntervallo)
	}
	//ordino a seconda dell'orario di fine degli intervalli
	sort.Slice(intervalli, func(i, j int) bool { return intervalli[i].fine < intervalli[j].fine })
	soluzione := []*intervallo{}
	somma, soluzione := Opt(len(intervalli)-1, soluzione, intervalli)
	fmt.Println(somma, soluzione)
	for i := 0; i < len(soluzione); i++ {
		fmt.Println(soluzione[i].inizio, soluzione[i].fine, soluzione[i].peso)
	}
	somma, soluzione = (Opt2(len(intervalli)-1, intervalli))
	for i := 0; i < len(soluzione); i++ {
		fmt.Println(soluzione[i].inizio, soluzione[i].fine, soluzione[i].peso)
	}
	fmt.Println(somma)
}

//dato un intervallo in posizione j restituisco l'indice del primo intervallo che non si sovrappone a quest'ultimo
//SE non c è return -1
func p(intervalli []*intervallo, j int) int {
	i := j - 1
	for i >= 0 {
		if !isSovrapposto(intervalli[i], intervalli[j]) {
			return i
		}
		i--
	}
	return i
}

//dati due intervalli restituisce true se si sovrappongono
func isSovrapposto(i1, i2 *intervallo) bool {
	if i2.inizio > i1.fine || i1.inizio > i2.fine {
		return false
	}
	return true
}

/*
per risolvere il problema degli intervalli, massimizzando il peso totale SE ho j intervalli ordinati per l'ora di fine,
per risolvere il problema j mi affido alla soluzione ottima di j-1
*/
func Opt(j int, S []*intervallo, intervalli []*intervallo) (somma int, soluzione []*intervallo) {
	if j == -1 {
		return 0, S
	}
	current := intervalli[j].peso
	toadd, newSoluzione := Opt(p(intervalli, j), S, intervalli)
	current += toadd
	back, soluzione := Opt(j-1, S, intervalli)
	if current > back {
		newSoluzione = append(newSoluzione, intervalli[j])
		return current, newSoluzione
	} else {
		return back, soluzione
	}
}
func Opt2(j int, intervalli []*intervallo) (somma int, soluzione []*intervallo) {
	M := make([]int, j+1)
	for i := 0; i <= j; i++ {
		toAdd := 0
		if p(intervalli, i) != -1 {
			toAdd = M[p(intervalli, i)]
		}
		current := intervalli[i].peso + toAdd
		back := 0
		if i-1 >= 0 {
			back = M[i-1]
		}
		if current > back {
			fmt.Println(current, back)
			M[i] = current
			soluzione = append(soluzione, intervalli[i])
		} else {
			M[i] = back
		}
	}
	return M[j], soluzione
}
