package main
import "fmt"

func swap (a *int, b *int){
*a,*b = *b,*a
}

func main(){
a,b := 5,6
fmt.Println(a,b)
swap(&a,&b)
fmt.Println(a,b)
}