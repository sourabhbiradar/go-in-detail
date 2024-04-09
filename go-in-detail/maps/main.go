package main

import "fmt"

func main() {
	// maps

	// Declaration and intizlization

	// way 1
	var map1 = map[int]string{ // map[key_type]value_type
		1: "a", // key-value seperated by colon ':'
		2: "b", // key:value,
		3: "c", // last value also has comma
	}
	fmt.Println("map1 :", map1)

	// way 2 using make()
	var map3 map[int]int = make(map[int]int, 3)
	map3[1] = 99
	map3[2] = 98
	map3[3] = 97
	map3[4] = 96
	fmt.Println("map3 :", map3)

	// way 3 using short hand
	names := map[string]string{ // key type : string & value type :string
		"Cristiano": "Ronaldo",
		"Virat":     "Kohli",
	}
	fmt.Println("Full Names :", names)

	// accessin value usin key
	fmt.Println("accessin value usin key :", names["Virat"])

	// adding elements to map
	names["Sunil"] = "Chettri"
	fmt.Println("map with new element :", names)

	// deleting elements
	delete(names, "Cristiano")
	fmt.Println("map with deleted element :", names)

	// checkin if key exists
	_, ok := names["Vishi"] // replace ["Vishi"] with ["Virat"]
	fmt.Println("Does key exist? :", ok)

	// array as value type
	list := map[int][3]string{
		1: [3]string{"two", "red", "roses"}, // optional [3]string
		2: {"one", "black", "pearl"},
		3: {"six", "pink", "stars"},
	}
	fmt.Println("value of type array :", list)

	// accessing single value
	fmt.Println("single value from value array :", list[3][2]) // stars

	// nested map
	date := make(map[int]map[string]string)

	date[1] = make(map[string]string)
	date[1]["June"] = "2016"
	date[1]["May"] = "1987"
	date[1]["August"] = "1947"

	date[2] = make(map[string]string)
	date[2]["Monday"] = "Jan"
	date[2]["Sunday"] = "Oct"
	date[2]["Friday"] = "Sept"

	fmt.Println("nested map :", date)

}