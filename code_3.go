package main

import "fmt"

func fibonacchi (){
	fmt.Println("this section is for printing fibonacchi number")
	var f1 int = 0
	var f2 int = 1 
	var i int = 1
	var a = [20]int {f1,f2}
	fmt.Printf("%v ",f1)
	for i = 2 ; i < 20 ; i++ {
       
	   a[i] = f1 + f2
	   fmt.Printf("%v ",f2)
	   f1 = a[i-1]
	   f2 = a[i]
    
	}
}