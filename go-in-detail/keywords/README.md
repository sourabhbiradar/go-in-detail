# KEYWORDS
Read here about Keywords or copy-paste keywords.go in IDE

There are 25 KEYWORDS in Go
 Following are 'KEYWORDS' in Go and their general use :
 
 1. 'var' : used to declare variable
 var name type
 
 2. 'const' : declares variable as consatant
 const var_name type
 
 3. 'map' : composite data type

  map[type1]type2
 
 4. 'if' : used to create if-else statement

 if condition {
 }else if{
 }else{
 }
 
 5. 'else' : used in 'if statements' ,if else,if else if.
 
 else
 if condition {
 }else if{
 }else{
 }
 
 6. 'for' : used to create for statement.  

 for i := 0; i < 5; i++ {
 }
 
 7. 'range' : used in for-range statement  iterating over slice, array, map, channel, etc.

 for i,v := range above_types{
 } 

 8. 'break' : used to BREAK out of loops

  for {
 break
 }
 
 9. 'continue' : used in loops to skip a particular instance and continue with itration

  for i := 0; i < 5; i++ {
 if i == 3 {
 continue
 }
 fmt.Println(i)
 }
 
 10. 'switch' : switch statement

 switch {
 case :
 default :
 }
 
 11. 'case' : (~conditions) used in switch and select 
 
 statements
  switch {
 case :
 }
 
 12. 'default' : used in switch statement ,as last case --> if none of the cases are true thn execute DEFAULT case

 switch {
 case :
 default :
 }
 
 13. 'fallthrough' : used in switch statement , after execution of suitable case fallthrough forces to execute rest cases

 switch {
 case :
 fallthrough
 case == true :
 fallthrough
 case :
 fallthrough
 default :
 }
 
 14. 'goto' : goto statement allows unconditional jump to a labeled statement

 	goto g
 h:
 fmt.Println(1)
 g:
 fmt.Println(2)
 goto h
 
 15. 'select' : switch statement for goroutines

 select {
 case <-ch :
 }
 
 16. 'func' : short for Function ,used to create a function

 func name()
 
 17. 'return' : used to return value from func() can be named or naked

 func() return_type{
 return value
 }
 
 18. 'defer' : forces func() associated with to execute at the last (LIFO) 

 defer func()
 
 19. 'package' : https://go.dev/blog/package-names 

 package main
 
 20. 'import' :  imports the specified package from the directory of $GOPATH (if no path is mentioned) or else from the mentioned directory

 import ''fmt''

 import (
 ''fmt''
 ''math/rand''
 )
 
 21. 'type' : used to create struct , interface and custom data type 


 type name struct{}
 type name interface{}
 type name any_type
 type name custom_type
 
 22. 'struct' : a user-defined type to store a collection of different fields into a single field

 type name struct{
 x type1
 s type2
 }
 
 23. 'interface' : umbrella for methods

  type name interface{
 method1()
 method(x type)return_type
 }
 
 24. 'go' : added before func() to make it goroutine

  go func()
 
 25. 'chan' : used to crerate channel 
 ch := make (chan type)