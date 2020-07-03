package main

import "fmt"

func bubblesort(numbers []int){
	n:= len(numbers)
	for i:=0; i<n; i++ {
		sweep(numbers)
	}
}

func sweep (numbers [] int){
firstIndex:=0
secondIndex:=1
n:= len(numbers)
for(secondIndex<n){
	var firstNumber int = numbers[firstIndex]
	var secondNumber int = numbers[secondIndex]

	if (firstNumber>secondNumber){
		numbers[firstIndex] = secondNumber;
		numbers[secondIndex] = firstNumber
	}
	firstIndex++
	secondIndex++
}

}
func main(){
var numbers [] int = [] int {4,2,3,1,3} 
fmt.Println(numbers)
bubblesort(numbers)
fmt.Println(numbers)



}