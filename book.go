package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	author1 := Author{Name: "J.K Williams", Age: 62}
	book1 := Book{Title: "Introduction to Algorithm", Author: author1}

	author2 := Author{Name: "Vladimir Nabokov", Age: 85}
	book2 := Book{Title: "Lolita", Author: author2}

	author3 := Author{Name: "Willian Faukner", Age: 75}
	book3 := Book{Title: "Palefire", Author: author3}

	author4 := Author{Name: "Sharlok Homes", Age: 65}
	book4 := Book{Title: "Sharlock Homes", Author: author4}

	fmt.Printf("%+v", book1)

	byteArray, err := json.MarshalIndent(book1, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	byteArray1, err := json.MarshalIndent(book2, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray1))

	byteArray2, err := json.MarshalIndent(book3, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray2))

	byteArray3, err := json.MarshalIndent(book4, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray3))
}
