package main

import "fmt"

func main() {
	age := 10           // int literal
	name := "Jack"      // string literal
	rightHanded := true // boolean literal

	fmt.Printf("%s is %d years old. Right handed:%t", name, age, rightHanded)
	fmt.Println()

	ageInTenYears:=age+10

	fmt.Printf("In ten years,%s will be %d years old",name,ageInTenYears)
	fmt.Println()
	
	isATeenager:=age>=13
	fmt.Println(name,"is a teemager:",isATeenager)
	// In total the expressions are : 9
	// Expression is a "something" which can be evaluated
	// to a single value during runtime
}
