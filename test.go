package main
import ("fmt")

func fact (a int) int{
    x:=1
	for i:=1; i<=a; i++ {
		x=x*i
	}
return(x)
}

func main(){
	fmt.Println(fact(4))
}


 