package main

import (
	"fmt"
)

type Address struct {
	flatNum int
	city    string
}

func (a Address) GetAddress() {
	fmt.Println("GetAddress() Address :", a)
}

type ID struct {
	name    string
	age     int
	Address // embedded 'Address' into 'ID' struct
	Company // embedded
}

func (i ID) GetDetails() {
	fmt.Println("GetDetails() Details :", i)
}

type Company struct {
	name   string
	salary int
}

func main() {
	// embedding and promotion
	fmt.Println("\tembedding and promotion")
	var id ID = ID{ // creating an instance
		name:    "Abc",
		age:     23,
		Address: Address{201, "Bluru"},
		Company: Company{"Airbus", 200000},
	}

	fmt.Println("Address :", id.Address)
	fmt.Println("Flat number :", id.flatNum) // promotion : as embedded 'Address' fields can be accessed by 'ID' instance
	fmt.Println("Name :", id.name)
	fmt.Println("Company name :", id.Company.name) // two field as 'name' so specify struct

	// calling methods
	id.GetDetails()
	id.GetAddress() // promotion : as GetAddress method is called using 'ID' instance

}