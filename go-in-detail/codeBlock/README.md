# code block and shadowing

# code block and scoping

Accessibility of variables declared within a block of code, (within a set of curly braces {}). 

Go has lexical block scope, meaning the scope of a variable is limited to the block in which it is declared.

Types of blocks scope :
a) Package-level scope : variables declared at package-level can be accessed anywhere in code.

b) Function level : variables declared in func() can be accessed anywhere in that function.

c) statement-level : accessible with that statement if{} , for{} etc.

d) block litral (anonyymous block) : accessable with that {}.

NOTE : variables declared 'outer block' are accessible within 'inner block' but reverse is not true.

# shadowing

Shadowing happens when you name inner vaiable same name as outer variable.

The inner variable takes priority over outer variable.

NOTE : Avoid shadowing to avoid confusion.



