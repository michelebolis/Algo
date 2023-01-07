package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(valuta(scanner.Text()))
	scanner.Scan()
	fmt.Println(converti(scanner.Text()))
}

/*implementa il precedente
algoritmo, utilizzando una pila. La funzione riceve una espressione in notazione postfissa
e restituisce il suo valore. es "5 3 - 2 *" output: 4
*/
func valuta(espressione string) int {
	var pila []int
	entry := strings.Split(espressione, " ")
	for i := 0; i < len(entry); i++ {
		var num int
		if unicode.IsNumber(rune(entry[i][0])) {
			num, _ = strconv.Atoi(entry[i])
		} else {
			op2 := pila[len(pila)-1]
			op1 := pila[len(pila)-2]
			pila = pila[:len(pila)-2]
			switch entry[i] {
			case "+":
				num = op2 + op1
			case "-":
				num = op2 - op1
			case "/":
				num = op2 / op1
			case "*":
				num = op2 * op1
			}
		}
		pila = append(pila, num)
	}
	return pila[0]
}

/*
riceve una espressione
in notazione infissa e restituisce l’espressione equivalente in notazione postfissa, usando
l’algoritmo qui sopra. es "( ( 5 - 3 ) * 2 )" output: 5 3 - 2 *
*/
func converti(espressione string) (s string) {
	var pila []byte
	entry := strings.Split(espressione, " ")
	for i := 0; i < len(entry); i++ {
		if unicode.IsNumber(rune(entry[i][0])) {
			s += " " + (entry[i])
		} else {
			switch entry[i] {
			case "(":
				continue
			case ")":
				s += " " + string(pila[len(pila)-1])
				pila = pila[:len(pila)-1]
			default:
				pila = append(pila, entry[i][0])
			}
		}
	}
	if len(pila) != 0 {
		for i := len(pila) - 1; i >= 0; i-- {
			s += " " + string(pila[i])
		}
	}
	return s[1:]
}
