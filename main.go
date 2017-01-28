package main

import "fmt"
import "time"

func main() {

	for i := 1; i <= 5; i++ {
		var x = getDate()
		echo(x)
	}

}

func getDate() time.Time {
	return time.Now()
	
}

func echo(x time.Time) {
	fmt.Println(x)
}