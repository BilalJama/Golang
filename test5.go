package main

import("fmt"
"sort"
)

type Car struct{
Brand string
Model string
Year int
}



func main(){
c1:= Car{"Fiat", "Pinto", 2002}
c2:= Car{"Fiat", "Pinto", 1932}
c3:= Car{"Fiat", "Pinto", 1992}
x:= []Car{}
x = append(x,c1)
x = append(x,c2)
x = append(x,c3)

sort.Slice(x, func(i, j int) bool{
	return x[i].Year < x[j].Year
	
	})
   fmt.Println(x)



}