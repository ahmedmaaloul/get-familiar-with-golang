package main

import "fmt"

func main(){
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)


	nameOne = "deutsch"
	nameThree = "bowser"
	fmt.Println(nameOne, nameTwo, nameThree)
	nameFour := "yoshi"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne,ageTwo,ageThree)

	var num int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256
	fmt.Println(num, numTwo, numThree)
	age:= 24
	name:= "Ahmed"
	var score float32 = 99.5
	fmt.Println(score)
	var scoreTwo float64 = 882348347284294729.7
	fmt.Println(scoreTwo)
	scoreThree := 1.4
	fmt.Println(scoreThree)
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line! \n")
	fmt.Println("goodbye ninjas !")
	fmt.Println("my age is", age, "yes my age is", age)
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %0.1f points! \n", 255.55)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
	ages:= [4]int{24,28,31,55}
	names:= [4]string{"drachkoun","drachkouna","akechismou","akechismeha"}
	fmt.Println(names, len(names))
	fmt.Println(ages,len(ages))
	names[1] = "yoshi hay hay"
	fmt.Println(names, len(names))


	
	var scores = []int{100,50,30}
	scores[2] = 25 
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne,rangeTwo,rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}