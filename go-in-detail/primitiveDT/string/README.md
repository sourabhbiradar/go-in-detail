# string

string data type

used for text , characters , symbols and emojis

value/data to be enclosed within double quotes " "

strings are stored as byte array internally and are immutable.

back ticks ` ` : used to store raw data 
note: not single quotes

- main.go covers basics of string data type
- string.go covers strings in depth

# strings in depth

Strings are stored as sequence of bytes.Hence, strings are kind of array.

Each byte can be accessed by index.

When dealin with unicodes or other lanuages its BETTER to convert string into rune/byte than type cast them.
OR use strings or utf8 packages.