# call by value 
# call by reference

or pass by value

func arguements are passed by value

func params are copy of actual variables. Hence ,functions work on copy & NOT on actual variables' value.

Another of passing arguements to func is , 'call by refrence'..

Here, pointer/address of variable is passed.

Even though pointer/address are passed , go creates COPY of pointer/address when passed to func.

Same with slice and map.

Appending a slice in func will NOT effect actual slice variable.

But deleting/appending map effects,so its NOT recommanded to pass maps as params.