# bits and bitwise operators

Bits are smallest unit of computer memory

Bit [0] is called unset and bit [1] is called set.

Bitwise operators are used for low-level bit manipulation operations.

Bitwise logical operators :

a) OR --> '|'
    0  0 --> 0
    1  0 --> 1
    0  1 --> 1
    1  1 --> 1
NOTE : atleast one bit should be set [1] then result is set [1]

b) AND --> '&'
   0  0 --> 0
   1  0 --> 0
   0  1 --> 0
   1  1 --> 1
NOTE : atleast one bit shud b unset [0] for result to b unset [0]

c) XOR --> '^'
   0  0 --> 0
   1  0 --> 1
   0  1 --> 1
   1  1 --> 0
NOTE : if both bit are same result is unset [0] and set [1] otherwise

d) AND NOT --> '&^'
    0  0 --> 0
    0  1 --> 0
    1  0 --> 1
    1  1 --> 1

e) INVERT/NOT --> '^'
   1 --> 0
   0 --> 1

f) Shift operators --> '<<' '>>'
   << --> Left shift add bits to integer value 
   >> --> Right shift removes bits
NOTE : 
