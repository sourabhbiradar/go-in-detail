# slices

Slices are basically dynamic arrays , meaning slices dont have fixed length.

NOTE  : size/length of slice is not declared otherwise Go will create array of that size.

Unlike arrays , elements can be added to slices.

Slices can be declared same as arrays. And also can be done using make() function.

len() and cap() functions are used to get size and capacity of slice.

Once the capacity is reached Go creates a new slice with more capacity and copies all data to it . Old one is discared.

To add elements to slice we use append() function.
Appending == Adding
WE can append slices to slice.

nil slice are different from empty slice.
nil slice has no underlying array.
empty slice has underlyin array with no value stored.

