package main

import "fmt"

func main() {
	var Moviename []string
	Moviename = []string{"One Piece Film Z", "Spirited Away", "My Neighbor Totoro"}

	Moviename = append(Moviename, "Ponyo", "Hippo")
	fmt.Println(Moviename)

	DelMovie := Moviename[:2]
	fmt.Println(DelMovie)
}
