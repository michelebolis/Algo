package main 
import(
	"math"
	"fmt"
)
func main(){
	grafo:=make(map[string]*node)
	grafo["A"]=new(node)
	grafo["B"]=new(node)
	grafo["C"]=new(node)
	grafo["D"]=new(node)
	grafo["E"]=new(node)
	grafo["F"]=new(node)
	grafo["G"]=new(node)
	archi:=make(map[string][]to_arco)
	archi["A"]=[]to_arco{to_arco{"B", 17}, to_arco{"D", 25}, to_arco{"C", 28}}
	archi["B"]=[]to_arco{to_arco{"A", 17}, to_arco{"D", 18}, to_arco{"E", 3}}
	archi["C"]=[]to_arco{to_arco{"A", 28}, to_arco{"D", 36}, to_arco{"F", 23}}
	archi["D"]=[]to_arco{to_arco{"B", 18}, to_arco{"A", 25}, to_arco{"C", 36}, to_arco{"F", 1}, to_arco{"G", 4}, to_arco{"E", 9}}
	archi["E"]=[]to_arco{to_arco{"B", 3}, to_arco{"D", 9}, to_arco{"G", 15}}
	archi["F"]=[]to_arco{to_arco{"C", 23}, to_arco{"G", 20}, to_arco{"D", 1}}
	archi["G"]=[]to_arco{to_arco{"D", 4}, to_arco{"F", 20}, to_arco{"E", 15}}
	newGrafo, newArchi:=Prim(grafo, archi)
	fmt.Println(newGrafo)
	fmt.Println(newArchi)
}	
type node struct{
	key string
	//campi ...
}
type to_arco struct{
	to string 
	cost int 
}
func Prim(grafo map[string]*node, archi map[string][]to_arco)(newGrafo map[string]*node, newArchi map[string][]to_arco){
	var coda[]to_arco 
	d:= make(map[string]int)
	vicino:=make(map[string]string) 
	for key,_:=range grafo{
		d[key]=math.MaxInt/2 
		coda=append(coda, to_arco{key,d[key]})
	}
	newGrafo=make(map[string]*node)
	newArchi=make(map[string][]to_arco)
	
	for len(coda)!=0{
		v:=coda[0].to
		coda=coda[1:]
		fix(coda, d)
		newGrafo[v]=grafo[v] //aggiungo a T il vertice y
		if d[v]!=math.MaxInt/2{
			x:=vicino[v] //aggiungo l'arco (x, y) con x=vicino[y]
			for _, uscente:=range archi[x]{
				if uscente.to==v{
					newArchi[x]=append(newArchi[x], uscente)
					newArchi[uscente.to]=append(newArchi[uscente.to], to_arco{x, uscente.cost})
					break
				}
			}
		}
		for _, uscente:=range archi[v]{
			z:=uscente.to
			_, ok:= newGrafo[z]
			if !ok && uscente.cost < d[z]{
				d[z]=uscente.cost 
				new_to_arco:=to_arco{z,uscente.cost}
				coda[find(coda, z)]=new_to_arco
				coda=fix(coda, d)
				vicino[z]=v
			}
		}
	}
	return 
}

func find(coda[]to_arco, key string)int{
	for i:=0;i<len(coda);i++{
		if coda[i].to==key{
			return i
		}
	}
	return -1
}
func fix(coda[]to_arco, d map[string]int)[]to_arco{
	posMin:=0
	min:=math.MaxInt
	for i:=0;i<len(coda);i++{
		if d[coda[i].to]<min{
			min=d[coda[i].to]
			posMin=i
		}
	}
	if posMin!=0{
		coda[0], coda[posMin]=coda[posMin], coda[0]
	}
	return coda
}