package main

import "fmt"

func main() {
	var Moviename []string
	Moviename = []string{"One Piece Film Z", "Spirited Away", "My Neighbor Totoro"}

	Moviename = append(Moviename, "Ponyo", "Hippo")
	fmt.Println(Moviename)

	fmt.Println(Moviename[:4])
	fmt.Println(Moviename[1:4])

	// DelMovie := Moviename[:2]
	// fmt.Println(DelMovie)
}
