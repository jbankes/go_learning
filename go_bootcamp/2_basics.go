package main

import "fmt"

func main() {
	name, location := "prince oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s", name, age, location)
}

// var (
//   name string = "prince oberyn"
//   age int = 32
//   location string = "Dorne"
// )

// can also be
// var (
//   name1, location1 string
//   age1 int
// )

// var name2 string
// var age2 int
// var location2 string

// var (
//   name = "prince oberyn"
//   age = 32
//   location = "Dorne"
// )

// var (
//   name, age, location = "prince oberyn", 32, "Dorne"
// )
