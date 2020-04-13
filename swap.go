package main

import "fmt"

func swap(x,y *int){
	x*y*=y*x*
}
func main (){
	x :=1
	y :=2
	fmt.Println("до",x,"до", y)
	swap(&x,&y)
	fmt.Println("после",x,"после", y)
}	