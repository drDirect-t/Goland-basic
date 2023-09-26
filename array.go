package main

import "fmt"

var Moviename [5]string

func main() {
	Moviename[0] = "The Fast and the Furious 3 Tokyo Drift"
	Moviename[1] = "One Piece Film Z"
	Moviename[2] = "Spirited Away"
	Moviename[3] = "My Neighbor Totoro"
	Moviename[4] = "Ponyo"

	fmt.Println("Movie: ", Moviename)
}
