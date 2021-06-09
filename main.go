package main

import (
	"bufio"
	"fmt"
	"github.com/Dukenbayev/onelab_home_work3/convert/atoi"
	"github.com/Dukenbayev/onelab_home_work3/convert/itoa"
	"github.com/Dukenbayev/onelab_home_work3/fibonacci"
	"github.com/Dukenbayev/onelab_home_work3/reverse"
	"os"
)

func main(){
	//reverse string
	fmt.Println("1. Reverse string. Input ur string:")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("Reverse string of %s: is %s", reverse.Reverse(text),text)

	//fibonacci check
	fmt.Println("2. Fibonacci Numbers:")
	generator := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(generator(), " ")

	}
	// String to Integer
	var string_number string

	fmt.Println("\n3. String to Integer. Input your string number:")
	fmt.Scan(&string_number)
	number,err := atoi.Atoi(string_number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Check type: %T and number: %v \n" ,number,number)

	//Integer to String
	var integer_number int = 412412

	fmt.Println("4. Integer to String.Input your integer number:")
	fmt.Scan(&integer_number)
	fmt.Printf("Check type: %T and number: %s" ,itoa.Itoa(integer_number),itoa.Itoa(integer_number))

	//sort_imports
}