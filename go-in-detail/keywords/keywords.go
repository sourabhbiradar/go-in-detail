package main

import "fmt"

func main() {
	fmt.Println(
		"There are 25 keywords in Go\n",
		"Following are 'KEYWORDS' in Go and their general use :\n",
		"\n",

		"var : used to declare variable\n",
		"var name type\n",
		"\n",

		"const : declares variable as consatant\n",
		"const var_name type\n",
		"\n",

		"map : data type\n ",
		"map[type1]type2\n",
		"\n",

		"if : used to create if-else statement\n",
		"if condition {\n",
		"}else if{\n",
		"}else{\n",
		"}\n",
		"\n",

		"else : used in 'if statements' ,if else,if else if else\n",
		"if condition {\n",
		"}else if{\n",
		"}else{\n",
		"}\n",
		"\n",

		"for : used to create for statement \n",
		"for i := 0; i < 5; i++ {\n",
		"}\n",
		"\n",

		"range : used in for-range statement ",
		"iterating over slice, array, map, channel, etc.\n",
		"for i,v:=range above_types{\n",
		"}",
		"\n",

		"break : used to BREAK out of loops\n ",
		"for {\n",
		"break\n",
		"}\n",
		"\n",

		"continue : used in loops to skip a particular instance and continue with itration\n ",
		"for i := 0; i < 5; i++ {\n",
		"if i == 3 {\n",
		"continue\n",
		"}\n",
		"fmt.Println(i)\n",
		"}\n",
		"\n",

		"switch : short version of if-else\n",
		"switch {\n",
		"case :\n",
		"default :\n",
		"}\n",
		"\n",

		"case : (~conditions) used in switch and select statements\n ",
		"switch {\n",
		"case :\n",
		"}\n",
		"\n",

		"default : used in switch statement ,as last case --> if none of the cases are true thn execute DEFAULT case\n",
		"switch {\n",
		"case :\n",
		"default :\n",
		"}\n",
		"\n",

		"fallthrough : used in switch statement , after execution of suitable case fallthrough forces to execute rest cases\n",
		"switch {\n",
		"case :\n",
		"fallthrough\n",
		"case == true :\n",
		"fallthrough\n",
		"case :\n",
		"fallthrough\n",
		"default :\n",
		"}\n",
		"\n",

		"goto : goto statement allows unconditional jump to a labeled statement\n",
		"	goto g\n",
		"h:\n",
		"fmt.Println(1)\n",
		"g:\n",
		"fmt.Println(2)\n",
		"goto h\n",
		"\n",

		"select : switch statement for goroutines \n",
		"select {\n",
		"case <-ch :\n",
		"}\n",
		"\n",

		"func : short for Function ,used to create a function \n",
		"func name()\n",
		"\n",

		"return : used to return value from func() can be named or naked\n",
		"func() return_type{\n",
		"return value\n",
		"}\n",
		"\n",

		"defer : forces func() associated with to execute at the last (LIFO) \n",
		"defer func()\n",
		"\n",

		"package : https://go.dev/blog/package-names \n",
		"package main\n",
		"\n",

		"import :  imports the specified package from the directory of $GOPATH (if no path is mentioned) or else from the mentioned directory\n",
		"import ''fmt''\n",
		"import (\n",
		"''fmt''\n",
		"''math/rand''\n",
		")\n",
		"\n",

		"type : used to create struct , interface and custom data type\n",
		"type name struct{}\n",
		"type name interface{}\n",
		"type name custom_type\n",
		"\n",

		"struct :a user-defined type to store a collection of different fields into a single field\n",
		"type name struct{\n",
		"x type1\n",
		"s type2\n",
		"}\n",
		"\n",

		"interface : umbrella for methods\n ",
		"type name interface{\n",
		"method1()\n",
		"method(x type)return_type\n",
		"}\n",
		"\n",

		"go : added before func() to make it goroutine\n ",
		"go func()\n",
		"\n",

		"chan : used to crerate channel \n",
		"ch := make (chan type)\n",
	)
}
