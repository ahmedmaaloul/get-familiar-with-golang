package main

import "fmt"


func main(){

	myBill := newBill("mario's bill")
	myBill.updateTip(10)
	myBill.addItem("fricasse",1)
	fmt.Println(myBill.format())
	
}