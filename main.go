package main

import (
	"bufio"
	"fmt"
	"github.com/Dukenbayev/onelab_home_work3/convert/atoi"
	"github.com/Dukenbayev/onelab_home_work3/fibonacci"
	"github.com/Dukenbayev/onelab_home_work3/reverse"
	"os"
)

func main(){
	//reverse string
	fmt.Println("Input ur string")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("1. Reverse string of %s: is %s", reverse.Reverse(text),text)

	//fibonacci check
	fmt.Println("2. Fibonacci Numbers:")
	generator := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(generator(), " ")

	}
	// String to Integer
	var string_number string

	fmt.Println("3. String to Integer.\n")
	fmt.Println("Input your string number:")
	fmt.Scan(&string_number)
	number,err := atoi.Atoi(string_number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Check type: %T and number: %v" ,number,number)

}