package main

import "fmt"

var phone = make(map[string]int)

func main() {

	//add data
	phone["Iphone 13"] = 24000
	phone["Iphone 15"] = 32000
	phone["Iphone 88"] = 240000
	fmt.Println("Phone :", phone)

	//delete data
	delete(phone, "Iphone 88")
	fmt.Println("Phone :", phone)

	//update data
	phone["Iphone 13"] = 25990
	fmt.Println("Phone :", phone)

	//default value for map
	PhoneName := map[string]string{"1": "Samsung", "2": "Iphone"}
	fmt.Println("PhoneName :", PhoneName)
}
