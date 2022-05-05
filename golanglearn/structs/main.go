package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	uye := person{}
	alex := person{
		lastName:  "sasf",
		firstName: "asas",
		contactInfo: struct{
			email: "asqwewe@a.adsa",
			zip:   23131,
		},
	}
	uye.lastName = "sasasasasasf"
	uye.firstName = "asasasas"
	uye.contactInfo.zip = 23
	// fmt.Println(alex)
	fmt.Printf("%+v", alex)
	fmt.Printf("%+v", uye)
}
