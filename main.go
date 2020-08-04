package main

import (
	"fmt"
)

func main() {
	// var name string = "aki"
	// var name = "aki"
	name := "aki"

	fmt.Printf("hello %v \n", name)
	fmt.Printf("hello %v again! \n", name)
}


// 数当てゲーム
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10) + 1
	count := 0

	for {
		var guess int

		fmt.Print("Your guess? ")
		fmt.Scanf("%v", &guess)

		count++

		if answer == guess {
			fmt.Printf("Bingo! It took %v guesses\n", count)
			break
		} else if answer > guess {
			fmt.Println("The answer is higher!")
		} else {
			fmt.Println("The anwer is lower!")
		}
	}
}
