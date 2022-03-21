package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	name  string
	skill string
}

func main() {
	var people [5]Person
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(people); i++ {
		fmt.Println("Enter the name of the person:")
		scanner.Scan()
		people[i].name = scanner.Text()
		fmt.Println("Enter the skill of the person:")
		scanner.Scan()
		people[i].skill = scanner.Text()
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(people[i])
	}

        for index,value:=range people{
            fmt.Println(index,value)
        }
}
