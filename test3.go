package main
import("fmt"
 "sort")


func dedup (fixed []int) []int {
table:= make(map[int] bool)
list:= [] int {}

for _, value:= range fixed {
	if table[value] == false{
	list = append(list, value)
	table[value] = true
	fmt.Println(table)
	}

}
return list
}

func main() {
	a := []int{100, 20, 20, 299, 20, 40, 70, 20, 90, 100}
	fmt.Println(a)
	fmt.Println(dedup(a))
	sort.Ints(a)
	fmt.Println(a)

}