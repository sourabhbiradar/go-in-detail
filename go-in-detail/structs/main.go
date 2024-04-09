package main

import "fmt"

// defining struct
type details struct { // struct name 'student'
	name  string
	class string // field name 'class' of type string
	ID    int
	marks []float32
}

func main() {
	// structs

	// declaration n intialization
	// explicit --> field name : value
	var student1 details = details{name: "Abc", class: "8th", ID: 2012022, marks: []float32{87, 76, 91}}
	fmt.Println("student1 details :", student1) // no field names when printed

	// implicit --> no field name
	student2 := details{"Xyz", "6th", 2014038, []float32{54, 81, 33}}
	fmt.Printf("student2 details : %+v\n", student2) // "%+v" prints values n field names

	// explicit vertical
	student3 := details{
		name:  "EFG",
		class: "3rd",
		ID:    2017043,
	}
	fmt.Println("student3 details :", student3)
	// since marks field is omitted last field is empty slice []

	// implicit vertical
	var student4 = details{
		"IJK",
		"9th",
		2011108,
		[]float32{92, 94, 99},
	}
	fmt.Println("student4 details :", student4)
	// you can NOT omit fields in implicit initialization

	// accessing fields of struct
	fmt.Println("ID of student2 :", student2.ID) // varible.field_name

	// editing fields
	student3.marks = []float32{34, 78, 65}
	fmt.Println("updated student3 marks field :", student3) // now no empty marks field

}