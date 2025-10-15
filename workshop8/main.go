package main

import "fmt"

func main(){
	menu := map[string]float64{
		"soup":				4.99,
		"pie": 				7.99,
		"salad":			6.99,
		"toffee pudding":	3.55, 
	}
	fmt.Println(menu)
	for k, v := range menu{
		fmt.Println(k, "-", v)
	}
	phonebook := map[int]string{
		228462423:	"Faker",
		234248394:	"Dragon",
		234234823:	"Skipi",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[234234823])
	phonebook[228462423]= "Ahmed yofaker"
	fmt.Println(phonebook)
	phonebook[234248394] = "yoshi hay hay, brainrot"
	fmt.Println(phonebook)
}