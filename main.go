package main

import "fmt"

func sayHowdie() {
	fmt.Println("He says howdie")

}

func robotRunner()  {

	click := 0
	distanceCount := 0
	sprintZone := 110
	for distanceCount <= sprintZone {
		click++
		fmt.Println("Click number =", click)
		distanceCount+=10
		fmt.Println("Distance covered =", distanceCount)

	}

}

func main() {

	/*{
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		sayHowdie()
	}*/

	robotRunner()
}
