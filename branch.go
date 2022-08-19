package main

import (
	"fmt"
	"os"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong grade: %d", score))
	case score < 60:
		g = "E"
	case score <= 70:
		g = "D"
	case score <= 80:
		g = "C"
	case score <= 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const (
		filename = "a.txt"
	)

	//contents, err := os.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	if contents, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	println(
		grade(75),
		grade(85),
		grade(100),
		grade(-3),
	)

}
