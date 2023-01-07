package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	corretto, err := verifica(scanner.Text())
	if !corretto {
		fmt.Println(err)
	} else {
		fmt.Println("File html ben formato")
	}
}

/*
per ogni tag di apertura esiste uno e un solo un tag di chiusura corrispondente
se due tag di apertura compaiono in un determinato ordine, i corrispondenti tag di chiusura
devono comparire nell’ordine opposto
es "<a> <b> </b> <c> <d> </d> </c> </a>" BEN formata
es "<a> <b> </b> </c>" errore in posizione 4
es "<a> <b> </b> <c> <d> </d>" sono rimasti aperti i seguenti tag: <a> <c>
*/
func verifica(espressione string) (bool, string) {
	var pila []string
	entry := strings.Split(espressione, " ")
	for i := 0; i < len(entry); i++ {
		if entry[i][1] == '/' {
			tagapertura := pila[len(pila)-1]
			pila = pila[:len(pila)-1]
			if tagapertura != entry[i][2:len(entry[i])-1] {
				return false, "errore in pos " + fmt.Sprintf("%d", i+1)
			}
		} else {
			pila = append(pila, entry[i][1:len(entry[i])-1])
		}
	}
	if len(pila) != 0 {
		var s string
		for i := 0; i < len(pila); i++ {
			s += " <" + pila[i] + ">"
		}
		return false, "sono rimasti aperti i seguenti tag: " + s[1:]
	}
	return true, ""
}
