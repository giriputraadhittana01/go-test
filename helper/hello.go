package helper

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello "+name)
}

func Calculate(x,y int, operator string) int{
	switch operator{
		case "+":
			return x+y
		case "-":
			return x-y
		case "/":
			return x/y
		case "*":
			return x*y
		case "%":
			return x%y
		default:
			return -1
	}  
}


// Awal huruf besar menandakan public
var Name = "Boba"
// Awal huruf kecil menandakan private
var version = "1.0.2"