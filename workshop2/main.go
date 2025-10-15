package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting,"hello"))	
	fmt.Println(strings.Contains(greeting,"bye"))
	fmt.Println(strings.ReplaceAll(greeting,"hello","hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(greeting)	
	fmt.Println(strings.Index(greeting, "hello"))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println("original string value =", greeting)
	ages:= []int{45, 35, 34,12,32,24,55,90,240,240, 28,31}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)
	names := []string{"yoshi","mario","peach", "bowser","luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names,"bowser"))
}