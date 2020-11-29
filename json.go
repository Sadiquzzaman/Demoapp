package main

import (
	"encoding/json"
	"fmt"
)

type Mall struct {
	MallName MallName `json:"MallName"`
	location string   `json:"location"`
}

type MallName struct {
	name  string `json:"name"`
	movie Movie  `json:"movie"`
}

type Movie struct {
	movieName    string `json:"name"`
	directorName string `json:"directorName"`
	showtime     string `json:"showtime"`
}

func main() {
	TeNet := Movie{movieName: "TeNet", directorName: "Nolan", showtime: "11.00 AM  &&  5.00 PM"}
	Dark := Movie{movieName: "Dark", directorName: "Nolan", showtime: "11.00 AM  &&  5.00 PM"}
	Fightclub := Movie{movieName: "Fight club", directorName: "KalaManik", showtime: "10.00 AM  &&  5.00 PM"}
	Joker := Movie{movieName: "Joker", directorName: "Edge", showtime: "10.00 AM  &&  5.00 PM"}

	Zamuna1 := MallName{name: "Zamuna Future Park", movie: TeNet}
	Bashundhara1 := MallName{name: "Bashundhara Shopping Mall", movie: Dark}
	Padma1 := MallName{name: "Padma Mall", movie: Fightclub}
	shimanto1 := MallName{name: "Shimanto Cafe", movie: Joker}

	zamuna := Mall{MallName: Zamuna1, location: "Badda"}
	Bashundhara := Mall{MallName: Bashundhara1, location: "Panthapath"}
	Padma := Mall{MallName: Padma1, location: "Dhanmondi"}
	Shimanto := Mall{MallName: shimanto1, location: "Gulshan"}

	fmt.Printf("%+v\n", zamuna)
	fmt.Printf("%+v\n", Bashundhara)
	fmt.Printf("%+v\n", Padma)
	fmt.Printf("%+v\n", Shimanto)

	byteArray1, err := json.MarshalIndent(zamuna, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray1))

	byteArray2, err := json.MarshalIndent(Bashundhara, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray2))

	byteArray3, err := json.MarshalIndent(Padma, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray3))

	byteArray4, err := json.MarshalIndent(Shimanto, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray4))

}
