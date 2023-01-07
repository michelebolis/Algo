package main
import(
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "os"
)
type listNode struct{
  item int
  next *listNode
}
type linkedList struct{
  head *listNode
}
func main(){
  var list linkedList
  file,_:=os.Open("in.txt")
  scanner:=bufio.NewScanner(file)
  for scanner.Scan(){
    riga:=strings.Trim(scanner.Text(), " ")
    if riga=="f"{
      return
    }
    if riga[0]=='+'{
      num,_:=strconv.Atoi(riga[1:])
      presente,_:=searchList(list, num)
      fmt.Println("il numero non è presente, lo aggiungo")
      if !presente{
        list=addNewNode(list, num)
      }
    }else if riga[0]=='-'{
      num,_:=strconv.Atoi(riga[1:])
      _,list=removeItem(list,num)
    }else if riga[0]=='?'{
      num,_:=strconv.Atoi(riga[1:])
      trovato,_:=searchList(list, num)
      if trovato{
        fmt.Println(num, "appartiene all'insieme")
      }else{
        fmt.Println(num, "non appartiene all'insieme")
      }
    }else if riga=="c"{
      n:=numLista(list)
      fmt.Println("L'insieme è composto da", n, "elementi")
    }else if riga=="p"{
      printList(list)
    }else if riga=="o"{
      printReverseList(list)
    }else if riga=="d"{
      list.head=nil
    }
  }
}

func newNode(num int )*listNode{
  return &listNode{num,nil}
}
func addNewNode(list linkedList, num int) linkedList{
  node:=newNode(num)
  node.next=list.head
  list.head=node
  return list
}
func printList(list linkedList){
  p:=list.head
  for p!=nil{
    fmt.Print(p.item, " ")
    p=p.next
  }
  fmt.Println()
}
func printReverseList(list linkedList){
  var s string
  p:=list.head
  for p!=nil {
    s= strconv.Itoa(p.item) + " " + s
    p=p.next
  }
  fmt.Println(s)
}
func searchList(list linkedList, num int) (bool, *listNode){
  p:=list.head
  for p!=nil{
    if p.item==num{
      return true, p
    }
    p=p.next
  }
  return false, nil
}
func removeItem(list linkedList, num int) (bool, linkedList){
  var prec *listNode
  prec=nil
  p:=list.head
  for p!=nil{
    if p.item==num{
      if prec==nil{
        list.head=p.next
        return true, list
      }
      prec.next=p.next
      return true, list
    }

    prec=p
    p=p.next
  }
  return false, list
}
func numLista(list linkedList)(n int) {
  p:=list.head
  for p!=nil{
    n++
    p=p.next
  }
  return
}
