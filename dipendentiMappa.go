package main
import(
  "fmt"
)
type organigramma map[string]gerarchie
type gerarchie struct{
  supervisore string
  subordinati[]string
}
func main(){
  
}
//Dato un certo dipendente, stampare l’elenco dei suoi subordinati.
func stampaSubordinati(dipendente string, organigramma organigramma){
  subo:=organigramma[dipendente].subordinati
  for i:=0;i<len(subo);i++{
    fmt.Println(subo[i])
  }
}
//Contare quanti sono i dipendenti che non hanno alcun subordinato.
func quantiSenzaSubordinati(organigramma organigramma)(i int){
  for _,gerarchia:=range organigramma{
    if len(gerarchia.subordinati)==0{
      i++
    }
  }
  return
}
//Dato un certo dipendente, individuare chi è il suo supervisore.
func supervisore(dipendente string, organigramma organigramma)string{
  return organigramma[dipendente].supervisore
}
/*
Dato un certo dipendente, stampare la lista dei dipendenti che si trovano sopra di lui gerarchicamente,
partendo dal suo supervisore e risalendo la gerarchia fino a un dipendente di
massimo livello.
*/
func stampaImpiegatiSopra(dipendente string, organigramma organigramma){
  sup:=organigramma[dipendente].supervisore
  if sup=="" {return}
  if organigramma[sup].supervisore!=""{
    stampaSubordinati(organigramma[sup].supervisore, organigramma)
    stampaImpiegatiSopra(sup, organigramma)
  }else{
    fmt.Println(sup)
  }
}
/*
Stampare l’elenco di tutti i dipendenti –non importa l’ordine–, indicando per ciascuno chi è
il suo supervisore (tranne nel caso di dipendenti di massimo livello).
*/
func stampaImpiegatiConSupervisore(organigramma organigramma){
  for dipendente, gerarchia:=range organigramma{
    fmt.Println(dipendente, ", supervisore:", gerarchia.supervisore)
  }
}
