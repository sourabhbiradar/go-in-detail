# integer data type 

integers in depth.

'int(value)' can also be used to convert other numberic data type or value to integer --> type casting

Usually 'int' is used in most cases.

Using appropriate data type can save up system memory so use integer data types accordingly.

Specific int types :

a) Unsigned integers : (NO NEGETIVE NUMBERS)

   uint8  --> 0 to 255
   uint16 --> 0 to 65535
   unit32 --> 0 to 4294967295
   unit64 --> 0 to 18446744073709551615

b) Signed integers : 

   int8  --> -128 to 127
   int16 --> -32768 to 32767 
   int32 --> -2147483648 to 2147483647
   int64 --> -9223372036854775808 to 9223372036854775807

c) Machine dependent :
   
   int --> int32 for 32-bit PC
           int64 for 64-bit 
   uint --> uint32 for 32-bit PC
            uint64 for 64-bit

d) Rune alias for int32 used to represent Unicode code points and characters.

e) Byte alias for uint8 represents raw data and used when working  with ASCII characters and byte data.
   