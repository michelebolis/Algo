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
	archi["D"]=[]to_arco{to_arco{"B", 16}, to_arco{"A", 25}, to_arco{"C", 36}, to_arco{"F", 1}, to_arco{"G", 4}, to_arco{"E", 9}}
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

/*
ALGORITMO Prim(Grafo G=(V, E, w)) -> Albero
Sia C una coda con priorità vuota
Sia d e vicino due array con indici di V
//Inizialmente
FOR EACH v appartenente V DO
	d[v] <- infinito //ogni vertice v ha priorità d[v]=infinito
	C.insert(v, infinito)
T <- (vuoto, vuoto)

//ad ogni passo
DO
	Y <- C.deleteMin() //preleva vertice y con d[y] minima
	VT <- VT U {y}
	IF d[y] != infinito THEN
		X <- vicino[y] //aggiungo a T il vertice y e l'arco (x, y) con x=vicino[y]
		ET <- ET U {(x, y)}
	FOR EACH (y, z) Є E DO
		IF z !Є VT AND w(y, z) < d[z] THEN
			d[z] <- w(y, z)
			C.changeKey(z, w(y, z))
			vicino[z] <- y
WHILE C!=vuoto
RETURN T
*/
func Prim(grafo map[string]*node, archi map[string][]to_arco)(newGrafo map[string]*node, newArchi map[string][]to_arco){
	var coda[]to_arco 
	d:= make(map[string]int)
	vicino:=make(map[string]string) 
	for key,_:=range grafo{
		d[key]=math.MaxInt/2 
		coda=append(coda, archi[key]...)
	}
	newGrafo=make(map[string]*node)
	newArchi=make(map[string][]to_arco)

	for len(coda)!=0{
		v:=coda[0].to //y
		coda=coda[1:]
		newGrafo[v]=grafo[v]
		if d[v]!=math.MaxInt/2{
			x:=vicino[v]//aggiungo a T il vertice y e l'arco (x, y) con x=vicino[y]
			/*
			for index, to_arco:=range archi[x]{
				if to_arco.to==v{
					for index2,to_arco2:=range newArchi[x]{
						if to_arco2.to==x{
							newArchi[x][index2]=(archi[x][index])
						}
					}
					
					break
				}

			}
			*/
		}
		for _, to_archi:=range archi[v]{
			z:=to_archi.to
			_, ok:= newGrafo[z]
			if !ok && to_archi.cost < d[z]{
				/*
				d[z] <- w(y, z)
				C.changeKey(z, w(y, z))
				vicino[z] <- y
				*/
				d[z]=to_archi.cost 
				coda[find(coda, z)]=to_arco{z,to_archi.cost}
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