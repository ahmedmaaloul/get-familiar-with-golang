package main

import (
	"fmt"
)
func main(){
	age:= 24
	fmt.Println(age <= 30)
	fmt.Println(age >= 30)
	fmt.Println(age == 24)
	fmt.Println(age != 30)

	if age < 20{
		fmt.Println("age is less than 30")
	} else if age < 25{
		fmt.Println("age is less than 25")
	} else {
		fmt.Println("age is not less than 25")
	}
	names := []string{"mario","luigi", "yoshi","peach","bowser"}
	for index, val := range names{
		if index == 1 {
			fmt.Println("continuning at pos", index)
			continue 
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, val)
	}
	
}